package configs

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     int
	Dbname   string
	Sslmode  string
}

func InitConfig(path string) (string, error) {
	config := Config{}
	yfile, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal(yfile, &config)

	if err != nil {
		log.Fatal(err)
	}
	initconfig := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", config.Username, config.Password, config.Host, config.Port, config.Dbname, config.Sslmode)
	return initconfig, err
}
