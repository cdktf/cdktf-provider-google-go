package identityplatformtenantinboundsamlconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IdentityPlatformTenantInboundSamlConfigConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// Human friendly display name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/identity_platform_tenant_inbound_saml_config#display_name IdentityPlatformTenantInboundSamlConfig#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// idp_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/identity_platform_tenant_inbound_saml_config#idp_config IdentityPlatformTenantInboundSamlConfig#idp_config}
	IdpConfig *IdentityPlatformTenantInboundSamlConfigIdpConfig `field:"required" json:"idpConfig" yaml:"idpConfig"`
	// The name of the InboundSamlConfig resource.
	//
	// Must start with 'saml.' and can only have alphanumeric characters,
	// hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
	// alphanumeric character, and have at least 2 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/identity_platform_tenant_inbound_saml_config#name IdentityPlatformTenantInboundSamlConfig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// sp_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/identity_platform_tenant_inbound_saml_config#sp_config IdentityPlatformTenantInboundSamlConfig#sp_config}
	SpConfig *IdentityPlatformTenantInboundSamlConfigSpConfig `field:"required" json:"spConfig" yaml:"spConfig"`
	// The name of the tenant where this inbound SAML config resource exists.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/identity_platform_tenant_inbound_saml_config#tenant IdentityPlatformTenantInboundSamlConfig#tenant}
	Tenant *string `field:"required" json:"tenant" yaml:"tenant"`
	// If this config allows users to sign in with the provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/identity_platform_tenant_inbound_saml_config#enabled IdentityPlatformTenantInboundSamlConfig#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/identity_platform_tenant_inbound_saml_config#id IdentityPlatformTenantInboundSamlConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/identity_platform_tenant_inbound_saml_config#project IdentityPlatformTenantInboundSamlConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/identity_platform_tenant_inbound_saml_config#timeouts IdentityPlatformTenantInboundSamlConfig#timeouts}
	Timeouts *IdentityPlatformTenantInboundSamlConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

