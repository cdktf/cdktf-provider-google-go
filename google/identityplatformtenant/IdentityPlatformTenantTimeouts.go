package identityplatformtenant


type IdentityPlatformTenantTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/identity_platform_tenant#create IdentityPlatformTenant#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/identity_platform_tenant#delete IdentityPlatformTenant#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/identity_platform_tenant#update IdentityPlatformTenant#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

