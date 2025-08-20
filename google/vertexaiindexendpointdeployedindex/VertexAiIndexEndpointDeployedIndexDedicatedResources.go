// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaiindexendpointdeployedindex


type VertexAiIndexEndpointDeployedIndexDedicatedResources struct {
	// machine_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/vertex_ai_index_endpoint_deployed_index#machine_spec VertexAiIndexEndpointDeployedIndex#machine_spec}
	MachineSpec *VertexAiIndexEndpointDeployedIndexDedicatedResourcesMachineSpec `field:"required" json:"machineSpec" yaml:"machineSpec"`
	// The minimum number of machine replicas this DeployedModel will be always deployed on.
	//
	// This value must be greater than or equal to 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/vertex_ai_index_endpoint_deployed_index#min_replica_count VertexAiIndexEndpointDeployedIndex#min_replica_count}
	MinReplicaCount *float64 `field:"required" json:"minReplicaCount" yaml:"minReplicaCount"`
	// The maximum number of replicas this DeployedModel may be deployed on when the traffic against it increases.
	//
	// If maxReplicaCount is not set, the default value is minReplicaCount
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/vertex_ai_index_endpoint_deployed_index#max_replica_count VertexAiIndexEndpointDeployedIndex#max_replica_count}
	MaxReplicaCount *float64 `field:"optional" json:"maxReplicaCount" yaml:"maxReplicaCount"`
}

