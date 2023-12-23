package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env            string         `yaml:"env" env-default:"local"`
	GRPC           GRPCConfig     `yaml:"grpc"`
	HTTP           HTTPConfig     `yaml:"http"`
	Postgres       PostgresConfig `yaml:"postgres"`
	MigrationsPath string
	TokenTTL       time.Duration `yaml:"token_ttl" env-default:"1h"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type HTTPConfig struct {
	Port int `yaml:"port"`
}

type PostgresConfig struct {
	Port    string `yaml:"port"`
	Host    string `yaml:"host"`
	DbName  string `yaml:"db_name"`
	User    string `yaml:"user"`
	Pwd     string `yaml:"password"`
	SslMode string `yaml:"sslmode"`
}

func InitConfiguration() *Config {
	configPath := getConfigPath()
	if configPath == "" {
		panic("empty config path")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("empty config path: " + err.Error())
	}

	return &cfg
}

func getConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	// если не получилось взять с флага --config, пытаемся взять из переменной окружения
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
