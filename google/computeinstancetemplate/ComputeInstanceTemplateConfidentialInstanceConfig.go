// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstancetemplate


type ComputeInstanceTemplateConfidentialInstanceConfig struct {
	// Specifies which confidential computing technology to use.
	//
	// This could be one of the following values: SEV, SEV_SNP.
	// If SEV_SNP, min_cpu_platform = "AMD Milan" is currently required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/compute_instance_template#confidential_instance_type ComputeInstanceTemplate#confidential_instance_type}
	ConfidentialInstanceType *string `field:"optional" json:"confidentialInstanceType" yaml:"confidentialInstanceType"`
	// Defines whether the instance should have confidential compute enabled. Field will be deprecated in a future release.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/compute_instance_template#enable_confidential_compute ComputeInstanceTemplate#enable_confidential_compute}
	EnableConfidentialCompute interface{} `field:"optional" json:"enableConfidentialCompute" yaml:"enableConfidentialCompute"`
}

