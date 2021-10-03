/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"
)

// GreetCmd represents the Greet command
var GreetCmd = &cobra.Command{
	Use:   "Greet",
	Short: "The Greet command Greets anyone or the author",
	Long: `The Greet command Greets anyone or the author`,
	Run: func(cmd *cobra.Command, args []string) {
		Greet(args[0])
	},
}

func Greet(name string) {
	fmt.Println("Hello " + name)
}

func init() {
	rootCmd.AddCommand(GreetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// GreetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// GreetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
