package main

import (
	"fmt"
	"github.com/jhole89/orbital/database"
	"log"
	"net/http"
)

var (
	conf  Config
	graph database.Graph
	err   error
)

func main() {
	conf.getConf()

	graph, err = database.GetGraph(conf.Database.Type, conf.Database.Endpoint)
	if err != nil {
		log.Println(err)
	}

	dh, err := createDataHandler()
	if err != nil {
		log.Println(err)
	}
	http.Handle("/data", dh)

	ah, err := createAdminHandler()
	if err != nil {
		log.Println(err)
	}
	http.Handle("/admin", ah)

	log.Printf("Launching Orbital API at http://127.0.0.1:%d\n", conf.Service.Port)
	if err = http.ListenAndServe(fmt.Sprintf(":%d", conf.Service.Port), nil); err != nil {
		log.Fatalln("Unable to serve.")
	}
}
