// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunservice


type CloudRunServiceTemplateSpecContainersEnvFromSecretRef struct {
	// local_object_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.3.0/docs/resources/cloud_run_service#local_object_reference CloudRunService#local_object_reference}
	LocalObjectReference *CloudRunServiceTemplateSpecContainersEnvFromSecretRefLocalObjectReference `field:"optional" json:"localObjectReference" yaml:"localObjectReference"`
	// Specify whether the Secret must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.3.0/docs/resources/cloud_run_service#optional CloudRunService#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

