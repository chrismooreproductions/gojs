/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

type Something struct {
	something *Something
}

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			e := fmt.Errorf("getting current directory: %v", err)
			fmt.Print(e)
			return
		}
		fmt.Println(dir)
		js1Path := fmt.Sprintf("%s/ts/dist", dir)
		fmt.Println(js1Path)
		scr1, err := exec.Command("node", js1Path).Output()
		if err != nil {
			e := fmt.Errorf("executing ts: %v", err)
			fmt.Println(e)
			return
		}
		js2Path := fmt.Sprintf("%s/ts2/dist", dir)
		cmd2 := exec.Command("node", js2Path)
		cmd2.Env = []string{fmt.Sprintf("MESSAGE=%s", string(scr1))}
		o, err := cmd2.Output()
		if err != nil {
			e := fmt.Errorf("executing ts2: %v", err)
			fmt.Println(e)
			return
		}
		fmt.Print(string(o))
		if err != nil {
			fmt.Print(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(execCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// execCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// execCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
