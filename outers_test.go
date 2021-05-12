package outers

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

type testLoadStruct struct {
	H string `yaml:"hello"`
}

func TestLoad(t *testing.T) {
	os.Setenv("OUTERS_DIR", "testdata")
	var s testLoadStruct
	err := Load("", "test", &s)
	require.NoError(t, err)
	require.Equal(t, "world", s.H)
}
