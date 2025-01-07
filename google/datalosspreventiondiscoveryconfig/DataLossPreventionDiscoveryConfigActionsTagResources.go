// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigActionsTagResources struct {
	// Whether applying a tag to a resource should lower the risk of the profile for that resource.
	//
	// For example, in conjunction with an [IAM deny policy](https://cloud.google.com/iam/docs/deny-overview), you can deny all principals a permission if a tag value is present, mitigating the risk of the resource. This also lowers the data risk of resources at the lower levels of the resource hierarchy. For example, reducing the data risk of a table data profile also reduces the data risk of the constituent column data profiles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/data_loss_prevention_discovery_config#lower_data_risk_to_low DataLossPreventionDiscoveryConfig#lower_data_risk_to_low}
	LowerDataRiskToLow interface{} `field:"optional" json:"lowerDataRiskToLow" yaml:"lowerDataRiskToLow"`
	// The profile generations for which the tag should be attached to resources.
	//
	// If you attach a tag to only new profiles, then if the sensitivity score of a profile subsequently changes, its tag doesn't change. By default, this field includes only new profiles. To include both new and updated profiles for tagging, this field should explicitly include both 'PROFILE_GENERATION_NEW' and 'PROFILE_GENERATION_UPDATE'. Possible values: ["PROFILE_GENERATION_NEW", "PROFILE_GENERATION_UPDATE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/data_loss_prevention_discovery_config#profile_generations_to_tag DataLossPreventionDiscoveryConfig#profile_generations_to_tag}
	ProfileGenerationsToTag *[]*string `field:"optional" json:"profileGenerationsToTag" yaml:"profileGenerationsToTag"`
	// tag_conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/data_loss_prevention_discovery_config#tag_conditions DataLossPreventionDiscoveryConfig#tag_conditions}
	TagConditions interface{} `field:"optional" json:"tagConditions" yaml:"tagConditions"`
}

