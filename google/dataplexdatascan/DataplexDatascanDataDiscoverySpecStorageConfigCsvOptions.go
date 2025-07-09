// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexdatascan


type DataplexDatascanDataDiscoverySpecStorageConfigCsvOptions struct {
	// The delimiter that is used to separate values. The default is ',' (comma).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/dataplex_datascan#delimiter DataplexDatascan#delimiter}
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// The character encoding of the data. The default is UTF-8.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/dataplex_datascan#encoding DataplexDatascan#encoding}
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// The number of rows to interpret as header rows that should be skipped when reading data rows.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/dataplex_datascan#header_rows DataplexDatascan#header_rows}
	HeaderRows *float64 `field:"optional" json:"headerRows" yaml:"headerRows"`
	// The character used to quote column values.
	//
	// Accepts '"' (double quotation mark) or ``` (single quotation mark). If unspecified, defaults to '"' (double quotation mark).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/dataplex_datascan#quote DataplexDatascan#quote}
	Quote *string `field:"optional" json:"quote" yaml:"quote"`
	// Whether to disable the inference of data types for CSV data. If true, all columns are registered as strings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/dataplex_datascan#type_inference_disabled DataplexDatascan#type_inference_disabled}
	TypeInferenceDisabled interface{} `field:"optional" json:"typeInferenceDisabled" yaml:"typeInferenceDisabled"`
}

