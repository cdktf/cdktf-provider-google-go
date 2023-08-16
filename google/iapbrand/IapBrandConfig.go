package iapbrand

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IapBrandConfig struct {
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
	// Application name displayed on OAuth consent screen.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iap_brand#application_title IapBrand#application_title}
	ApplicationTitle *string `field:"required" json:"applicationTitle" yaml:"applicationTitle"`
	// Support email displayed on the OAuth consent screen.
	//
	// Can be either a
	// user or group email. When a user email is specified, the caller must
	// be the user with the associated email address. When a group email is
	// specified, the caller can be either a user or a service account which
	// is an owner of the specified group in Cloud Identity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iap_brand#support_email IapBrand#support_email}
	SupportEmail *string `field:"required" json:"supportEmail" yaml:"supportEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iap_brand#id IapBrand#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iap_brand#project IapBrand#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iap_brand#timeouts IapBrand#timeouts}
	Timeouts *IapBrandTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

