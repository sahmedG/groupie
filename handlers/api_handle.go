package ConcertAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func SendRequest(link string, pointer interface{}) {
	res, err := http.Get(link)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(responseData, &pointer)
}
