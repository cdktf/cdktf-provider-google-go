package identityplatformprojectdefaultconfig


type IdentityPlatformProjectDefaultConfigSignInAnonymous struct {
	// Whether anonymous user auth is enabled for the project or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/identity_platform_project_default_config#enabled IdentityPlatformProjectDefaultConfig#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

