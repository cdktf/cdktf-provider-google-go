// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastreamstream


type DatastreamStreamDestinationConfigBigqueryDestinationConfigBlmtConfig struct {
	// The Cloud Storage bucket name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/datastream_stream#bucket DatastreamStream#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The bigquery connection. Format: '{project}.{location}.{name}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/datastream_stream#connection_name DatastreamStream#connection_name}
	ConnectionName *string `field:"required" json:"connectionName" yaml:"connectionName"`
	// The file format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/datastream_stream#file_format DatastreamStream#file_format}
	FileFormat *string `field:"required" json:"fileFormat" yaml:"fileFormat"`
	// The table format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/datastream_stream#table_format DatastreamStream#table_format}
	TableFormat *string `field:"required" json:"tableFormat" yaml:"tableFormat"`
	// The root path inside the Cloud Storage bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/datastream_stream#root_path DatastreamStream#root_path}
	RootPath *string `field:"optional" json:"rootPath" yaml:"rootPath"`
}

