package computeinstancefromtemplate


type ComputeInstanceFromTemplateNetworkInterfaceAliasIpRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#ip_cidr_range ComputeInstanceFromTemplate#ip_cidr_range}.
	IpCidrRange *string `field:"optional" json:"ipCidrRange" yaml:"ipCidrRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#subnetwork_range_name ComputeInstanceFromTemplate#subnetwork_range_name}.
	SubnetworkRangeName *string `field:"optional" json:"subnetworkRangeName" yaml:"subnetworkRangeName"`
}

