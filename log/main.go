package main

import (
	// "log"
	// "os"
	log "github.com/Sirupsen/logrus"
	"github.com/logmatic/logmatic-go"
)

// func log2file() {
// 	f,err :=os.OpenFile("test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
// 	if err !=nil{
// 		log.Fatal(err)
//   	}
 
//  	defer f.Close()
//  	log.SetOutput(f)
// 	log.Println("==========works==============")
// }

func main() {
	// use JSONFormatter
	log.SetFormatter(&logmatic.JSONFormatter{})
	
	contextLogger := log.WithFields(log.Fields{
		"common": "XXX common content XXX",
		"other": "YYY special context YYY",
	})

	contextLogger.Info("AAAAAAAAAAAA")
	contextLogger.Info("BBBBBBBBBBBB")

	
	// // log an event as usual with logrus
	// log.WithFields(log.Fields{"string": "foo", "int": 1, "float": 1.1 }).Info("My first ssl event from golang")
}