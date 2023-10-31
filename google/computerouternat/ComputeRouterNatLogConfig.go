// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computerouternat


type ComputeRouterNatLogConfig struct {
	// Indicates whether or not to export logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.4.0/docs/resources/compute_router_nat#enable ComputeRouterNat#enable}
	Enable interface{} `field:"required" json:"enable" yaml:"enable"`
	// Specifies the desired filtering of logs on this NAT. Possible values: ["ERRORS_ONLY", "TRANSLATIONS_ONLY", "ALL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.4.0/docs/resources/compute_router_nat#filter ComputeRouterNat#filter}
	Filter *string `field:"required" json:"filter" yaml:"filter"`
}

