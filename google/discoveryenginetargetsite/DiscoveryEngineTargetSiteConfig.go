// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryenginetargetsite

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DiscoveryEngineTargetSiteConfig struct {
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
	// The unique id of the data store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/discovery_engine_target_site#data_store_id DiscoveryEngineTargetSite#data_store_id}
	DataStoreId *string `field:"required" json:"dataStoreId" yaml:"dataStoreId"`
	// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/discovery_engine_target_site#location DiscoveryEngineTargetSite#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The user provided URI pattern from which the 'generated_uri_pattern' is generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/discovery_engine_target_site#provided_uri_pattern DiscoveryEngineTargetSite#provided_uri_pattern}
	ProvidedUriPattern *string `field:"required" json:"providedUriPattern" yaml:"providedUriPattern"`
	// If set to false, a uri_pattern is generated to include all pages whose address contains the provided_uri_pattern.
	//
	// If set to true, an uri_pattern
	// is generated to try to be an exact match of the provided_uri_pattern or
	// just the specific page if the provided_uri_pattern is a specific one.
	// provided_uri_pattern is always normalized to generate the URI pattern to
	// be used by the search engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/discovery_engine_target_site#exact_match DiscoveryEngineTargetSite#exact_match}
	ExactMatch interface{} `field:"optional" json:"exactMatch" yaml:"exactMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/discovery_engine_target_site#id DiscoveryEngineTargetSite#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/discovery_engine_target_site#project DiscoveryEngineTargetSite#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/discovery_engine_target_site#timeouts DiscoveryEngineTargetSite#timeouts}
	Timeouts *DiscoveryEngineTargetSiteTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The possible target site types. Possible values: ["INCLUDE", "EXCLUDE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/discovery_engine_target_site#type DiscoveryEngineTargetSite#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

