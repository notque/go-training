package main

import (
  "log"
  "net/http"
  "time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
  tm := time.Now().Format(time.RFC1123)
  w.Write([]byte("The time is: " + tm))
}

// AuditingHandler implements CADF Taxonomy on requests, and creates
// an Audit Entry in RabbitMQ
func AuditingHandler(next http.Handler) http.Handler {
  return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
    // Our middleware logic goes here...
    //setup interception of response metadata
	  writer := responseWriter{original: w}

	  //forward request to actual handler
	  l.handler.ServeHTTP(&writer, r)
    next.ServeHTTP(rw, request)
  })
}

func main() {
  mux := http.NewServeMux()

  // Convert the timeHandler function to a HandlerFunc type
  th := http.HandlerFunc(timeHandler)
  // And add it to the ServeMux
  mux.Handle("/time", th)

  log.Println("Listening...")
  http.ListenAndServe(":3000", mux)
}