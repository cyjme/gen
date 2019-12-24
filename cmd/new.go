package cmd

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"gen/util"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "create a new api project",
	Long:  `create a new api project`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Println("At least one arg, Project name must be provided")
			return
		}
		projectName := args[0]

		command := exec.Command("git", "clone", "--depth", "1", "https://gen-template.git", projectName)
		var stdBuffer bytes.Buffer

		mw := io.MultiWriter(os.Stdout, &stdBuffer)
		command.Stdout = mw
		command.Stderr = mw

		if err := command.Run(); err != nil {
			cmd.Printf("Error while running git clone: %s", err)
		}

		command = exec.Command("rm", "-rf", "./"+projectName+"/.git")
		if err := command.Run(); err != nil {
			cmd.Printf("Error while running git clone: %s", err)
		}

		files, err := GetAllFiles("./" + projectName)

		for _, file := range files {
			replaceProjectName(file, projectName)

			util.FormatSourceCode(file)
		}
		replaceProjectName("./"+projectName+"/go.mod", projectName)

		if err != nil {
			cmd.Printf("Error while set project name", err)
		}

		cmd.Println("\033[1;31m new project ok")
	},
}

func replaceProjectName(filePath string, projectName string) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("Error while read project file", err)
	}
	newFile := strings.Replace(string(file), "{{projectName}}", projectName, -1)

	f, err := os.Create(filePath)
	if err != nil {
		log.Println("Error while read project file", err)
	}
	f.Write([]byte(newFile))
}

func GetAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)

	for _, fi := range dir {
		if fi.IsDir() {
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetAllFiles(dirPth + PthSep + fi.Name())
		} else {
			ok := strings.HasSuffix(fi.Name(), ".go")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	// read child file
	for _, table := range dirs {
		temp, _ := GetAllFiles(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
