package core

import (
	"gopkg.in/yaml.v2"
	"io/fs"
	"io/ioutil"
	"log"
	"materialsSort/config"
	"materialsSort/global"
)

const ConfigFile = "settings.yaml"

// InitConfig 初始化配置文件
func InitConfig() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)

	if err != nil {
		log.Fatalf("config file read error: %v", err)
	}
	err = yaml.Unmarshal(yamlConf, &c) // 解析配置文件 反序列化

	if err != nil {
		log.Fatalf("config file parse error: %v", err)
	}

	log.Println("config yamlFile load Init success.")

	global.Config = c
}

// SetYaml 将配置写入配置文件
func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config) // 将config转化为yaml格式
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	global.Log.Info("配置文件修改成功")
	return nil
}
