// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package scceventthreatdetectioncustommodule


type SccEventThreatDetectionCustomModuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/scc_event_threat_detection_custom_module#create SccEventThreatDetectionCustomModule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/scc_event_threat_detection_custom_module#delete SccEventThreatDetectionCustomModule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/scc_event_threat_detection_custom_module#update SccEventThreatDetectionCustomModule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

