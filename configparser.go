package configparser

import (
	"fmt"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Version string      `yaml:"version"`
	Pro     interface{} `yaml:"pro"`
	Pre     interface{} `yaml:"pre"`
	Dev     interface{} `yaml:"dev"`
	Test    interface{} `yaml:"test"`
	Default interface{} `yaml:"default"`
}

const (
	ErrorEnvNotfound = "env [%s] not found"
	KeyFieldTag      = "yaml"
)

func LoadConfigFromFile(filePath string, env string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return LoadConfigFromMemory(data, env)
}

func LoadConfigFromMemory(configure []byte, env string) (string, error) {
	var config Config
	err := yaml.Unmarshal(configure, &config)
	if err != nil {
		return "", err
	}

	configReflectType := reflect.TypeOf(config)
	for i := 0; i < configReflectType.NumField(); i++ {
		structTag := configReflectType.Field(i).Tag.Get(KeyFieldTag)
		if structTag == env {
			envConfigReflect := reflect.ValueOf(config).Field(i).Interface()
			defauleConfigReflectType := config.Default
			if envConfigReflect == nil && defauleConfigReflectType == nil {
				return "", fmt.Errorf(ErrorEnvNotfound, env)
			}
			if envConfigReflect == nil {
				defaultConf, err := yaml.Marshal(config.Default)
				if err != nil {
					return "", err
				}
				return string(defaultConf), nil
			}
			if defauleConfigReflectType == nil {
				envConf, err := yaml.Marshal(reflect.ValueOf(config).Field(i).Interface())
				if err != nil {
					return "", err
				}
				return string(envConf), nil
			}
			merged := mergeMapStringInterface(reflect.ValueOf(config).Field(i).Interface().(map[string]interface{}), config.Default.(map[string]interface{}))
			mergedConf, err := yaml.Marshal(merged)
			if err != nil {
				return "", err
			}
			return string(mergedConf), nil
		}
	}

	return "", fmt.Errorf(ErrorEnvNotfound, env)
}

func mergeMapStringInterface(cover map[string]interface{}, base map[string]interface{}) map[string]interface{} {
	for k, v := range cover {
		switch v.(type) {
		case map[string]interface{}:
			if base[k] == nil {
				base[k] = v
			} else {
				mergeMapStringInterface(v.(map[string]interface{}), base[k].(map[string]interface{}))
			}
		default:
			base[k] = v
		}
	}
	return base
}
