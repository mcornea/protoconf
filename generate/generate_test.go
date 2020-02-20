package generate

import (
	"go/build"
	"log"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	gopath := build.Default.GOPATH
	log.Println(gopath)
	gopathabs, _ := filepath.Abs(gopath)
	gopathArr := filepath.SplitList(gopathabs)
	gopathArr = append(gopathArr, "src")
	gopath = filepath.Join(gopathArr...)
	log.Println(gopath)

	// p, err := NewPackageImporter("github.com/prometheus/prometheus/config", "Job", gopath)
	p, err := NewPackageImporter("github.com/hashicorp/nomad/api", "Job", gopath)
	assert.NoError(t, err)

	p.Visit()
	t.Fail()
}
