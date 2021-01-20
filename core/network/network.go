package network

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

const (
	cmdName = "network"
	cmdDesc = "Perform network operations: start|destroy|create|delete."
)

type Network struct {
	Name    string
	ProcCmd []*exec.Cmd
	Context *context.Context
}

func (n *Network) addCmdProc(cmd *exec.Cmd) {
	n.ProcCmd = append(n.ProcCmd, cmd)
}

var networkCmd = &cobra.Command{
	Use:   cmdName,
	Short: fmt.Sprint(cmdDesc),
	Long:  fmt.Sprint(cmdDesc),
}

//----------------------------------------------------------------------------------
// Cmd() returns the cobra command for Network
//----------------------------------------------------------------------------------
func Cmd() *cobra.Command {
	networkCmd.AddCommand(startCmd())
	networkCmd.AddCommand(createCmd())
	//networkCmd.AddCommand(destroyCmd())
	return networkCmd
}
