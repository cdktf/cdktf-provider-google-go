// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastreamstream


type DatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemasOracleTablesOracleColumns struct {
	// Column name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/datastream_stream#column DatastreamStream#column}
	Column *string `field:"optional" json:"column" yaml:"column"`
	// The Oracle data type. Full data types list can be found here: https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/Data-Types.html.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/datastream_stream#data_type DatastreamStream#data_type}
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
}

