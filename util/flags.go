package util

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

type Flag struct {
	Name        string
	Short       string
	Description string
	Value       interface{}
}

type CommandFlag []Flag

func (flags CommandFlag) Register(cmd *cobra.Command) *cobra.Command {
	for _, i := range flags {
		viper.BindEnv(strings.ToUpper(i.Name))
		switch i.Value.(type) {
		case int:
			cmd.Flags().IntP(i.Name, i.Short, i.Value.(int), i.Description)
		case bool:
			cmd.Flags().BoolP(i.Name, i.Short, i.Value.(bool), i.Description)
		case string:
			cmd.Flags().StringP(i.Name, i.Short, i.Value.(string), i.Description)
		}
		viper.BindPFlag(strings.ToUpper(i.Name), cmd.Flags().Lookup(i.Name))
	}
	return cmd
}

func ImposeChannelFlag() {
	name := "channels"
	//exit program if channels do not make sense
	////TODO: Exit on channels not multiple of 8, and make it work
	if viper.GetInt(name) == 0 {
		log.Println(name + " flag is mandatory!")
		os.Exit(0)
	}
}
