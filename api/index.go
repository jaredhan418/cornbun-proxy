package handler
 
import (
  "fmt"
  "net/http"

  "os"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  pwd, err  := os.Getwd()
	if err != nil {
		return
	}
  fmt.Fprintf(w, pwd)
}