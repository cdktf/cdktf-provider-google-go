// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoringalertpolicy


type MonitoringAlertPolicyConditionsConditionSqlBooleanTest struct {
	// The name of the column containing the boolean value.
	//
	// If the value in a row is
	// NULL, that row is ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/monitoring_alert_policy#column MonitoringAlertPolicy#column}
	Column *string `field:"required" json:"column" yaml:"column"`
}

