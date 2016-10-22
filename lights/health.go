package lights

import (
	"net/http"

	"time"
	"encoding/json"
)

type HealthSerial struct {
	handler     http.Handler
	allowedHost string
}

func HealthHandler(handler http.Handler) *HealthSerial {
	return &HealthSerial{handler: handler}
}

func (s *HealthSerial) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	this := struct {
		Time    time.Time;
		Message string;
	}{
		Time:time.Now(),
		Message:"OK",
	}
	Tjson, _ := json.Marshal(this)
	w.WriteHeader(200)
	w.Write(Tjson)

}



