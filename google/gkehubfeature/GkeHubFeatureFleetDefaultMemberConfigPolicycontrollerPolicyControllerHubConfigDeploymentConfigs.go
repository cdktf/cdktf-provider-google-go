// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeature


type GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigDeploymentConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.25.0/docs/resources/gke_hub_feature#component GkeHubFeature#component}.
	Component *string `field:"required" json:"component" yaml:"component"`
	// container_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.25.0/docs/resources/gke_hub_feature#container_resources GkeHubFeature#container_resources}
	ContainerResources *GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResources `field:"optional" json:"containerResources" yaml:"containerResources"`
	// Pod affinity configuration. Possible values: ["AFFINITY_UNSPECIFIED", "NO_AFFINITY", "ANTI_AFFINITY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.25.0/docs/resources/gke_hub_feature#pod_affinity GkeHubFeature#pod_affinity}
	PodAffinity *string `field:"optional" json:"podAffinity" yaml:"podAffinity"`
	// pod_toleration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.25.0/docs/resources/gke_hub_feature#pod_toleration GkeHubFeature#pod_toleration}
	PodToleration interface{} `field:"optional" json:"podToleration" yaml:"podToleration"`
	// Pod replica count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.25.0/docs/resources/gke_hub_feature#replica_count GkeHubFeature#replica_count}
	ReplicaCount *float64 `field:"optional" json:"replicaCount" yaml:"replicaCount"`
}

