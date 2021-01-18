package util

import (
	"gopkg.in/yaml.v2"
	"log"
)

//----------------------------------------------------------------------------------
// MergeStructs()
//----------------------------------------------------------------------------------
// Merges two structs by marshaling the override struct into bytes as yaml, then
// unmarshalling the bytes into the template struct via the yaml encoder.
// `template` is the struct to be overridden, the `override` struct overrides.
//----------------------------------------------------------------------------------
func MergeStructs (template, override interface{}, msgArgs ...string) {
	d, err := yaml.Marshal(override)
	if err != nil {
		log.Fatalf("(%s)\nerror marshaling struct for merge: %v", msgArgs, err)
	}

	err = yaml.Unmarshal(d, template)
	if err != nil {
		log.Fatalf("(%s)\nerror unmarshing override struct into template for merge: %v", msgArgs, err)
	}
}

