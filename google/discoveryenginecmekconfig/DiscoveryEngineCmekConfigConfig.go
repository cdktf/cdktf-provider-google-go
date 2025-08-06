// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryenginecmekconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DiscoveryEngineCmekConfigConfig struct {
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
	// KMS key resource name which will be used to encrypt resources 'projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{keyId}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/discovery_engine_cmek_config#kms_key DiscoveryEngineCmekConfig#kms_key}
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
	// The geographic location where the CMEK config should reside. The value can only be one of "us" and "eu".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/discovery_engine_cmek_config#location DiscoveryEngineCmekConfig#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The unique id of the cmek config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/discovery_engine_cmek_config#cmek_config_id DiscoveryEngineCmekConfig#cmek_config_id}
	CmekConfigId *string `field:"optional" json:"cmekConfigId" yaml:"cmekConfigId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/discovery_engine_cmek_config#id DiscoveryEngineCmekConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/discovery_engine_cmek_config#project DiscoveryEngineCmekConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Set the following CmekConfig as the default to be used for child resources if one is not specified.
	//
	// The default value is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/discovery_engine_cmek_config#set_default DiscoveryEngineCmekConfig#set_default}
	SetDefault interface{} `field:"optional" json:"setDefault" yaml:"setDefault"`
	// single_region_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/discovery_engine_cmek_config#single_region_keys DiscoveryEngineCmekConfig#single_region_keys}
	SingleRegionKeys interface{} `field:"optional" json:"singleRegionKeys" yaml:"singleRegionKeys"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/discovery_engine_cmek_config#timeouts DiscoveryEngineCmekConfig#timeouts}
	Timeouts *DiscoveryEngineCmekConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

