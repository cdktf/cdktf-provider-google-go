package computeinstance


type ComputeInstanceNetworkInterfaceIpv6AccessConfig struct {
	// The service-level to be provided for IPv6 traffic when the subnet has an external subnet.
	//
	// Only PREMIUM tier is valid for IPv6
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance#network_tier ComputeInstance#network_tier}
	NetworkTier *string `field:"required" json:"networkTier" yaml:"networkTier"`
	// The domain name to be used when creating DNSv6 records for the external IPv6 ranges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance#public_ptr_domain_name ComputeInstance#public_ptr_domain_name}
	PublicPtrDomainName *string `field:"optional" json:"publicPtrDomainName" yaml:"publicPtrDomainName"`
}

