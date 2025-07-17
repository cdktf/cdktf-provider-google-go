// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spannerinstanceconfig


type SpannerInstanceConfigReplicas struct {
	// If true, this location is designated as the default leader location where leader replicas are placed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/spanner_instance_config#default_leader_location SpannerInstanceConfigA#default_leader_location}
	DefaultLeaderLocation interface{} `field:"optional" json:"defaultLeaderLocation" yaml:"defaultLeaderLocation"`
	// The location of the serving resources, e.g. "us-central1".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/spanner_instance_config#location SpannerInstanceConfigA#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Indicates the type of replica.  See the [replica types documentation](https://cloud.google.com/spanner/docs/replication#replica_types) for more details. Possible values: ["READ_WRITE", "READ_ONLY", "WITNESS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/spanner_instance_config#type SpannerInstanceConfigA#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

