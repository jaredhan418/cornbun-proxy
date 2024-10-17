package handler
 
import (
  "fmt"
  "net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  params := r.URL.Query();
	vin := params.Get("vin")

  fmt.Fprintf(w, vin)
}