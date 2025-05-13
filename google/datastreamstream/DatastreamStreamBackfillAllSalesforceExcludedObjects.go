// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastreamstream


type DatastreamStreamBackfillAllSalesforceExcludedObjects struct {
	// objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/datastream_stream#objects DatastreamStream#objects}
	Objects interface{} `field:"required" json:"objects" yaml:"objects"`
}

