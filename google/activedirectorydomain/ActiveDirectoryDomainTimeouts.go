package activedirectorydomain


type ActiveDirectoryDomainTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/active_directory_domain#create ActiveDirectoryDomain#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/active_directory_domain#delete ActiveDirectoryDomain#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/active_directory_domain#update ActiveDirectoryDomain#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

