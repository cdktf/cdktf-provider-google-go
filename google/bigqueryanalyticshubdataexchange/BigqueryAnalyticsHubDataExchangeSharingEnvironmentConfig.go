// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigqueryanalyticshubdataexchange


type BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfig struct {
	// dcr_exchange_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/bigquery_analytics_hub_data_exchange#dcr_exchange_config BigqueryAnalyticsHubDataExchange#dcr_exchange_config}
	DcrExchangeConfig *BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigDcrExchangeConfig `field:"optional" json:"dcrExchangeConfig" yaml:"dcrExchangeConfig"`
	// default_exchange_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/bigquery_analytics_hub_data_exchange#default_exchange_config BigqueryAnalyticsHubDataExchange#default_exchange_config}
	DefaultExchangeConfig *BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigDefaultExchangeConfig `field:"optional" json:"defaultExchangeConfig" yaml:"defaultExchangeConfig"`
}

