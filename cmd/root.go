package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github-manager/api"
	"github.com/spf13/viper"
	homedir "github.com/mitchellh/go-homedir"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use: "github-manager",
	Short: "A cli application to manage our GitHub Enterprise instance",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	client := NewClient()
	
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.github-manager.yaml)")
}

type config struct {
	Token string
	Enterprise_id string
}

var conf config
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".passparams")
	}

	viper.Set("token", os.Getenv("GITHUB_TOKEN"))
	viper.Set("enterprise_id", os.Getenv("GITHUB_ENTERPRISEID"))

	err := viper.Unmarshal(&conf)
	if err != nil { panic("Unable to decode") }
}

