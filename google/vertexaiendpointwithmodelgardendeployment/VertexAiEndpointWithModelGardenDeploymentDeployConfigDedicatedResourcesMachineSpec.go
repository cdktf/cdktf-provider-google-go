// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaiendpointwithmodelgardendeployment


type VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpec struct {
	// The number of accelerators to attach to the machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#accelerator_count VertexAiEndpointWithModelGardenDeployment#accelerator_count}
	AcceleratorCount *float64 `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// Possible values: ACCELERATOR_TYPE_UNSPECIFIED NVIDIA_TESLA_K80 NVIDIA_TESLA_P100 NVIDIA_TESLA_V100 NVIDIA_TESLA_P4 NVIDIA_TESLA_T4 NVIDIA_TESLA_A100 NVIDIA_A100_80GB NVIDIA_L4 NVIDIA_H100_80GB NVIDIA_H100_MEGA_80GB NVIDIA_H200_141GB NVIDIA_B200 TPU_V2 TPU_V3 TPU_V4_POD TPU_V5_LITEPOD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#accelerator_type VertexAiEndpointWithModelGardenDeployment#accelerator_type}
	AcceleratorType *string `field:"optional" json:"acceleratorType" yaml:"acceleratorType"`
	// The type of the machine.
	//
	// See the [list of machine types supported for
	// prediction](https://cloud.google.com/vertex-ai/docs/predictions/configure-compute#machine-types)
	//
	// See the [list of machine types supported for custom
	// training](https://cloud.google.com/vertex-ai/docs/training/configure-compute#machine-types).
	//
	// For DeployedModel this field is optional, and the default
	// value is 'n1-standard-2'. For BatchPredictionJob or as part of
	// WorkerPoolSpec this field is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#machine_type VertexAiEndpointWithModelGardenDeployment#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// The number of nodes per replica for multihost GPU deployments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#multihost_gpu_node_count VertexAiEndpointWithModelGardenDeployment#multihost_gpu_node_count}
	MultihostGpuNodeCount *float64 `field:"optional" json:"multihostGpuNodeCount" yaml:"multihostGpuNodeCount"`
	// reservation_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#reservation_affinity VertexAiEndpointWithModelGardenDeployment#reservation_affinity}
	ReservationAffinity *VertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesMachineSpecReservationAffinity `field:"optional" json:"reservationAffinity" yaml:"reservationAffinity"`
	// The topology of the TPUs. Corresponds to the TPU topologies available from GKE. (Example: tpu_topology: "2x2x1").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#tpu_topology VertexAiEndpointWithModelGardenDeployment#tpu_topology}
	TpuTopology *string `field:"optional" json:"tpuTopology" yaml:"tpuTopology"`
}

