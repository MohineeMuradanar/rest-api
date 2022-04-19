// Package controller Student API.
//
// this application is provide student information is using go code to define an Rest API
//
//     Schemes: http, https
//     Host: localhost:3000
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
// swagger:meta
package controller

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//Configuration struct
type Configuration struct {
	Port             string // port no
	ConnectionString string // connection string
	Database         string // database name
	Collection       string // collection
}

/*ReadConfig Reading the configs from  db.properties
 */
func ReadConfig() Configuration {
	var configfile = "config.properties"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Configuration
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	//log.Print(config.Index)
	return config
}

func Handler() {
	config := ReadConfig()
	var port = ":" + config.Port

	router := mux.NewRouter()
	corsObj := handlers.AllowedOrigins([]string{"*"})
	http.Handle("/", router)

	router.HandleFunc("/home", home)
	router.HandleFunc("/students", createNewStudent).Methods("POST")
	router.HandleFunc("/students/{id}", getById).Methods("GET")
	router.HandleFunc("/students", getAll).Methods("GET")

	router.HandleFunc("/students/{id}", updateStudent_Patch).Methods("PATCH")
	router.HandleFunc("/students/{id}", replaceStudent_Put).Methods("PUT")

	router.HandleFunc("/students/{id}", deleteStudent).Methods("DELETE")

	router.HandleFunc("/check", check).Methods(http.MethodGet)

	// log.Fatal(http.ListenAndServe(":3001", router))

	fmt.Printf("application listening port %s\n", port)
	http.ListenAndServe(port, handlers.CORS(corsObj)(router))

}


