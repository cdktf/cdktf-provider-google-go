// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkebackuprestoreplan


type GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilter struct {
	// group_kinds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.15.0/docs/resources/gke_backup_restore_plan#group_kinds GkeBackupRestorePlan#group_kinds}
	GroupKinds interface{} `field:"optional" json:"groupKinds" yaml:"groupKinds"`
	// This is a JSONPath expression that matches specific fields of candidate resources and it operates as a filtering parameter (resources that are not matched with this expression will not be candidates for transformation).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.15.0/docs/resources/gke_backup_restore_plan#json_path GkeBackupRestorePlan#json_path}
	JsonPath *string `field:"optional" json:"jsonPath" yaml:"jsonPath"`
	// (Filtering parameter) Any resource subject to transformation must be contained within one of the listed Kubernetes Namespace in the Backup.
	//
	// If this field is not provided, no namespace filtering will
	// be performed (all resources in all Namespaces, including all
	// cluster-scoped resources, will be candidates for transformation).
	// To mix cluster-scoped and namespaced resources in the same rule,
	// use an empty string ("") as one of the target namespaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.15.0/docs/resources/gke_backup_restore_plan#namespaces GkeBackupRestorePlan#namespaces}
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
}

