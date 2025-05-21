// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigActionsPubSubNotificationPubsubCondition struct {
	// expressions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/data_loss_prevention_discovery_config#expressions DataLossPreventionDiscoveryConfig#expressions}
	Expressions *DataLossPreventionDiscoveryConfigActionsPubSubNotificationPubsubConditionExpressions `field:"optional" json:"expressions" yaml:"expressions"`
}

