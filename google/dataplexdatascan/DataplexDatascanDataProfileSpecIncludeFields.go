// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexdatascan


type DataplexDatascanDataProfileSpecIncludeFields struct {
	// Expected input is a list of fully qualified names of fields as in the schema.
	//
	// Only top-level field names for nested fields are supported.
	// For instance, if 'x' is of nested field type, listing 'x' is supported but 'x.y.z' is not supported. Here 'y' and 'y.z' are nested fields of 'x'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/dataplex_datascan#field_names DataplexDatascan#field_names}
	FieldNames *[]*string `field:"optional" json:"fieldNames" yaml:"fieldNames"`
}

