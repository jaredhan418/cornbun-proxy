package handler
 
import (
  "fmt"
  "net/http"
	
	"cornbun-proxy/vehicle"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  params := r.URL.Query();
	vin := params.Get("vin")
	token := params.Get("token")

	vehicle.Unlock(vin, token)
  fmt.Fprintf(w, vin)
}