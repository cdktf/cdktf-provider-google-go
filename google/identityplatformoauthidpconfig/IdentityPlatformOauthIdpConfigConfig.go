package identityplatformoauthidpconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IdentityPlatformOauthIdpConfigConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The client id of an OAuth client.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/identity_platform_oauth_idp_config#client_id IdentityPlatformOauthIdpConfig#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// For OIDC Idps, the issuer identifier.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/identity_platform_oauth_idp_config#issuer IdentityPlatformOauthIdpConfig#issuer}
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// The name of the OauthIdpConfig. Must start with 'oidc.'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/identity_platform_oauth_idp_config#name IdentityPlatformOauthIdpConfig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The client secret of the OAuth client, to enable OIDC code flow.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/identity_platform_oauth_idp_config#client_secret IdentityPlatformOauthIdpConfig#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Human friendly display name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/identity_platform_oauth_idp_config#display_name IdentityPlatformOauthIdpConfig#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// If this config allows users to sign in with the provider.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/identity_platform_oauth_idp_config#enabled IdentityPlatformOauthIdpConfig#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/identity_platform_oauth_idp_config#id IdentityPlatformOauthIdpConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/identity_platform_oauth_idp_config#project IdentityPlatformOauthIdpConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/identity_platform_oauth_idp_config#timeouts IdentityPlatformOauthIdpConfig#timeouts}
	Timeouts *IdentityPlatformOauthIdpConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

