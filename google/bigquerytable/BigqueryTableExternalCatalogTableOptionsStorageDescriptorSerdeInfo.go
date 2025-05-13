// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerytable


type BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfo struct {
	// Specifies a fully-qualified class name of the serialization library that is responsible for the translation of data between table representation and the underlying low-level input and output format structures.
	//
	// The maximum length is 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/bigquery_table#serialization_library BigqueryTable#serialization_library}
	SerializationLibrary *string `field:"required" json:"serializationLibrary" yaml:"serializationLibrary"`
	// Name of the SerDe. The maximum length is 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/bigquery_table#name BigqueryTable#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Key-value pairs that define the initialization parameters for the serialization library. Maximum size 10 Kib.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/bigquery_table#parameters BigqueryTable#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

