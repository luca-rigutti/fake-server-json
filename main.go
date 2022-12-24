package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"io/ioutil"
	"log"
	"path/filepath"
	"encoding/json"
)

func main() {
	readJsonFolder()


	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		  fmt.Printf("server closed\n")
	  } else if err != nil {
		  fmt.Printf("error starting server: %s\n", err)
		  os.Exit(1)
	  }
}


func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func readJsonFolder() {
	filesJson := []string {};
	files, err := ioutil.ReadDir("RouteJson")
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
		if(!file.IsDir() && filepath.Ext(file.Name())==".json"){
        	
			filesJson = append(filesJson, file.Name())
		}
    }

	for _,nameFileJson := range filesJson {
		fmt.Println(nameFileJson)
		readJsonFile(nameFileJson)
	}




}

func readJsonFile(filename string){
	    // Open our jsonFile
		jsonFile, err := os.Open("RouteJson/"+filename)
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successfully Opened "+filename)
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()
	
		byteValue, _ := ioutil.ReadAll(jsonFile)
	
		var result map[string]interface{}
		json.Unmarshal([]byte(byteValue), &result)
	
		fmt.Println(result["url"])
		fmt.Println(result["response"])
		
}