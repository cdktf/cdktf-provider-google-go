package computeinstancetemplate


type ComputeInstanceTemplateNetworkInterfaceAliasIpRange struct {
	// The IP CIDR range represented by this alias IP range.
	//
	// This IP CIDR range must belong to the specified subnetwork and cannot contain IP addresses reserved by system or used by other network interfaces. At the time of writing only a netmask (e.g. /24) may be supplied, with a CIDR format resulting in an API error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_instance_template#ip_cidr_range ComputeInstanceTemplate#ip_cidr_range}
	IpCidrRange *string `field:"required" json:"ipCidrRange" yaml:"ipCidrRange"`
	// The subnetwork secondary range name specifying the secondary range from which to allocate the IP CIDR range for this alias IP range.
	//
	// If left unspecified, the primary range of the subnetwork will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_instance_template#subnetwork_range_name ComputeInstanceTemplate#subnetwork_range_name}
	SubnetworkRangeName *string `field:"optional" json:"subnetworkRangeName" yaml:"subnetworkRangeName"`
}

