package cloudidentitygroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudIdentityGroupConfig struct {
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
	// group_key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_identity_group#group_key CloudIdentityGroup#group_key}
	GroupKey *CloudIdentityGroupGroupKey `field:"required" json:"groupKey" yaml:"groupKey"`
	// The labels that apply to the Group.
	//
	// Must not contain more than one entry. Must contain the entry
	// 'cloudidentity.googleapis.com/groups.discussion_forum': '' if the Group is a Google Group or
	// 'system/groups/external': '' if the Group is an external-identity-mapped group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_identity_group#labels CloudIdentityGroup#labels}
	Labels *map[string]*string `field:"required" json:"labels" yaml:"labels"`
	// The resource name of the entity under which this Group resides in the Cloud Identity resource hierarchy.
	//
	// Must be of the form identitysources/{identity_source_id} for external-identity-mapped
	// groups or customers/{customer_id} for Google Groups.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_identity_group#parent CloudIdentityGroup#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// An extended description to help users determine the purpose of a Group. Must not be longer than 4,096 characters.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_identity_group#description CloudIdentityGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of the Group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_identity_group#display_name CloudIdentityGroup#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_identity_group#id CloudIdentityGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The initial configuration options for creating a Group.
	//
	// See the
	// [API reference](https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups/create#initialgroupconfig)
	// for possible values. Default value: "EMPTY" Possible values: ["INITIAL_GROUP_CONFIG_UNSPECIFIED", "WITH_INITIAL_OWNER", "EMPTY"]
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_identity_group#initial_group_config CloudIdentityGroup#initial_group_config}
	InitialGroupConfig *string `field:"optional" json:"initialGroupConfig" yaml:"initialGroupConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_identity_group#timeouts CloudIdentityGroup#timeouts}
	Timeouts *CloudIdentityGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

