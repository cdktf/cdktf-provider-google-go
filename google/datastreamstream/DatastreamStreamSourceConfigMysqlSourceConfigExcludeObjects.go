// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastreamstream


type DatastreamStreamSourceConfigMysqlSourceConfigExcludeObjects struct {
	// mysql_databases block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/datastream_stream#mysql_databases DatastreamStream#mysql_databases}
	MysqlDatabases interface{} `field:"required" json:"mysqlDatabases" yaml:"mysqlDatabases"`
}

