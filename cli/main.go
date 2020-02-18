package main

import (
	"gofup/upload"
	"log"
	"os"
)

const (
	PORT    = "15223"   // port to open
	ADDRESS = "0.0.0.0" //"0.0.0.0" // localhost
	FILEDIR = "myfile"  // file folder
	SERVER  = "recv"    // recieve a file
	CLIENT  = "send"    // send a file
)

func main() {
	if len(os.Args) <= 1 || len(os.Args) > 2 {
		log.Println("Commands:")
		log.Println("	send - As client")
		log.Println("	recv - As server")
		return
	}

	app, err := upload.NewFUpload(FILEDIR, ADDRESS+":"+PORT)
	if err != nil {
		log.Fatal(err)
	}

	switch os.Args[1] {
	case CLIENT:
		log.Println("[ CLIENT ] >>> ")
		err := app.Send()
		if err != nil {
			log.Fatal(err)
		}
		return
	case SERVER:
		log.Println("[ SERVER ] <<<")
		err := app.Receive()
		if err != nil {
			log.Fatal(err)
		}
		return
	default:
		log.Println("Unknown!")
	}

}
