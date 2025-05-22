// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfunctionsfunction


type CloudfunctionsFunctionSourceRepository struct {
	// The URL pointing to the hosted repository where the function is defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/cloudfunctions_function#url CloudfunctionsFunction#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

