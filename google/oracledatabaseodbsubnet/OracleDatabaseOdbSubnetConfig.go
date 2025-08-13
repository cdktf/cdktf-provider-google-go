// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oracledatabaseodbsubnet

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OracleDatabaseOdbSubnetConfig struct {
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
	// The CIDR range of the subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/oracle_database_odb_subnet#cidr_range OracleDatabaseOdbSubnet#cidr_range}
	CidrRange *string `field:"required" json:"cidrRange" yaml:"cidrRange"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/oracle_database_odb_subnet#location OracleDatabaseOdbSubnet#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/oracle_database_odb_subnet#odbnetwork OracleDatabaseOdbSubnet#odbnetwork}
	Odbnetwork *string `field:"required" json:"odbnetwork" yaml:"odbnetwork"`
	// The ID of the OdbSubnet to create.
	//
	// This value is restricted
	// to (^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$) and must be a maximum of 63
	// characters in length. The value must start with a letter and end with
	// a letter or a number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/oracle_database_odb_subnet#odb_subnet_id OracleDatabaseOdbSubnet#odb_subnet_id}
	OdbSubnetId *string `field:"required" json:"odbSubnetId" yaml:"odbSubnetId"`
	// Purpose of the subnet. Possible values: CLIENT_SUBNET BACKUP_SUBNET.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/oracle_database_odb_subnet#purpose OracleDatabaseOdbSubnet#purpose}
	Purpose *string `field:"required" json:"purpose" yaml:"purpose"`
	// Whether or not to allow Terraform to destroy the instance.
	//
	// Unless this field is set to false in Terraform state, a terraform destroy or terraform apply that would delete the instance will fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/oracle_database_odb_subnet#deletion_protection OracleDatabaseOdbSubnet#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/oracle_database_odb_subnet#id OracleDatabaseOdbSubnet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels or tags associated with the resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/oracle_database_odb_subnet#labels OracleDatabaseOdbSubnet#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/oracle_database_odb_subnet#project OracleDatabaseOdbSubnet#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/oracle_database_odb_subnet#timeouts OracleDatabaseOdbSubnet#timeouts}
	Timeouts *OracleDatabaseOdbSubnetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

