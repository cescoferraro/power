package cmd

import (
	"bufio"
	"fmt"
	"github.com/cescoferraro/power/util"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/exec"
)

var ExposeCmd = &cobra.Command{
	Use:   "expose",
	Short: "A brief description of your command",
	Long:  `A loooooooonger description of your command.`,
	Run: func(cobracmd *cobra.Command, args []string) {
		RunNgrok(viper.GetString("port"))
	},
}

func init() {
	flags := util.CommandFlag{
		util.Flag{
			Name:        "port",
			Short:       "p",
			Description: "A descriptio about this cool flag",
			Value:       5000},
		util.Flag{
			Name:        "ngrok",
			Description: "A ngrok.io token",
			Value:       "youneedavalidngroktokenbaby"},
	}
	RootCmd.AddCommand(flags.Register(ExposeCmd))

}

func RunNgrok(port string) {
	d := color.New(color.FgCyan, color.Bold)
	d.Println("NGROK LOGS ARE IN THIS COLOR!")
	// docker build current directory
	cmdName := "ngrok"
	cmdArgs := []string{
		"http",
		"--subdomain",
		viper.GetString("device"),
		"--log",
		"stdout",
		"--authtoken",
		viper.GetString("ngrok"),
		"--log-level",
		"info",
		port,
	}
	executeShellCmd(cmdName, cmdArgs)
}

func executeShellCmd(cmdName string, cmdArgs []string) {
	cmd := exec.Command(cmdName, cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			d := color.New(color.FgCyan, color.Bold)
			d.Printf("%s\n", scanner.Text())
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
}
