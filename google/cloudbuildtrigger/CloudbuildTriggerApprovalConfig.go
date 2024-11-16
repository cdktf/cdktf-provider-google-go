// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbuildtrigger


type CloudbuildTriggerApprovalConfig struct {
	// Whether or not approval is needed.
	//
	// If this is set on a build, it will become pending when run,
	// and will need to be explicitly approved to start.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/cloudbuild_trigger#approval_required CloudbuildTrigger#approval_required}
	ApprovalRequired interface{} `field:"optional" json:"approvalRequired" yaml:"approvalRequired"`
}

