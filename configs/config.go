package configs

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Yscream/go-form-reg/pkg/models"
	"gopkg.in/yaml.v2"
)

func GetConfig() string {

	yfile, err := ioutil.ReadFile("./cmd/config.yml")

	if err != nil {
		log.Fatal(err)
	}

	c := models.Config{}

	err = yaml.Unmarshal(yfile, &c)

	if err != nil {
		log.Fatal(err)
	}
	result := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.Username, c.Password, c.Host, c.Port, c.Dbname, c.Sslmode)
	return result
}
