// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeploytarget


type ClouddeployTargetAssociatedEntities struct {
	// The name for the key in the map for which this object is mapped to in the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/clouddeploy_target#entity_id ClouddeployTarget#entity_id}
	EntityId *string `field:"required" json:"entityId" yaml:"entityId"`
	// anthos_clusters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/clouddeploy_target#anthos_clusters ClouddeployTarget#anthos_clusters}
	AnthosClusters interface{} `field:"optional" json:"anthosClusters" yaml:"anthosClusters"`
	// gke_clusters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/clouddeploy_target#gke_clusters ClouddeployTarget#gke_clusters}
	GkeClusters interface{} `field:"optional" json:"gkeClusters" yaml:"gkeClusters"`
}

