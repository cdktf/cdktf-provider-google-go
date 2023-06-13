package identityplatformconfig


type IdentityPlatformConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/identity_platform_config#create IdentityPlatformConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/identity_platform_config#delete IdentityPlatformConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/identity_platform_config#update IdentityPlatformConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

