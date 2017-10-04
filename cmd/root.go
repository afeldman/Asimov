package cmd

import (
       "log"
       "path"

       homedir "github.com/mitchellh/go-homedir"

       "github.com/spf13/cobra"
       "github.com/spf13/viper"
)

const ASIMOV_NAME = "asimov"

var (
    Asimov = &cobra.Command{
    	   Use:   ASIMOV_NAME,
    	   Short: "Asimov is a backup tool for Robots using FTP.",
    	   Long:  `A fast and flexible backup tool for robot FTP Server.
    	      	   Each Robot using FTP can be backuped.`,
}

     confFile string
     )

func init() {
     cobra.OnInitialize(initConfig)

     Asimov.PersistentFlags().StringVar(&confFile,"config","${HOME}/.config/asimov/asimov.config.yml","config file")

     Asimov.AddCommand(versionCmd)

}

func Execute() {
     Asimov.Execute()
}

func initConfig() {
     if confFile != "" {
     	viper.SetConfigFile(confFile)
     } else {
       home, err := homedir.Dir()
       if err != nil {
       	  log.Fatal(err)
       }

       viper.AddConfigPath(path.Join(home,".config","asimov"))
       viper.SetConfigName("asimov.conf")
     }

     viper.AutomaticEnv()

     if err := viper.ReadInConfig(); err != nil {
     	log.Println(err)
     }
}