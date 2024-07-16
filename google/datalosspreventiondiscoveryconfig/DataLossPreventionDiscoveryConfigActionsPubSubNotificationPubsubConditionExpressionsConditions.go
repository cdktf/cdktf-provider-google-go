// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigActionsPubSubNotificationPubsubConditionExpressionsConditions struct {
	// The minimum data risk score that triggers the condition. Possible values: ["HIGH", "MEDIUM_OR_HIGH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/data_loss_prevention_discovery_config#minimum_risk_score DataLossPreventionDiscoveryConfig#minimum_risk_score}
	MinimumRiskScore *string `field:"optional" json:"minimumRiskScore" yaml:"minimumRiskScore"`
	// The minimum sensitivity level that triggers the condition. Possible values: ["HIGH", "MEDIUM_OR_HIGH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/data_loss_prevention_discovery_config#minimum_sensitivity_score DataLossPreventionDiscoveryConfig#minimum_sensitivity_score}
	MinimumSensitivityScore *string `field:"optional" json:"minimumSensitivityScore" yaml:"minimumSensitivityScore"`
}

