package tagstagkey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TagsTagKeyConfig struct {
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
	// Input only. The resource name of the new TagKey's parent. Must be of the form organizations/{org_id}.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/tags_tag_key#parent TagsTagKey#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Input only.
	//
	// The user friendly name for a TagKey. The short name should be unique for TagKeys within the same tag namespace.
	//
	// The short name must be 1-63 characters, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/tags_tag_key#short_name TagsTagKey#short_name}
	ShortName *string `field:"required" json:"shortName" yaml:"shortName"`
	// User-assigned description of the TagKey. Must not exceed 256 characters.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/tags_tag_key#description TagsTagKey#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/tags_tag_key#id TagsTagKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/tags_tag_key#timeouts TagsTagKey#timeouts}
	Timeouts *TagsTagKeyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

