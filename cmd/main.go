package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gojou/goofjson/cmd/routes"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

type myData struct {
	Conf struct {
		URI  string `yaml:"dburi"`
		User string `yaml:"user"`
		Pwd  string `yaml:"password"`
	}
}

func readConf(filename string) (*myData, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &myData{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return c, nil
}
func main() {
	c, err := readConf("cmd/conf.yaml")
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
	routes.Routing(r)

	// Critical to work on AppEngine
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
	return err

}
