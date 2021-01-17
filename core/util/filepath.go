package util

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path/filepath"
)

//----------------------------------------------------------------------------------
// GetPath()
//----------------------------------------------------------------------------------
// GetPath allows configuration strings that specify a (config-file) relative path
//
// For example: Assume our config is located in /etc/hyperledger/fabric/core.yaml with
// a key "msp.configPath" = "msp/config.yaml".
//
// This function will return:
//      GetPath("msp.configPath") -> /etc/hyperledger/fabric/msp/config.yaml
//
//----------------------------------------------------------------------------------
func GetPath(key string) string {
	p := viper.GetString(key)
	if p == "" {
		return ""
	}

	return TranslatePath(filepath.Dir(viper.ConfigFileUsed()), p)
}

//----------------------------------------------------------------------------------
// TranslatePath()
//----------------------------------------------------------------------------------
// Translates a relative path into a fully qualified path relative to the config
// file that specified it.  Absolute paths are passed unscathed.
//----------------------------------------------------------------------------------
func TranslatePath(base, p string) string {
	if filepath.IsAbs(p) {
		return p
	}

	return filepath.Join(base, p)
}

//----------------------------------------------------------------------------------
// FileOrDirectoryExists()
//----------------------------------------------------------------------------------
// Checks to see if a file or a directory exists at the given path. Returns a bool
// or error if operation did not go as expected.
//----------------------------------------------------------------------------------
func FileOrDirectoryExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return false, err
}

//----------------------------------------------------------------------------------
// ReadInYamlData()
//----------------------------------------------------------------------------------
// Reads in from a filepath and returns the byte data for that file
//----------------------------------------------------------------------------------
func ReadInYamlData(filePath string) ([]byte, error) {
	if filePath == "" {
		return nil, fmt.Errorf("filename cannot be empty")
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf( "error reading from filepath: %s\n%v", filePath, err)
	}

	return data, nil
}