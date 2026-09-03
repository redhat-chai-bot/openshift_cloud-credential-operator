package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/openshift/cloud-credential-operator/pkg/cmd/provisioning/aws"
	"github.com/openshift/cloud-credential-operator/pkg/cmd/provisioning/azure"
	"github.com/openshift/cloud-credential-operator/pkg/cmd/provisioning/gcp"
	"github.com/openshift/cloud-credential-operator/pkg/cmd/provisioning/ibmcloud"
	"github.com/openshift/cloud-credential-operator/pkg/cmd/provisioning/nutanix"
	"github.com/openshift/cloud-credential-operator/pkg/version"
)

func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "ccoctl",
		Short: "OpenShift credentials provisioning tool",
	}

	if v := version.Get().String(); len(v) == 0 {
		rootCmd.Version = "<unknown>"
	} else {
		rootCmd.Version = v
	}

	rootCmd.AddCommand(aws.NewAWSCmd())
	rootCmd.AddCommand(gcp.NewGCPCmd())
	rootCmd.AddCommand(ibmcloud.NewIBMCloudCmd())
	rootCmd.AddCommand(nutanix.NewNutanixCmd())
	rootCmd.AddCommand(azure.NewAzureCmd())

	return rootCmd
}

func main() {
	if err := newRootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}
