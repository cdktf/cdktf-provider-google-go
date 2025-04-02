// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chronicledataaccesslabel


type ChronicleDataAccessLabelTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/chronicle_data_access_label#create ChronicleDataAccessLabel#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/chronicle_data_access_label#delete ChronicleDataAccessLabel#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/chronicle_data_access_label#update ChronicleDataAccessLabel#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

