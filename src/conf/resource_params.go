package conf

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"time"
)

var (
	GlobalConfig *Config
)

type Config struct {
	APPName string       `toml:"app_name"`
	SConfig ServerConfig `toml:"server"`
	//LogConfig           LogConfig            `toml:"golog"`
	RiskDBConfig DBconfig `toml:"riskdb"`
	//CacheConfig     map[string]Redisconfig     `toml:"cache"`
}

type ServerConfig struct {
	MaxCpu       int `toml:"max_cpu"`
	HttpPort     int `toml:"http_port"`
	PprofPort    int `toml:"pprof_port"`
	RestPort     int `toml:"rest_port"`
	WriteTimeout int `toml:"write_timeout"`
	ReadTimeout  int `toml:"read_timeout"`
}

type LogConfig struct {
	Level      string `toml:"level"`
	Console    int    `toml:"console"`
	Dir        string `toml:"dir"`
	Filename   string `toml:"filename"`
	ReserveNum int    `toml:"reserve_num"`
	Suffix     string `toml:"suffix"`
	Colorfull  int    `toml:"colorfull"`
}

type Redisconfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	PoolSize int    `toml:"pool_size"`
	MaxIdle  int    `toml:"max_idle"`

	IdleTimeout    time.Duration `toml:"idle_timeout"`
	ConnectTimeout time.Duration `toml:"connect_timeout"`
	ReadTimeout    time.Duration `toml:"read_timeout"`
	WriteTimeout   time.Duration `toml:"write_timeout"`
}

type DBconfig struct {
	Usr      string `toml:"user"`
	Pwd      string `toml:"pwd"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	DBname   string `toml:"db_name"`
	MaxIdle  int    `toml:"max_idle"`
	MaxOpen  int    `toml:"max_open"`
	PoolSize int    `toml:"pool_size"`
}

func InitConfig(filename string) {
	if GlobalConfig == nil {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			panic("read config file failed " + filename + " " + err.Error())
		}
		if _, err := toml.Decode(string(data), &GlobalConfig); err != nil {
			panic("toml decode failed " + err.Error())
		}
	}
}
