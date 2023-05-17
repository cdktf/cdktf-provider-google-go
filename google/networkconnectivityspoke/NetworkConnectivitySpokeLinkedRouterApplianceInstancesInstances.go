package networkconnectivityspoke


type NetworkConnectivitySpokeLinkedRouterApplianceInstancesInstances struct {
	// The IP address on the VM to use for peering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/network_connectivity_spoke#ip_address NetworkConnectivitySpoke#ip_address}
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// The URI of the virtual machine resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/network_connectivity_spoke#virtual_machine NetworkConnectivitySpoke#virtual_machine}
	VirtualMachine *string `field:"optional" json:"virtualMachine" yaml:"virtualMachine"`
}

