// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sccmanagementfoldersecurityhealthanalyticscustommodule


type SccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelector struct {
	// The resource types to run the detector on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/scc_management_folder_security_health_analytics_custom_module#resource_types SccManagementFolderSecurityHealthAnalyticsCustomModule#resource_types}
	ResourceTypes *[]*string `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
}

