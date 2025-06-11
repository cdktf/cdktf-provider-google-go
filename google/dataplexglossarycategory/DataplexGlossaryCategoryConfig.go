// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexglossarycategory

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataplexGlossaryCategoryConfig struct {
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
	// The location where the glossary category should reside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/dataplex_glossary_category#location DataplexGlossaryCategory#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The immediate parent of the GlossaryCategory in the resource-hierarchy.
	//
	// It can either be a Glossary or a Category. Format: projects/{projectId}/locations/{locationId}/glossaries/{glossaryId} OR projects/{projectId}/locations/{locationId}/glossaries/{glossaryId}/categories/{categoryId}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/dataplex_glossary_category#parent DataplexGlossaryCategory#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// The category id for creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/dataplex_glossary_category#category_id DataplexGlossaryCategory#category_id}
	CategoryId *string `field:"optional" json:"categoryId" yaml:"categoryId"`
	// The user-mutable description of the GlossaryCategory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/dataplex_glossary_category#description DataplexGlossaryCategory#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User friendly display name of the GlossaryCategory.
	//
	// This is user-mutable. This will be same as the categoryId, if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/dataplex_glossary_category#display_name DataplexGlossaryCategory#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The glossary id for creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/dataplex_glossary_category#glossary_id DataplexGlossaryCategory#glossary_id}
	GlossaryId *string `field:"optional" json:"glossaryId" yaml:"glossaryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/dataplex_glossary_category#id DataplexGlossaryCategory#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User-defined labels for the GlossaryCategory.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/dataplex_glossary_category#labels DataplexGlossaryCategory#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/dataplex_glossary_category#project DataplexGlossaryCategory#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/dataplex_glossary_category#timeouts DataplexGlossaryCategory#timeouts}
	Timeouts *DataplexGlossaryCategoryTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

