package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

/*
http.Server
	type Server struct {
		Addr string
		Handler Handler // ←ここにインタフェースhttp.Handlerを満たすHandlerを代入する
		TLSConfig *tls.Config
		ReadTimeout time.Duration
		ReadHeaderTimeout time.Duration
		WriteTimeout time.Duration
		IdleTimeout time.Duration
		MaxHeaderBytes int
		TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
		ConnState func(net.Conn, ConnState)
		ErrorLog *log.Logger
		BaseContext func(net.Listener) context.Context
		ConnContext func(ctx context.Context, c net.Conn) context.Context
	}
*/

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello go</body></html>\n")
}

func main() {
	httpServer := &http.Server{
		Addr:    ":18888",
	}
	http.HandleFunc("/", handler)
	log.Println("start http listening :18888")
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}