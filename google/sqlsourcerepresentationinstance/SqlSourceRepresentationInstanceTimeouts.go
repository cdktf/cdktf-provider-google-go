// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqlsourcerepresentationinstance


type SqlSourceRepresentationInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/sql_source_representation_instance#create SqlSourceRepresentationInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/sql_source_representation_instance#delete SqlSourceRepresentationInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

