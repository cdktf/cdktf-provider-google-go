// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastreamconnectionprofile


type DatastreamConnectionProfilePrivateConnectivity struct {
	// A reference to a private connection resource. Format: 'projects/{project}/locations/{location}/privateConnections/{name}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.15.0/docs/resources/datastream_connection_profile#private_connection DatastreamConnectionProfile#private_connection}
	PrivateConnection *string `field:"required" json:"privateConnection" yaml:"privateConnection"`
}

