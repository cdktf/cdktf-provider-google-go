// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sccmanagementfoldersecurityhealthanalyticscustommodule


type SccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputProperties struct {
	// Name of the property for the custom output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/scc_management_folder_security_health_analytics_custom_module#name SccManagementFolderSecurityHealthAnalyticsCustomModule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// value_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/scc_management_folder_security_health_analytics_custom_module#value_expression SccManagementFolderSecurityHealthAnalyticsCustomModule#value_expression}
	ValueExpression *SccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputPropertiesValueExpression `field:"optional" json:"valueExpression" yaml:"valueExpression"`
}

