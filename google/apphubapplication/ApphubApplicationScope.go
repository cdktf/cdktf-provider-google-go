// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apphubapplication


type ApphubApplicationScope struct {
	// Required. Scope Type.   Possible values: REGIONAL Possible values: ["REGIONAL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/apphub_application#type ApphubApplication#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

