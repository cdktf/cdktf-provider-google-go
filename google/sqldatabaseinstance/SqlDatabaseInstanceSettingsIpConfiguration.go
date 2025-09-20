// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqldatabaseinstance


type SqlDatabaseInstanceSettingsIpConfiguration struct {
	// The name of the allocated ip range for the private ip CloudSQL instance.
	//
	// For example: "google-managed-services-default". If set, the instance ip will be created in the allocated range. The range name must comply with RFC 1035. Specifically, the name must be 1-63 characters long and match the regular expression [a-z]([-a-z0-9]*[a-z0-9])?.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/sql_database_instance#allocated_ip_range SqlDatabaseInstance#allocated_ip_range}
	AllocatedIpRange *string `field:"optional" json:"allocatedIpRange" yaml:"allocatedIpRange"`
	// authorized_networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/sql_database_instance#authorized_networks SqlDatabaseInstance#authorized_networks}
	AuthorizedNetworks interface{} `field:"optional" json:"authorizedNetworks" yaml:"authorizedNetworks"`
	// The custom subject alternative names for an instance with "CUSTOMER_MANAGED_CAS_CA" as the "server_ca_mode".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/sql_database_instance#custom_subject_alternative_names SqlDatabaseInstance#custom_subject_alternative_names}
	CustomSubjectAlternativeNames *[]*string `field:"optional" json:"customSubjectAlternativeNames" yaml:"customSubjectAlternativeNames"`
	// Whether Google Cloud services such as BigQuery are allowed to access data in this Cloud SQL instance over a private IP connection.
	//
	// SQLSERVER database type is not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/sql_database_instance#enable_private_path_for_google_cloud_services SqlDatabaseInstance#enable_private_path_for_google_cloud_services}
	EnablePrivatePathForGoogleCloudServices interface{} `field:"optional" json:"enablePrivatePathForGoogleCloudServices" yaml:"enablePrivatePathForGoogleCloudServices"`
	// Whether this Cloud SQL instance should be assigned a public IPV4 address.
	//
	// At least ipv4_enabled must be enabled or a private_network must be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/sql_database_instance#ipv4_enabled SqlDatabaseInstance#ipv4_enabled}
	Ipv4Enabled interface{} `field:"optional" json:"ipv4Enabled" yaml:"ipv4Enabled"`
	// The VPC network from which the Cloud SQL instance is accessible for private IP.
	//
	// For example, projects/myProject/global/networks/default. Specifying a network enables private IP. At least ipv4_enabled must be enabled or a private_network must be configured. This setting can be updated, but it cannot be removed after it is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/sql_database_instance#private_network SqlDatabaseInstance#private_network}
	PrivateNetwork *string `field:"optional" json:"privateNetwork" yaml:"privateNetwork"`
	// psc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/sql_database_instance#psc_config SqlDatabaseInstance#psc_config}
	PscConfig interface{} `field:"optional" json:"pscConfig" yaml:"pscConfig"`
	// Specify how the server certificate's Certificate Authority is hosted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/sql_database_instance#server_ca_mode SqlDatabaseInstance#server_ca_mode}
	ServerCaMode *string `field:"optional" json:"serverCaMode" yaml:"serverCaMode"`
	// The resource name of the server CA pool for an instance with "CUSTOMER_MANAGED_CAS_CA" as the "server_ca_mode".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/sql_database_instance#server_ca_pool SqlDatabaseInstance#server_ca_pool}
	ServerCaPool *string `field:"optional" json:"serverCaPool" yaml:"serverCaPool"`
	// Specify how SSL connection should be enforced in DB connections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/sql_database_instance#ssl_mode SqlDatabaseInstance#ssl_mode}
	SslMode *string `field:"optional" json:"sslMode" yaml:"sslMode"`
}

