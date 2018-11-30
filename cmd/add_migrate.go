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
	"gen/cmd/file"
	"github.com/spf13/cobra"
)

// migrateCmd represents the api command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "add migrate",
	Long: `add migrate`,
	Run: func(cmd *cobra.Command, args []string) {
		migrateFile := file.NewMigrateFile(modelFlags.name)
		migrateFile.Write()
		cmd.Println("migrate file write success")
	},
}

func init() {
	addCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.
	//migrateCmd.PersistentFlags().StringP("model", "m","", "the api use model")
	migrateCmd.Flags().StringVarP(&modelFlags.name, "model", "m", "user", "the api use which model")
	migrateCmd.Flags().StringSliceVarP(&modelFlags.fields, "fields", "f", []string{"name:string","email:string"}, "fields is the model's field slice")
	migrateCmd.Flags().BoolVarP(&modelFlags.needBase, "base", "b", true, "the model needs base ID/CreatedAt/UpdatedAt?")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
