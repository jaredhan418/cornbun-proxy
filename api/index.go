package handler
 
import (
  "fmt"
  "net/http"
  "io/ioutil"

  "os"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  pwd, err  := os.Getwd()
	if err != nil {
		return
	}

  files, err  := ioutil.ReadDir(pwd)
	if err != nil {
		return
	}

	fileNames := ""

  for _, file := range files {
		fileNames = fileNames + file.Name() + "##"
	}

  fmt.Fprintf(w, "<html><body><h1>Hello from Go API!</h1>" + "<div>"+ pwd + fileNames + "</div></body></html>")
}