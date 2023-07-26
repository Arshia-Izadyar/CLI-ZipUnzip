package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	util "zipU/utils"

	"github.com/spf13/cobra"
)

var File string

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:     "code",
	Short:   "use code to open inziped file with VsCode",
	Example: "zipU code {Filename}",
	Long:    ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if File == "" && len(args) < 1 {
			return errors.New("Add one args")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var fileName string

		if File != "" {
			fileName = File
		} else {
			fileName = args[0]
		}

		fmt.Println(fileName)
		fileExist, err := util.FileExists(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "File Wrong %+v", err)
			return
		}
		if fileExist {
			fileName, err = filepath.Abs(fileName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error Happened %+v", err)
			}
			fmt.Println(fileName)

		} else {
			fmt.Fprintf(os.Stderr, "File does not exist %+v", err)
			return
		}

		pwd, err := os.Getwd()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error Happened on pwd %+v", err)
		}

		err = util.Unzip(fileName, pwd)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error Happened on unzip %+v", err)
		}
		os.Chdir(util.FilenameWithoutExtension(fileName))

		pwd, err = os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error Happened on pwd %+v", err)
		}

		command := exec.Command("code", pwd)
		err = command.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error Happened on exec the command %+v", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(codeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// codeCmd.PersistentFlags().StringVarP("foo", "", "A help for foo")
	codeCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "A File name to unzip and open in IDE")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// codeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
