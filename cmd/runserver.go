package cmd

import (
	"github.com/cescoferraro/power/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"os"

	"github.com/cescoferraro/power/lights"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var jwt string

var RunserverCmd = &cobra.Command{
	Use:   "runserver",
	Short: "A brief description of your command",
	Long:  `A loooooooonger description of your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("jwt", jwt)
		util.RunIfVerbose(util.PrintViperConfig)
		go PING()
		go RunNgrok(viper.GetString("port"))
		util.ImposeChannelFlag()
		ROUTER := mux.NewRouter()
		lights_router := lights.Routes(ROUTER)
		logs := handlers.LoggingHandler(os.Stdout, lights_router)
		http.ListenAndServe("0.0.0.0:"+viper.GetString("port"), logs)
	},
}

func init() {
	hardwareUnique, e := ioutil.ReadFile("/var/lib/dbus/machine-id")
	if e != nil {
		log.Println(e)
	}

	viper.SetDefault("device", strings.Trim(string(hardwareUnique), "\n"))
	viper.SetDefault("url", "http://"+strings.Trim(string(hardwareUnique), "\n")+".ngrok.io")
	flags := util.CommandFlag{
		util.Flag{
			Name:        "verbose",
			Short:       "v",
			Description: "A descriptio about this cool flag",
			Value:       true},
		util.Flag{
			Name:        "port",
			Short:       "p",
			Description: "A descriptio about this cool flag",
			Value:       5000},
		util.Flag{
			Name:        "ping",
			Description: "A descriptio about this cool flag",
			Value:       true},
		util.Flag{
			Name:        "channels",
			Short:       "c",
			Description: "A descriptio about this cool flag",
			Value:       8},
		util.Flag{
			Name:        "serial_port",
			Description: "A descriptio about this cool flag",
			Value:       "/dev/ttyACM0"},
		util.Flag{
			Name:        "owner",
			Short:       "o",
			Description: "A descriptio about this cool flag",
			Value:       "francescoaferraro@gmail.com"},
		util.Flag{
			Name:        "api",
			Description: "A descriptio about this cool flag",
			Value:       "https://api.cescoferraro.xyz/iot/devices"},
		util.Flag{
			Name:        "env",
			Description: "A descriptio about this cool flag",
			Value:       "Development"},
		util.Flag{
			Name:        "dev-api",
			Description: "A descriptio about this cool flag",
			Value:       "http://localhost:9000/iot/devices"},
		util.Flag{
			Name:        "ngrok-token",
			Description: "A descriptio about this cool flag",
			Value:       "***REMOVED***"},
		util.Flag{
			Name:        "ping-interval",
			Description: "The interval  between liveprobe check with the api",
			Value:       12},
	}

	RootCmd.AddCommand(flags.Register(RunserverCmd))

}
