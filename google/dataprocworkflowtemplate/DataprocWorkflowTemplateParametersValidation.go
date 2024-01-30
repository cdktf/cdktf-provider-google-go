// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocworkflowtemplate


type DataprocWorkflowTemplateParametersValidation struct {
	// regex block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/dataproc_workflow_template#regex DataprocWorkflowTemplate#regex}
	Regex *DataprocWorkflowTemplateParametersValidationRegex `field:"optional" json:"regex" yaml:"regex"`
	// values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/dataproc_workflow_template#values DataprocWorkflowTemplate#values}
	Values *DataprocWorkflowTemplateParametersValidationValues `field:"optional" json:"values" yaml:"values"`
}

