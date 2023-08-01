package containerazurenodepool


type ContainerAzureNodePoolConfigSshConfig struct {
	// The SSH public key data for VMs managed by Anthos.
	//
	// This accepts the authorized_keys file format used in OpenSSH according to the sshd(8) manual page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_azure_node_pool#authorized_key ContainerAzureNodePool#authorized_key}
	AuthorizedKey *string `field:"required" json:"authorizedKey" yaml:"authorizedKey"`
}

