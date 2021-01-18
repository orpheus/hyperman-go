package util

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

//----------------------------------------------------------------------------------
// MergeStructs()
//----------------------------------------------------------------------------------
// Merges two structs by marshaling the override struct into bytes as yaml, then
// unmarshalling the bytes into the template struct via the yaml encoder.
// `template` is the struct to be overridden, the `override` struct overrides.
//----------------------------------------------------------------------------------
func MergeStructs (template, override interface{}, msgArgs ...string) {
	// todo: add error handling
	// check to see if structs are the same, if they are return error or panic

	d, err := yaml.Marshal(override)
	if err != nil {
		log.Fatalf("(%s)\nerror marshaling struct for merge: %v", msgArgs, err)
	}

	err = yaml.Unmarshal(d, template)
	if err != nil {
		log.Fatalf("(%s)\nerror unmarshing override struct into template for merge: %v", msgArgs, err)
	}
}

//----------------------------------------------------------------------------------
// UnmarshalYaml()
//----------------------------------------------------------------------------------
// Unmarshal a yaml file given a filePath into a config struct. If reading in the
// yaml file unmarshalling into the config errors, Fatal will be called. os.exit(1)
//----------------------------------------------------------------------------------
func UnmarshalYaml (filePath string, config interface{}, msgArgs ...string) {
	yamlBytes, err := ReadInYamlData(filePath)
	if err != nil {
		log.Fatalf("(%v)\nfailed to read in yaml from path: %v\n%v", msgArgs, filePath, err)
	}

	err = yaml.Unmarshal(yamlBytes, config)
	if err != nil {
		log.Fatalf("(%v)\nerror unmarshalling yaml into config: %v\n%v", msgArgs, filePath, err)
	}
}


//----------------------------------------------------------------------------------
// MarshalAndWriteYaml()
//----------------------------------------------------------------------------------
// Marshals a struct config into bytes via a yaml encoder and then writes the bytes
// to the given filePath using ioutil. If either of these operations fails, Fatal
// will be called, exiting the program with os.exit(1).
//----------------------------------------------------------------------------------
func MarshalAndWriteYaml (config interface{}, filePath string, perm os.FileMode, msgArgs ...string) {
	d, err := yaml.Marshal(config)
	if err != nil {
		log.Fatalf("(%v)\nerror marshalling config to bytes from: %v\n%v", msgArgs, filePath, err)
	}

	err = ioutil.WriteFile(filePath, d, perm)
	if err != nil {
		log.Fatalf("(%v)\nioutil error writing yaml to: %v\n%v", msgArgs, filePath, err)
	}
}