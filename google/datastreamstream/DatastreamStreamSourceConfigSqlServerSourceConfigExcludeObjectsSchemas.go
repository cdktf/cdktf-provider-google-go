// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastreamstream


type DatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemas struct {
	// Schema name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/datastream_stream#schema DatastreamStream#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// tables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/datastream_stream#tables DatastreamStream#tables}
	Tables interface{} `field:"optional" json:"tables" yaml:"tables"`
}

