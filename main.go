package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	method string
	url    string
)

// I could've done it by just crul <Link> > <FileName>
// but I like it

func main() {
	var req, reqErr = http.NewRequest(method, url, nil)

	if reqErr != nil {
		log.Println("Error During Req")
		os.Exit(2)
	}

	var client *http.Client = &http.Client{}

	var result, resultErr = client.Do(req)

	if resultErr != nil {
		log.Println("resultErr")
		os.Exit(2)
	}

	var data, dataErr = ioutil.ReadAll(result.Body)
	if dataErr != nil {
		log.Println("decode Err")
		os.Exit(2)
	}

	var file, fileErr = os.Create("v.mp4")

	if fileErr != nil {
		log.Println("File Create Err")
		os.Exit(2)
	}
	var donebytes, byteErr = file.Write(data)
	if byteErr != nil {
		log.Println("File Write Err")
		os.Exit(2)
	}
	log.Println(donebytes)
}
