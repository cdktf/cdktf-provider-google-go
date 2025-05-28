// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oracledatabasecloudvmcluster


type OracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptions struct {
	// Indicates whether diagnostic collection is enabled for the VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/oracle_database_cloud_vm_cluster#diagnostics_events_enabled OracleDatabaseCloudVmCluster#diagnostics_events_enabled}
	DiagnosticsEventsEnabled interface{} `field:"optional" json:"diagnosticsEventsEnabled" yaml:"diagnosticsEventsEnabled"`
	// Indicates whether health monitoring is enabled for the VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/oracle_database_cloud_vm_cluster#health_monitoring_enabled OracleDatabaseCloudVmCluster#health_monitoring_enabled}
	HealthMonitoringEnabled interface{} `field:"optional" json:"healthMonitoringEnabled" yaml:"healthMonitoringEnabled"`
	// Indicates whether incident logs and trace collection are enabled for the VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/oracle_database_cloud_vm_cluster#incident_logs_enabled OracleDatabaseCloudVmCluster#incident_logs_enabled}
	IncidentLogsEnabled interface{} `field:"optional" json:"incidentLogsEnabled" yaml:"incidentLogsEnabled"`
}

