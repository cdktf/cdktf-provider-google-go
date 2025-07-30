// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbuildtrigger


type CloudbuildTriggerWebhookConfig struct {
	// Resource name for the secret required as a URL parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/cloudbuild_trigger#secret CloudbuildTrigger#secret}
	Secret *string `field:"required" json:"secret" yaml:"secret"`
}

