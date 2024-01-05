package main

import(
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	Get()
	GetQuery()
	Head()
}

// $ curl http://localhost:18888
func Get() {
	resp, err := http.Get("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	log.Println("Status:", resp.Status)
	log.Println("StatusCode:", resp.StatusCode)
	log.Println("Headers:", resp.Header)
}

// $ curl -G --data-urlencode "query=hello world" http://localhost:18888
func GetQuery() {
	values := url.Values{
		"query": {"hello world"},
	}
	resp, _ := http.Get("http://localhost:18888" + "?" + values.Encode())
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}

// $ curl --head http://localhost:18888
func Head() {
	resp, err := http.Head("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
	log.Println("Headers:", resp.Header)
}