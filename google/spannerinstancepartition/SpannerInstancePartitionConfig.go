// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spannerinstancepartition

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpannerInstancePartitionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the instance partition's configuration (similar to a region) which defines the geographic placement and replication of data in this instance partition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/spanner_instance_partition#config SpannerInstancePartition#config}
	Config *string `field:"required" json:"config" yaml:"config"`
	// The descriptive name for this instance partition as it appears in UIs.
	//
	// Must be unique per project and between 4 and 30 characters in length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/spanner_instance_partition#display_name SpannerInstancePartition#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The instance to create the instance partition in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/spanner_instance_partition#instance SpannerInstancePartition#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// A unique identifier for the instance partition, which cannot be changed after the instance partition is created.
	//
	// The name must be between 2 and 64 characters
	// and match the regular expression [a-z][a-z0-9\\-]{0,61}[a-z0-9].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/spanner_instance_partition#name SpannerInstancePartition#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/spanner_instance_partition#id SpannerInstancePartition#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The number of nodes allocated to this instance partition.
	//
	// One node equals
	// 1000 processing units. Exactly one of either node_count or processing_units
	// must be present.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/spanner_instance_partition#node_count SpannerInstancePartition#node_count}
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
	// The number of processing units allocated to this instance partition. Exactly one of either node_count or processing_units must be present.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/spanner_instance_partition#processing_units SpannerInstancePartition#processing_units}
	ProcessingUnits *float64 `field:"optional" json:"processingUnits" yaml:"processingUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/spanner_instance_partition#project SpannerInstancePartition#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/spanner_instance_partition#timeouts SpannerInstancePartition#timeouts}
	Timeouts *SpannerInstancePartitionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

