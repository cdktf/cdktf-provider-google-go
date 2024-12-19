// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mlenginemodel


type MlEngineModelDefaultVersion struct {
	// The name specified for the version when it was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/ml_engine_model#name MlEngineModel#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

