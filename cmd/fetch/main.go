package main

import (
	"fmt"
	"github.com/leozqi/fetch/internal/manager"
	"github.com/urfave/cli/v2"
	"os"
)

func fromCmd(cCtx *cli.Context) error {
	if cCtx.Args().Len() < 1 {
		fmt.Println("Give target name")
		os.Exit(1)
	} else if cCtx.Args().Len() < 2 {
		fmt.Println("Give repo URL")
		os.Exit(1)
	}

	state := manager.LoadState()
	err := manager.AddSource(
		state,
		cCtx.Args().First(),
		cCtx.Args().Get(1),
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Added new package source:", cCtx.Args().Get(0), "@", cCtx.Args().Get(1))
	err = manager.SaveState(state)

	return err
}

func refreshCmd(cCtx *cli.Context) error {
	state := manager.LoadState()
	fmt.Println("refreshing packages from sources:")

	for name, url := range state.Sources {
		fmt.Println(name, "@", url)
	}

	oldVersion := map[string]string{}
	for k, v := range state.CurrentVersion {
		oldVersion[k] = v
	}

	err := manager.RefreshSources(state)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("refreshed packages")

	for name, version := range state.CurrentVersion {
		if version != oldVersion[name] {
			fmt.Println("updated", name, oldVersion[name], "->", version)
		}
	}
	return nil
}

func main() {
	app := &cli.App{
		Name:  "fetch",
		Usage: "The universal package manager",
		Commands: []*cli.Command{
			{
				Name:   "refresh",
				Usage:  "Refresh list of packages",
				Action: refreshCmd,
			},
			{
				Name:  "install",
				Usage: "Install a package",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Installed ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:  "update",
				Usage: "Update an preinstalled package",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Updated ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:  "uninstall",
				Usage: "Uninstall an installed package",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Removed ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:  "list",
				Usage: "Lists installed packages",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Removed ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:  "view",
				Usage: "View all files associated with package on computer",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Viewed", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:  "depends",
				Usage: "view all dependencies of the package",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Depends", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:   "from",
				Usage:  "add new source repository for packages",
				Action: fromCmd,
			},
			{
				Name:  "sources",
				Usage: "view all sources for packages",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Added new source", cCtx.Args().First())
					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}
