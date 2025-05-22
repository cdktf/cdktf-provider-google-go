// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerytable


type BigqueryTableRangePartitioning struct {
	// The field used to determine how to create a range-based partition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/bigquery_table#field BigqueryTable#field}
	Field *string `field:"required" json:"field" yaml:"field"`
	// range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/bigquery_table#range BigqueryTable#range}
	Range *BigqueryTableRangePartitioningRange `field:"required" json:"range" yaml:"range"`
}

