// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chronicleretrohunt


type ChronicleRetrohuntTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/chronicle_retrohunt#create ChronicleRetrohunt#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/chronicle_retrohunt#delete ChronicleRetrohunt#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

