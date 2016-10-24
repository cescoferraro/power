package util

import (
	"github.com/fatih/color"
	"github.com/spf13/viper"
	"log"
	"strings"
	"sort"
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
