// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package composerenvironment


type ComposerEnvironmentConfigWorkloadsConfigTriggerer struct {
	// The number of triggerers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/composer_environment#count ComposerEnvironment#count}
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// CPU request and limit for a single Airflow triggerer replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/composer_environment#cpu ComposerEnvironment#cpu}
	Cpu *float64 `field:"required" json:"cpu" yaml:"cpu"`
	// Memory (GB) request and limit for a single Airflow triggerer replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/composer_environment#memory_gb ComposerEnvironment#memory_gb}
	MemoryGb *float64 `field:"required" json:"memoryGb" yaml:"memoryGb"`
}

