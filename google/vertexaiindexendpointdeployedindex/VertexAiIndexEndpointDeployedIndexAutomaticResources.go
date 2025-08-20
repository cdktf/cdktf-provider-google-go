// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaiindexendpointdeployedindex


type VertexAiIndexEndpointDeployedIndexAutomaticResources struct {
	// The maximum number of replicas this DeployedModel may be deployed on when the traffic against it increases.
	//
	// If maxReplicaCount is not set, the default value is minReplicaCount. The max allowed replica count is 1000.
	//
	// The maximum number of replicas this DeployedModel may be deployed on when the traffic against it increases. If the requested value is too large, the deployment will error, but if deployment succeeds then the ability to scale the model to that many replicas is guaranteed (barring service outages). If traffic against the DeployedModel increases beyond what its replicas at maximum may handle, a portion of the traffic will be dropped. If this value is not provided, a no upper bound for scaling under heavy traffic will be assume, though Vertex AI may be unable to scale beyond certain replica number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/vertex_ai_index_endpoint_deployed_index#max_replica_count VertexAiIndexEndpointDeployedIndex#max_replica_count}
	MaxReplicaCount *float64 `field:"optional" json:"maxReplicaCount" yaml:"maxReplicaCount"`
	// The minimum number of replicas this DeployedModel will be always deployed on.
	//
	// If minReplicaCount is not set, the default value is 2 (we don't provide SLA when minReplicaCount=1).
	//
	// If traffic against it increases, it may dynamically be deployed onto more replicas up to [maxReplicaCount](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/AutomaticResources#FIELDS.max_replica_count), and as traffic decreases, some of these extra replicas may be freed. If the requested value is too large, the deployment will error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/vertex_ai_index_endpoint_deployed_index#min_replica_count VertexAiIndexEndpointDeployedIndex#min_replica_count}
	MinReplicaCount *float64 `field:"optional" json:"minReplicaCount" yaml:"minReplicaCount"`
}

