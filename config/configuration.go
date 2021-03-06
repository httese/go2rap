package config

import (
	"io/ioutil"
	"encoding/json"
)

//配置文件结构
type CfgJson struct {
	Servers    []Server
	Conditions []Condition
}

//服务器代理配置
type Server struct {
	Name  string
	Host  string
	Proxy string
}

//条件转发配置
type Condition struct {
	ServerA string
	ServerB string
	//rap(mock平台)前缀路径，例如：/mockjsdata/1
	PrefixPath string
	Path       []string
}

//读取并解析配置文件
func ReadCfg(filename string) (*CfgJson, error) {

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cfgJson CfgJson

	if err := json.Unmarshal(bytes, &cfgJson); err != nil {
		return nil, err
	}

	return &cfgJson, nil
}
