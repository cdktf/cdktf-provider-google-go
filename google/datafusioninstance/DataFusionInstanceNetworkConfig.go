package datafusioninstance


type DataFusionInstanceNetworkConfig struct {
	// The IP range in CIDR notation to use for the managed Data Fusion instance nodes.
	//
	// This range must not overlap with any other ranges used in the Data Fusion instance network.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_fusion_instance#ip_allocation DataFusionInstance#ip_allocation}
	IpAllocation *string `field:"required" json:"ipAllocation" yaml:"ipAllocation"`
	// Name of the network in the project with which the tenant project will be peered for executing pipelines.
	//
	// In case of shared VPC where the network resides in another host
	// project the network should specified in the form of projects/{host-project-id}/global/networks/{network}
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_fusion_instance#network DataFusionInstance#network}
	Network *string `field:"required" json:"network" yaml:"network"`
}

