// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apihubhostprojectregistration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApihubHostProjectRegistrationConfig struct {
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
	// Required.
	//
	// Immutable. Google cloud project name in the format: "projects/abc" or "projects/123".
	// As input, project name with either project id or number are accepted.
	// As output, this field will contain project number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/apihub_host_project_registration#gcp_project ApihubHostProjectRegistration#gcp_project}
	GcpProject *string `field:"required" json:"gcpProject" yaml:"gcpProject"`
	// Required.
	//
	// The ID to use for the Host Project Registration, which will become the
	// final component of the host project registration's resource name. The ID
	// must be the same as the Google cloud project specified in the
	// host_project_registration.gcp_project field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/apihub_host_project_registration#host_project_registration_id ApihubHostProjectRegistration#host_project_registration_id}
	HostProjectRegistrationId *string `field:"required" json:"hostProjectRegistrationId" yaml:"hostProjectRegistrationId"`
	// Part of 'parent'. See documentation of 'projectsId'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/apihub_host_project_registration#location ApihubHostProjectRegistration#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/apihub_host_project_registration#id ApihubHostProjectRegistration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/apihub_host_project_registration#project ApihubHostProjectRegistration#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/apihub_host_project_registration#timeouts ApihubHostProjectRegistration#timeouts}
	Timeouts *ApihubHostProjectRegistrationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

