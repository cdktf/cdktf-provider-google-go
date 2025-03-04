// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbuildtrigger


type CloudbuildTriggerBuildStepVolumes struct {
	// Name of the volume to mount.
	//
	// Volume names must be unique per build step and must be valid names for
	// Docker volumes. Each named volume must be used by at least two build steps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/cloudbuild_trigger#name CloudbuildTrigger#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Path at which to mount the volume.
	//
	// Paths must be absolute and cannot conflict with other volume paths on
	// the same build step or with certain reserved volume paths.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/cloudbuild_trigger#path CloudbuildTrigger#path}
	Path *string `field:"required" json:"path" yaml:"path"`
}

