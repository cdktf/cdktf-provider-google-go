package computesubnetwork


type ComputeSubnetworkSecondaryIpRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.63.1/docs/resources/compute_subnetwork#ip_cidr_range ComputeSubnetwork#ip_cidr_range}.
	IpCidrRange *string `field:"optional" json:"ipCidrRange" yaml:"ipCidrRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.63.1/docs/resources/compute_subnetwork#range_name ComputeSubnetwork#range_name}.
	RangeName *string `field:"optional" json:"rangeName" yaml:"rangeName"`
}

