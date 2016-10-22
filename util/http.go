package util

import (
	"strconv"
	"net/http"
	"github.com/fatih/color"
	"github.com/spf13/viper"
	"log"
)

func GetBodySize(r *http.Request) int {
	bodySize, err := strconv.Atoi(r.Header["Content-Length"][0]);
	if err != nil {
		LogIfVerbose("Problem with Content-Lenght")
	}
	return bodySize
}

type Adapter func(http.Handler) http.Handler
// Adapt h with all specified adapters.
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

func NEWLogIfVerbose(cor color.Attribute, block, logg string) {
	if viper.GetBool("verbose") {
		red := color.New(cor).SprintFunc()
		log.Printf("[%s] "+logg, red(block))
	}
}

func HttpAssertError(w http.ResponseWriter,r *http.Request, err error, code int, origin string) {
	defer r.Body.Close()
	NEWLogIfVerbose(color.FgBlack, origin, err.Error())
	http.Error(w, http.StatusText(code), code)
	return
}



func EnableCORS() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if origin := r.Header.Get("Origin"); origin != "" {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers",
					"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			}
			// Stop here if its Preflighted OPTIONS request
			if r.Method == "OPTIONS" {
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

