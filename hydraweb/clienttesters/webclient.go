package main

import(
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main(){
	//requestbin no longer is public so will not use this url for the current test
	url := ""
	resp, err := http.Get(url)
	inspectResponse(resp, err)

	//sets up anonymous struct to test json marshal for rest below
	data, err := json.Marshal(struct{
		X int
		Y float32
	}{X:4, Y:3.8})
	if err != nil{
		log.Fatal("Error occured while marshaling json", err)
	}

	//basic post request
	// url, body , io.reader
	resp, err = http.Post(url, "application/json", bytes.NewReader(data))
	inspectResponse(resp, err)

		//example client with a timeout feature to stop stale connections
	client := http.Client{
		Timeout : 3 * time.Second,
	}
	client.Get(url)

	//basic put request
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil{
		log.Fatal(err)
	}
	req.Header.Add("x-testheader", "learning go header")
	req.Header.Set("User-Agent", "Go learning HTTP/1.1")
	resp, err = client.Do(req)
	inspectResponse(resp, err)

	//this is a get request to ipfy, I dont want to go it, but this is how a get request should possibly look in basic form
	resp, err = http.Get("")
	if err != nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()
	v := struct{
		IP string `json:"ip"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&v)
	if err != nil{
		log.Fatal(err)
	}
	log.Println(v.IP)


}

func inspectResponse(resp *http.Response, err error){
	if err != nil{
		log.Fatal("Error occured wile marshaling json ", err)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Fatal("Error occured while trying to read http response")
	}
	log.Println(string(b))
}