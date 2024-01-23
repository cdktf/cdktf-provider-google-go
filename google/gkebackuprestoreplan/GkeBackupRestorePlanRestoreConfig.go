// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkebackuprestoreplan


type GkeBackupRestorePlanRestoreConfig struct {
	// If True, restore all namespaced resources in the Backup. Setting this field to False will result in an error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/gke_backup_restore_plan#all_namespaces GkeBackupRestorePlan#all_namespaces}
	AllNamespaces interface{} `field:"optional" json:"allNamespaces" yaml:"allNamespaces"`
	// Defines the behavior for handling the situation where cluster-scoped resources being restored already exist in the target cluster.
	//
	// This MUST be set to a value other than 'CLUSTER_RESOURCE_CONFLICT_POLICY_UNSPECIFIED'
	// if 'clusterResourceRestoreScope' is anyting other than 'noGroupKinds'.
	// See https://cloud.google.com/kubernetes-engine/docs/add-on/backup-for-gke/reference/rest/v1/RestoreConfig#clusterresourceconflictpolicy
	// for more information on each policy option. Possible values: ["USE_EXISTING_VERSION", "USE_BACKUP_VERSION"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/gke_backup_restore_plan#cluster_resource_conflict_policy GkeBackupRestorePlan#cluster_resource_conflict_policy}
	ClusterResourceConflictPolicy *string `field:"optional" json:"clusterResourceConflictPolicy" yaml:"clusterResourceConflictPolicy"`
	// cluster_resource_restore_scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/gke_backup_restore_plan#cluster_resource_restore_scope GkeBackupRestorePlan#cluster_resource_restore_scope}
	ClusterResourceRestoreScope *GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScope `field:"optional" json:"clusterResourceRestoreScope" yaml:"clusterResourceRestoreScope"`
	// excluded_namespaces block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/gke_backup_restore_plan#excluded_namespaces GkeBackupRestorePlan#excluded_namespaces}
	ExcludedNamespaces *GkeBackupRestorePlanRestoreConfigExcludedNamespaces `field:"optional" json:"excludedNamespaces" yaml:"excludedNamespaces"`
	// Defines the behavior for handling the situation where sets of namespaced resources being restored already exist in the target cluster.
	//
	// This MUST be set to a value other than 'NAMESPACED_RESOURCE_RESTORE_MODE_UNSPECIFIED'
	// if the 'namespacedResourceRestoreScope' is anything other than 'noNamespaces'.
	// See https://cloud.google.com/kubernetes-engine/docs/add-on/backup-for-gke/reference/rest/v1/RestoreConfig#namespacedresourcerestoremode
	// for more information on each mode. Possible values: ["DELETE_AND_RESTORE", "FAIL_ON_CONFLICT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/gke_backup_restore_plan#namespaced_resource_restore_mode GkeBackupRestorePlan#namespaced_resource_restore_mode}
	NamespacedResourceRestoreMode *string `field:"optional" json:"namespacedResourceRestoreMode" yaml:"namespacedResourceRestoreMode"`
	// Do not restore any namespaced resources if set to "True". Specifying this field to "False" is not allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/gke_backup_restore_plan#no_namespaces GkeBackupRestorePlan#no_namespaces}
	NoNamespaces interface{} `field:"optional" json:"noNamespaces" yaml:"noNamespaces"`
	// selected_applications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/gke_backup_restore_plan#selected_applications GkeBackupRestorePlan#selected_applications}
	SelectedApplications *GkeBackupRestorePlanRestoreConfigSelectedApplications `field:"optional" json:"selectedApplications" yaml:"selectedApplications"`
	// selected_namespaces block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/gke_backup_restore_plan#selected_namespaces GkeBackupRestorePlan#selected_namespaces}
	SelectedNamespaces *GkeBackupRestorePlanRestoreConfigSelectedNamespaces `field:"optional" json:"selectedNamespaces" yaml:"selectedNamespaces"`
	// transformation_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/gke_backup_restore_plan#transformation_rules GkeBackupRestorePlan#transformation_rules}
	TransformationRules interface{} `field:"optional" json:"transformationRules" yaml:"transformationRules"`
	// Specifies the mechanism to be used to restore volume data.
	//
	// This should be set to a value other than 'NAMESPACED_RESOURCE_RESTORE_MODE_UNSPECIFIED'
	// if the 'namespacedResourceRestoreScope' is anything other than 'noNamespaces'.
	// If not specified, it will be treated as 'NO_VOLUME_DATA_RESTORATION'.
	// See https://cloud.google.com/kubernetes-engine/docs/add-on/backup-for-gke/reference/rest/v1/RestoreConfig#VolumeDataRestorePolicy
	// for more information on each policy option. Possible values: ["RESTORE_VOLUME_DATA_FROM_BACKUP", "REUSE_VOLUME_HANDLE_FROM_BACKUP", "NO_VOLUME_DATA_RESTORATION"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/gke_backup_restore_plan#volume_data_restore_policy GkeBackupRestorePlan#volume_data_restore_policy}
	VolumeDataRestorePolicy *string `field:"optional" json:"volumeDataRestorePolicy" yaml:"volumeDataRestorePolicy"`
}

