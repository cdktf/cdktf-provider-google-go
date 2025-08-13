// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesRegex struct {
	// Pattern defining the regular expression. Its syntax (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/data_loss_prevention_job_trigger#pattern DataLossPreventionJobTrigger#pattern}
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
	// The index of the submatch to extract as findings.
	//
	// When not specified, the entire match is returned. No more than 3 may be included.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/data_loss_prevention_job_trigger#group_indexes DataLossPreventionJobTrigger#group_indexes}
	GroupIndexes *[]*float64 `field:"optional" json:"groupIndexes" yaml:"groupIndexes"`
}

