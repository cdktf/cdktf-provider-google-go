// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerydatapolicydatapolicy


type BigqueryDatapolicyDataPolicyDataMaskingPolicy struct {
	// The available masking rules. Learn more here: https://cloud.google.com/bigquery/docs/column-data-masking-intro#masking_options. Possible values: ["SHA256", "ALWAYS_NULL", "DEFAULT_MASKING_VALUE", "LAST_FOUR_CHARACTERS", "FIRST_FOUR_CHARACTERS", "EMAIL_MASK", "DATE_YEAR_MASK"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/bigquery_datapolicy_data_policy#predefined_expression BigqueryDatapolicyDataPolicy#predefined_expression}
	PredefinedExpression *string `field:"optional" json:"predefinedExpression" yaml:"predefinedExpression"`
	// The name of the BigQuery routine that contains the custom masking routine, in the format of projects/{projectNumber}/datasets/{dataset_id}/routines/{routine_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/bigquery_datapolicy_data_policy#routine BigqueryDatapolicyDataPolicy#routine}
	Routine *string `field:"optional" json:"routine" yaml:"routine"`
}

