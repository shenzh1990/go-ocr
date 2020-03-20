package settings

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

var config = flag.String("f", "./pkg/settings/config.yaml", "config file")

type Config struct {
	Fcoin    FcoinConfig `yaml:"Fcoin"`
	DingDing DingConfig  `yaml:"DingDing"`
	WeiXin   WxConfig    `yaml:"WeiXin"`
	Base     BaseConfig  `yaml:"Base"`
	Db       DbConfig    `yaml:"Db"`
	Redis    RedisConfig `yaml:"Redis"`
}
type BaseConfig struct {
	RunMode      string        `yaml:"RunMode"`
	HTTPPort     int           `yaml:"HTTPPort"`
	ReadTimeout  time.Duration `yaml:"ReadTimeout"`
	WriteTimeout time.Duration `yaml:"WriteTimeout"`
	PageSize     int           `yaml:"PageSize"`
	JwtSecret    string        `yaml:"JwtSecret"`
}
type DbConfig struct {
	DriverName string `yaml:"DriverName"`
	DBUrl      string `yaml:"DBUrl"`
}
type WxConfig struct {
	AppKey    string `yaml:"AppKey"`
	AppSecret string `yaml:"AppSecret"`
	Url       string `yaml:"Url"`
}

type DingConfig struct {
	AppKey    string `yaml:"AppKey"`
	AppSecret string `yaml:"AppSecret"`
	ChatId    string `yaml:"ChatId"`
	Url       string `yaml:"Url"`
	//manager5970
}
type FcoinConfig struct {
	FcoinKey    string `yaml:"FcoinKey"`
	FcoinSecret string `yaml:"FcoinSecret"`
}

type RedisConfig struct {
	RedisHost string `yaml:"RedisHost"`
	RedisDB   string `yaml:"RedisDB"`
	RedisPwd  string `yaml:"RedisPwd"`
	Timeout   int64  `yaml:"Timeout"`

	PoolMaxIdle     int   `yaml:"PoolMaxIdle"`
	PoolMaxActive   int   `yaml:"PoolMaxActive"`
	PoolIdleTimeout int64 `yaml:"PoolIdleTimeout"`
	PoolWait        bool  `yaml:"PoolWait"`
}

func (c *Config) getConf(filepath string) *Config {
	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
func init() {
	BitConfig.getConf(*config)
	LoadBase()
}

var (
	BitConfig Config
	RunMode   string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

func LoadBase() {
	RunMode = BitConfig.Base.RunMode
	HTTPPort = BitConfig.Base.HTTPPort
	ReadTimeout = BitConfig.Base.ReadTimeout * time.Second
	WriteTimeout = BitConfig.Base.WriteTimeout * time.Second
	JwtSecret = BitConfig.Base.JwtSecret
	PageSize = BitConfig.Base.PageSize
}
