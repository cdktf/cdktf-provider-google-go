// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigActionsPubSubNotificationPubsubConditionExpressions struct {
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/data_loss_prevention_discovery_config#conditions DataLossPreventionDiscoveryConfig#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// The operator to apply to the collection of conditions Possible values: ["OR", "AND"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/data_loss_prevention_discovery_config#logical_operator DataLossPreventionDiscoveryConfig#logical_operator}
	LogicalOperator *string `field:"optional" json:"logicalOperator" yaml:"logicalOperator"`
}

