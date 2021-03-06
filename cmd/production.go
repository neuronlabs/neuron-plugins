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
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/neuronlabs/neuron-extensions/internal/modules"
)

// productionModeCmd represents the productionMode command
var productionModeCmd = &cobra.Command{
	Use:   "production",
	Short: "Sets back the develop-mode submodules go.mod files for production ready state",
	Long: `This function sets back the go.mod files from develop-mode for all the submodules provided in the arguments. 
It should be executed before new commits, so that the modules wouldn't have relative path replacements.' 
By default neuron-extensions directory is assumed to be the current working directory. It could be set using 'neuron-extensions' flag.
The function could set the production-mode for all submodules (using -all flag) or 
for the selected submodule by providing it as the argument.`,
	Run: runProductionMode,
}

func init() {
	modulesCmd.AddCommand(productionModeCmd)
}

// neuron-extensions modules develop-mode . codec/json
func runProductionMode(cmd *cobra.Command, args []string) {
	path, err := getExtensionsPath()
	if err != nil {
		os.Exit(1)
	}

	neuronPath, err := modulesCmd.PersistentFlags().GetString("neuron-path")
	if err != nil {
		fmt.Printf("Err: %v", err)
		os.Exit(1)
	}
	if neuronPath == "" {
		neuronPath = filepath.Clean(path + "/../neuron")
	}

	var subModules []*modules.SubModule
	all, err := modulesCmd.PersistentFlags().GetBool("all")
	if err != nil {
		fmt.Printf("Err: %v\n", err)
		os.Exit(1)
	}

	if all {
		subModules, err = modules.ListSubmodules(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		if len(args) == 0 {
			fmt.Printf("No modules provided to set on develop-mode. Provide submodules as an additional arguments or set the 'all' flag.")
			os.Exit(1)
		}
		for _, arg := range args {
			modulePath := arg
			moduleName := arg
			if strings.Contains(arg, extensionsPackage) {
				modulePath = strings.TrimPrefix(modulePath, extensionsPackage)
			} else {
				moduleName = extensionsPackage + "/" + moduleName
			}
			if modulePath[0] != '/' {
				modulePath = "/" + modulePath
			}
			modulePath = path + modulePath
			subModules = []*modules.SubModule{{Path: modulePath, ModuleName: moduleName}}
		}
	}

	for _, subModule := range subModules {
		if subModule.ModuleName == extensionsPackage {
			continue
		}
		if err = modules.SetProductionMode(subModule); err != nil {
			if err == modules.ErrModuleAlreadyDevelopment {
				continue
			}
			fmt.Printf("Replacing development module failed: %v\n", err)
			os.Exit(1)
		}
	}
}
