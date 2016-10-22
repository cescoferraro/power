package lights

import (
	"net/http"
	"encoding/json"
	"github.com/cescoferraro/power/models"
	"github.com/cescoferraro/power/util"
)

type Status struct {
	handler     http.Handler
	allowedHost string
}
type StatusResponse map[string]bool

func StatusHandler(handler http.Handler) *Status {
	return &Status{handler: handler}
}

func (s *Status) ServeHTTP(w http.ResponseWriter, r *http.Request) {


	serial, err :=models.NewSerialDao()
	if err != nil {
		util.HttpAssertError(w, r, err, http.StatusBadRequest, "POWER/STATUS"); return
	}


	status, err := serial.CurrentState()
	if err != nil {
		util.HttpAssertError(w, r, err, http.StatusBadRequest, "POWER/STATUS"); return
	}
	go serial.Free()

	var response StatusResponse = status
	text, err := json.Marshal(response)
	w.WriteHeader(200)
	w.Write(text)

}



