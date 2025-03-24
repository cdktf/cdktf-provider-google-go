// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeature


type GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGit struct {
	// Type of secret configured for access to the Git repo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/gke_hub_feature#secret_type GkeHubFeature#secret_type}
	SecretType *string `field:"required" json:"secretType" yaml:"secretType"`
	// The Google Cloud Service Account Email used for auth when secretType is gcpServiceAccount.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/gke_hub_feature#gcp_service_account_email GkeHubFeature#gcp_service_account_email}
	GcpServiceAccountEmail *string `field:"optional" json:"gcpServiceAccountEmail" yaml:"gcpServiceAccountEmail"`
	// URL for the HTTPS Proxy to be used when communicating with the Git repo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/gke_hub_feature#https_proxy GkeHubFeature#https_proxy}
	HttpsProxy *string `field:"optional" json:"httpsProxy" yaml:"httpsProxy"`
	// The path within the Git repository that represents the top level of the repo to sync.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/gke_hub_feature#policy_dir GkeHubFeature#policy_dir}
	PolicyDir *string `field:"optional" json:"policyDir" yaml:"policyDir"`
	// The branch of the repository to sync from. Default: master.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/gke_hub_feature#sync_branch GkeHubFeature#sync_branch}
	SyncBranch *string `field:"optional" json:"syncBranch" yaml:"syncBranch"`
	// The URL of the Git repository to use as the source of truth.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/gke_hub_feature#sync_repo GkeHubFeature#sync_repo}
	SyncRepo *string `field:"optional" json:"syncRepo" yaml:"syncRepo"`
	// Git revision (tag or hash) to check out. Default HEAD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/gke_hub_feature#sync_rev GkeHubFeature#sync_rev}
	SyncRev *string `field:"optional" json:"syncRev" yaml:"syncRev"`
	// Period in seconds between consecutive syncs. Default: 15.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/gke_hub_feature#sync_wait_secs GkeHubFeature#sync_wait_secs}
	SyncWaitSecs *string `field:"optional" json:"syncWaitSecs" yaml:"syncWaitSecs"`
}

