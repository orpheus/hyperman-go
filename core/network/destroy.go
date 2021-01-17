package network

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

func destroyCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "destroy",
		Short: "Destroy a network.",
		Long:  `Kill all running processes in a fabric network.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				return fmt.Errorf("trailing args detected")
			}
			fmt.Println("Destroy network")
			return nil
		},
	}
}

//func KillProcesses (commands []*exec.Cmd) {
//	for _, cmd := range commands {
//		KillProcess(cmd, "")
//	}
//}


func KillProcess (cmd *exec.Cmd, name string) {
	if err := cmd.Process.Kill(); err != nil {
		log.Fatalf("Failed to kill process: %s\n%v", name, err)
	}
	log.Printf("Succesfully killed process: %s", name)
}
