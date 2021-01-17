package network

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

const (
	netFuncName = "network"
	netCmdDes   = "Perform network operations: start|destroy|create|delete."
)

type Network struct {
	Name string
	ProcCmd []*exec.Cmd
	Context *context.Context
}

func (n *Network) addCmdProc (cmd *exec.Cmd) {
	n.ProcCmd = append(n.ProcCmd, cmd)
}

var networkCmd = &cobra.Command{
	Use:   netFuncName,
	Short: fmt.Sprint(netCmdDes),
	Long:  fmt.Sprint(netCmdDes),
}

// Cmd returns the cobra command for Network
func Cmd() *cobra.Command {
	networkCmd.AddCommand(startCmd())
	networkCmd.AddCommand(createCmd())
	//networkCmd.AddCommand(destroyCmd())
	return networkCmd
}



