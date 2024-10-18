package handler
 
import (
  "fmt"
  "net/http"

  "os"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, os.Getwd())
}