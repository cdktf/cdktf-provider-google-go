// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package edgecontainercluster


type EdgecontainerClusterMaintenancePolicyWindowRecurringWindow struct {
	// An RRULE (https://tools.ietf.org/html/rfc5545#section-3.8.5.3) for how this window recurs. They go on for the span of time between the start and end time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.34.0/docs/resources/edgecontainer_cluster#recurrence EdgecontainerCluster#recurrence}
	Recurrence *string `field:"optional" json:"recurrence" yaml:"recurrence"`
	// window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.34.0/docs/resources/edgecontainer_cluster#window EdgecontainerCluster#window}
	Window *EdgecontainerClusterMaintenancePolicyWindowRecurringWindowWindow `field:"optional" json:"window" yaml:"window"`
}

