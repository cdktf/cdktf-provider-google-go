// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxtool


type DialogflowCxToolDataStoreSpecDataStoreConnections struct {
	// The full name of the referenced data store. Formats: projects/{project}/locations/{location}/collections/{collection}/dataStores/{dataStore} projects/{project}/locations/{location}/dataStores/{dataStore}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/dialogflow_cx_tool#data_store DialogflowCxTool#data_store}
	DataStore *string `field:"optional" json:"dataStore" yaml:"dataStore"`
	// The type of the connected data store. See [DataStoreType](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/DataStoreConnection#datastoretype) for valid values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/dialogflow_cx_tool#data_store_type DialogflowCxTool#data_store_type}
	DataStoreType *string `field:"optional" json:"dataStoreType" yaml:"dataStoreType"`
	// The document processing mode for the data store connection.
	//
	// Should only be set for PUBLIC_WEB and UNSTRUCTURED data stores. If not set it is considered as DOCUMENTS, as this is the legacy mode.
	// See [DocumentProcessingMode](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/DataStoreConnection#documentprocessingmode) for valid values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/dialogflow_cx_tool#document_processing_mode DialogflowCxTool#document_processing_mode}
	DocumentProcessingMode *string `field:"optional" json:"documentProcessingMode" yaml:"documentProcessingMode"`
}

