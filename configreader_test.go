package config

import (
	"bytes"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tomlMask = []byte(`
A=1
C=1

[table]
name = "Mask"`)

var tomlBase = []byte(`
A=2
B=1

[table]
name = "Base"`)

var config Config

func init() {

	arr := [2]*viper.Viper{Build(tomlMask), Build(tomlBase)}
	config = NewFromVipers(arr[:])
}

func TestTopLevel(t *testing.T) {

	assert.Equal(t, 1, config.GetInt("A"))
	assert.Equal(t, 1, config.GetInt("B"))
	assert.Equal(t, 1, config.GetInt("C"))
}

func TestNestedLevel(t *testing.T) {

	assert.Equal(t, "Mask", config.GetString("table.name"))
}

func Build(config []byte) *viper.Viper {

	r := bytes.NewReader(config)
	v := viper.New()
	v.SetConfigType("toml")
	v.ReadConfig(r)
	return v
}
