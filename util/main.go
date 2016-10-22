package util

import (
	"log"
	"net/http"
	"github.com/spf13/viper"
)

func LogIfVerbose(logg  interface{}) {
	if viper.GetBool("verbose") {
		log.Println(logg)
	}
}

func RunIfVerbose(logg func()) {
	if viper.GetBool("verbose") {
		logg()
	}
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func LogIfError(err error) {
	if err != nil {
		log.Println(err)
	}
}
func PrintIf(err error) {
	if err != nil {
		log.Println(err)
	}
}
func PrintRequestHeaders(r *http.Request) {
	for k, v := range r.Header {
		log.Println("key:", k, "value:", v)
	}
}

func PrintViperConfig() {

	// TODO: HANDLE NESTED YAMLS BETTER
	keys := viper.AllKeys()
	log.Println("***********VIPER*************")
	for _, key := range keys {
		log.Println(key + ": " + viper.GetString(key))

	}
	log.Println("*****************************")
	return
}
