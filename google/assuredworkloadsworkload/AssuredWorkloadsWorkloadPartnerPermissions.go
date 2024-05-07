// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package assuredworkloadsworkload


type AssuredWorkloadsWorkloadPartnerPermissions struct {
	// Optional. Allow partner to view violation alerts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.28.0/docs/resources/assured_workloads_workload#assured_workloads_monitoring AssuredWorkloadsWorkload#assured_workloads_monitoring}
	AssuredWorkloadsMonitoring interface{} `field:"optional" json:"assuredWorkloadsMonitoring" yaml:"assuredWorkloadsMonitoring"`
	// Allow the partner to view inspectability logs and monitoring violations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.28.0/docs/resources/assured_workloads_workload#data_logs_viewer AssuredWorkloadsWorkload#data_logs_viewer}
	DataLogsViewer interface{} `field:"optional" json:"dataLogsViewer" yaml:"dataLogsViewer"`
	// Optional. Allow partner to view access approval logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.28.0/docs/resources/assured_workloads_workload#service_access_approver AssuredWorkloadsWorkload#service_access_approver}
	ServiceAccessApprover interface{} `field:"optional" json:"serviceAccessApprover" yaml:"serviceAccessApprover"`
}

