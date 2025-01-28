// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryenginedatastore


type DiscoveryEngineDataStoreDocumentProcessingConfigChunkingConfigLayoutBasedChunkingConfig struct {
	// The token size limit for each chunk. Supported values: 100-500 (inclusive). Default value: 500.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/discovery_engine_data_store#chunk_size DiscoveryEngineDataStore#chunk_size}
	ChunkSize *float64 `field:"optional" json:"chunkSize" yaml:"chunkSize"`
	// Whether to include appending different levels of headings to chunks from the middle of the document to prevent context loss.
	//
	// Default value: False.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/discovery_engine_data_store#include_ancestor_headings DiscoveryEngineDataStore#include_ancestor_headings}
	IncludeAncestorHeadings interface{} `field:"optional" json:"includeAncestorHeadings" yaml:"includeAncestorHeadings"`
}

