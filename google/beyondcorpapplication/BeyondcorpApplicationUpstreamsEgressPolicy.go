// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package beyondcorpapplication


type BeyondcorpApplicationUpstreamsEgressPolicy struct {
	// Required. List of regions where the application sends traffic to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/beyondcorp_application#regions BeyondcorpApplication#regions}
	Regions *[]*string `field:"required" json:"regions" yaml:"regions"`
}

