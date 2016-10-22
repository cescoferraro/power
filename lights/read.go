package lights

import (
	"net/http"
	"github.com/cescoferraro/power/models"
	"github.com/cescoferraro/power/util"
	"github.com/gorilla/mux"
	"time"
	"encoding/json"
	"strconv"
)

type ReadSerial struct {
	handler     http.Handler
	allowedHost string
}

func ReadSerialHandler(handler http.Handler) *ReadSerial {
	return &ReadSerial{handler: handler}
}

func (s *ReadSerial) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ss, err := models.NewSerialDao()
	if err != nil {
		util.HttpAssertError(w, r, err, http.StatusBadRequest, "POWER/ACTIONS"); return
	}


	cmd, err := ss.GetReadCommand(mux.Vars(r))
	if err != nil {
		util.HttpAssertError(w, r, err, http.StatusBadRequest, "POWER/ACTIONS"); return
	}

	_, err = ss.Port.Write([]byte(cmd))
	if err != nil {
		util.HttpAssertError(w, r, err, http.StatusBadRequest, "POWER/ACTIONS"); return
	}


	reason, err := ss.ReadFromSerial(mux.Vars(r))
	if err != nil {
		util.HttpAssertError(w, r, err, http.StatusBadRequest, "POWER/ACTIONS"); return
	}
	go ss.Free()


	this := struct {
		Time    time.Time;
		Message string;
	}{
		Time:time.Now(),
		Message:strconv.FormatBool(reason),
	}
	Tjson, _ := json.Marshal(this)
	w.WriteHeader(200)
	w.Write(Tjson)

}



