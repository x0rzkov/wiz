/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/alexkreidler/wiz/executor"
	"github.com/alexkreidler/wiz/server"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

func remove(l []string, item string) []string {
	for i, other := range l {
		if other == item {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}

// executorCmd represents the executor command
var executorCmd = &cobra.Command{
	Use:   "executor",
	Short: "Starts the Wiz Executor serving the Processor API",
	//	Long: `A longer description that spans multiple lines and likely contains examples
	//and usage of using your command. For example:
	//
	//Cobra is a CLI library for Go that empowers applications.
	//This application is a tool to generate the needed files
	//to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if daemon {
			args := remove(os.Args, "-d")

			c := exec.Command(args[0], args[1:]...)
			//c.Stdout = os.Stdout
			err := c.Start()
			if err != nil {
				log.Fatal(err)
			}

			log.Println("Started daemon with command:", c, ". Running on PID", c.Process.Pid)

			return
		}

		log.Println("Starting server on port", port)
		s := server.NewServer(executor.NewProcessorExecutor())

		err := s.Run(port)
		if err != nil {
			fmt.Println(err)
		}
	},
}

var port string
var daemon bool

func init() {
	rootCmd.AddCommand(executorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// executorCmd.PersistentFlags().String("foo", "", "A help for foo")

	executorCmd.Flags().StringVarP(&port, "port", "p", ":8080", "Sets the port that the executor serves the Processor API on")
	executorCmd.Flags().BoolVarP(&daemon, "daemon", "d", false, "If enabled, runs the executor as a daemon, simply spawns as a separate process.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// executorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
