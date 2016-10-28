package util

import (
	"github.com/fatih/color"
	"github.com/spf13/viper"
	"log"
	"strings"
	"sort"
	"net"
	"crypto/sha1"
	"encoding/base64"
)

func LogIfVerbose(logg interface{}) {
	if viper.GetBool("verbose") {
		log.Println(logg)
	}
}

func RunIfVerbose(logg func()) {
	if viper.GetBool("verbose") {
		logg()
	}
}

func PrintViperConfig() {
	keys := viper.AllKeys()
	sort.Strings(keys)
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()
	log.Println("***********VIPER*************")
	for _, key := range keys {
		log.Printf("%s %s %s\n", red(strings.ToUpper(key)), yellow(key), white(viper.GetString(key)))
	}
	log.Println("*****************************")
	return
}


func UniqueID() string{
	var MAC string
	ifs, _ := net.Interfaces()
	for _, v := range ifs {
		if v.Name == "eth0" {
			MAC = v.HardwareAddr.String()
		}
	}
	hasher := sha1.New()

	hasher.Write([]byte(MAC))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	removeEquals := func(r rune) rune {
		if strings.IndexRune("=", r) < 0 {
			return r
		}
		return -1
	}
	thing := strings.Map(removeEquals, sha)
	return thing
}
