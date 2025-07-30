// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sccmanagementprojectsecurityhealthanalyticscustommodule


type SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelector struct {
	// The resource types to run the detector on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/scc_management_project_security_health_analytics_custom_module#resource_types SccManagementProjectSecurityHealthAnalyticsCustomModule#resource_types}
	ResourceTypes *[]*string `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
}

