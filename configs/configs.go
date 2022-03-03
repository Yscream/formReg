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

func GetConfig(path string) (string, error) {

	yfile, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err)
	}

	c := Config{}

	err = yaml.Unmarshal(yfile, &c)

	if err != nil {
		log.Fatal(err)
	}
	result := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.Username, c.Password, c.Host, c.Port, c.Dbname, c.Sslmode)
	return result, err
}
