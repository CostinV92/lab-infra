# labtool

CLI app to wrap ansible playbooks controling the lab infrastructure.

### Dependencies
- [Cobra](https://github.com/spf13/cobra) 
- [Viper](https://github.com/spf13/viper)

### Directory structure
```
cmd/
 |- root.go
 |- subcommand1/
     |- subcommand1.go
     |- leaf1.go
     |- leaf2.go
 |- subcommand2/
     |- subcommand2.go
     |- leaf1.go
 |- subcommand3/
     |- subcommand3.go
```
Detailed info related to the structure [here](https://github.com/spf13/cobra/blob/main/site/content/user_guide.md#organizing-subcommands)

### Add a new command
1. run `cobra-cli add [new_command_name]`
2. create folders, and the move the newly created `new_command_name.go` to its hirarchical place
3. add desired flags to the `init()` function of the newly created command
4. every new command MUST call `utilcmd.InitRunCmd()` in its `Run()/PreRun()` function
    - this initializes the `exec.Cmd` to be run and loads the config file
5. call `AddCommand(new_command)` in the `init()` function of the newly created command's parent
6. call `utilcmd.AddExtraVar()`and `utilcmd.SetPlaybook()` accordingly

See existing commands for details.