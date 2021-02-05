package organization

import (
	"fmt"
	"github.com/orpheus/hyperspace/core/network"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

const (
	cmdName = "org"
	cmdDesc = "Organization CRUD abilities."
)

var cmd = &cobra.Command{
	Use:   cmdName,
	Short: fmt.Sprint(cmdDesc),
	Long:  fmt.Sprint(cmdDesc),
}

//----------------------------------------------------------------------------------
// Cmd() returns the cobra command for Network
//----------------------------------------------------------------------------------
func Cmd() *cobra.Command {
	cmd.AddCommand(createCmd())
	return cmd
}

func createCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Creates organizations.",
		Long:  `Create fabric organizations with generated crypto.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return createOrganization()
		},
	}
}

const defYaml = `
tls: true
nodes:
	- type: peer
	
	  
`

func createOrganization () error {
	// read from config
		// if no config was passed, read from current working directory
		// use a viper to read in config?
		// no, read in straight from struct into a config struct

	// toDo: need a binary collector.

	// for now, just use the global binary and
	// hardcode a yaml.
		// scratch: hardcode everything inside this fn first then delegate

	// CA
		// what should this define?
		// what's the minimal I'll allow,
		// what's the maximal

		// toDo: need cert-storage & adapters

	generateTlsCert()

	// Crypto
		// read in configs paths
		// generate crypto using command script and binary

	return nil
}

//----------------------------------------------------------------------------------
// generateTlsCert()
//----------------------------------------------------------------------------------
// Starts a CA server with tls enabled to generate a tls-cert.pem
//----------------------------------------------------------------------------------
func generateTlsCert () {
	// in the path I call this binary, let's create the tls
	// we are starting and killing the server just to get it
	// to create the tls.perm. It won't create this even with
	// the --tls.enabled flag when using the init command
	cmd := exec.Command("fabric-ca-server",
		"start",
		"-b",
		"admin:adminp",
		"--tls.enabled",
	)
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	log.Println("Starting fabric-ca-server with tls.")
	err := cmd.Start()
	if err != nil {
		log.Fatalf("Failed to run cmd: %v", err)
	}
	log.Printf("Waiting for server to generate tls-cert.pem.")
	// is there a better way to write this?
	for _, err = os.Stat("tls-cert.pem"); os.IsNotExist(err); {
		_, err = os.Stat("tls-cert.pem")
	}
	log.Println("Generated tls-cert.pem. Killing server...")
	// kill server after we generated our cert
	network.KillProcess(cmd, "fabric-ca-tls-server")
}