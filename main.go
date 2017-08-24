package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/hokaccha/go-prettyjson"
)

func createHandler(rescode int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rq, _ := httputil.DumpRequest(r, false)
		fmt.Println(string(rq))
		body, _ := ioutil.ReadAll(r.Body)
		ct := r.Header.Get("Content-Type")
		if strings.Contains(ct, "json") {
			body, _ = prettyjson.Format(body)
		}
		fmt.Println(string(body))
		w.WriteHeader(rescode)
	})

}

func main() {
	address := flag.String("address", "localhost:9999", "listen address")
	rsp := flag.Int("responsecode", 200, "the responsecode")

	flag.Parse()

	http.Handle("/", createHandler(*rsp))

	log.Fatal(http.ListenAndServe(*address, nil))
}
