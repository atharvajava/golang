/* http.Handle to handle web requests
 */
package main

import (
    "fmt"
    "net/http"
)

//The function handler is of the type http.HandlerFunc. It takes an http.ResponseWriter and an http.Request as its arguments.
func handler(w http.ResponseWriter, r *http.Request) {
	//An http.ResponseWriter value assembles the HTTP server's response; by writing to it, we send data to the HTTP client.
	//An http.Request is a data structure that represents the client HTTP request.
	//r.URL.Path is the path component of the request URL.
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	//The trailing [1:] means "create a sub-slice of Path from the 1st character to the end." 
	//This drops the leading "/" from the path name.
}

//begins with call to handle func
func main() {
    http.HandleFunc("/", handler) // This tells the http package to handle all request to the web root ("/")
	// with handler


    http.ListenAndServe(":3000", nil)
	//The above specifies to listen to which port no.
}
