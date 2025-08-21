// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexglossaryterm

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataplexGlossaryTermConfig struct {
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
	// The location where the glossary term should reside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/dataplex_glossary_term#location DataplexGlossaryTerm#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The immediate parent of the GlossaryTerm in the resource-hierarchy.
	//
	// It can either be a Glossary or a Term. Format: projects/{projectId}/locations/{locationId}/glossaries/{glossaryId} OR projects/{projectId}/locations/{locationId}/glossaries/{glossaryId}/terms/{termId}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/dataplex_glossary_term#parent DataplexGlossaryTerm#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// The user-mutable description of the GlossaryTerm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/dataplex_glossary_term#description DataplexGlossaryTerm#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User friendly display name of the GlossaryTerm.
	//
	// This is user-mutable. This will be same as the termId, if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/dataplex_glossary_term#display_name DataplexGlossaryTerm#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The glossary id for creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/dataplex_glossary_term#glossary_id DataplexGlossaryTerm#glossary_id}
	GlossaryId *string `field:"optional" json:"glossaryId" yaml:"glossaryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/dataplex_glossary_term#id DataplexGlossaryTerm#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User-defined labels for the GlossaryTerm.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/dataplex_glossary_term#labels DataplexGlossaryTerm#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/dataplex_glossary_term#project DataplexGlossaryTerm#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The term id for creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/dataplex_glossary_term#term_id DataplexGlossaryTerm#term_id}
	TermId *string `field:"optional" json:"termId" yaml:"termId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/dataplex_glossary_term#timeouts DataplexGlossaryTerm#timeouts}
	Timeouts *DataplexGlossaryTermTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

