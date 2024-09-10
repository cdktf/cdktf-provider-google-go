// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerytable


type BigqueryTableExternalDataConfigurationAvroOptions struct {
	// If sourceFormat is set to "AVRO", indicates whether to interpret logical types as the corresponding BigQuery data type (for example, TIMESTAMP), instead of using the raw type (for example, INTEGER).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/bigquery_table#use_avro_logical_types BigqueryTable#use_avro_logical_types}
	UseAvroLogicalTypes interface{} `field:"required" json:"useAvroLogicalTypes" yaml:"useAvroLogicalTypes"`
}

