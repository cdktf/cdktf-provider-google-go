// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastreamstream


type DatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjects struct {
	// schemas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/datastream_stream#schemas DatastreamStream#schemas}
	Schemas interface{} `field:"required" json:"schemas" yaml:"schemas"`
}

