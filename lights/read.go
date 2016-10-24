package lights

import (
	"encoding/json"
	"github.com/cescoferraro/power/util"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

type ReadSerial struct {
	handler     http.Handler
	allowedHost string
}

func ReadSerialHandler(handler http.Handler) *ReadSerial {
	return &ReadSerial{handler: handler}
}

func (s *ReadSerial) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("start")
	ss, err := NewSerialDao()
	if err != nil {
		util.HttpAssertError(w, r, err, http.StatusBadRequest, "POWER/ACTIONS")
		return
	}
	log.Println("afterdao")

	cmd, err := ss.GetReadCommand(mux.Vars(r))
	if err != nil {
		util.HttpAssertError(w, r, err, http.StatusBadRequest, "POWER/ACTIONS")
		return
	}
	log.Println(cmd)

	_, err = ss.Port.Write([]byte(cmd))
	if err != nil {
		util.HttpAssertError(w, r, err, http.StatusBadRequest, "POWER/ACTIONS")
		return
	}

	log.Println("between")

	reason, err := ss.ReadFromSerial(mux.Vars(r))
	if err != nil {
		util.HttpAssertError(w, r, err, http.StatusBadRequest, "POWER/ACTIONS")
		return
	}

	this := struct {
		Time    time.Time
		Message string
	}{
		Time:    time.Now(),
		Message: strconv.FormatBool(reason),
	}
	go ss.Free()
	Tjson, _ := json.Marshal(this)
	w.WriteHeader(200)
	w.Write(Tjson)

}
