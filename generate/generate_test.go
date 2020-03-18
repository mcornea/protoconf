package generate

import (
	"log"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var gopath string

func TestGenerate(t *testing.T) {
	// gopath := build.Default.GOPATH
	log.Println(gopath)
	gopathabs, err := filepath.Abs(gopath)
	assert.NoError(t, err)
	gopathArr := filepath.SplitList(gopathabs)
	// gopathArr = append(gopathArr, "src")
	gopath = filepath.Join(gopathArr...)
	log.Println(gopath)

	// p, err := NewPackageImporter("github.com/prometheus/prometheus/config", "Job", gopath)
	// p, err := NewPackageImporter("github.com/hashicorp/nomad/api", "Job", gopath)
	p, err := NewPackageImporter("github.com/hashicorp/nomad/nomad/structs", "Job", gopath)
	// p, err := NewPackageImporter("testdata", "Job", gopath, "GOOS=linux", "GOARCH=amd64", "GOCACHE=/tmp/gocache", "GOPATH="+gopath)
	assert.NoError(t, err)

	p.Visit()
	t.Fail()
}
