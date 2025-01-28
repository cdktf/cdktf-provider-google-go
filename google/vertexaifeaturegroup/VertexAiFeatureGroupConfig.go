// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaifeaturegroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VertexAiFeatureGroupConfig struct {
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
	// big_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/vertex_ai_feature_group#big_query VertexAiFeatureGroup#big_query}
	BigQuery *VertexAiFeatureGroupBigQuery `field:"optional" json:"bigQuery" yaml:"bigQuery"`
	// The description of the FeatureGroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/vertex_ai_feature_group#description VertexAiFeatureGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/vertex_ai_feature_group#id VertexAiFeatureGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels with user-defined metadata to organize your FeatureGroup.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/vertex_ai_feature_group#labels VertexAiFeatureGroup#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The resource name of the Feature Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/vertex_ai_feature_group#name VertexAiFeatureGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/vertex_ai_feature_group#project VertexAiFeatureGroup#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region of feature group. eg us-central1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/vertex_ai_feature_group#region VertexAiFeatureGroup#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/vertex_ai_feature_group#timeouts VertexAiFeatureGroup#timeouts}
	Timeouts *VertexAiFeatureGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

