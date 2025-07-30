// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexdatascan


type DataplexDatascanDataDiscoverySpecStorageConfigJsonOptions struct {
	// The character encoding of the data. The default is UTF-8.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dataplex_datascan#encoding DataplexDatascan#encoding}
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// Whether to disable the inference of data types for JSON data.
	//
	// If true, all columns are registered as their primitive types (strings, number, or boolean).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dataplex_datascan#type_inference_disabled DataplexDatascan#type_inference_disabled}
	TypeInferenceDisabled interface{} `field:"optional" json:"typeInferenceDisabled" yaml:"typeInferenceDisabled"`
}

