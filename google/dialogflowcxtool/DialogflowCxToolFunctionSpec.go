// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxtool


type DialogflowCxToolFunctionSpec struct {
	// Optional.
	//
	// The JSON schema is encapsulated in a [google.protobuf.Struct](https://protobuf.dev/reference/protobuf/google.protobuf/#struct) to describe the input of the function.
	// This input is a JSON object that contains the function's parameters as properties of the object
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/dialogflow_cx_tool#input_schema DialogflowCxTool#input_schema}
	InputSchema *string `field:"optional" json:"inputSchema" yaml:"inputSchema"`
	// Optional.
	//
	// The JSON schema is encapsulated in a [google.protobuf.Struct](https://protobuf.dev/reference/protobuf/google.protobuf/#struct) to describe the output of the function.
	// This output is a JSON object that contains the function's parameters as properties of the object
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/dialogflow_cx_tool#output_schema DialogflowCxTool#output_schema}
	OutputSchema *string `field:"optional" json:"outputSchema" yaml:"outputSchema"`
}

