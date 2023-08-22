package organizationiamcustomrole

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationIamCustomRoleConfig struct {
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
	// The numeric ID of the organization in which you want to create a custom role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/organization_iam_custom_role#org_id OrganizationIamCustomRole#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// The names of the permissions this role grants when bound in an IAM policy.
	//
	// At least one permission must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/organization_iam_custom_role#permissions OrganizationIamCustomRole#permissions}
	Permissions *[]*string `field:"required" json:"permissions" yaml:"permissions"`
	// The role id to use for this role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/organization_iam_custom_role#role_id OrganizationIamCustomRole#role_id}
	RoleId *string `field:"required" json:"roleId" yaml:"roleId"`
	// A human-readable title for the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/organization_iam_custom_role#title OrganizationIamCustomRole#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// A human-readable description for the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/organization_iam_custom_role#description OrganizationIamCustomRole#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/organization_iam_custom_role#id OrganizationIamCustomRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The current launch stage of the role. Defaults to GA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/organization_iam_custom_role#stage OrganizationIamCustomRole#stage}
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
}

