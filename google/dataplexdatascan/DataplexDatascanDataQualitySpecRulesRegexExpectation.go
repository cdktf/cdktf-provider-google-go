// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexdatascan


type DataplexDatascanDataQualitySpecRulesRegexExpectation struct {
	// A regular expression the column value is expected to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/dataplex_datascan#regex DataplexDatascan#regex}
	Regex *string `field:"required" json:"regex" yaml:"regex"`
}

