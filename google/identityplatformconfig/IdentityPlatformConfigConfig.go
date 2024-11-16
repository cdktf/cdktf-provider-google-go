// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityplatformconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IdentityPlatformConfigConfig struct {
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
	// List of domains authorized for OAuth redirects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/identity_platform_config#authorized_domains IdentityPlatformConfig#authorized_domains}
	AuthorizedDomains *[]*string `field:"optional" json:"authorizedDomains" yaml:"authorizedDomains"`
	// Whether anonymous users will be auto-deleted after a period of 30 days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/identity_platform_config#autodelete_anonymous_users IdentityPlatformConfig#autodelete_anonymous_users}
	AutodeleteAnonymousUsers interface{} `field:"optional" json:"autodeleteAnonymousUsers" yaml:"autodeleteAnonymousUsers"`
	// blocking_functions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/identity_platform_config#blocking_functions IdentityPlatformConfig#blocking_functions}
	BlockingFunctions *IdentityPlatformConfigBlockingFunctions `field:"optional" json:"blockingFunctions" yaml:"blockingFunctions"`
	// client block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/identity_platform_config#client IdentityPlatformConfig#client}
	Client *IdentityPlatformConfigClient `field:"optional" json:"client" yaml:"client"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/identity_platform_config#id IdentityPlatformConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// mfa block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/identity_platform_config#mfa IdentityPlatformConfig#mfa}
	Mfa *IdentityPlatformConfigMfa `field:"optional" json:"mfa" yaml:"mfa"`
	// monitoring block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/identity_platform_config#monitoring IdentityPlatformConfig#monitoring}
	Monitoring *IdentityPlatformConfigMonitoring `field:"optional" json:"monitoring" yaml:"monitoring"`
	// multi_tenant block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/identity_platform_config#multi_tenant IdentityPlatformConfig#multi_tenant}
	MultiTenant *IdentityPlatformConfigMultiTenant `field:"optional" json:"multiTenant" yaml:"multiTenant"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/identity_platform_config#project IdentityPlatformConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// quota block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/identity_platform_config#quota IdentityPlatformConfig#quota}
	Quota *IdentityPlatformConfigQuota `field:"optional" json:"quota" yaml:"quota"`
	// sign_in block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/identity_platform_config#sign_in IdentityPlatformConfig#sign_in}
	SignIn *IdentityPlatformConfigSignIn `field:"optional" json:"signIn" yaml:"signIn"`
	// sms_region_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/identity_platform_config#sms_region_config IdentityPlatformConfig#sms_region_config}
	SmsRegionConfig *IdentityPlatformConfigSmsRegionConfig `field:"optional" json:"smsRegionConfig" yaml:"smsRegionConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/identity_platform_config#timeouts IdentityPlatformConfig#timeouts}
	Timeouts *IdentityPlatformConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

