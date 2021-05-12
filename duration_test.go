package outers

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"testing"
	"time"
)

type testStructD struct {
	D Duration `yaml:"d"`
}

func TestDuration_Duration(t *testing.T) {
	d := Duration(time.Second * 3)
	td := testStructD{
		D: d,
	}
	buf, _ := yaml.Marshal(td)
	assert.Equal(t, "d: 3s\n", string(buf))
	var td1 testStructD
	_ = yaml.Unmarshal(buf, &td1)
	assert.Equal(t, time.Second*3, td1.D.Unwrap())
	err := yaml.Unmarshal([]byte(`{}`), &td1)
	assert.NoError(t, err)
}
