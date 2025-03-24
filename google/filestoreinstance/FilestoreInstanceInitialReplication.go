// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package filestoreinstance


type FilestoreInstanceInitialReplication struct {
	// replicas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/filestore_instance#replicas FilestoreInstance#replicas}
	Replicas interface{} `field:"optional" json:"replicas" yaml:"replicas"`
	// The replication role. Default value: "STANDBY" Possible values: ["ROLE_UNSPECIFIED", "ACTIVE", "STANDBY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/filestore_instance#role FilestoreInstance#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

