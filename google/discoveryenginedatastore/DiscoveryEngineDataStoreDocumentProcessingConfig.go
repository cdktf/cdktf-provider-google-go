// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryenginedatastore


type DiscoveryEngineDataStoreDocumentProcessingConfig struct {
	// chunking_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/discovery_engine_data_store#chunking_config DiscoveryEngineDataStore#chunking_config}
	ChunkingConfig *DiscoveryEngineDataStoreDocumentProcessingConfigChunkingConfig `field:"optional" json:"chunkingConfig" yaml:"chunkingConfig"`
	// default_parsing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/discovery_engine_data_store#default_parsing_config DiscoveryEngineDataStore#default_parsing_config}
	DefaultParsingConfig *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfig `field:"optional" json:"defaultParsingConfig" yaml:"defaultParsingConfig"`
	// parsing_config_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/discovery_engine_data_store#parsing_config_overrides DiscoveryEngineDataStore#parsing_config_overrides}
	ParsingConfigOverrides interface{} `field:"optional" json:"parsingConfigOverrides" yaml:"parsingConfigOverrides"`
}

