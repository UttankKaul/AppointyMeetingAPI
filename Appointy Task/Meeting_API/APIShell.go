package main

import (
	"fmt"
	"os"
)

//import "go.mongodb.org/mongo-driver/mongo"
//import "go.mongodb.org/mongo-driver/mongo/options" // Use these packages for MangoDB
//import "go.mongodb.org/mongo-driver/mongo/readpref"

func particpants() { //particpantsinformation
	var CusName string // Customer Name(paticpant Name)
	var Cusemail string
	var RSVPCheck string
	fmt.Println("Enter you Email:")
	fmt.Scanln(&Cusemail)
	fmt.Println("Enter your name")
	fmt.Scanln(&CusName)
	fmt.Scanln("Applied for RSVP ? :")
	fmt.Scanln(&RSVPCheck)

	fmt.Println("Particpant Email is : ", Cusemail)
	fmt.Println("Particpant name is :", CusName)

	if RSVPCheck == "Yes" {
		fmt.Println("The Particpant has applied for RSVP")
	} else {
		fmt.Println("The Particpant has not applied for RSVP")

	}

}

func connectToMeeting() { //InfoCollection

	fmt.Println("Connction... estalished... ")
	var Logintime int
	var CusStartTime int
	var MeetingID string
	var NoofParticpants int
	var Title string
	var CusEndTime int
	var CusTimeStamp int

	fmt.Println("Hi! Welcome to the Appointy API. Thank you for joining the experiece!.....")

	// data can be entred by JSON DB file  also
	jsonFile, debug := os.Open("Database.json")

	if debug != nil {
		fmt.Println(debug)
	}
	fmt.Println("Successfully Opened the Database.json file...")

	fmt.Println("Meeting ID:")
	fmt.Scanln(&MeetingID)
	fmt.Println("No of particpants:")
	fmt.Scanln(&NoofParticpants)
	fmt.Println("Titile:")
	fmt.Scanln(&Title)
	fmt.Println("Enter your Login Time: ") // in 24 Hours format For ex.. 1345 == 1:45PM
	fmt.Scanln(&Logintime)
	fmt.Println("Enter Staring Time:")
	fmt.Scanln(&CusStartTime)
	fmt.Println("Enter EndTime")
	fmt.Scanln(&CusEndTime)
	fmt.Println("Enter TS")
	fmt.Scanln(&CusTimeStamp)
	//fmt.Println("\n\n\n")

	fmt.Println("Meeting ID : ", MeetingID)
	fmt.Println("No of people in Meeting :", NoofParticpants)
	fmt.Println("Title of meeting is:", Title)
	fmt.Println("StartTime :  ", CusStartTime)
	fmt.Println("EndTime : ", CusEndTime)
	fmt.Println("TimeStamp :", CusTimeStamp)

	defer jsonFile.Close() // File is now closed

	particpants()

}

func main() {

	connectToMeeting()
}

//HTTP/1.1 200 OK
//Content-Type: application/vnd.api+json
