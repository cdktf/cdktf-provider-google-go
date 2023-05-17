package identityplatformprojectdefaultconfig


type IdentityPlatformProjectDefaultConfigSignInPhoneNumber struct {
	// Whether phone number auth is enabled for the project or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/identity_platform_project_default_config#enabled IdentityPlatformProjectDefaultConfig#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A map of <test phone number, fake code> that can be used for phone auth testing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/identity_platform_project_default_config#test_phone_numbers IdentityPlatformProjectDefaultConfig#test_phone_numbers}
	TestPhoneNumbers *map[string]*string `field:"optional" json:"testPhoneNumbers" yaml:"testPhoneNumbers"`
}

