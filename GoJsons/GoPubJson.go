package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)
func main(){
	http.HandleFunc("/",serveRest)
	http.ListenAndServe("localhost:8181",nil)
}
func serveRest(w http.ResponseWriter, r *http.Request){
	response,err:=publishJson()
	if err!=nil {
		panic(err)
	}
	fmt.Fprintf(w,string(response))
}

func publishJson()([]byte,error){
	cars:=make(map[string]int)
	cars["2Door"]=20
	cars["4Door"]=32
	trucks:=make(map[string] int)
	trucks["1ton"]=14
	trucks["2ton"]=4

	d:=Data{cars,trucks}
	p:=Payload{d}

	return json.Marshall(p,"","")
}

type Payload struct{
	Envelope Data
}

type Data struct{
	CarByDoor Cars
	TrucksByTon Trucks
}

type Cars map[string] int
type Trucks map[string]int