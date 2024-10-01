// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexdatascan


type DataplexDatascanDataQualitySpecRulesStatisticRangeExpectation struct {
	// column statistics. Possible values: ["STATISTIC_UNDEFINED", "MEAN", "MIN", "MAX"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/dataplex_datascan#statistic DataplexDatascan#statistic}
	Statistic *string `field:"required" json:"statistic" yaml:"statistic"`
	// The maximum column statistic value allowed for a row to pass this validation.
	//
	// At least one of minValue and maxValue need to be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/dataplex_datascan#max_value DataplexDatascan#max_value}
	MaxValue *string `field:"optional" json:"maxValue" yaml:"maxValue"`
	// The minimum column statistic value allowed for a row to pass this validation.
	//
	// At least one of minValue and maxValue need to be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/dataplex_datascan#min_value DataplexDatascan#min_value}
	MinValue *string `field:"optional" json:"minValue" yaml:"minValue"`
	// Whether column statistic needs to be strictly lesser than ('<') the maximum, or if equality is allowed.
	//
	// Only relevant if a maxValue has been defined. Default = false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/dataplex_datascan#strict_max_enabled DataplexDatascan#strict_max_enabled}
	StrictMaxEnabled interface{} `field:"optional" json:"strictMaxEnabled" yaml:"strictMaxEnabled"`
	// Whether column statistic needs to be strictly greater than ('>') the minimum, or if equality is allowed.
	//
	// Only relevant if a minValue has been defined. Default = false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/dataplex_datascan#strict_min_enabled DataplexDatascan#strict_min_enabled}
	StrictMinEnabled interface{} `field:"optional" json:"strictMinEnabled" yaml:"strictMinEnabled"`
}

