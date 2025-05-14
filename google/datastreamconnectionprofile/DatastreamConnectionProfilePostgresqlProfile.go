// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastreamconnectionprofile


type DatastreamConnectionProfilePostgresqlProfile struct {
	// Database for the PostgreSQL connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/datastream_connection_profile#database DatastreamConnectionProfile#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Hostname for the PostgreSQL connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/datastream_connection_profile#hostname DatastreamConnectionProfile#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Username for the PostgreSQL connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/datastream_connection_profile#username DatastreamConnectionProfile#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Password for the PostgreSQL connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/datastream_connection_profile#password DatastreamConnectionProfile#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Port for the PostgreSQL connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/datastream_connection_profile#port DatastreamConnectionProfile#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// A reference to a Secret Manager resource name storing the user's password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/datastream_connection_profile#secret_manager_stored_password DatastreamConnectionProfile#secret_manager_stored_password}
	SecretManagerStoredPassword *string `field:"optional" json:"secretManagerStoredPassword" yaml:"secretManagerStoredPassword"`
}

