package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	c, err := readConf("conf.yaml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", c.Conf.URI)
	fmt.Printf("%v\n", c.Conf.User)
	fmt.Printf("%v\n", c.Conf.Pwd)

	e := startApp()
	if e != nil {
		log.Fatal(e)
	}
}

func startApp() error {

	_, err := (fmt.Println("Hello world!"))

	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	routing(r)

	// Critical to work on AppEngine
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
	return err

}
