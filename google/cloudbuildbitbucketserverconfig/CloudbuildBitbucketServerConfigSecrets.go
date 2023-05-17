package cloudbuildbitbucketserverconfig


type CloudbuildBitbucketServerConfigSecrets struct {
	// The resource name for the admin access token's secret version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloudbuild_bitbucket_server_config#admin_access_token_version_name CloudbuildBitbucketServerConfig#admin_access_token_version_name}
	AdminAccessTokenVersionName *string `field:"required" json:"adminAccessTokenVersionName" yaml:"adminAccessTokenVersionName"`
	// The resource name for the read access token's secret version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloudbuild_bitbucket_server_config#read_access_token_version_name CloudbuildBitbucketServerConfig#read_access_token_version_name}
	ReadAccessTokenVersionName *string `field:"required" json:"readAccessTokenVersionName" yaml:"readAccessTokenVersionName"`
	// Immutable.
	//
	// The resource name for the webhook secret's secret version. Once this field has been set, it cannot be changed.
	// Changing this field will result in deleting/ recreating the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloudbuild_bitbucket_server_config#webhook_secret_version_name CloudbuildBitbucketServerConfig#webhook_secret_version_name}
	WebhookSecretVersionName *string `field:"required" json:"webhookSecretVersionName" yaml:"webhookSecretVersionName"`
}

