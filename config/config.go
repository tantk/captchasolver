package config

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"path/filepath"
	"runtime"
)

//Configuration :
type Configuration struct {
	RESTport   string
}
var (
	_, b, _, _ = runtime.Caller(0)
	// Root folder of this project
	Root = filepath.Join(filepath.Dir(b), "../")
)

//GetConfig gets env var from a json file
func GetConfig(params ...string) Configuration {
	configuration := Configuration{}
	fileName := fmt.Sprintf(Root + "/config/conf.json")
	gonfig.GetConf(fileName, &configuration)
	return configuration
}