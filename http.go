package main

import (
	"io"
	"log"
	"net/http"

	"github.com/rs/cors"
)

// ConnectHandler listens for POST requests on /connect.
// A valid request should have an encoded WebRTC offer as its body.
func ConnectHandler() (http.Handler, error) {
	mux := http.NewServeMux()
	mux.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			//TODO: log the error
			Logger.Infof("Got an http request with bad method %q\n", r.Method)
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		offer := make([]byte, 4096)
		l, e := r.Body.Read(offer)
		if e != io.EOF {
			Logger.Errorf("Failed to read http request body: %q", e)
		}
		Logger.Infof("Got a valid POST request with offer: %q", string(offer[:l]))
		peer := Listen(string(offer[:l]))
		// reply with server's key
		w.Write(peer.Answer)
	})
	return cors.Default().Handler(mux), nil
}

func HTTPGo(address string) (e error) {
	h, e := ConnectHandler()
	if e != nil {
		log.Fatal(e)
		return
	}
	http.ListenAndServe(address, h)
	return
}