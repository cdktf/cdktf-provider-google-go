// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelarmortemplate


type ModelArmorTemplateFilterConfigRaiSettingsRaiFilters struct {
	// Possible values: SEXUALLY_EXPLICIT HATE_SPEECH HARASSMENT DANGEROUS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/model_armor_template#filter_type ModelArmorTemplate#filter_type}
	FilterType *string `field:"required" json:"filterType" yaml:"filterType"`
	// Possible values: LOW_AND_ABOVE MEDIUM_AND_ABOVE HIGH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/model_armor_template#confidence_level ModelArmorTemplate#confidence_level}
	ConfidenceLevel *string `field:"optional" json:"confidenceLevel" yaml:"confidenceLevel"`
}

