// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeecontrolplaneaccess


type ApigeeControlPlaneAccessTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/apigee_control_plane_access#create ApigeeControlPlaneAccess#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/apigee_control_plane_access#delete ApigeeControlPlaneAccess#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/apigee_control_plane_access#update ApigeeControlPlaneAccess#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

