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
	"fmt"
	"github.com/AlecAivazis/survey"
	"github.com/spf13/cobra"
)

var qs = []*survey.Question{
	{
		Name:     "name",
		Prompt:   &survey.Input{Message: "What is your name?"},
		Validate: survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "color",
		Prompt: &survey.Select{
			Message: "Choose a color:",
			Options: []string{"red", "blue", "green"},
			Default: "red",
		},
	},
	{
		Name: "age",
		Prompt:   &survey.Input{Message: "How old are you?"},
	},
}

// questionCmd represents the question command
var questionCmd = &cobra.Command{
	Use:   "question",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		answers := struct {
			Name          string                  // survey will match the question and field names
			FavoriteColor string `survey:"color"` // or you can tag fields to match a specific name
			Age           int                     // if the types don't match exactly, survey will try to convert for you
		}{}

		// perform the questions
		err := survey.Ask(qs, &answers)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Printf("%s chose %s.", answers.Name, answers.FavoriteColor)
	},
}

func init() {
	rootCmd.AddCommand(questionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// questionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// questionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
