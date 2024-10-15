// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigActionsTagResourcesTagConditionsTag struct {
	// The namespaced name for the tag value to attach to resources.
	//
	// Must be in the format '{parent_id}/{tag_key_short_name}/{short_name}', for example, "123456/environment/prod".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/data_loss_prevention_discovery_config#namespaced_value DataLossPreventionDiscoveryConfig#namespaced_value}
	NamespacedValue *string `field:"optional" json:"namespacedValue" yaml:"namespacedValue"`
}

