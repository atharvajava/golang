//Using net/http to serve wiki pages
package main

import(
	"fmt"
	"io/ioutil"
	"net/http"
)

//  viewHandler that will allow users to view a wiki page. It will handle URLs prefixed with "/view/".
//this function extracts the page title from r.URL.Path, the path component of the request URL.
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
// The Path is re-sliced with [len("/view/"):] to drop the leading "/view/" component of the request path. 
//This is because the path will invariably begin with "/view/", which is not part of the page's title.
    p, _ := loadPage(title)
	//The function then loads the page data, formats the page with a string of simple HTML, and writes it 
	//to w, the http.ResponseWriter.
	//Again, note the use of _ to ignore the error return value from loadPage. This is done here for simplicity
	// and generally considered bad practice. We will attend to this later.
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
//To use this handler, we rewrite our main function to initialize http using the viewHandler 
//to handle any requests under the path /view/.
func main() {
    http.HandleFunc("/view/", viewHandler)
    http.ListenAndServe(":3000", nil)
}