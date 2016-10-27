package cmd

import (
	"github.com/cescoferraro/power/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"os"

	"github.com/cescoferraro/power/lights"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

var jwt string
var version string

var RunserverCmd = &cobra.Command{
	Use:   "runserver",
	Short: "A brief description of your command",
	Long:  `A loooooooonger description of your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("jwt", jwt)

		viper.Set("version", version)
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
	sha:= util.UniqueID()
	viper.SetDefault("device", sha)
	viper.SetDefault("url", "http://"+sha+".ngrok.io")

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
			Name:        "serial-port",
			Description: "A descriptio about this cool flag",
			Value:       "/dev/ttyACM0"},
		util.Flag{
			Name:        "owner",
			Short:       "o",
			Description: "A descriptio about this cool flag",
			Value:       "public"},
		util.Flag{
			Name:        "env",
			Description: "A descriptio about this cool flag",
			Value:       "prod"},
		util.Flag{
			Name:        "api",
			Description: "A descriptio about this cool flag",
			Value:       "https://api.cescoferraro.xyz/iot/devices"},
		util.Flag{
			Name:        "dev-api",
			Description: "A descriptio about this cool flag",
			Value:       "http://localhost:9000/iot/devices"},
		util.Flag{
			Name:        "ngrok",
			Description: "A ngrok.io token",
			Value:       "youneedavalidngroktokenbaby"},
		util.Flag{
			Name:        "ping-interval",
			Description: "The interval  this app send a ping request to the api. If --env=dev request go to --dev-api",
			Value:       12},
	}

	RootCmd.AddCommand(flags.Register(RunserverCmd))

}
