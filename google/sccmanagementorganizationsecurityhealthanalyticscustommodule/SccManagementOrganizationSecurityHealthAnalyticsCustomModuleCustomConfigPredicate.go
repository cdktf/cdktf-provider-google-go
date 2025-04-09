// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sccmanagementorganizationsecurityhealthanalyticscustommodule


type SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigPredicate struct {
	// Textual representation of an expression in Common Expression Language syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/scc_management_organization_security_health_analytics_custom_module#expression SccManagementOrganizationSecurityHealthAnalyticsCustomModule#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Description of the expression.
	//
	// This is a longer text which describes the
	// expression, e.g. when hovered over it in a UI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/scc_management_organization_security_health_analytics_custom_module#description SccManagementOrganizationSecurityHealthAnalyticsCustomModule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/scc_management_organization_security_health_analytics_custom_module#location SccManagementOrganizationSecurityHealthAnalyticsCustomModule#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/scc_management_organization_security_health_analytics_custom_module#title SccManagementOrganizationSecurityHealthAnalyticsCustomModule#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

