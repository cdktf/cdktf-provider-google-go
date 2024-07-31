// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeaturemembership


type GkeHubFeatureMembershipConfigmanagementConfigSyncGit struct {
	// The GCP Service Account Email used for auth when secretType is gcpServiceAccount.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.1/docs/resources/gke_hub_feature_membership#gcp_service_account_email GkeHubFeatureMembership#gcp_service_account_email}
	GcpServiceAccountEmail *string `field:"optional" json:"gcpServiceAccountEmail" yaml:"gcpServiceAccountEmail"`
	// URL for the HTTPS proxy to be used when communicating with the Git repo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.1/docs/resources/gke_hub_feature_membership#https_proxy GkeHubFeatureMembership#https_proxy}
	HttpsProxy *string `field:"optional" json:"httpsProxy" yaml:"httpsProxy"`
	// The path within the Git repository that represents the top level of the repo to sync.
	//
	// Default: the root directory of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.1/docs/resources/gke_hub_feature_membership#policy_dir GkeHubFeatureMembership#policy_dir}
	PolicyDir *string `field:"optional" json:"policyDir" yaml:"policyDir"`
	// Type of secret configured for access to the Git repo.
	//
	// Must be one of ssh, cookiefile, gcenode, token, gcpserviceaccount or none. The validation of this is case-sensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.1/docs/resources/gke_hub_feature_membership#secret_type GkeHubFeatureMembership#secret_type}
	SecretType *string `field:"optional" json:"secretType" yaml:"secretType"`
	// The branch of the repository to sync from. Default: master.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.1/docs/resources/gke_hub_feature_membership#sync_branch GkeHubFeatureMembership#sync_branch}
	SyncBranch *string `field:"optional" json:"syncBranch" yaml:"syncBranch"`
	// The URL of the Git repository to use as the source of truth.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.1/docs/resources/gke_hub_feature_membership#sync_repo GkeHubFeatureMembership#sync_repo}
	SyncRepo *string `field:"optional" json:"syncRepo" yaml:"syncRepo"`
	// Git revision (tag or hash) to check out. Default HEAD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.1/docs/resources/gke_hub_feature_membership#sync_rev GkeHubFeatureMembership#sync_rev}
	SyncRev *string `field:"optional" json:"syncRev" yaml:"syncRev"`
	// Period in seconds between consecutive syncs. Default: 15.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.1/docs/resources/gke_hub_feature_membership#sync_wait_secs GkeHubFeatureMembership#sync_wait_secs}
	SyncWaitSecs *string `field:"optional" json:"syncWaitSecs" yaml:"syncWaitSecs"`
}

