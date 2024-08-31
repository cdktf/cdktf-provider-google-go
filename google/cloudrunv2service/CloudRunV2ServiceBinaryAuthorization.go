// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2service


type CloudRunV2ServiceBinaryAuthorization struct {
	// If present, indicates to use Breakglass using this justification.
	//
	// If useDefault is False, then it must be empty. For more information on breakglass, see https://cloud.google.com/binary-authorization/docs/using-breakglass
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/cloud_run_v2_service#breakglass_justification CloudRunV2Service#breakglass_justification}
	BreakglassJustification *string `field:"optional" json:"breakglassJustification" yaml:"breakglassJustification"`
	// The path to a binary authorization policy. Format: projects/{project}/platforms/cloudRun/{policy-name}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/cloud_run_v2_service#policy CloudRunV2Service#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// If True, indicates to use the default project's binary authorization policy. If False, binary authorization will be disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/cloud_run_v2_service#use_default CloudRunV2Service#use_default}
	UseDefault interface{} `field:"optional" json:"useDefault" yaml:"useDefault"`
}

