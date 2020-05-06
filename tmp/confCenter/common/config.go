package common

import (
	"encoding/json"
	"io/ioutil"
)

var conf *Config

// Config 配置信息
type Config struct {
	WebAddr string       `json:"web_addr"`
	Etcd    *EtcdConfig  `json:"etcd"`
	Mysql   *MysqlConfig `json:"mysql"`
	App     *AppConfig   `json:"app"`
	Micro	*MicroConfig `json:"micro"`
}

// EtcdConfig Etcd配置信息
type EtcdConfig struct {
	Endpoints   []string `json:"endpoints"`
	DialTimeout int      `json:"dial_timeout"`
}

// MysqlConfig Mysql配置信息
type MysqlConfig struct {
	Addr     string `json:"addr"`
	User     string `json:"user"`
	Password string `json:"password"`
	Db       string `json:"db"`
}

// AppConfig 应用配置信息
type AppConfig struct {
	Mysql *MysqlConfig `json:"mysql"`
}

// MicroConfig 微服务名和版本信息
type MicroConfig struct {
	Name string `json:"name"`
	Version string `json:"version"`
}

// InitConfig 初始化配置
func InitConfig(cfgFile string) error {
	var err error
	conf, err = readConf(cfgFile)
	if err != nil {
		return err
	}

	return nil
}

// readConf 读取配置文件
func readConf(cfgFile string) (*Config, error) {
	f, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		return nil, err
	}

	data := &Config{}
	err = json.Unmarshal(f, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetMicroConfig 获取微服务信息
func GetMicroConfig() *MicroConfig {
	return conf.Micro
}

// GetEtcdConfig 获取Etcd配置
func GetEtcdConfig() *EtcdConfig {
	return conf.Etcd
}

// GetWebAddr 获取web地址
func GetWebAddr() string {
	return conf.WebAddr
}

// GetMysqlConfig 获取mysql配置
func GetMysqlConfig() *MysqlConfig {
	return conf.Mysql
}

// GetAppMysqlConfig 获取应用需要mysql配置
func GetAppMysqlConfig() *MysqlConfig {
	if conf.App == nil {
		return nil
	}
	return conf.App.Mysql
}