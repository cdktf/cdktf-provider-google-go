// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alloydbinstance


type AlloydbInstanceClientConnectionConfig struct {
	// Configuration to enforce connectors only (ex: AuthProxy) connections to the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/alloydb_instance#require_connectors AlloydbInstance#require_connectors}
	RequireConnectors interface{} `field:"optional" json:"requireConnectors" yaml:"requireConnectors"`
	// ssl_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/alloydb_instance#ssl_config AlloydbInstance#ssl_config}
	SslConfig *AlloydbInstanceClientConnectionConfigSslConfig `field:"optional" json:"sslConfig" yaml:"sslConfig"`
}

