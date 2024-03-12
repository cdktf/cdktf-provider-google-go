// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package migrationcenterpreferenceset


type MigrationCenterPreferenceSetVirtualMachinePreferencesSoleTenancyPreferencesNodeTypes struct {
	// Name of the Sole Tenant node. Consult https://cloud.google.com/compute/docs/nodes/sole-tenant-nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.20.0/docs/resources/migration_center_preference_set#node_name MigrationCenterPreferenceSet#node_name}
	NodeName *string `field:"optional" json:"nodeName" yaml:"nodeName"`
}

