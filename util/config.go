package util

import (
	"fmt"
	"github.com/spf13/viper"
)

func SpawnHyperspaceViper (paths ...string) *viper.Viper {
	v := viper.New()
	v.SetConfigName("hyperspace")
	v.SetConfigType("yaml")
	for _, path := range paths {
		v.AddConfigPath(path)
		// adding this so we can run from the cmdscripts dir
		v.AddConfigPath(fmt.Sprintf("../%s", path))
	}
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return v
}

