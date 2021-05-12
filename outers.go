package outers

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	EnvOutersDir  = "OUTERS_DIR"
	KeyDefaultKey = "default"
)

func Load(key, kind string, out interface{}) (err error) {
	if key == "" {
		key = KeyDefaultKey
	}
	var (
		baseDirs   = []string{"."}
		subDirs    = []string{".", "conf", "config"}
		extensions = []string{".yaml", ".yml"}
	)
	if envDir := strings.TrimSpace(os.Getenv(EnvOutersDir)); envDir != "" {
		baseDirs = append([]string{envDir}, baseDirs...)
	}
	if execPath, _ := os.Executable(); execPath != "" {
		baseDirs = append(baseDirs, filepath.Dir(execPath))
	}
	for _, baseDir := range baseDirs {
		for _, subDir := range subDirs {
			for _, ext := range extensions {
				file := filepath.Join(baseDir, subDir, key+"."+kind+ext)
				var buf []byte
				if buf, err = ioutil.ReadFile(file); err != nil {
					if os.IsNotExist(err) {
						err = nil
						continue
					}
				}
				if buf, err = Render(buf); err != nil {
					return
				}
				if err = yaml.Unmarshal(buf, out); err != nil {
					return
				}
				return
			}
		}
	}
	err = fmt.Errorf("missing yaml config file for %s.%s", key, kind)
	return
}
