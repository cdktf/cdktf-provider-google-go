// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastreamprivateconnection


type DatastreamPrivateConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.11.0/docs/resources/datastream_private_connection#create DatastreamPrivateConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.11.0/docs/resources/datastream_private_connection#delete DatastreamPrivateConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

