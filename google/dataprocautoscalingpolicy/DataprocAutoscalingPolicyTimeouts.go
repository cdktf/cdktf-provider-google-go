// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocautoscalingpolicy


type DataprocAutoscalingPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.6.0/docs/resources/dataproc_autoscaling_policy#create DataprocAutoscalingPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.6.0/docs/resources/dataproc_autoscaling_policy#delete DataprocAutoscalingPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.6.0/docs/resources/dataproc_autoscaling_policy#update DataprocAutoscalingPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

