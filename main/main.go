package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	"log"
)

type Data struct {
	Data []Record `json:"data"`
	Msg string `json:"msg"`
	StateCode int `json:"stateCode"`
	Success bool `json:"success"`
}

type Record struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Phone int `json:"phone"`
	Addr string `json:"addr"`
}



func main(){
	b, err := ioutil.ReadFile("E:/test.json")

	if err != nil{
		fmt.Print(err)
	} else{
		//str := string(b)
		t := &Data{}
		err := json.Unmarshal(b, t)
		if err != nil{
			fmt.Println(err)
		}
		log.Print(*t)
	}

}