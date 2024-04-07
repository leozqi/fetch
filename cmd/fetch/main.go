package main

import (
	"fmt"
	"github.com/leozqi/fetch/internal/gitpipe"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "fetch",
		Usage: "The universal package manager",
		Commands: []*cli.Command{
			{
				Name:  "refresh",
				Usage: "Refresh list of packages",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Refreshed")
					err := gitpipe.GetOrigin("https://github.com/leozqi/fetch-index.git")
					if err != nil {
						fmt.Println(err)
					}
					return nil
				},
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
				Name:  "from",
				Usage: "add new target repository for packages",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Added new remote", cCtx.Args().First())
					return nil
				},
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
