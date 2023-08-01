package networksecurityaddressgroup


type NetworkSecurityAddressGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_security_address_group#create NetworkSecurityAddressGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_security_address_group#delete NetworkSecurityAddressGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_security_address_group#update NetworkSecurityAddressGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

