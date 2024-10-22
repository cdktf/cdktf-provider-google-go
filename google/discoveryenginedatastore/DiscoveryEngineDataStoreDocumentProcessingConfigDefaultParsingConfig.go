// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryenginedatastore


type DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfig struct {
	// digital_parsing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/discovery_engine_data_store#digital_parsing_config DiscoveryEngineDataStore#digital_parsing_config}
	DigitalParsingConfig *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigDigitalParsingConfig `field:"optional" json:"digitalParsingConfig" yaml:"digitalParsingConfig"`
	// layout_parsing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/discovery_engine_data_store#layout_parsing_config DiscoveryEngineDataStore#layout_parsing_config}
	LayoutParsingConfig *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigLayoutParsingConfig `field:"optional" json:"layoutParsingConfig" yaml:"layoutParsingConfig"`
	// ocr_parsing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/discovery_engine_data_store#ocr_parsing_config DiscoveryEngineDataStore#ocr_parsing_config}
	OcrParsingConfig *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOcrParsingConfig `field:"optional" json:"ocrParsingConfig" yaml:"ocrParsingConfig"`
}

