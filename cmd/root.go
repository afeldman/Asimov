package cmd

import (
       "fmt"
       "log"
       "path"

       homedir "github.com/mitchellh/go-homedir"

       "github.com/spf13/cobra"
       "github.com/spf13/viper"

  //     "robot"
)

var (
    Asimov = &cobra.Command{
    	   Use:   "asimov",
    	   Short: "Asimov is a backup tool for Robots using FTP.",
    	   Long:  `A fast and flexible backup tool for robot FTP Server.
    	      	   Each Robot using FTP can be backuped.`,
}

     confFile string
     )

func init() {
     cobra.OnInitialize(initConfig)

     Asimov.PersistentFlags().StringVar(&confFile,"config","","config file (default $HOME/.config/asimov/asimov.yaml)")

     Asimov.AddCommand(version)
     Asimov.AddCommand(add)
}

func Execute() {
     Asimov.Execute()
}

func initConfig() {
     if confFile != "" {
     	bfg.Config_Path = confFile
     	viper.SetConfigFile(confFile)
     } else {
       home, err := homedir.Dir()
       if err != nil {
       	  log.Fatal(err)
       }

       bfg.Config_Path = path.Join(home,".config","asimov","asimov.yaml")

       viper.AddConfigPath(path.Join(home,".config","asimov"))
       viper.SetConfigName("asimov")
     }

     viper.AutomaticEnv()

     if err := viper.ReadInConfig(); err != nil {
     	log.Println(err)
     } else {
        if err := viper.Unmarshal(&bfg); err != nil {
	   log.Fatal("unable to decode into the Backup Structure, %v", err)
	}
     }
}