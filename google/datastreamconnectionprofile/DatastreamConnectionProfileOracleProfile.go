// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastreamconnectionprofile


type DatastreamConnectionProfileOracleProfile struct {
	// Database for the Oracle connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/datastream_connection_profile#database_service DatastreamConnectionProfile#database_service}
	DatabaseService *string `field:"required" json:"databaseService" yaml:"databaseService"`
	// Hostname for the Oracle connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/datastream_connection_profile#hostname DatastreamConnectionProfile#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Username for the Oracle connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/datastream_connection_profile#username DatastreamConnectionProfile#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Connection string attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/datastream_connection_profile#connection_attributes DatastreamConnectionProfile#connection_attributes}
	ConnectionAttributes *map[string]*string `field:"optional" json:"connectionAttributes" yaml:"connectionAttributes"`
	// Password for the Oracle connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/datastream_connection_profile#password DatastreamConnectionProfile#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Port for the Oracle connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/datastream_connection_profile#port DatastreamConnectionProfile#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// A reference to a Secret Manager resource name storing the user's password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/datastream_connection_profile#secret_manager_stored_password DatastreamConnectionProfile#secret_manager_stored_password}
	SecretManagerStoredPassword *string `field:"optional" json:"secretManagerStoredPassword" yaml:"secretManagerStoredPassword"`
}

