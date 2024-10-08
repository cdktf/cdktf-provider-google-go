// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package endpointsserviceconsumersiammember

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EndpointsServiceConsumersIamMemberConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/endpoints_service_consumers_iam_member#consumer_project EndpointsServiceConsumersIamMember#consumer_project}.
	ConsumerProject *string `field:"required" json:"consumerProject" yaml:"consumerProject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/endpoints_service_consumers_iam_member#member EndpointsServiceConsumersIamMember#member}.
	Member *string `field:"required" json:"member" yaml:"member"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/endpoints_service_consumers_iam_member#role EndpointsServiceConsumersIamMember#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/endpoints_service_consumers_iam_member#service_name EndpointsServiceConsumersIamMember#service_name}.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/endpoints_service_consumers_iam_member#condition EndpointsServiceConsumersIamMember#condition}
	Condition *EndpointsServiceConsumersIamMemberCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/endpoints_service_consumers_iam_member#id EndpointsServiceConsumersIamMember#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

