package utils

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func Resolve() map[string]string{
	value := make(map[string]string)
	b ,err :=ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	json.Unmarshal(b,&value)
	return value
}

