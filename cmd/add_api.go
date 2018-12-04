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
	"bytes"
	"io"
	"os"
	"os/exec"

	"github.com/cyjme/gen/cmd/file"
	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "add api, include: controller model migrate router",
	Long:  `add api, include: controller model migrate router`,
	Run: func(cmd *cobra.Command, args []string) {
		modelFile := file.NewModelFile(modelFlags.name, modelFlags.fields, modelFlags.needBase)
		modelFile.Write()

		migrateFile := file.NewMigrateFile(modelFlags.name)
		migrateFile.Write()

		controllerFile := file.NewControllerFile(modelFlags.name)
		controllerFile.Write()

		routerFile := file.NewRouterFile(modelFlags.name)
		routerFile.Write()

		command := exec.Command("swag", "init")
		var stdBuffer bytes.Buffer

		mw := io.MultiWriter(os.Stdout, &stdBuffer)
		command.Stdout = mw
		command.Stderr = mw

		if err := command.Run(); err != nil {
			cmd.Printf("Error while running swag init: %s", err)
		}
	},
}

type ModelFlags struct {
	name     string
	fields   []string
	needBase bool
}

var modelFlags ModelFlags

func init() {
	addCmd.AddCommand(apiCmd)

	// Here you will define your flags and configuration settings.
	//apiCmd.PersistentFlags().StringP("model", "m","", "the api use model")
	apiCmd.Flags().StringVarP(&modelFlags.name, "model", "m", "user", "the api use which model")
	apiCmd.Flags().StringSliceVarP(&modelFlags.fields, "fields", "f", []string{"name:string", "email:string"}, "fields is the model's field slice")
	apiCmd.Flags().BoolVarP(&modelFlags.needBase, "base", "b", true, "the model needs base ID/CreatedAt/UpdatedAt?")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
