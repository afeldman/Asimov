package cmd

import (
       "log"
       "path"

       homedir "github.com/mitchellh/go-homedir"

       "github.com/spf13/cobra"
       "github.com/spf13/viper"
)

var (
    Asimov = &cobra.Command{
    	   Use:   "asimov",
    	   Short: "Asimov is a backup tool for Robots using FTP.",
    	   Long:  `
	   	  ASIMOV a simple FANUC Robot Backup 
	   	  ==================================

To backup a FANUC Robot the FTP server on the FANUC controller can be used.
To make it simple to download all the files in an repository this tool helps.
This tool provides different kind of backups and a configuration setting to
make it easy to configurate the list and the destination where to backup and
which Robot controller to backup.

AUTHOR:
  Anton Feldmann <anton.feldmann@gmail.com>

`,
}

     confFile string
     )

func init() {
     cobra.OnInitialize(initConfig)

     Asimov.PersistentFlags().StringVar(&confFile,"config","","config file (default $HOME/.config/asimov/asimov.yaml)")
     Asimov.PersistentFlags().StringP("target","d","Backup Path","Path to the backup Directory")

     viper.BindPFlag("target", Asimov.PersistentFlags().Lookup("target"))

     viper.SetDefault("target","./backup")

     bfg.Destination = viper.GetString("target")

     Asimov.AddCommand(version)
     Asimov.AddCommand(add)
     Asimov.AddCommand(remove)
     Asimov.AddCommand(backup_root)
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