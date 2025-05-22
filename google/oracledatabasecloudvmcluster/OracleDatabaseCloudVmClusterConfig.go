// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oracledatabasecloudvmcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OracleDatabaseCloudVmClusterConfig struct {
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
	// CIDR range of the backup subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/oracle_database_cloud_vm_cluster#backup_subnet_cidr OracleDatabaseCloudVmCluster#backup_subnet_cidr}
	BackupSubnetCidr *string `field:"required" json:"backupSubnetCidr" yaml:"backupSubnetCidr"`
	// Network settings. CIDR to use for cluster IP allocation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/oracle_database_cloud_vm_cluster#cidr OracleDatabaseCloudVmCluster#cidr}
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// The ID of the VM Cluster to create.
	//
	// This value is restricted
	// to (^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$) and must be a maximum of 63
	// characters in length. The value must start with a letter and end with
	// a letter or a number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/oracle_database_cloud_vm_cluster#cloud_vm_cluster_id OracleDatabaseCloudVmCluster#cloud_vm_cluster_id}
	CloudVmClusterId *string `field:"required" json:"cloudVmClusterId" yaml:"cloudVmClusterId"`
	// The name of the Exadata Infrastructure resource on which VM cluster resource is created, in the following format: projects/{project}/locations/{region}/cloudExadataInfrastuctures/{cloud_extradata_infrastructure}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/oracle_database_cloud_vm_cluster#exadata_infrastructure OracleDatabaseCloudVmCluster#exadata_infrastructure}
	ExadataInfrastructure *string `field:"required" json:"exadataInfrastructure" yaml:"exadataInfrastructure"`
	// Resource ID segment making up resource 'name'. See documentation for resource type 'oracledatabase.googleapis.com/DbNode'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/oracle_database_cloud_vm_cluster#location OracleDatabaseCloudVmCluster#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of the VPC network. Format: projects/{project}/global/networks/{network}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/oracle_database_cloud_vm_cluster#network OracleDatabaseCloudVmCluster#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Whether Terraform will be prevented from destroying the cluster.
	//
	// Deleting this cluster via terraform destroy or terraform apply will only succeed if this field is false in the Terraform state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/oracle_database_cloud_vm_cluster#deletion_protection OracleDatabaseCloudVmCluster#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// User friendly name for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/oracle_database_cloud_vm_cluster#display_name OracleDatabaseCloudVmCluster#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/oracle_database_cloud_vm_cluster#id OracleDatabaseCloudVmCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels or tags associated with the VM Cluster.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/oracle_database_cloud_vm_cluster#labels OracleDatabaseCloudVmCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/oracle_database_cloud_vm_cluster#project OracleDatabaseCloudVmCluster#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/oracle_database_cloud_vm_cluster#properties OracleDatabaseCloudVmCluster#properties}
	Properties *OracleDatabaseCloudVmClusterProperties `field:"optional" json:"properties" yaml:"properties"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/oracle_database_cloud_vm_cluster#timeouts OracleDatabaseCloudVmCluster#timeouts}
	Timeouts *OracleDatabaseCloudVmClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

