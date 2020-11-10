package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "http://requestbin.net/r/1jl9i0u1"
	//comment chunk begin
	resp, err := http.Get(url)
	// inspectResponse(resp, err)

	// data, err := json.Marshal(struct {
	// 	X int
	// 	Y float32
	// }{X: 4, Y: 3.8})
	// if err != nil {
	// 	log.Fatal("Error occured while marshaling json ", err)
	// }

	// resp, err = http.Post(url, "application/json", bytes.NewReader(data))
	// inspectResponse(resp, err)
	//comment chunk end
	//comment chunk begin
	// client := http.Client{
	// 	Timeout: 3 * time.Second,
	// }
	// client.Get(url)
	//comment chunk end
	//comment chunk begin
	// req, err := http.NewRequest("PUT", "http://requestbin.net/r/1jl9i0u1", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// req.Header.Add("x-testheader", "learning go header")
	// req.Header.Set("User-Agent", "Go learning HTTP/1.1")
	// resp, err = client.Do(req)
	// inspectResponse(resp, err)
	//comment chunk end
	//comment chunk begin
	resp, err = http.Get("https://api.ipify.org?format=json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	v := struct {
		IP string `json:"ip"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&v)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(v.IP)
	//comment chunk end
}

func inspectResponse(resp *http.Response, err error) {
	if err != nil {
		log.Fatal("Error occured while marshaling json ", err)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error occured while trying to read http response body ", err)
	}
	log.Println(string(b))
}
