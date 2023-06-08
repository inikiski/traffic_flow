package envx

import (
	"github.com/labstack/gommon/log"
	"gopkg.in/yaml.v3"
	"os"
)

func ReadYamlConfig(v interface{}, fileName ...string) error {
	file := "config.yaml"
	if len(fileName) > 0 {
		file = fileName[0]
	}
	dataBytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Read config.yaml failed: %s", err.Error())
		return err
	}
	err = yaml.Unmarshal(dataBytes, v)
	if err != nil {
		log.Fatalf("Failed parse config: %s", err.Error())
		return err
	}
	return nil
}

func MustReadYamlConfig(v interface{}, fileName ...string) {
	if err := ReadYamlConfig(v, fileName...); err != nil {
		log.Fatal(err)
	}
}
