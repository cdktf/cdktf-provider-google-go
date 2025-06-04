// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunservice


type CloudRunServiceTemplateSpecVolumesCsi struct {
	// Unique name representing the type of file system to be created.
	//
	// Cloud Run supports the following values:
	//   * gcsfuse.run.googleapis.com: Mount a Google Cloud Storage bucket using GCSFuse. This driver requires the
	//     run.googleapis.com/execution-environment annotation to be unset or set to "gen2"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_service#driver CloudRunService#driver}
	Driver *string `field:"required" json:"driver" yaml:"driver"`
	// If true, all mounts created from this volume will be read-only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_service#read_only CloudRunService#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Driver-specific attributes.
	//
	// The following options are supported for available drivers:
	//   * gcsfuse.run.googleapis.com
	//     * bucketName: The name of the Cloud Storage Bucket that backs this volume. The Cloud Run Service identity must have access to this bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_service#volume_attributes CloudRunService#volume_attributes}
	VolumeAttributes *map[string]*string `field:"optional" json:"volumeAttributes" yaml:"volumeAttributes"`
}

