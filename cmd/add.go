// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"gen/cmd/vars"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add (api/controller/model/router/migrate)",
	Long:  `add (api/controller/model/router/migrate)`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("At least one arg")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	pwd, err := os.Getwd()
	if err != nil {
		log.Println("get pwd error", err)
	}
	pwdSlice := strings.Split(pwd, "/")
	vars.ProjectName = pwdSlice[len(pwdSlice)-1]

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
