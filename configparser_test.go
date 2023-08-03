package configparser

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

type PostgresSqlConnConfigs struct {
	WritePool PostgresSqlConnPoolConf `yaml:"write_pool"`
	ReadPool  PostgresSqlConnPoolConf `yaml:"read_pool"`
}

type PostgresSqlConnPoolConf struct {
	Url                    *string `yaml:"url"`
	Port                   *string `yaml:"port"`
	UserName               *string `yaml:"user_name"`
	Password               *string `yaml:"password"`
	DbName                 *string `yaml:"dbname"`
	MaxIdleConn            *int    `yaml:"max_idle_Conns"`
	MaxOpenConn            *int    `yaml:"max_open_conns"`
	ConnMaxLifetimeSeconds *int64  `yaml:"conn_max_lifetime_seconds"`
	UserA                  *User   `yaml:"user"`
}
type User struct {
	Username *string `yaml:"username"`
	Password *string `yaml:"password"`
}

func TestLoadConfigFromFile(t *testing.T) {
	workPath, errGetWd := os.Getwd()
	assert.Nil(t, errGetWd)
	confDirPath := path.Join(workPath, "test_data/")

	t.Run("valid.yaml", func(t *testing.T) {
		t.Parallel()
		fileName := strings.Split(t.Name(), "/")[1]
		filePath := path.Join(confDirPath, fileName)

		_, err := LoadConfigFromFile(filePath, "not_exist")
		assert.NotNil(t, err)

		confs, err := LoadConfigFromFile(filePath, "pro")
		assert.Nil(t, err)

		var config PostgresSqlConnConfigs
		err = yaml.Unmarshal([]byte(confs), &config)
		assert.Nil(t, err)
	})

	t.Run("not_exist.yaml", func(t *testing.T) {
		t.Parallel()
		fileName := strings.Split(t.Name(), "/")[1]
		filePath := path.Join(confDirPath, fileName)
		_, err := LoadConfigFromFile(filePath, "pro")
		assert.NotNil(t, err)
	})

	t.Run("valid_ex_is_nil", func(t *testing.T) {
		confFileName := "valid_ex_is_nil.yaml"
		confFilePath := path.Join(confDirPath, confFileName)
		_, err := LoadConfigFromFile(confFilePath, "pro")
		assert.NoError(t, err)
	})

	t.Run("valid_default_is_nil", func(t *testing.T) {
		confFileName := "valid_default_is_nil.yaml"
		confFilePath := path.Join(confDirPath, confFileName)
		_, err := LoadConfigFromFile(confFilePath, "pro")
		assert.NoError(t, err)
	})

	t.Run("invalid_empty_file", func(t *testing.T) {
		confFileName := "invalid_empty_file.yaml"
		confFilePath := path.Join(confDirPath, confFileName)
		_, err := LoadConfigFromFile(confFilePath, "pro")
		assert.NotNil(t, err)
	})

	t.Run("valid_empty_file", func(t *testing.T) {
		confFileName := "valid_ex_more_then_default.yaml"
		confFilePath := path.Join(confDirPath, confFileName)
		_, err := LoadConfigFromFile(confFilePath, "pro")
		assert.NoError(t, err)
	})
}
