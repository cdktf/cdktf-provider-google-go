// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaideploymentresourcepool


type VertexAiDeploymentResourcePoolDedicatedResourcesAutoscalingMetricSpecs struct {
	// The resource metric name. Supported metrics: For Online Prediction: * 'aiplatform.googleapis.com/prediction/online/accelerator/duty_cycle' * 'aiplatform.googleapis.com/prediction/online/cpu/utilization'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/vertex_ai_deployment_resource_pool#metric_name VertexAiDeploymentResourcePool#metric_name}
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The target resource utilization in percentage (1% - 100%) for the given metric;
	//
	// once the real usage deviates from the target by a certain percentage, the machine replicas change. The default value is 60 (representing 60%) if not provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/vertex_ai_deployment_resource_pool#target VertexAiDeploymentResourcePool#target}
	Target *float64 `field:"optional" json:"target" yaml:"target"`
}

