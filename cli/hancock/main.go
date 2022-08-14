// Copyright (c) 2020-present Cloud <cloud@txthinking.com>
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of version 3 of the GNU General Public
// License as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/txthinking/hancock"
	"github.com/urfave/cli/v2"
)

func main() {
	cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
		fmt.Println("")
		fmt.Println("Note:")
		fmt.Println("")
		fmt.Println("    When adding an instance, the user must be allowed to execute sudo without a password!!!")
		fmt.Println("    nami and joker are automatically installed when you run command for the first time on instance.")
		fmt.Println("")
		fmt.Println("Add instance")
		fmt.Println("")
		fmt.Println("    $ hancock add --name mylinux --server 1.2.3.4:22 --user root --password mypassword")
		fmt.Println("    $ hancock add --name mylinux --server 1.2.3.4:22 --user root --key ./path/to/mykey.pem")
		fmt.Println("    $ hancock add -n mylinux -s 1.2.3.4:22 -u root -p mypassword")
		fmt.Println("    $ hancock add -n mylinux -s 1.2.3.4:22 -u root -k ./path/to/mykey.pem")
		fmt.Println("")
		fmt.Println("List instances")
		fmt.Println("")
		fmt.Println("    $ hancock list")
		fmt.Println("")
		fmt.Println("Remove instance")
		fmt.Println("")
		fmt.Println("    $ hancock mylinux remove")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("Run nami on an instance")
		fmt.Println("")
		fmt.Println("    $ hancock mylinux nami install brook")
		fmt.Println("    $ hancock mylinux nami list")
		fmt.Println("")
		fmt.Println("Run joker on an instance")
		fmt.Println("")
		fmt.Println("    $ hancock mylinux joker brook server --listen :9999 --password hello")
		fmt.Println("    $ hancock mylinux joker last")
		fmt.Println("    $ hancock mylinux joker list")
		fmt.Println("    $ hancock mylinux joker stop 1234")
		fmt.Println("    $ hancock mylinux joker log 1234")
		fmt.Println("")
		fmt.Println("Run command and wait output on an instance")
		fmt.Println("")
		fmt.Println("    $ hancock mylinux echo hello")
		fmt.Println("    $ hancock mylinux sleep 3 '&&' echo hello")
		fmt.Println("")
		fmt.Println("Start command and do not wait output on an instance")
		fmt.Println("")
		fmt.Println("    $ hancock mylinux start echo hello")
		fmt.Println("    $ hancock mylinux start sleep 3 '&&' echo hello")
		fmt.Println("")
		fmt.Println("Upload your own command")
		fmt.Println("")
		fmt.Println("    $ hancock mylinux upload ./path/to/command")
		fmt.Println("")
	}
	app := cli.NewApp()
	app.Name = "hancock"
	app.Version = "20220814"
	app.Usage = "Deploy and run command on remote instances with built-in nami, joker"
	app.Authors = []*cli.Author{
		{
			Name:  "Cloud",
			Email: "cloud@txthinking.com",
		},
	}
	app.Copyright = "https://github.com/txthinking/hancock"
	app.Commands = []*cli.Command{
		&cli.Command{
			Name:  "add",
			Usage: "Add or overwrite instance",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "name",
					Aliases: []string{"n"},
					Usage:   "Give a name to instance, like: mylinux",
				},
				&cli.StringFlag{
					Name:    "server",
					Aliases: []string{"s"},
					Usage:   "SSH address off instance, like: 1.2.3.4:22",
				},
				&cli.StringFlag{
					Name:    "user",
					Aliases: []string{"u"},
					Usage:   "SSH user",
				},
				&cli.StringFlag{
					Name:    "password",
					Aliases: []string{"p"},
					Usage:   "SSH password, must provide one of key or password",
				},
				&cli.StringFlag{
					Name:    "key",
					Aliases: []string{"k"},
					Usage:   "SSH private key file path, must provide one of key or password, like: ./path/to/key.pem",
				},
			},
			Action: func(c *cli.Context) error {
				h, err := hancock.NewHancock()
				if err != nil {
					return err
				}
				defer h.Close()
				if c.String("name") == "" || c.String("server") == "" || c.String("user") == "" || (c.String("password") == "" && c.String("key") == "") {
					cli.ShowCommandHelp(c, "add")
					return nil
				}
				return h.Add(c.String("name"), c.String("server"), c.String("user"), c.String("password"), c.String("key"))
			},
		},
	}
	if len(os.Args) == 2 && os.Args[1] == "list" {
		h, err := hancock.NewHancock()
		if err != nil {
			log.Println(err)
			os.Exit(1)
			return
		}
		defer h.Close()
		h.PrintAll()
		return
	}
	if len(os.Args) == 3 && os.Args[2] == "remove" {
		h, err := hancock.NewHancock()
		if err != nil {
			log.Println(err)
			os.Exit(1)
			return
		}
		defer h.Close()
		if err := h.Remove(os.Args[1]); err != nil {
			log.Println(err)
			os.Exit(1)
			return
		}
		return
	}
	if len(os.Args) == 4 && os.Args[2] == "upload" {
		h, err := hancock.NewHancock()
		if err != nil {
			log.Println(err)
			os.Exit(1)
			return
		}
		defer h.Close()
		if err := h.Upload(os.Args[1], os.Args[3]); err != nil {
			log.Println(err)
			os.Exit(1)
			return
		}
		return
	}
	if len(os.Args) > 3 && os.Args[2] == "joker" && (os.Args[3] != "list" && os.Args[3] != "last" && os.Args[3] != "stop" && os.Args[3] != "log" && os.Args[3] != "help" && os.Args[3] != "version" && os.Args[3] != "--help" && os.Args[3] != "-h" && os.Args[3] != "--version" && os.Args[3] != "-v") {
		h, err := hancock.NewHancock()
		if err != nil {
			log.Println(err)
			os.Exit(1)
			return
		}
		defer h.Close()
		if err := h.Start(append([]string{os.Args[1]}, os.Args[2:]...)); err != nil {
			log.Println(err)
			os.Exit(1)
			return
		}
		return
	}
	if len(os.Args) > 3 && os.Args[2] == "start" {
		h, err := hancock.NewHancock()
		if err != nil {
			log.Println(err)
			os.Exit(1)
			return
		}
		defer h.Close()
		if err := h.Start(append([]string{os.Args[1]}, os.Args[3:]...)); err != nil {
			log.Println(err)
			os.Exit(1)
			return
		}
		return
	}
	if len(os.Args) > 2 && os.Args[1] != "add" {
		h, err := hancock.NewHancock()
		if err != nil {
			log.Println(err)
			os.Exit(1)
			return
		}
		defer h.Close()
		if err := h.Run(append([]string{os.Args[1]}, os.Args[2:]...)); err != nil {
			log.Println(err)
			os.Exit(1)
			return
		}
		return
	}
	if err := app.Run(os.Args); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
