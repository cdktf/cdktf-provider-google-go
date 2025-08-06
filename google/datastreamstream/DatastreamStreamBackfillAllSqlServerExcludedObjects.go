// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastreamstream


type DatastreamStreamBackfillAllSqlServerExcludedObjects struct {
	// schemas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/datastream_stream#schemas DatastreamStream#schemas}
	Schemas interface{} `field:"required" json:"schemas" yaml:"schemas"`
}

