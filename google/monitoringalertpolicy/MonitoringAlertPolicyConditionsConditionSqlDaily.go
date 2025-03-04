// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoringalertpolicy


type MonitoringAlertPolicyConditionsConditionSqlDaily struct {
	// The number of days between runs.
	//
	// Must be greater than or equal
	// to 1 day and less than or equal to 30 days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/monitoring_alert_policy#periodicity MonitoringAlertPolicy#periodicity}
	Periodicity *float64 `field:"required" json:"periodicity" yaml:"periodicity"`
	// execution_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/monitoring_alert_policy#execution_time MonitoringAlertPolicy#execution_time}
	ExecutionTime *MonitoringAlertPolicyConditionsConditionSqlDailyExecutionTime `field:"optional" json:"executionTime" yaml:"executionTime"`
}

