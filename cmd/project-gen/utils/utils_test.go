package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseTemplateJSON(t *testing.T) {
	rv, err := parseTemplateJSON("./testdata/kratos.json")
	assert.NoError(t, err)
	fmt.Printf("%v\n", rv)
}

func TestGenerate(t *testing.T) {
	err := GenerateFromTemplateProject("testdata", "/tmp/new_project")
	assert.NoError(t, err)
}

func TestClone(t *testing.T) {
	err := CloneProjectTemplate("https://github.com/go-kratos/proposal", "/tmp/pp")
	assert.NoError(t, err)
}