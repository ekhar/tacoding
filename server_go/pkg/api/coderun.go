package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
 Output     string `json:"stdout"`
 StatusCode   string `json:"stderr"`
 Memory string    `json:"output"`
 CpuTime string `json:"code"`
}

func jdoodle_call() []byte{
	print("calling API")

	CLIENT_ID := "e87471b6fcd9bd57ee9498aa21f9a028"
	CLIENT_SECRET := "b2dd432576cc25c58f3f0ab05f1ac2183910581c451b28b6fcbdd75994e33b6d"

	script := "print('helloworld')"
	language := "python3"
	versionIndex :="3"
	stdin:=""


	jsonData := map[string]string{"clientId":  CLIENT_ID, "clientSecret": CLIENT_SECRET, "script": script, "stdin": stdin, "language": language,"versionIndex": versionIndex}
	

	jsonValue, _ := json.Marshal(jsonData)
    request, _ := http.NewRequest("POST","https://api.jdoodle.com/v1/execute", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
	    fmt.Printf("The HTTP request failed with error %s\n", err)
		return nil
	} else {
	    data, _ := ioutil.ReadAll(response.Body)
	    fmt.Println(string(data))
		return data
	}

}