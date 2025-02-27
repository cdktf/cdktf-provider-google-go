// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apphubworkload

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApphubWorkloadConfig struct {
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
	// Part of 'parent'.  Full resource name of a parent Application. Example: projects/{HOST_PROJECT_ID}/locations/{LOCATION}/applications/{APPLICATION_ID}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/apphub_workload#application_id ApphubWorkload#application_id}
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Immutable. The resource name of the original discovered workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/apphub_workload#discovered_workload ApphubWorkload#discovered_workload}
	DiscoveredWorkload *string `field:"required" json:"discoveredWorkload" yaml:"discoveredWorkload"`
	// Part of 'parent'.  Full resource name of a parent Application. Example: projects/{HOST_PROJECT_ID}/locations/{LOCATION}/applications/{APPLICATION_ID}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/apphub_workload#location ApphubWorkload#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The Workload identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/apphub_workload#workload_id ApphubWorkload#workload_id}
	WorkloadId *string `field:"required" json:"workloadId" yaml:"workloadId"`
	// attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/apphub_workload#attributes ApphubWorkload#attributes}
	Attributes *ApphubWorkloadAttributes `field:"optional" json:"attributes" yaml:"attributes"`
	// User-defined description of a Workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/apphub_workload#description ApphubWorkload#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User-defined name for the Workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/apphub_workload#display_name ApphubWorkload#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/apphub_workload#id ApphubWorkload#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/apphub_workload#project ApphubWorkload#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/apphub_workload#timeouts ApphubWorkload#timeouts}
	Timeouts *ApphubWorkloadTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

