package cmd

import (
	"github.com/spf13/cobra"
	"github.com/cescoferraro/power/util"
	"github.com/spf13/viper"
	"strings"
	"io/ioutil"
	"os/exec"
	"fmt"
	"os"
	"bufio"
)

var ExposeCmd = &cobra.Command{
	Use:   "expose",
	Short: "A brief description of your command",
	Long:  `A loooooooonger description of your command.`,
	Run: func(cobracmd *cobra.Command, args []string) {
		hardwareUnique, _ := ioutil.ReadFile("/var/lib/dbus/machine-id")
		viper.SetDefault("device", strings.Trim(string(hardwareUnique), "\n"))

		util.LogIfVerbose("verbose enabled")
		util.LogIfVerbose(string(hardwareUnique))
		// docker build current directory
		cmdName := "ngrok"

		cmdArgs := []string{
			"http",
			"--subdomain",
			strings.Trim(string(hardwareUnique), "\n"),
			"--log",
			"stdout",
			"--authtoken",
			"***REMOVED***",
			"--log-level",
			"info",
			viper.GetString("port"),
		}

		cmd := exec.Command(cmdName, cmdArgs...)
		cmdReader, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
			os.Exit(1)
		}

		scanner := bufio.NewScanner(cmdReader)
		go func() {
			for scanner.Scan() {
				fmt.Printf("%s\n", scanner.Text())
			}
		}()

		err = cmd.Start()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
			os.Exit(1)
		}

		err = cmd.Wait()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
			os.Exit(1)
		}


	},
}

func init() {

	flags := util.CommandFlag{
		util.Flag{
			Name:"port",
			Short: "p",
			Description: "A descriptio about this cool flag",
			Value:5000},
	}

	RootCmd.AddCommand(flags.Register(ExposeCmd))

}

