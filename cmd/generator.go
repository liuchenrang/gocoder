// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"coder/bootstrap"
	"coder/config"
	"coder/generator"
	"coder/providers"
	"github.com/spf13/cobra"
	"os"
)

// generatorCmd represents the generator command
var generatorCmd = &cobra.Command{
	Use:   "generator",
	Short: "A brief description of your command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		
		config.Project = config.NewServerConfig(cfgFile)
		bootstrap.GetApp().Register(providers.NewDatabase())
		bootstrap.Bootstrap()
		
	
		creatorType, _ := cmd.PersistentFlags().GetString("type")
		creatorName, _ := cmd.PersistentFlags().GetString("name")
		tplPath, _ := cmd.PersistentFlags().GetString("t")
		output, _ := cmd.PersistentFlags().GetString("o")
		preView, _ := cmd.PersistentFlags().GetBool("p")
		
		newGenerator := generator.NewGenerator(tplPath)
		instance := newGenerator.GetInstance(creatorType)
		instance.SetPreView(preView)
		instance.SetCreateName(creatorName)
		instance.SetOutputPath(output)
		instance.Render(instance.(generator.IData))
	},
}


func init() {
	rootCmd.AddCommand(generatorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	dir, _ := os.Getwd()
	generatorCmd.PersistentFlags().String("type", "", "type for model service or controller")
	generatorCmd.PersistentFlags().String("name", "", "name of model service or controller")
	generatorCmd.PersistentFlags().String("o", "", "output path of model service or controller")
	
	generatorCmd.PersistentFlags().String("t", dir + "/generator/templates", "template path of model service or controller")
	generatorCmd.PersistentFlags().String("m", dir + "/models", " model path")
	generatorCmd.PersistentFlags().Bool("p",false, "view result")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generatorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
