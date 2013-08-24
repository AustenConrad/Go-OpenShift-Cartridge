package hello

import (
	"fmt"
	"net/http"
)

// Writes "hello [client ip address]" to the web page and to the log.
func Ip(rw http.ResponseWriter, req *http.Request) {

	// Write to the web page.
	fmt.Fprintln(rw, "Hello "+req.Header.Get("X-Forwarded-For"))

	// Write to the log.
	fmt.Println("Served client: "+req.Header.Get("X-Forwarded-For"))

}
