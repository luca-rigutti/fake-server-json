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

type M map[string]interface{}

func main() {
	jsons := readJsonFolder()

	for _, j := range jsons {
		jsonBytes, err := json.MarshalIndent(j["response"],"","   ")
		    if err != nil {
        log.Fatal(err)
    }
		http.HandleFunc(j["url"].(string), testHandler(string(jsonBytes)))//(fmt.Sprintf("%v",j["response"])))
	}
	
	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		  fmt.Printf("server closed\n")
	  } else if err != nil {
		  fmt.Printf("error starting server: %s\n", err)
		  os.Exit(1)
	  }
}

func testHandler(name string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("got / request\n")
		io.WriteString(w, name)
    }
}


func readJsonFolder() []M {
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

	var responses []M

	for _,nameFileJson := range filesJson {
		fmt.Println(nameFileJson)
		responses = append(responses, readJsonFile(nameFileJson))
	}

	return responses

}

func readJsonFile(filename string) map[string]interface{} {
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

		return result
		
}