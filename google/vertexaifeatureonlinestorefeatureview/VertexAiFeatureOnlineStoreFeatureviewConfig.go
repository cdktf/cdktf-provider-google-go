// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaifeatureonlinestorefeatureview

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VertexAiFeatureOnlineStoreFeatureviewConfig struct {
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
	// The name of the FeatureOnlineStore to use for the featureview.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/vertex_ai_feature_online_store_featureview#feature_online_store VertexAiFeatureOnlineStoreFeatureview#feature_online_store}
	FeatureOnlineStore *string `field:"required" json:"featureOnlineStore" yaml:"featureOnlineStore"`
	// big_query_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/vertex_ai_feature_online_store_featureview#big_query_source VertexAiFeatureOnlineStoreFeatureview#big_query_source}
	BigQuerySource *VertexAiFeatureOnlineStoreFeatureviewBigQuerySource `field:"optional" json:"bigQuerySource" yaml:"bigQuerySource"`
	// feature_registry_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/vertex_ai_feature_online_store_featureview#feature_registry_source VertexAiFeatureOnlineStoreFeatureview#feature_registry_source}
	FeatureRegistrySource *VertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySource `field:"optional" json:"featureRegistrySource" yaml:"featureRegistrySource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/vertex_ai_feature_online_store_featureview#id VertexAiFeatureOnlineStoreFeatureview#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A set of key/value label pairs to assign to this FeatureView.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/vertex_ai_feature_online_store_featureview#labels VertexAiFeatureOnlineStoreFeatureview#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Name of the FeatureView.
	//
	// This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/vertex_ai_feature_online_store_featureview#name VertexAiFeatureOnlineStoreFeatureview#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/vertex_ai_feature_online_store_featureview#project VertexAiFeatureOnlineStoreFeatureview#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region for the resource. It should be the same as the featureonlinestore region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/vertex_ai_feature_online_store_featureview#region VertexAiFeatureOnlineStoreFeatureview#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// sync_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/vertex_ai_feature_online_store_featureview#sync_config VertexAiFeatureOnlineStoreFeatureview#sync_config}
	SyncConfig *VertexAiFeatureOnlineStoreFeatureviewSyncConfig `field:"optional" json:"syncConfig" yaml:"syncConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/vertex_ai_feature_online_store_featureview#timeouts VertexAiFeatureOnlineStoreFeatureview#timeouts}
	Timeouts *VertexAiFeatureOnlineStoreFeatureviewTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

