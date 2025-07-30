// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apihubapihubinstance


type ApihubApiHubInstanceConfigA struct {
	// Optional.
	//
	// The Customer Managed Encryption Key (CMEK) used for data encryption.
	// The CMEK name should follow the format of
	// 'projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)',
	// where the location must match the instance location.
	// If the CMEK is not provided, a GMEK will be created for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/apihub_api_hub_instance#cmek_key_name ApihubApiHubInstance#cmek_key_name}
	CmekKeyName *string `field:"optional" json:"cmekKeyName" yaml:"cmekKeyName"`
	// Optional. If true, the search will be disabled for the instance. The default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/apihub_api_hub_instance#disable_search ApihubApiHubInstance#disable_search}
	DisableSearch interface{} `field:"optional" json:"disableSearch" yaml:"disableSearch"`
	// Optional.
	//
	// Encryption type for the region. If the encryption type is CMEK, the
	// cmek_key_name must be provided. If no encryption type is provided,
	// GMEK will be used.
	// Possible values:
	// ENCRYPTION_TYPE_UNSPECIFIED
	// GMEK
	// CMEK
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/apihub_api_hub_instance#encryption_type ApihubApiHubInstance#encryption_type}
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// Optional. The name of the Vertex AI location where the data store is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/apihub_api_hub_instance#vertex_location ApihubApiHubInstance#vertex_location}
	VertexLocation *string `field:"optional" json:"vertexLocation" yaml:"vertexLocation"`
}

