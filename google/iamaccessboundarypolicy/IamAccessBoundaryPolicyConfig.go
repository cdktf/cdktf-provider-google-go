package iamaccessboundarypolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamAccessBoundaryPolicyConfig struct {
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
	// The name of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/iam_access_boundary_policy#name IamAccessBoundaryPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The attachment point is identified by its URL-encoded full resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/iam_access_boundary_policy#parent IamAccessBoundaryPolicy#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/iam_access_boundary_policy#rules IamAccessBoundaryPolicy#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// The display name of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/iam_access_boundary_policy#display_name IamAccessBoundaryPolicy#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/iam_access_boundary_policy#id IamAccessBoundaryPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/iam_access_boundary_policy#timeouts IamAccessBoundaryPolicy#timeouts}
	Timeouts *IamAccessBoundaryPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

