// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apphubworkload


type ApphubWorkloadAttributesCriticality struct {
	// Criticality type. Possible values: ["MISSION_CRITICAL", "HIGH", "MEDIUM", "LOW"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/apphub_workload#type ApphubWorkload#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

