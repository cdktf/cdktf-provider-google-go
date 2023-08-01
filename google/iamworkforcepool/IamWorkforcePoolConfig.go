package iamworkforcepool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamWorkforcePoolConfig struct {
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
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/iam_workforce_pool#location IamWorkforcePool#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Immutable. The resource name of the parent. Format: 'organizations/{org-id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/iam_workforce_pool#parent IamWorkforcePool#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// The name of the pool.
	//
	// The ID must be a globally unique string of 6 to 63 lowercase letters,
	// digits, or hyphens. It must start with a letter, and cannot have a trailing hyphen.
	// The prefix 'gcp-' is reserved for use by Google, and may not be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/iam_workforce_pool#workforce_pool_id IamWorkforcePool#workforce_pool_id}
	WorkforcePoolId *string `field:"required" json:"workforcePoolId" yaml:"workforcePoolId"`
	// A user-specified description of the pool. Cannot exceed 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/iam_workforce_pool#description IamWorkforcePool#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the pool is disabled.
	//
	// You cannot use a disabled pool to exchange tokens,
	// or use existing tokens to access resources. If the pool is re-enabled, existing tokens grant access again.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/iam_workforce_pool#disabled IamWorkforcePool#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// A user-specified display name of the pool in Google Cloud Console. Cannot exceed 32 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/iam_workforce_pool#display_name IamWorkforcePool#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/iam_workforce_pool#id IamWorkforcePool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Duration that the Google Cloud access tokens, console sign-in sessions, and 'gcloud' sign-in sessions from this pool are valid.
	//
	// Must be greater than 15 minutes (900s) and less than 12 hours (43200s).
	// If 'sessionDuration' is not configured, minted credentials have a default duration of one hour (3600s).
	// A duration in seconds with up to nine fractional digits, ending with ''s''. Example: "'3.5s'".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/iam_workforce_pool#session_duration IamWorkforcePool#session_duration}
	SessionDuration *string `field:"optional" json:"sessionDuration" yaml:"sessionDuration"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/iam_workforce_pool#timeouts IamWorkforcePool#timeouts}
	Timeouts *IamWorkforcePoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

