// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryenginedatastore


type DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigLayoutParsingConfig struct {
	// If true, the LLM based annotation is added to the image during parsing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/discovery_engine_data_store#enable_image_annotation DiscoveryEngineDataStore#enable_image_annotation}
	EnableImageAnnotation interface{} `field:"optional" json:"enableImageAnnotation" yaml:"enableImageAnnotation"`
	// If true, the LLM based annotation is added to the table during parsing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/discovery_engine_data_store#enable_table_annotation DiscoveryEngineDataStore#enable_table_annotation}
	EnableTableAnnotation interface{} `field:"optional" json:"enableTableAnnotation" yaml:"enableTableAnnotation"`
	// List of HTML classes to exclude from the parsed content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/discovery_engine_data_store#exclude_html_classes DiscoveryEngineDataStore#exclude_html_classes}
	ExcludeHtmlClasses *[]*string `field:"optional" json:"excludeHtmlClasses" yaml:"excludeHtmlClasses"`
	// List of HTML elements to exclude from the parsed content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/discovery_engine_data_store#exclude_html_elements DiscoveryEngineDataStore#exclude_html_elements}
	ExcludeHtmlElements *[]*string `field:"optional" json:"excludeHtmlElements" yaml:"excludeHtmlElements"`
	// List of HTML ids to exclude from the parsed content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/discovery_engine_data_store#exclude_html_ids DiscoveryEngineDataStore#exclude_html_ids}
	ExcludeHtmlIds *[]*string `field:"optional" json:"excludeHtmlIds" yaml:"excludeHtmlIds"`
	// Contains the required structure types to extract from the document. Supported values: 'shareholder-structure'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/discovery_engine_data_store#structured_content_types DiscoveryEngineDataStore#structured_content_types}
	StructuredContentTypes *[]*string `field:"optional" json:"structuredContentTypes" yaml:"structuredContentTypes"`
}

