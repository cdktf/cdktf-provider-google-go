// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2job


type CloudRunV2JobTemplateTemplateVolumesGcs struct {
	// Name of the cloud storage bucket to back the volume.
	//
	// The resource service account must have permission to access the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/cloud_run_v2_job#bucket CloudRunV2Job#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// If true, mount this volume as read-only in all mounts. If false, mount this volume as read-write.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/cloud_run_v2_job#read_only CloudRunV2Job#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

