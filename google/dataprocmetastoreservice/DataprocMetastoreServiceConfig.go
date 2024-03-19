// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocmetastoreservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocMetastoreServiceConfig struct {
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
	// The ID of the metastore service.
	//
	// The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
	// and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
	// 3 and 63 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/dataproc_metastore_service#service_id DataprocMetastoreService#service_id}
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
	// The database type that the Metastore service stores its data. Default value: "MYSQL" Possible values: ["MYSQL", "SPANNER"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/dataproc_metastore_service#database_type DataprocMetastoreService#database_type}
	DatabaseType *string `field:"optional" json:"databaseType" yaml:"databaseType"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/dataproc_metastore_service#encryption_config DataprocMetastoreService#encryption_config}
	EncryptionConfig *DataprocMetastoreServiceEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// hive_metastore_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/dataproc_metastore_service#hive_metastore_config DataprocMetastoreService#hive_metastore_config}
	HiveMetastoreConfig *DataprocMetastoreServiceHiveMetastoreConfig `field:"optional" json:"hiveMetastoreConfig" yaml:"hiveMetastoreConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/dataproc_metastore_service#id DataprocMetastoreService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User-defined labels for the metastore service.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/dataproc_metastore_service#labels DataprocMetastoreService#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The location where the metastore service should reside. The default value is 'global'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/dataproc_metastore_service#location DataprocMetastoreService#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/dataproc_metastore_service#maintenance_window DataprocMetastoreService#maintenance_window}
	MaintenanceWindow *DataprocMetastoreServiceMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// metadata_integration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/dataproc_metastore_service#metadata_integration DataprocMetastoreService#metadata_integration}
	MetadataIntegration *DataprocMetastoreServiceMetadataIntegration `field:"optional" json:"metadataIntegration" yaml:"metadataIntegration"`
	// The relative resource name of the VPC network on which the instance can be accessed.
	//
	// It is specified in the following form:
	//
	// "projects/{projectNumber}/global/networks/{network_id}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/dataproc_metastore_service#network DataprocMetastoreService#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/dataproc_metastore_service#network_config DataprocMetastoreService#network_config}
	NetworkConfig *DataprocMetastoreServiceNetworkConfig `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// The TCP port at which the metastore service is reached. Default: 9083.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/dataproc_metastore_service#port DataprocMetastoreService#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/dataproc_metastore_service#project DataprocMetastoreService#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The release channel of the service. If unspecified, defaults to 'STABLE'. Default value: "STABLE" Possible values: ["CANARY", "STABLE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/dataproc_metastore_service#release_channel DataprocMetastoreService#release_channel}
	ReleaseChannel *string `field:"optional" json:"releaseChannel" yaml:"releaseChannel"`
	// scaling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/dataproc_metastore_service#scaling_config DataprocMetastoreService#scaling_config}
	ScalingConfig *DataprocMetastoreServiceScalingConfig `field:"optional" json:"scalingConfig" yaml:"scalingConfig"`
	// telemetry_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/dataproc_metastore_service#telemetry_config DataprocMetastoreService#telemetry_config}
	TelemetryConfig *DataprocMetastoreServiceTelemetryConfig `field:"optional" json:"telemetryConfig" yaml:"telemetryConfig"`
	// The tier of the service. Possible values: ["DEVELOPER", "ENTERPRISE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/dataproc_metastore_service#tier DataprocMetastoreService#tier}
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/dataproc_metastore_service#timeouts DataprocMetastoreService#timeouts}
	Timeouts *DataprocMetastoreServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

