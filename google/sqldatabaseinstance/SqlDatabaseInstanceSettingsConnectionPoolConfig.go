// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqldatabaseinstance


type SqlDatabaseInstanceSettingsConnectionPoolConfig struct {
	// Whether Managed Connection Pool is enabled for this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/sql_database_instance#connection_pooling_enabled SqlDatabaseInstance#connection_pooling_enabled}
	ConnectionPoolingEnabled interface{} `field:"optional" json:"connectionPoolingEnabled" yaml:"connectionPoolingEnabled"`
	// flags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/sql_database_instance#flags SqlDatabaseInstance#flags}
	Flags interface{} `field:"optional" json:"flags" yaml:"flags"`
}

