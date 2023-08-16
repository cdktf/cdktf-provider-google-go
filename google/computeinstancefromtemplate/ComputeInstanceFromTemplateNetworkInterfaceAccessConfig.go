package computeinstancefromtemplate


type ComputeInstanceFromTemplateNetworkInterfaceAccessConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#nat_ip ComputeInstanceFromTemplate#nat_ip}.
	NatIp *string `field:"optional" json:"natIp" yaml:"natIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#network_tier ComputeInstanceFromTemplate#network_tier}.
	NetworkTier *string `field:"optional" json:"networkTier" yaml:"networkTier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#public_ptr_domain_name ComputeInstanceFromTemplate#public_ptr_domain_name}.
	PublicPtrDomainName *string `field:"optional" json:"publicPtrDomainName" yaml:"publicPtrDomainName"`
}

