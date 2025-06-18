// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastreamstream


type DatastreamStreamBackfillAllSalesforceExcludedObjectsObjects struct {
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/datastream_stream#fields DatastreamStream#fields}
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// Name of object in Salesforce Org.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/datastream_stream#object_name DatastreamStream#object_name}
	ObjectName *string `field:"optional" json:"objectName" yaml:"objectName"`
}

