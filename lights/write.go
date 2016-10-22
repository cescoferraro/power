package lights

import (
	"net/http"
	"github.com/cescoferraro/power/models"
	"github.com/cescoferraro/power/util"
	"github.com/gorilla/mux"
	"time"
	"encoding/json"
	"log"
)

type Serial struct {
	handler     http.Handler
	allowedHost string
}

func SerialHandler(handler http.Handler) *Serial {
	return &Serial{handler: handler}
}

func (s *Serial) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serial, err := models.NewSerialDao()

	if err != nil {
		util.HttpAssertError(w, r, err, http.StatusBadRequest, "POWER/ACTIONS"); return
	}



	cmd, err := serial.GetWriteCommand(mux.Vars(r))
	if err != nil {
		util.HttpAssertError(w, r, err, http.StatusBadRequest, "POWER/ACTIONS"); return
	}

	log.Println([]byte(cmd))
	_, err = serial.Port.Write([]byte(cmd))
	if err != nil {
		util.HttpAssertError(w, r, err, http.StatusBadRequest, "POWER/ACTIONS"); return
	}
	go serial.Free()


	this := struct {
		Time    time.Time;
		Message string;
	}{
		Time:time.Now(),
		Message:"Success",
	}
	Tjson, _ := json.Marshal(this)
	w.WriteHeader(200)
	w.Write(Tjson)

}




