package handler
 
import (
  "net/http"
	"os"
	"fmt"
	
	"cornbun-proxy/vehicle"
)
 
func Unlock(w http.ResponseWriter, r *http.Request) {
  params := r.URL.Query();
	vin := params.Get("vin")
	token := params.Get("token")

	tmpFile, err := os.CreateTemp("", "private")
	if err!= nil {
		fmt.Println(err)
		return
	}

	privateKey := `
-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIEr3nyLjulknyqkNsmgMQNxyggUgHSVSyC3EjdZdU/QHoAoGCCqGSM49
AwEHoUQDQgAEVf9iFUHygS4sAJDCFiY/lqJDhMueGWrktchxAKQtP4bdhHxHUyWl
zhyqnEhyS8AhPuXrgJ0+c7I8L18M82R6lQ==
-----END EC PRIVATE KEY-----
`

	tmpFile.WriteString(privateKey)

	vehicle.Unlock(vin, token, tmpFile.Name())

	fmt.Fprintf(w, "success unlock")

	defer tmpFile.Close()
}