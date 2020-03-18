package generate

import (
	"fmt"
	"go/token"
	"go/types"
	"log"
	"path/filepath"
	"strings"

	"github.com/fatih/structtag"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/builder"
	"github.com/jhump/protoreflect/desc/protoprint"
	"golang.org/x/tools/go/packages"
)

// PackageImporter is a tool to walk through go structs in a package and write them out into a proto file.
type PackageImporter struct {
	pkgs               []*packages.Package
	pkgID              string
	rootMsg            string
	protoFilesRegistry *protoFilesRegistry
	errors             []error
}

var wkt = make(map[string]*builder.FileBuilder)

func init() {
	durationFile, err := desc.LoadFileDescriptor("google/protobuf/duration.proto")
	if err != nil {
		log.Fatal(err)
	}
	durationBuilder, err := builder.FromFile(durationFile)
	if err != nil {
		log.Fatal(err)
	}
	wkt["duration"] = durationBuilder

	structFile, err := desc.LoadFileDescriptor("google/protobuf/struct.proto")
	if err != nil {
		log.Fatal(err)
	}
	structBuilder, err := builder.FromFile(structFile)
	if err != nil {
		log.Fatal(err)
	}
	wkt["struct"] = structBuilder
}

// NewPackageImporter creates a new PackageImporter
func NewPackageImporter(pkg, rootMsg, goSrcPath string, env ...string) (*PackageImporter, error) {
	fset := token.NewFileSet()
	cfg := &packages.Config{
		Dir:  goSrcPath,
		Mode: packages.LoadSyntax,
		Fset: fset,
		Env: env,
	}
	packages, err := packages.Load(cfg, filepath.Join(goSrcPath, pkg))
	if err != nil {
		return nil, err
	}
	log.Println(packages)
	return &PackageImporter{
		pkgID:              pkg,
		pkgs:               packages,
		rootMsg:            rootMsg,
		protoFilesRegistry: newProtoFilesRegistry(),
		errors:             []error{},
	}, nil
}

func goFieldTypeToProtoFieldType(x string) *builder.FieldType {
	x = strings.Replace(x, "*", "", -1)
	switch x {
	case "bool":
		return builder.FieldTypeBool()
	case "byte":
		return builder.FieldTypeBytes()
	case "float":
		return builder.FieldTypeFloat()
	case "float64":
		return builder.FieldTypeFloat()
	case "int":
		return builder.FieldTypeInt32()
	case "int8":
		return builder.FieldTypeInt32()
	case "int32":
		return builder.FieldTypeInt32()
	case "int64":
		return builder.FieldTypeInt64()
	case "uint":
		return builder.FieldTypeUInt32()
	case "uint8":
		return builder.FieldTypeUInt32()
	case "uint16":
		return builder.FieldTypeUInt32()
	case "uint32":
		return builder.FieldTypeUInt32()
	case "uint64":
		return builder.FieldTypeUInt64()
	case "string":
		return builder.FieldTypeString()
	case "time.Duration":
		return builder.FieldTypeMessage(wkt["duration"].GetMessage("Duration"))
	case "interface{}":
		return builder.FieldTypeMessage(wkt["struct"].GetMessage("Value"))
	case "error":
		return nil
	}
	return builder.FieldTypeMessage(wkt["struct"].GetMessage("Value"))
}

func (p *PackageImporter) goFieldToProtoField(f *types.Var, tagstr string) *builder.FieldBuilder {
	t1 := f.Type()
	pkg := f.Pkg()
	b := builder.NewField(f.Name(), builder.FieldTypeString())

	if tag, err := structtag.Parse(tagstr); err == nil {
		if y, e := tag.Get("yaml"); e == nil {
			b.SetJsonName(y.Name)
			if y.Name != "" {
				b.SetName(y.Name)
			}
		}
		if j, e := tag.Get("json"); e == nil {
			// b.SetName(j.Name)
			b.SetJsonName(j.Name)
		}
	}
	if strings.HasPrefix(f.Type().String(), "func") {
		return nil
	}
	t := goFieldTypeToProtoFieldType(t1.Underlying().String())

	// When a type is map
	if s, ok := t1.Underlying().(*types.Map); ok {
		t1 = s.Elem()
		key := goFieldTypeToProtoFieldType(s.Key().Underlying().String())
		val := goFieldTypeToProtoFieldType(s.Elem().Underlying().String())
		if key != nil && val != nil {
			b = builder.NewMapField(f.Name(), key, val)
		}
		log.Println(pkg.Name(), f.Name(), "detected as map of", t1.String())
	}

	// When a type is an Slice
	if s, ok := t1.Underlying().(*types.Slice); ok {
		t1 = s.Elem()
		t = goFieldTypeToProtoFieldType(s.Elem().String())
		b.SetRepeated()
		log.Println(pkg.Name(), f.Name(), "detected as slice of", t1.String())
	}
	// When a type is an array
	if s, ok := t1.Underlying().(*types.Array); ok {
		t1 = s.Elem()
		t = goFieldTypeToProtoFieldType(s.Elem().String())
		b.SetRepeated()
		log.Println(pkg.Name(), f.Name(), "detected as array of", t1.String())
	}

	// When a type is a pointer
	if s, ok := t1.Underlying().(*types.Pointer); ok {
		t1 = s.Elem()
		log.Println(pkg.Name(), f.Name(), "detected as pointer of", t1.String())
	}

	// When a type is another struct
	if _, ok := t1.Underlying().(*types.Struct); ok {
		a := strings.Split(t1.String(), ".")
		structName := a[len(a)-1]
		pkgID := strings.Join(a[:len(a)-1], ".")
		a = strings.Split(pkgID, "/")
		pkgName := a[len(a)-1]
		file := p.protoFilesRegistry.getProtoFile(pkgName, pkgID)
		msg := file.GetMessage(structName)
		if msg == nil {
			p.errors = append(p.errors, fmt.Errorf("could not load message for %v %v %v", pkgID, structName, t1))
			return nil
		}
		t = builder.FieldTypeMessage(msg)
		log.Println(pkg.Name(), f.Name(), "detected as struct of", t1.String())
	}

	if b.IsMap() {
		return b
	}
	if t != nil {
		log.Println(pkg.Name(), f.Name(), "detected as", t1.String())
		b.SetType(t)
		return b
	}
	p.errors = append(p.errors, fmt.Errorf("could not understand field for %v %v, underlying: %v", f.Name(), f.Type(), t1.Underlying()))
	return nil
}

type protoFilesRegistry struct {
	reg map[string]map[string]*builder.FileBuilder
}

func newProtoFilesRegistry() *protoFilesRegistry {
	return &protoFilesRegistry{
		reg: make(map[string]map[string]*builder.FileBuilder),
	}
}

func (p *protoFilesRegistry) nameFor(pkgName, pkgID string) string {
	if pkg, ok := p.reg[pkgName]; ok {
		pkgLen := len(pkg)
		if pkgLen > 0 {
			a := strings.Split(pkgID, "/")
			ret := strings.Join(a, "_")
			if len(a) > pkgLen {
				ret = strings.Join(a[len(a)-pkgLen:], "_")
			}
			return ret
		}
	}
	return pkgName
}
func (p protoFilesRegistry) getProtoFile(pkgName, pkgID string) *builder.FileBuilder {
	if pkg, ok := p.reg[pkgName]; ok {
		if file, ok := pkg[pkgID]; ok {
			return file
		}
		for _pkgID, file := range pkg {
			file.SetName(p.nameFor(pkgName, _pkgID) + ".proto")
			file.SetPackageName(p.nameFor(pkgName, _pkgID))
		}
	} else {
		p.reg[pkgName] = make(map[string]*builder.FileBuilder)
	}
	file := builder.NewFile(p.nameFor(pkgName, pkgID) + ".proto")
	file.SetPackageName(p.nameFor(pkgName, pkgID))
	file.SetOptions(&descriptor.FileOptions{GoPackage: &pkgID})
	file.SetProto3(true)
	p.reg[pkgName][pkgID] = file
	return file
}

func (p *PackageImporter) pre(pkg *packages.Package) bool {
	file := p.protoFilesRegistry.getProtoFile(pkg.Name, pkg.ID)
	for _, t := range pkg.TypesInfo.Defs {
		if t == nil {
			continue
		}
		if !t.Exported() {
			continue
		}
		if _, ok := t.Type().Underlying().(*types.Struct); ok {
			msg := builder.NewMessage(t.Name())
			file.TryAddMessage(msg)
		}
	}
	return true
}

func (p *PackageImporter) post(pkg *packages.Package) {
	if pkg.ID != p.pkgID {
		return
	}
	file := p.protoFilesRegistry.getProtoFile(pkg.Name, pkg.ID)
	for _, t := range pkg.TypesInfo.Defs {
		if t == nil {
			continue
		}
		if !t.Exported() {
			continue
		}
		if s, ok := t.Type().Underlying().(*types.Struct); ok {
			log.Println("*******", pkg.ID, t.Id(), "******")
			msg := file.GetMessage(t.Name())
			if msg == nil {
				log.Println("could not find message for", t.Name(), file.GetName())
				continue
			}
			for i := 0; i < s.NumFields(); i++ {
				f := s.Field(i)
				tag := s.Tag(i)
				if f.IsField() && f.Exported() {
					bf := p.goFieldToProtoField(f, tag)
					if bf != nil {
						msg.TryAddField(bf)
					}
				}
			}
		}
	}
	print(file)
}

// Visit will run the packages visit logic
func (p *PackageImporter) Visit() {
	packages.Visit(p.pkgs, p.pre, p.post)
	for _, err := range p.errors {
		log.Println(err)
	}
}

func print(b *builder.FileBuilder) {
	p := &protoprint.Printer{}
	desc, err := b.Build()
	if err != nil {
		log.Fatal(err)
	}
	str, err := p.PrintProtoToString(desc)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(str)
}
