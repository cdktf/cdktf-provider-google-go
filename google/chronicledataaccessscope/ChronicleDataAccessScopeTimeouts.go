// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chronicledataaccessscope


type ChronicleDataAccessScopeTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/chronicle_data_access_scope#create ChronicleDataAccessScope#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/chronicle_data_access_scope#delete ChronicleDataAccessScope#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/chronicle_data_access_scope#update ChronicleDataAccessScope#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

