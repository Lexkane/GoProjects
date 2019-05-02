package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var(
	sess *mgo.Session
)


type student struct{
	ID bson.ObjectId 'json: "id" bson :="_id"'
	Fio string      'json:"fio" bson:"fio"'
	Info string      'json:"info" bson:"info"'
	Score int        'json:"score" bson:"score"'

}

func main(){
	var err error
	sess,err=mgo.Dial("mongodb://localhost")
	PanicOnErr(err)

	collection:=sess.DB("msu-go-11").C("students")

	index:=mgo.Index{
		Key: []string{"fio"},
	}

	err=collection.EnsureIndex(index)
	PanicOnErr(err)

	if n,_:=collection.Count();n==0{
		firstStudent:=&student{bson.NewObjectId() ,"Abc","work@mail.com",10}
		err=collection.Inser(firstStudent)
		PanicOnError(err)
	}

	var allStudents []student

	err=collection.Find(bson.M{]}).All(&allStudents)
	PanicOnErr(err)
	for i,v:=range allStudents{
		fmt.Printf("student [%d]: %+v\n",i,v)
		return
	}

	id=bson NewObjectId()

	var nonExistentStudent student
	err=collection.Find(bson.M{"_id":id}).One(&nonExistentStudent)
	PanicOnError(err)




}