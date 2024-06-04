// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apphubservice


type ApphubServiceAttributes struct {
	// business_owners block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/apphub_service#business_owners ApphubService#business_owners}
	BusinessOwners interface{} `field:"optional" json:"businessOwners" yaml:"businessOwners"`
	// criticality block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/apphub_service#criticality ApphubService#criticality}
	Criticality *ApphubServiceAttributesCriticality `field:"optional" json:"criticality" yaml:"criticality"`
	// developer_owners block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/apphub_service#developer_owners ApphubService#developer_owners}
	DeveloperOwners interface{} `field:"optional" json:"developerOwners" yaml:"developerOwners"`
	// environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/apphub_service#environment ApphubService#environment}
	Environment *ApphubServiceAttributesEnvironment `field:"optional" json:"environment" yaml:"environment"`
	// operator_owners block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/apphub_service#operator_owners ApphubService#operator_owners}
	OperatorOwners interface{} `field:"optional" json:"operatorOwners" yaml:"operatorOwners"`
}

