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

type outBytes struct {
	bytes []byte
}

func (so *outBytes) Write(p []byte) (n int, err error) {
	so.bytes = append(so.bytes, p...)
	return os.Stdout.Write(p)
}

// exec2Cmd represents the exec2 command
var exec2Cmd = &cobra.Command{
	Use:   "exec2",
	Short: "Stream data between exec commands",
	Long: `Stream data between exec commands via stdin/out`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			e := fmt.Errorf("getting current directory: %v", err)
			fmt.Print(e)
			return
		}
		
		var ob outBytes
		ts1Path := fmt.Sprintf("%s/ts1/dist", dir)
		cmd1 := exec.Command("node", ts1Path)
		cmd1.Stdin = os.Stdin
		cmd1.Stdout = &ob
		cmd1.Stderr = os.Stderr
		if err = cmd1.Run(); err != nil {
			fmt.Printf("running cmd1: %v\n", err)
			return
		}
		
		ts2Path := fmt.Sprintf("%s/ts2/dist", dir)
		cmd2 := exec.Command("node", ts2Path)
		cmd2in, err := cmd2.StdinPipe()
		if err != nil {
			fmt.Printf("creating StdinPipe for cmd2: %v\n", err)
			return
		}
		cmd2.Stdout = os.Stdout
		cmd2.Stderr = os.Stderr

		go func() {
			defer cmd2in.Close()
			if _, err := cmd2in.Write(ob.bytes); err != nil {
				fmt.Printf("writing to cmd2 Stdin: %v\n", err)
				return
			}
		}()

		if err = cmd2.Run(); err != nil {
			fmt.Printf("running cmd2: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(exec2Cmd)
}
