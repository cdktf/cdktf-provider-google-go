// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocworkflowtemplate


type DataprocWorkflowTemplateTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dataproc_workflow_template#create DataprocWorkflowTemplate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dataproc_workflow_template#delete DataprocWorkflowTemplate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dataproc_workflow_template#update DataprocWorkflowTemplate#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

