package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//InputReq struct
type InputReq struct {
	Name string `json:"name"`
}

//OutputRes struct
type OutputRes struct {
	Greeting string `json:"greeting"`
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

/*POSTOperation for POST operation
sample REST client input:
{
"name":"foo"
}
*/
func POSTOperation(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	inptreq := InputReq{}
	json.NewDecoder(req.Body).Decode(&inptreq)

	resp := OutputRes{
		Greeting: "Hello, " + inptreq.Name + "!",
	}

	resjson, _ := json.Marshal(resp)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(201)
	fmt.Fprintf(rw, "%s", resjson)
}

func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)
	mux.POST("/hello", POSTOperation)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
