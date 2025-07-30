// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasemigrationservicemigrationjob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseMigrationServiceMigrationJobConfig struct {
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
	// The name of the destination connection profile resource in the form of projects/{project}/locations/{location}/connectionProfiles/{destinationConnectionProfile}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/database_migration_service_migration_job#destination DatabaseMigrationServiceMigrationJob#destination}
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// The ID of the migration job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/database_migration_service_migration_job#migration_job_id DatabaseMigrationServiceMigrationJob#migration_job_id}
	MigrationJobId *string `field:"required" json:"migrationJobId" yaml:"migrationJobId"`
	// The name of the source connection profile resource in the form of projects/{project}/locations/{location}/connectionProfiles/{sourceConnectionProfile}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/database_migration_service_migration_job#source DatabaseMigrationServiceMigrationJob#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// The type of the migration job. Possible values: ["ONE_TIME", "CONTINUOUS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/database_migration_service_migration_job#type DatabaseMigrationServiceMigrationJob#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The migration job display name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/database_migration_service_migration_job#display_name DatabaseMigrationServiceMigrationJob#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// dump_flags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/database_migration_service_migration_job#dump_flags DatabaseMigrationServiceMigrationJob#dump_flags}
	DumpFlags *DatabaseMigrationServiceMigrationJobDumpFlags `field:"optional" json:"dumpFlags" yaml:"dumpFlags"`
	// The path to the dump file in Google Cloud Storage, in the format: (gs://[BUCKET_NAME]/[OBJECT_NAME]).
	//
	// This field and the "dump_flags" field are mutually exclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/database_migration_service_migration_job#dump_path DatabaseMigrationServiceMigrationJob#dump_path}
	DumpPath *string `field:"optional" json:"dumpPath" yaml:"dumpPath"`
	// The type of the data dump. Supported for MySQL to CloudSQL for MySQL migrations only. Possible values: ["LOGICAL", "PHYSICAL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/database_migration_service_migration_job#dump_type DatabaseMigrationServiceMigrationJob#dump_type}
	DumpType *string `field:"optional" json:"dumpType" yaml:"dumpType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/database_migration_service_migration_job#id DatabaseMigrationServiceMigrationJob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The resource labels for migration job to use to annotate any related underlying resources such as Compute Engine VMs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/database_migration_service_migration_job#labels DatabaseMigrationServiceMigrationJob#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The location where the migration job should reside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/database_migration_service_migration_job#location DatabaseMigrationServiceMigrationJob#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// performance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/database_migration_service_migration_job#performance_config DatabaseMigrationServiceMigrationJob#performance_config}
	PerformanceConfig *DatabaseMigrationServiceMigrationJobPerformanceConfig `field:"optional" json:"performanceConfig" yaml:"performanceConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/database_migration_service_migration_job#project DatabaseMigrationServiceMigrationJob#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// reverse_ssh_connectivity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/database_migration_service_migration_job#reverse_ssh_connectivity DatabaseMigrationServiceMigrationJob#reverse_ssh_connectivity}
	ReverseSshConnectivity *DatabaseMigrationServiceMigrationJobReverseSshConnectivity `field:"optional" json:"reverseSshConnectivity" yaml:"reverseSshConnectivity"`
	// static_ip_connectivity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/database_migration_service_migration_job#static_ip_connectivity DatabaseMigrationServiceMigrationJob#static_ip_connectivity}
	StaticIpConnectivity *DatabaseMigrationServiceMigrationJobStaticIpConnectivity `field:"optional" json:"staticIpConnectivity" yaml:"staticIpConnectivity"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/database_migration_service_migration_job#timeouts DatabaseMigrationServiceMigrationJob#timeouts}
	Timeouts *DatabaseMigrationServiceMigrationJobTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// vpc_peering_connectivity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/database_migration_service_migration_job#vpc_peering_connectivity DatabaseMigrationServiceMigrationJob#vpc_peering_connectivity}
	VpcPeeringConnectivity *DatabaseMigrationServiceMigrationJobVpcPeeringConnectivity `field:"optional" json:"vpcPeeringConnectivity" yaml:"vpcPeeringConnectivity"`
}

