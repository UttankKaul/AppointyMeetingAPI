package main

//package bson
import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type partpantsInfo struct {
	CusName   string
	CusID     string // all datatypes and names are same as of APISheel Script
	Cusemail  string
	RSVPCheck string
}

type meetingInfoDB struct {
	Logintime       int
	CusStartTime    int
	MeetingID       string
	NoofParticpants int
	Title           string
	CusEndTime      int
	CusTimeStamp    int
}

func main() {
	//Mango.connect()
	//var N int
	clientOptions := options.Client().ApplyURI("mongodb://localhost:8080") //Connecting to the local Host. Open your WebBrowser....
	client, debug := mongo.Connect(context.TODO(), clientOptions)
	collection := client.Database("DataBase").Collection("insertMeetingDB")
	if debug != nil { //Check For Connection
		log.Fatal(debug) // Return errors When not conneted....
	} else {
		fmt.Println("Connected to the MangoDB database..............")

	}

	//mongoimport --db DataBase --collection InsertMeetingDB --file Database.json   // Use this to import data to the MangoDB

	ParticpantInfoDB := partpantsInfo{"Uttank", "appointy", "abc@gmail.com", "Yes"} //values taken by user(particpants) will be entered/inserted
	InfoDB := meetingInfoDB{945, 1005, "ABC1002", 60, "DBMS", 1005, 200}            //values taken for the meeting will be inserted(example)
	addInfoDB := []interface{}{ParticpantInfoDB, InfoDB}
	insertMeetingDB, debug2 := collection.InsertMany(context.TODO(), addInfoDB)

	/*for(i int ; i<N;i++)
	{
	if ParticpantInfo.RSVPCheck == "Yes" || meetingInfo.CusStartTime(i) = meetingInfo.CusStartTime(i+1){
		fmt.Println("ERROR!")
	} else {
		//fmt.Println("No Race Conditions...")
	}

	}*/
	//insertMeetingDB, debug3 := collection.InsertOne(context.TODO(), InfoDB)

	if debug2 != nil {
		log.Fatal(debug2) // Check for insertion of data
	}

	fmt.Println("Values Added to DB are :", insertMeetingDB.InsertedIDs)

	debug = client.Disconnect(context.TODO())
	fmt.Println("DB is now closed..........")

}
