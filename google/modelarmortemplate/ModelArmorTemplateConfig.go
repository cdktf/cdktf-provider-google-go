// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelarmortemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ModelArmorTemplateConfig struct {
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
	// filter_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/model_armor_template#filter_config ModelArmorTemplate#filter_config}
	FilterConfig *ModelArmorTemplateFilterConfig `field:"required" json:"filterConfig" yaml:"filterConfig"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/model_armor_template#location ModelArmorTemplate#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Id of the requesting object If auto-generating Id server-side, remove this field and template_id from the method_signature of Create RPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/model_armor_template#template_id ModelArmorTemplate#template_id}
	TemplateId *string `field:"required" json:"templateId" yaml:"templateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/model_armor_template#id ModelArmorTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/model_armor_template#labels ModelArmorTemplate#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/model_armor_template#project ModelArmorTemplate#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// template_metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/model_armor_template#template_metadata ModelArmorTemplate#template_metadata}
	TemplateMetadata *ModelArmorTemplateTemplateMetadata `field:"optional" json:"templateMetadata" yaml:"templateMetadata"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/model_armor_template#timeouts ModelArmorTemplate#timeouts}
	Timeouts *ModelArmorTemplateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

