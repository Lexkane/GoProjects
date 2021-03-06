package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
)
type Payload struct{
	Envelope Data
}

type Data struct{
	CarByDoor Cars
	TrucksByTon Trucks
}

type Cars map[string] int
type Trucks map[string]int

func main(){
	url:="http://localhost:4000"
	res,err:=http.Get(url)

	if err!=nil{
		panic(err)
	}
	defer res.Body.Close()

	body,err:=ioutil.ReadAll(res.Body)
	if err!=nil{
		panic(err)
	}
	var p Payload
	err=json.Unmarshal(body,&p)
	if err!=nil{
		panic(err)
	}
	fmt.Println(p.Envelope.CarByDoor,"\n",p.Envelope.TrucksByTon)


}