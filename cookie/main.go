package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",handler)
	// 8080ポートで起動
	http.ListenAndServe(":8080",nil)
}

func handler(w http.ResponseWriter,r *http.Request) {
	w.Header().Add("Set-Cookie", "VISIT=TRUE")
	if _,ok := r.Header["Cookie"]; ok {
		fmt.Fprint(w, "<html><body>2回以降</body></html>\n")
	}else {
		fmt.Fprint(w, "<html><body>初回</body></html>\n")
	}
}

