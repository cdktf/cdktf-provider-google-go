// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package blockchainnodeengineblockchainnodes

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BlockchainNodeEngineBlockchainNodesConfig struct {
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
	// ID of the requesting object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/blockchain_node_engine_blockchain_nodes#blockchain_node_id BlockchainNodeEngineBlockchainNodes#blockchain_node_id}
	BlockchainNodeId *string `field:"required" json:"blockchainNodeId" yaml:"blockchainNodeId"`
	// Location of Blockchain Node being created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/blockchain_node_engine_blockchain_nodes#location BlockchainNodeEngineBlockchainNodes#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// User-provided key-value pairs Possible values: ["ETHEREUM"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/blockchain_node_engine_blockchain_nodes#blockchain_type BlockchainNodeEngineBlockchainNodes#blockchain_type}
	BlockchainType *string `field:"optional" json:"blockchainType" yaml:"blockchainType"`
	// ethereum_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/blockchain_node_engine_blockchain_nodes#ethereum_details BlockchainNodeEngineBlockchainNodes#ethereum_details}
	EthereumDetails *BlockchainNodeEngineBlockchainNodesEthereumDetails `field:"optional" json:"ethereumDetails" yaml:"ethereumDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/blockchain_node_engine_blockchain_nodes#id BlockchainNodeEngineBlockchainNodes#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User-provided key-value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/blockchain_node_engine_blockchain_nodes#labels BlockchainNodeEngineBlockchainNodes#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/blockchain_node_engine_blockchain_nodes#project BlockchainNodeEngineBlockchainNodes#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/blockchain_node_engine_blockchain_nodes#timeouts BlockchainNodeEngineBlockchainNodes#timeouts}
	Timeouts *BlockchainNodeEngineBlockchainNodesTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

