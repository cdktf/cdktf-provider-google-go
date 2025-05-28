// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxflow


type DialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnections struct {
	// The full name of the referenced data store. Formats: projects/{project}/locations/{location}/collections/{collection}/dataStores/{dataStore} projects/{project}/locations/{location}/dataStores/{dataStore}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/dialogflow_cx_flow#data_store DialogflowCxFlow#data_store}
	DataStore *string `field:"optional" json:"dataStore" yaml:"dataStore"`
	// The type of the connected data store.
	//
	// * PUBLIC_WEB: A data store that contains public web content.
	// * UNSTRUCTURED: A data store that contains unstructured private data.
	// * STRUCTURED: A data store that contains structured data (for example FAQ). Possible values: ["PUBLIC_WEB", "UNSTRUCTURED", "STRUCTURED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/dialogflow_cx_flow#data_store_type DialogflowCxFlow#data_store_type}
	DataStoreType *string `field:"optional" json:"dataStoreType" yaml:"dataStoreType"`
	// The document processing mode for the data store connection.
	//
	// Should only be set for PUBLIC_WEB and UNSTRUCTURED data stores. If not set it is considered as DOCUMENTS, as this is the legacy mode.
	// * DOCUMENTS: Documents are processed as documents.
	// * CHUNKS: Documents are converted to chunks. Possible values: ["DOCUMENTS", "CHUNKS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/dialogflow_cx_flow#document_processing_mode DialogflowCxFlow#document_processing_mode}
	DocumentProcessingMode *string `field:"optional" json:"documentProcessingMode" yaml:"documentProcessingMode"`
}

