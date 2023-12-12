// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastreamstream


type DatastreamStreamSourceConfigOracleSourceConfigIncludeObjects struct {
	// oracle_schemas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.9.0/docs/resources/datastream_stream#oracle_schemas DatastreamStream#oracle_schemas}
	OracleSchemas interface{} `field:"required" json:"oracleSchemas" yaml:"oracleSchemas"`
}

