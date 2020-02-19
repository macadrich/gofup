package main

import (
	"gofup/fhandler"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const (
	PORT    = "15223"   // port to open
	ADDRESS = "0.0.0.0" //"0.0.0.0" // localhost
	FILEDIR = "myfile"  // file folder
	SERVER  = "recv"    // recieve a file
	CLIENT  = "send"    // send a file
)

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func main() {

	checkArgs := func(args []string) bool {
		if len(args) <= 2 || len(args) > 3 {
			log.Println("Commands:")
			log.Println("   send | recv ")
			log.Println("   path - file folder")
			return false
		}
		return true
	}

	if !checkArgs(os.Args) {
		return
	}

	app, err := fhandler.NewFUpload(ADDRESS + ":" + PORT)
	if err != nil {
		log.Fatal(err)
	}

	switch os.Args[1] {
	case CLIENT:
		log.Println("[ CLIENT ] >>> ")
		app.SetFileFolder(os.Args[2])
		err := app.Send()
		if err != nil {
			log.Fatal(err)
		}
		return
	case SERVER:
		log.Println("[ SERVER ] <<<")
		app.SetFileFolder(os.Args[2])
		err := app.Receive()
		if err != nil {
			log.Fatal(err)
		}
		return
	default:
		log.Println("Unknown!")
	}

	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-exit)

}
