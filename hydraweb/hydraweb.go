package hydraweb

import (
	"fmt"
	"net/http"

	"github.com/Hydra/hlogger"

	//"time"8
	"log"
)

func Run() {
	http.HandleFunc("/", sroot)
	//handlerFunction which we made
	http.Handle("/testhandle", newHandler())
	// http://localhost:8080/testquery?key1=1&key2=1
	http.HandleFunc("/testquery", queryTestHandler)
	//newHandler is now default
	http.ListenAndServe(":8080", nil)

	/*
		//gives the endpoints and sets the timeouts
			server := &http.Server{
				Addr:         ":8080",
				Handler:      newHandler(),
				ReadTimeout:  5 * time.Second,
				WriteTimeout: 5 * time.Second,
			}
			server.ListenAndServe()
	*/
}

func queryTestHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("Forms", r.Form)
	q := r.URL.Query()
	message := fmt.Sprintf("Query map: %v \n", q)

	v1, v2 := q.Get("key1"), q.Get("key2")
	if v1 == v2 {
		message = message + fmt.Sprintf("V1 and V2 are equal %s \n", v1)
	} else {
		message = message + fmt.Sprintf("V1 is equal %s, V2 is equal %s \n", v1, v2)
	}
	fmt.Fprint(w, message)
}

func sroot(w http.ResponseWriter, r *http.Request) {
	logger := hlogger.GetInstance()
	fmt.Fprint(w, "Welcome to the Hydra software system")
	logger.Println("Received an http Get request on root url")
}
