package identityplatformtenant

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IdentityPlatformTenantConfig struct {
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
	// Human friendly display name of the tenant.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/identity_platform_tenant#display_name IdentityPlatformTenant#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Whether to allow email/password user authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/identity_platform_tenant#allow_password_signup IdentityPlatformTenant#allow_password_signup}
	AllowPasswordSignup interface{} `field:"optional" json:"allowPasswordSignup" yaml:"allowPasswordSignup"`
	// Whether authentication is disabled for the tenant.
	//
	// If true, the users under
	// the disabled tenant are not allowed to sign-in. Admins of the disabled tenant
	// are not able to manage its users.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/identity_platform_tenant#disable_auth IdentityPlatformTenant#disable_auth}
	DisableAuth interface{} `field:"optional" json:"disableAuth" yaml:"disableAuth"`
	// Whether to enable email link user authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/identity_platform_tenant#enable_email_link_signin IdentityPlatformTenant#enable_email_link_signin}
	EnableEmailLinkSignin interface{} `field:"optional" json:"enableEmailLinkSignin" yaml:"enableEmailLinkSignin"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/identity_platform_tenant#id IdentityPlatformTenant#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/identity_platform_tenant#project IdentityPlatformTenant#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/identity_platform_tenant#timeouts IdentityPlatformTenant#timeouts}
	Timeouts *IdentityPlatformTenantTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

