// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mlenginemodel


type MlEngineModelTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/ml_engine_model#create MlEngineModel#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/ml_engine_model#delete MlEngineModel#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/ml_engine_model#update MlEngineModel#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

