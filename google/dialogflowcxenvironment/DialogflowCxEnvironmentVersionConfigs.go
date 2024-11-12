// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxenvironment


type DialogflowCxEnvironmentVersionConfigs struct {
	// Format: projects/{{project}}/locations/{{location}}/agents/{{agent}}/flows/{{flow}}/versions/{{version}}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.0/docs/resources/dialogflow_cx_environment#version DialogflowCxEnvironment#version}
	Version *string `field:"required" json:"version" yaml:"version"`
}

