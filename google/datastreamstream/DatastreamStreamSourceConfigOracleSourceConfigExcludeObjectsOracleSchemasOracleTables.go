// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastreamstream


type DatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemasOracleTables struct {
	// Table name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/datastream_stream#table DatastreamStream#table}
	Table *string `field:"required" json:"table" yaml:"table"`
	// oracle_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/datastream_stream#oracle_columns DatastreamStream#oracle_columns}
	OracleColumns interface{} `field:"optional" json:"oracleColumns" yaml:"oracleColumns"`
}

