// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeaturemembership


type GkeHubFeatureMembershipConfigmanagementConfigSync struct {
	// Enables the installation of ConfigSync.
	//
	// If set to true, ConfigSync resources will be created and the other ConfigSync fields will be applied if exist. If set to false, all other ConfigSync fields will be ignored, ConfigSync resources will be deleted. If omitted, ConfigSync resources will be managed depends on the presence of the git or oci field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/gke_hub_feature_membership#enabled GkeHubFeatureMembership#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// git block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/gke_hub_feature_membership#git GkeHubFeatureMembership#git}
	Git *GkeHubFeatureMembershipConfigmanagementConfigSyncGit `field:"optional" json:"git" yaml:"git"`
	// The Email of the Google Cloud Service Account (GSA) used for exporting Config Sync metrics to Cloud Monitoring.
	//
	// The GSA should have the Monitoring Metric Writer(roles/monitoring.metricWriter) IAM role. The Kubernetes ServiceAccount `default` in the namespace `config-management-monitoring` should be bound to the GSA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/gke_hub_feature_membership#metrics_gcp_service_account_email GkeHubFeatureMembership#metrics_gcp_service_account_email}
	MetricsGcpServiceAccountEmail *string `field:"optional" json:"metricsGcpServiceAccountEmail" yaml:"metricsGcpServiceAccountEmail"`
	// oci block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/gke_hub_feature_membership#oci GkeHubFeatureMembership#oci}
	Oci *GkeHubFeatureMembershipConfigmanagementConfigSyncOci `field:"optional" json:"oci" yaml:"oci"`
	// Set to true to enable the Config Sync admission webhook to prevent drifts.
	//
	// If set to `false`, disables the Config Sync admission webhook and does not prevent drifts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/gke_hub_feature_membership#prevent_drift GkeHubFeatureMembership#prevent_drift}
	PreventDrift interface{} `field:"optional" json:"preventDrift" yaml:"preventDrift"`
	// Specifies whether the Config Sync Repo is in "hierarchical" or "unstructured" mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/gke_hub_feature_membership#source_format GkeHubFeatureMembership#source_format}
	SourceFormat *string `field:"optional" json:"sourceFormat" yaml:"sourceFormat"`
}

