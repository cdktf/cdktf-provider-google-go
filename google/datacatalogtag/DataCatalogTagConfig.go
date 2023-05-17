package datacatalogtag

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCatalogTagConfig struct {
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
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_catalog_tag#fields DataCatalogTag#fields}
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// The resource name of the tag template that this tag uses. Example: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId} This field cannot be modified after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_catalog_tag#template DataCatalogTag#template}
	Template *string `field:"required" json:"template" yaml:"template"`
	// Resources like Entry can have schemas associated with them.
	//
	// This scope allows users to attach tags to an
	// individual column based on that schema.
	//
	// For attaching a tag to a nested column, use '.' to separate the column names. Example:
	// 'outer_column.inner_column'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_catalog_tag#column DataCatalogTag#column}
	Column *string `field:"optional" json:"column" yaml:"column"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_catalog_tag#id DataCatalogTag#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the parent this tag is attached to.
	//
	// This can be the name of an entry or an entry group. If an entry group, the tag will be attached to
	// all entries in that group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_catalog_tag#parent DataCatalogTag#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_catalog_tag#timeouts DataCatalogTag#timeouts}
	Timeouts *DataCatalogTagTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

