# ASIMOV - Basic FANUC Robot Backup

**asimov** is a basic FTP Backup tool. Using **asimov** different commands are available

1. *help* the **asimov help screen**
2. *version* the version of the backup tool **asimov**
3. *del* delete a robot from the list of available robots
4. *add* add a robot to the list of robots
5. *backup* this command start the backup process of all robots in the configuration list

The configuration list is set to *$HOME/.config/asimov/asimov.yaml*. If you like to use a different configuration file please use the *--config* flag. The *config* flag is available in each subcommand

```
Usage:
  asimov [command]

Available Commands:
  add         Add a Robot to the Robot List
  backup      Asimov start backup. There are different possibilities to make the backup
  del         Remove a Robot from the Robot List
  help        Help about any command
  version     print version number of Asimov

Flags:
      --config string   config file (default $HOME/.config/asimov/asimov.yaml)
  -h, --help            help for asimov
  -d, --target string   Path to the backup Directory (default "Backup Path")

Use "asimov [command] --help" for more information about a command.
```

## add a robot in **asimov**
To add a robot to the configuration list us the `asimov add [name] [ip-address] [flags]` command. It is assumed that the port *21* is default. This can not be changed for the moment. The *ip-addess* is going to be checkt, if the given *ip-address* is valid because the user might have a writing mistake. The order of attribuis is important and can not be changed.  
Each name of a robot is unique. So if you have a name given twice the *ip-address* of the first is set with the *ip-address* of the second robot. The check for the *ip-address* is comming :wink:

## delete a robot in **asimov**
to remove a robot from the list of robots in the configuration file use ``asimov del [name] [flags]``. If the robot name is in the list of robots, this command deletets the robot from the list.

## backup you robots using **asimov**
To give you, the user a chose of your will, **asimov** provides different aliases for backup

- `asimov backup [command]`
- `asimov bkg [command]`
- `asimov back [command]`
- `asimov bak [command]`

does all the same. The different commands for the backup are working totally different :wink:

1. *all* the all command or simply *\** backup all files of the robot available in the standard ftp directory
2. *app* download all app data
3. *bin* backup all binary files
4. *vison* backup all vision files

# Development
To develop this project you are going to need

1. robot
2. [go-homedir](http://github.com/mitchellh/go-homedir) `go get github.com/mitchellh/go-homedir`
3. [cobra](http://github.com/spf13/cobra) `go get github.com/spf13/cobra`
4. [viper](http://github.com/spf13/viper) `go get github.com/spf13/viper`
5. [yaml.v2](http://gopkg.in/yaml.v2) `go get gopkg.in/yaml.v2` 

feel free to make changes :nerd:
