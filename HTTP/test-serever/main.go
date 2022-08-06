package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Println("start http listening: 18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}


//HTTP/0.9を疑似的に試す
//$ curl --http1.0 http://localhost:18888/greeting

//HTTP/1.0
//$ curl -v --http1.0 http://localhost:18888/greeting

//ヘッダーの送信
//$ curl --http1.0 -H "X-Test:Hello" http://localhost:18888/greeting