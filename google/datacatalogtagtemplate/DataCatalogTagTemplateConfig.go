package datacatalogtagtemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCatalogTagTemplateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_catalog_tag_template#fields DataCatalogTagTemplate#fields}
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// The id of the tag template to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_catalog_tag_template#tag_template_id DataCatalogTagTemplate#tag_template_id}
	TagTemplateId *string `field:"required" json:"tagTemplateId" yaml:"tagTemplateId"`
	// The display name for this template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_catalog_tag_template#display_name DataCatalogTagTemplate#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// This confirms the deletion of any possible tags using this template.
	//
	// Must be set to true in order to delete the tag template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_catalog_tag_template#force_delete DataCatalogTagTemplate#force_delete}
	ForceDelete interface{} `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_catalog_tag_template#id DataCatalogTagTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_catalog_tag_template#project DataCatalogTagTemplate#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Template location region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_catalog_tag_template#region DataCatalogTagTemplate#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_catalog_tag_template#timeouts DataCatalogTagTemplate#timeouts}
	Timeouts *DataCatalogTagTemplateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

