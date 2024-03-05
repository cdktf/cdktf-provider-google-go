// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeploytarget


type ClouddeployTargetMultiTarget struct {
	// Required. The target_ids of this multiTarget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/clouddeploy_target#target_ids ClouddeployTarget#target_ids}
	TargetIds *[]*string `field:"required" json:"targetIds" yaml:"targetIds"`
}

