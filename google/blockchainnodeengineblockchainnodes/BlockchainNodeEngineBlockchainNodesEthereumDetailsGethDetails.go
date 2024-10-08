// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package blockchainnodeengineblockchainnodes


type BlockchainNodeEngineBlockchainNodesEthereumDetailsGethDetails struct {
	// Blockchain garbage collection modes. Only applicable when NodeType is FULL or ARCHIVE. Possible values: ["FULL", "ARCHIVE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/blockchain_node_engine_blockchain_nodes#garbage_collection_mode BlockchainNodeEngineBlockchainNodes#garbage_collection_mode}
	GarbageCollectionMode *string `field:"optional" json:"garbageCollectionMode" yaml:"garbageCollectionMode"`
}

