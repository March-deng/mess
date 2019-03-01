package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "cong",
	Short: "cong is a set of service developed by dengcong",
	Long:  `cong is a set of service developed by dengcong`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("call the rootcmd")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println("execute root cmd error")
		os.Exit(1)
	}
}

func onInitiate() {

}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.Println("read home dir error:", err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".cong")
	}
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.shortener.yaml)")
	rootCmd.PersistentFlags().Bool("verbose", false, "Make the operation more talkative")
	viper.BindPFlag("config.verbose", rootCmd.PersistentFlags().Lookup("verbose")) // nolint: errcheck

}
