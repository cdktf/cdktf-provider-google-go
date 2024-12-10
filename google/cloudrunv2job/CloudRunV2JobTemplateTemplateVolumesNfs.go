// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2job


type CloudRunV2JobTemplateTemplateVolumesNfs struct {
	// Hostname or IP address of the NFS server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/cloud_run_v2_job#server CloudRunV2Job#server}
	Server *string `field:"required" json:"server" yaml:"server"`
	// Path that is exported by the NFS server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/cloud_run_v2_job#path CloudRunV2Job#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// If true, mount this volume as read-only in all mounts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/cloud_run_v2_job#read_only CloudRunV2Job#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

