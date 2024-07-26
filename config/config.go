package config

import (
	"github.com/bytedance/gopkg/util/logger"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/tundrawork/powcha/biz/altcha/algorithm"
)

var (
	k    = koanf.New(".")
	path = "config.yaml"
	Conf Config
)

type Config struct {
	Altcha Altcha `yaml:"Altcha"`
}

type Altcha struct {
	Algorithm  string `yaml:"Algorithm"`
	Complexity int    `yaml:"Complexity"`
}

func Init() {
	if err := k.Load(file.Provider(path), yaml.Parser()); err != nil {
		logger.Fatalf("error loading config: %v", err)
	}

	if err := k.Unmarshal("", &Conf); err != nil {
		logger.Fatalf("error unmarshalling config: %v", err)
	}

	_, ok := algorithm.AlgorithmFromString(Conf.Altcha.Algorithm)
	if !ok {
		logger.Fatalf("invalid config value: algorithm=%v", Conf.Altcha.Algorithm)
	}
	if Conf.Altcha.Complexity < 100000 {
		logger.Fatalf("invalid config value: complexity=%v", Conf.Altcha.Complexity)
	}
}
