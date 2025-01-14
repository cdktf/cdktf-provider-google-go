// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaiendpoint


type VertexAiEndpointPredictRequestResponseLoggingConfig struct {
	// bigquery_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.16.0/docs/resources/vertex_ai_endpoint#bigquery_destination VertexAiEndpoint#bigquery_destination}
	BigqueryDestination *VertexAiEndpointPredictRequestResponseLoggingConfigBigqueryDestination `field:"optional" json:"bigqueryDestination" yaml:"bigqueryDestination"`
	// If logging is enabled or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.16.0/docs/resources/vertex_ai_endpoint#enabled VertexAiEndpoint#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Percentage of requests to be logged, expressed as a fraction in range(0,1].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.16.0/docs/resources/vertex_ai_endpoint#sampling_rate VertexAiEndpoint#sampling_rate}
	SamplingRate *float64 `field:"optional" json:"samplingRate" yaml:"samplingRate"`
}

