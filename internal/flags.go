package internal

import (
	"flag"
	"fmt"
	"os"
)

type Flag struct {
	*flag.FlagSet
	cmdUsage string
}

const (
	Src          = "src"
	Des          = "des"
	Mode         = "mode"
	defaultMode  = "-s"
	defaultValue = ""
)

var (
	Cmd         *Flag
	Subcommands map[string]*Flag
)

func init() {
	getCmd := &Flag{
		FlagSet:  flag.NewFlagSet("get", flag.ExitOnError),
		cmdUsage: "Get the remote file(src) and download to local file(des).",
	}
	getCmd.String(Src, defaultValue, "(required) the remote file on chunk server.")
	getCmd.String(Des, defaultValue, "(required) the local file.")

	addCmd := &Flag{
		FlagSet:  flag.NewFlagSet("add", flag.ExitOnError),
		cmdUsage: "Put the local file(src) and upload to remote file(des).",
	}
	addCmd.String(Src, defaultValue, "(required) the local file.")
	addCmd.String(Des, defaultValue, "(required) the remote path on chunk server.")

	mkdirCmd := &Flag{
		FlagSet:  flag.NewFlagSet("mkdir", flag.ExitOnError),
		cmdUsage: "Make a directory at target path.",
	}
	mkdirCmd.String(Des, defaultValue, "(required) the remote path on chunk server.")

	removeCmd := &Flag{
		FlagSet:  flag.NewFlagSet("remove", flag.ExitOnError),
		cmdUsage: "Remove the remote file(des).",
	}
	removeCmd.String(Des, defaultValue, "(required) the remote file.")

	moveCmd := &Flag{
		FlagSet:  flag.NewFlagSet("move", flag.ExitOnError),
		cmdUsage: "Move the remote file(src) to another remote file(des).",
	}
	moveCmd.String(Src, defaultValue, "(required) the remote file on chunk server.")
	moveCmd.String(Des, defaultValue, "(required) the remote file that src moved to.")

	listCmd := &Flag{
		FlagSet:  flag.NewFlagSet("list", flag.ExitOnError),
		cmdUsage: "List the all files in the remote Directory(des).",
	}
	listCmd.String(Des, defaultValue, "(required) the remote Directory.")
	listCmd.String(Mode, defaultMode, "(optional) the read mode, l means latest, s means stale(about 50ms delay).")

	statCmd := &Flag{
		FlagSet:  flag.NewFlagSet("stat", flag.ExitOnError),
		cmdUsage: "Get the specified file's information.",
	}
	statCmd.String(Des, defaultValue, "(required) the remote file.")
	statCmd.String(Mode, defaultMode, "(optional) the read mode, l means latest, s means stale(about 50ms delay).")

	renameCmd := &Flag{
		FlagSet:  flag.NewFlagSet("rename", flag.ExitOnError),
		cmdUsage: "Rename the specified file to a new name.",
	}
	renameCmd.String(Src, defaultValue, "(required) the specified file path.")
	renameCmd.String(Des, defaultValue, "(required) the new name.")

	// 注册
	Subcommands = map[string]*Flag{
		getCmd.Name():    getCmd,
		addCmd.Name():    addCmd,
		mkdirCmd.Name():  mkdirCmd,
		removeCmd.Name(): removeCmd,
		moveCmd.Name():   moveCmd,
		listCmd.Name():   listCmd,
		statCmd.Name():   statCmd,
		renameCmd.Name(): renameCmd,
	}

	if len(os.Args) < 2 {
		ShowUsage(Subcommands)
	}

	Cmd = Subcommands[os.Args[1]]
	if Cmd == nil {
		ShowUsage(Subcommands)
	}

	err := Cmd.Parse(os.Args[2:])
	if err != nil {
		ShowUsage(Subcommands)
	}

}

func ShowUsage(subcommands map[string]*Flag) {
	fmt.Printf("Usage: ./Cmd.exe COMMAND\n\n")
	for _, v := range subcommands {
		fmt.Printf("%s %s\n", v.Name(), v.cmdUsage)
		v.PrintDefaults()
		fmt.Println()
	}
	os.Exit(2)
}
