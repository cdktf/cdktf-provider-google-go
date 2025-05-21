// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryenginechatengine

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DiscoveryEngineChatEngineConfig struct {
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
	// chat_engine_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/discovery_engine_chat_engine#chat_engine_config DiscoveryEngineChatEngine#chat_engine_config}
	ChatEngineConfig *DiscoveryEngineChatEngineChatEngineConfig `field:"required" json:"chatEngineConfig" yaml:"chatEngineConfig"`
	// The collection ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/discovery_engine_chat_engine#collection_id DiscoveryEngineChatEngine#collection_id}
	CollectionId *string `field:"required" json:"collectionId" yaml:"collectionId"`
	// The data stores associated with this engine.
	//
	// Multiple DataStores in the same Collection can be associated here. All listed DataStores must be 'SOLUTION_TYPE_CHAT'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/discovery_engine_chat_engine#data_store_ids DiscoveryEngineChatEngine#data_store_ids}
	DataStoreIds *[]*string `field:"required" json:"dataStoreIds" yaml:"dataStoreIds"`
	// The display name of the engine. Should be human readable. UTF-8 encoded string with limit of 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/discovery_engine_chat_engine#display_name DiscoveryEngineChatEngine#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The ID to use for chat engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/discovery_engine_chat_engine#engine_id DiscoveryEngineChatEngine#engine_id}
	EngineId *string `field:"required" json:"engineId" yaml:"engineId"`
	// Location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/discovery_engine_chat_engine#location DiscoveryEngineChatEngine#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// common_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/discovery_engine_chat_engine#common_config DiscoveryEngineChatEngine#common_config}
	CommonConfig *DiscoveryEngineChatEngineCommonConfig `field:"optional" json:"commonConfig" yaml:"commonConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/discovery_engine_chat_engine#id DiscoveryEngineChatEngine#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The industry vertical that the chat engine registers.
	//
	// Vertical on Engine has to match vertical of the DataStore linked to the engine. Default value: "GENERIC" Possible values: ["GENERIC"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/discovery_engine_chat_engine#industry_vertical DiscoveryEngineChatEngine#industry_vertical}
	IndustryVertical *string `field:"optional" json:"industryVertical" yaml:"industryVertical"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/discovery_engine_chat_engine#project DiscoveryEngineChatEngine#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/discovery_engine_chat_engine#timeouts DiscoveryEngineChatEngine#timeouts}
	Timeouts *DiscoveryEngineChatEngineTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

