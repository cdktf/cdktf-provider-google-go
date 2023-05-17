package datacatalogpolicytag

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCatalogPolicyTagConfig struct {
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
	// User defined name of this policy tag.
	//
	// It must: be unique within the parent
	// taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces;
	// not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_catalog_policy_tag#display_name DataCatalogPolicyTag#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Taxonomy the policy tag is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_catalog_policy_tag#taxonomy DataCatalogPolicyTag#taxonomy}
	Taxonomy *string `field:"required" json:"taxonomy" yaml:"taxonomy"`
	// Description of this policy tag.
	//
	// It must: contain only unicode characters, tabs,
	// newlines, carriage returns and page breaks; and be at most 2000 bytes long when
	// encoded in UTF-8. If not set, defaults to an empty description.
	// If not set, defaults to an empty description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_catalog_policy_tag#description DataCatalogPolicyTag#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_catalog_policy_tag#id DataCatalogPolicyTag#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource name of this policy tag's parent policy tag.
	//
	// If empty, it means this policy tag is a top level policy tag.
	// If not set, defaults to an empty string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_catalog_policy_tag#parent_policy_tag DataCatalogPolicyTag#parent_policy_tag}
	ParentPolicyTag *string `field:"optional" json:"parentPolicyTag" yaml:"parentPolicyTag"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_catalog_policy_tag#timeouts DataCatalogPolicyTag#timeouts}
	Timeouts *DataCatalogPolicyTagTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

