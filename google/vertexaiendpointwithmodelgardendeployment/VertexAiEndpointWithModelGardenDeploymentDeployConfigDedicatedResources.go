// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaiendpointwithmodelgardendeployment


type VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResources struct {
	// machine_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#machine_spec VertexAiEndpointWithModelGardenDeployment#machine_spec}
	MachineSpec *VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpec `field:"required" json:"machineSpec" yaml:"machineSpec"`
	// The minimum number of machine replicas that will be always deployed on.
	//
	// This value must be greater than or equal to 1.
	//
	// If traffic increases, it may dynamically be deployed onto more replicas,
	// and as traffic decreases, some of these extra replicas may be freed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#min_replica_count VertexAiEndpointWithModelGardenDeployment#min_replica_count}
	MinReplicaCount *float64 `field:"required" json:"minReplicaCount" yaml:"minReplicaCount"`
	// autoscaling_metric_specs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#autoscaling_metric_specs VertexAiEndpointWithModelGardenDeployment#autoscaling_metric_specs}
	AutoscalingMetricSpecs interface{} `field:"optional" json:"autoscalingMetricSpecs" yaml:"autoscalingMetricSpecs"`
	// The maximum number of replicas that may be deployed on when the traffic against it increases.
	//
	// If the requested value is too large, the deployment
	// will error, but if deployment succeeds then the ability to scale to that
	// many replicas is guaranteed (barring service outages). If traffic increases
	// beyond what its replicas at maximum may handle, a portion of the traffic
	// will be dropped. If this value is not provided, will use
	// min_replica_count as the default value.
	//
	// The value of this field impacts the charge against Vertex CPU and GPU
	// quotas. Specifically, you will be charged for (max_replica_count *
	// number of cores in the selected machine type) and (max_replica_count *
	// number of GPUs per replica in the selected machine type).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#max_replica_count VertexAiEndpointWithModelGardenDeployment#max_replica_count}
	MaxReplicaCount *float64 `field:"optional" json:"maxReplicaCount" yaml:"maxReplicaCount"`
	// Number of required available replicas for the deployment to succeed.
	//
	// This field is only needed when partial deployment/mutation is
	// desired. If set, the deploy/mutate operation will succeed once
	// available_replica_count reaches required_replica_count, and the rest of
	// the replicas will be retried. If not set, the default
	// required_replica_count will be min_replica_count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#required_replica_count VertexAiEndpointWithModelGardenDeployment#required_replica_count}
	RequiredReplicaCount *float64 `field:"optional" json:"requiredReplicaCount" yaml:"requiredReplicaCount"`
	// If true, schedule the deployment workload on [spot VMs](https://cloud.google.com/kubernetes-engine/docs/concepts/spot-vms).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#spot VertexAiEndpointWithModelGardenDeployment#spot}
	Spot interface{} `field:"optional" json:"spot" yaml:"spot"`
}

