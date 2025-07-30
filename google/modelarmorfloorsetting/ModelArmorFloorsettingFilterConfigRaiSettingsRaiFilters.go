// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelarmorfloorsetting


type ModelArmorFloorsettingFilterConfigRaiSettingsRaiFilters struct {
	// Possible values: SEXUALLY_EXPLICIT HATE_SPEECH HARASSMENT DANGEROUS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/model_armor_floorsetting#filter_type ModelArmorFloorsetting#filter_type}
	FilterType *string `field:"required" json:"filterType" yaml:"filterType"`
	// Possible values: LOW_AND_ABOVE MEDIUM_AND_ABOVE HIGH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/model_armor_floorsetting#confidence_level ModelArmorFloorsetting#confidence_level}
	ConfidenceLevel *string `field:"optional" json:"confidenceLevel" yaml:"confidenceLevel"`
}

