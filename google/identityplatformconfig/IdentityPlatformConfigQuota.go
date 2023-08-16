package identityplatformconfig


type IdentityPlatformConfigQuota struct {
	// sign_up_quota_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/identity_platform_config#sign_up_quota_config IdentityPlatformConfig#sign_up_quota_config}
	SignUpQuotaConfig *IdentityPlatformConfigQuotaSignUpQuotaConfig `field:"optional" json:"signUpQuotaConfig" yaml:"signUpQuotaConfig"`
}

