// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoringalertpolicy


type MonitoringAlertPolicyDocumentationLinks struct {
	// A short display name for the link.
	//
	// The display name must not be empty or exceed 63 characters. Example: "playbook".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/monitoring_alert_policy#display_name MonitoringAlertPolicy#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The url of a webpage.
	//
	// A url can be templatized by using variables in the path or the query parameters. The total length of a URL should not exceed 2083 characters before and after variable expansion. Example: "https://my_domain.com/playbook?name=${resource.name}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/monitoring_alert_policy#url MonitoringAlertPolicy#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

