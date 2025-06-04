// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxpage


type DialogflowCxPageKnowledgeConnectorSettingsDataStoreConnections struct {
	// The full name of the referenced data store. Formats: projects/{project}/locations/{location}/collections/{collection}/dataStores/{dataStore} projects/{project}/locations/{location}/dataStores/{dataStore}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/dialogflow_cx_page#data_store DialogflowCxPage#data_store}
	DataStore *string `field:"optional" json:"dataStore" yaml:"dataStore"`
	// The type of the connected data store.
	//
	// * PUBLIC_WEB: A data store that contains public web content.
	// * UNSTRUCTURED: A data store that contains unstructured private data.
	// * STRUCTURED: A data store that contains structured data (for example FAQ). Possible values: ["PUBLIC_WEB", "UNSTRUCTURED", "STRUCTURED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/dialogflow_cx_page#data_store_type DialogflowCxPage#data_store_type}
	DataStoreType *string `field:"optional" json:"dataStoreType" yaml:"dataStoreType"`
	// The document processing mode for the data store connection.
	//
	// Should only be set for PUBLIC_WEB and UNSTRUCTURED data stores. If not set it is considered as DOCUMENTS, as this is the legacy mode.
	// * DOCUMENTS: Documents are processed as documents.
	// * CHUNKS: Documents are converted to chunks. Possible values: ["DOCUMENTS", "CHUNKS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/dialogflow_cx_page#document_processing_mode DialogflowCxPage#document_processing_mode}
	DocumentProcessingMode *string `field:"optional" json:"documentProcessingMode" yaml:"documentProcessingMode"`
}

