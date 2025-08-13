// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoringalertpolicy


type MonitoringAlertPolicyConditionsConditionSql struct {
	// The Log Analytics SQL query to run, as a string.
	//
	// The query must
	// conform to the required shape. Specifically, the query must not try to
	// filter the input by time.  A filter will automatically be applied
	// to filter the input so that the query receives all rows received
	// since the last time the query was run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/monitoring_alert_policy#query MonitoringAlertPolicy#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// boolean_test block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/monitoring_alert_policy#boolean_test MonitoringAlertPolicy#boolean_test}
	BooleanTest *MonitoringAlertPolicyConditionsConditionSqlBooleanTest `field:"optional" json:"booleanTest" yaml:"booleanTest"`
	// daily block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/monitoring_alert_policy#daily MonitoringAlertPolicy#daily}
	Daily *MonitoringAlertPolicyConditionsConditionSqlDaily `field:"optional" json:"daily" yaml:"daily"`
	// hourly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/monitoring_alert_policy#hourly MonitoringAlertPolicy#hourly}
	Hourly *MonitoringAlertPolicyConditionsConditionSqlHourly `field:"optional" json:"hourly" yaml:"hourly"`
	// minutes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/monitoring_alert_policy#minutes MonitoringAlertPolicy#minutes}
	Minutes *MonitoringAlertPolicyConditionsConditionSqlMinutes `field:"optional" json:"minutes" yaml:"minutes"`
	// row_count_test block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/monitoring_alert_policy#row_count_test MonitoringAlertPolicy#row_count_test}
	RowCountTest *MonitoringAlertPolicyConditionsConditionSqlRowCountTest `field:"optional" json:"rowCountTest" yaml:"rowCountTest"`
}

