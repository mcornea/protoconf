package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/dynamic"
	"protoconf.com/consts"
	protoconfvalue "protoconf.com/datatypes/proto/v1/protoconfvalue"
	"protoconf.com/protostdlib"
	"protoconf.com/protostdlib/secret"
)

// ReadConfig reads a materialized config
func ReadConfig(protoconfRoot string, configName string) (*protoconfvalue.ProtoconfValue, error) {
	filename := filepath.Join(protoconfRoot, consts.CompiledConfigPath, configName+consts.CompiledConfigExtension)

	configReader, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening config file, file=%s", filename)
	}
	defer configReader.Close()

	type configJSONType struct {
		ProtoFile string
	}
	var configJSON configJSONType
	if err = json.NewDecoder(configReader).Decode(&configJSON); err != nil {
		return nil, err
	}

	parser := &protoparse.Parser{ImportPaths: []string{filepath.Join(protoconfRoot, consts.SrcPath)}, ProtoStdLib: protostdlib.ProtoStdLib}
	descriptors, err := parser.ParseFiles(configJSON.ProtoFile)
	if err != nil {
		return nil, fmt.Errorf("error parsing proto file, file=%s err=%v", configJSON.ProtoFile, err)
	}
	anyResolver := dynamic.AnyResolver(nil, descriptors[0])

	if _, err = configReader.Seek(0, 0); err != nil {
		return nil, err
	}

	protoconfValue := &protoconfvalue.ProtoconfValue{}
	um := jsonpb.Unmarshaler{AnyResolver: anyResolver}
	if err = um.Unmarshal(configReader, protoconfValue); err != nil {
		return nil, fmt.Errorf("error unmarshaling, err=%s", err)
	}

	if err = updateSecrets(protoconfValue, anyResolver); err != nil {
		return nil, err
	}

	return protoconfValue, nil
}

func MessageFQN(msg descriptor.Message) string {
	fileDesc, protoDesc := descriptor.ForMessage(msg)
	fqn := protoDesc.GetName()
	if fileDesc.GetPackage() != "" {
		fqn = fileDesc.GetPackage() + "." + fqn
	}
	return fqn
}

func updateSecrets(protoconfValue *protoconfvalue.ProtoconfValue, anyResolver jsonpb.AnyResolver) error {
	name, err := ptypes.AnyMessageName(protoconfValue.Value)
	if err != nil {
		return err
	}

	value, err := anyResolver.Resolve(name)
	if err != nil {
		return err
	}

	message, err := dynamic.AsDynamicMessage(value)
	if err != nil {
		return err
	}

	secretFQN := MessageFQN(&secret.Secret{})
	visitor := func(pos int, len int, msgDesc *desc.MessageDescriptor) {
		if msgDesc.GetFullyQualifiedName() == secretFQN {
			protoconfValue.Secrets = append(protoconfValue.Secrets, &protoconfvalue.SecretMetadata{Pos: int32(pos), Len: int32(len)})
		}
	}

	d := decoder{msgDesc: message.GetMessageDescriptor(), visitor: visitor}
	if err := d.Unmarshal(protoconfValue.Value.Value); err != nil {
		return err
	}

	return nil
}
