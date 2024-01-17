// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocworkflowtemplate


type DataprocWorkflowTemplateParametersValidationValues struct {
	// Required. List of allowed values for the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.12.0/docs/resources/dataproc_workflow_template#values DataprocWorkflowTemplate#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

