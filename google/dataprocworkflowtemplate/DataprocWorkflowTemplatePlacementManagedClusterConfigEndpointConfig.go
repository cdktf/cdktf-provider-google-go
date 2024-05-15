// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocworkflowtemplate


type DataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfig struct {
	// Optional. If true, enable http access to specific ports on the cluster from external sources. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.1/docs/resources/dataproc_workflow_template#enable_http_port_access DataprocWorkflowTemplate#enable_http_port_access}
	EnableHttpPortAccess interface{} `field:"optional" json:"enableHttpPortAccess" yaml:"enableHttpPortAccess"`
}

