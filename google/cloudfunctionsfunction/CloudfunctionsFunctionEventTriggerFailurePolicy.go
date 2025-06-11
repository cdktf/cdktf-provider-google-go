// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfunctionsfunction


type CloudfunctionsFunctionEventTriggerFailurePolicy struct {
	// Whether the function should be retried on failure. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/cloudfunctions_function#retry CloudfunctionsFunction#retry}
	Retry interface{} `field:"required" json:"retry" yaml:"retry"`
}

