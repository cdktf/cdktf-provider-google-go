package identityplatformprojectdefaultconfig


type IdentityPlatformProjectDefaultConfigSignIn struct {
	// Whether to allow more than one account to have the same email.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/identity_platform_project_default_config#allow_duplicate_emails IdentityPlatformProjectDefaultConfig#allow_duplicate_emails}
	AllowDuplicateEmails interface{} `field:"optional" json:"allowDuplicateEmails" yaml:"allowDuplicateEmails"`
	// anonymous block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/identity_platform_project_default_config#anonymous IdentityPlatformProjectDefaultConfig#anonymous}
	Anonymous *IdentityPlatformProjectDefaultConfigSignInAnonymous `field:"optional" json:"anonymous" yaml:"anonymous"`
	// email block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/identity_platform_project_default_config#email IdentityPlatformProjectDefaultConfig#email}
	Email *IdentityPlatformProjectDefaultConfigSignInEmail `field:"optional" json:"email" yaml:"email"`
	// phone_number block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/identity_platform_project_default_config#phone_number IdentityPlatformProjectDefaultConfig#phone_number}
	PhoneNumber *IdentityPlatformProjectDefaultConfigSignInPhoneNumber `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
}

