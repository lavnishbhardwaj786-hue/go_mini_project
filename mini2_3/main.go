package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	PORT       string
	Mongo_name string
	Mongo_uri  string
	Myown      string
}
type jsonfileconfig struct {
	PORT       string `json:"Port"`
	Mongo_name string `json:"Mongo_name"`
	Mongo_uri  string `json:"Mongo_uri"`
	Myown      string `json:"Myown"`
}

func getenv(name string) (string, error) {
	value := os.Getenv(name)
	if value == "" {
		return "", errors.New("the value is empty")
	}
	return value, nil
}
func load() (config, error) {
	err := godotenv.Load()
	if err != nil {
		return config{}, fmt.Errorf("not able to load the env to process.env")
	}
	Mongo_name, err := getenv("Mongo_name")
	if err != nil {
		return config{}, fmt.Errorf("not able to load the value Mongo_name")
	}
	Port, err := getenv("PORT")
	if err != nil {
		return config{}, fmt.Errorf("not able to load the value Mongo_name")
	}
	Mongo_uri, err := getenv("Mongo_uri")
	if err != nil {
		return config{}, fmt.Errorf("not able to load the value port")
	}

	Myown, err := getenv("Myown")
	if err != nil {
		return config{}, fmt.Errorf("not able to load the value Myown")
	}
	return config{
		Mongo_name: Mongo_name,
		Mongo_uri:  Mongo_uri,
		Myown:      Myown,
		PORT:       Port,
	}, nil
}

func readfile() ([]byte, error) {
	data, err := os.ReadFile("information.json")
	if err != nil {
		fmt.Println("not able to read the file ", err)
	}
	return data, err
}
func (j *jsonfileconfig) readjson() error {
	jsondata, err := readfile()
	if err != nil {
		fmt.Println("not able to read the file ", err)
	}
	err = json.Unmarshal(jsondata, j)
	if err != nil {
		fmt.Println("not able to convert to json", err)
	}
	return err
}

func main() {
	var filejson jsonfileconfig
	filejson.readjson()
	envvalue, err := load()
	if err != nil {
		fmt.Println("this is the error of not workign load functino ", err)
	}

	if filejson.Mongo_name == envvalue.Mongo_name {
		fmt.Println(filejson.Mongo_name)
	} else {
		fmt.Println(envvalue.Mongo_name)
	}
	if filejson.Mongo_uri == envvalue.Mongo_uri {
		fmt.Println(filejson.Mongo_uri)
	} else {
		fmt.Println(envvalue.Mongo_uri)
	}
	if filejson.PORT == envvalue.PORT {
		fmt.Println(filejson.PORT)
	} else {
		fmt.Println(envvalue.PORT)
	}
	if filejson.Myown == envvalue.Myown {
		fmt.Println(filejson.Myown)
	} else {
		fmt.Println(envvalue.Myown)
	}

}
