/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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

// modulesCmd represents the modules command
var modulesCmd = &cobra.Command{
	Use:   "modules",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("modules called")
	},
}

const extensionsPackage = "github.com/neuronlabs/neuron-extensions"

func init() {
	rootCmd.AddCommand(modulesCmd)

	modulesCmd.PersistentFlags().StringP("extensions-path", "e", "", "Sets the neuron-extensions module os path.")
	modulesCmd.PersistentFlags().StringP("neuron-path", "n", "", "Sets the path to the github.com/neuronlabs/neuron module")
	modulesCmd.PersistentFlags().BoolP("all", "a", false, "Affects all submodules")
}