//Filename:cmd/api/entries.go
package main

import	(
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)
//create entry handler for the POST /v1/entries endpoint
func (app *application) createEntryHandler (w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Create a new entry...")
}
//create entry handler for the GET /v1/entries/:id endpoint
func (app *application) showEntryHandler (w http.ResponseWriter, r *http.Request){
	//use the '*ParamsfromContext())' function to get the request context as a slice
	params := httprouter.ParamsFromContext(r.Context())
	//get the value of the 'id' parameter
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	//display the school id
	fmt.Fprintln(w, "Show the details for the entry %d\n", id)
}