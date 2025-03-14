package config

// toml 방식 쓰기 위해서 go get github.com/naoina/toml 씀
import (
	"github.com/naoina/toml"
	"os"
)

// env -> toml

type Config struct {
	Server struct {
		Port string
	}
}

func NewConfig(filePath string) *Config {
	c := new(Config)

	if file, err := os.Open(filePath); err != nil {
		panic(err)
	} else if err = toml.NewDecoder(file).Decode(c); err != nil {
		panic(err)
	} else {
		return c
	}
}
