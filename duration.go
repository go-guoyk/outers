package outers

import (
	"gopkg.in/yaml.v3"
	"time"
)

type Duration time.Duration

func (d *Duration) UnmarshalYAML(value *yaml.Node) (err error) {
	var s string
	if err = value.Decode(&s); err != nil {
		return err
	}
	var raw time.Duration
	if raw, err = time.ParseDuration(s); err != nil {
		return
	}
	*d = Duration(raw)
	return
}

func (d Duration) MarshalYAML() (interface{}, error) {
	return d.Unwrap().String(), nil
}

func (d Duration) Unwrap() time.Duration {
	return time.Duration(d)
}
