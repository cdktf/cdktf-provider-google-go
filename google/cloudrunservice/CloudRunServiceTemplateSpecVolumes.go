// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunservice


type CloudRunServiceTemplateSpecVolumes struct {
	// Volume's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/cloud_run_service#name CloudRunService#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// csi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/cloud_run_service#csi CloudRunService#csi}
	Csi *CloudRunServiceTemplateSpecVolumesCsi `field:"optional" json:"csi" yaml:"csi"`
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/cloud_run_service#nfs CloudRunService#nfs}
	Nfs *CloudRunServiceTemplateSpecVolumesNfs `field:"optional" json:"nfs" yaml:"nfs"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/cloud_run_service#secret CloudRunService#secret}
	Secret *CloudRunServiceTemplateSpecVolumesSecret `field:"optional" json:"secret" yaml:"secret"`
}

