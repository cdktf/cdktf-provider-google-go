// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocworkflowtemplate


type DataprocWorkflowTemplateParametersValidationRegex struct {
	// Required.
	//
	// RE2 regular expressions used to validate the parameter's value. The value must match the regex in its entirety (substring matches are not sufficient).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/dataproc_workflow_template#regexes DataprocWorkflowTemplate#regexes}
	Regexes *[]*string `field:"required" json:"regexes" yaml:"regexes"`
}

