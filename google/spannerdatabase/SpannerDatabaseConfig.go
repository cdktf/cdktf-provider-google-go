// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spannerdatabase

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpannerDatabaseConfig struct {
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
	// The instance to create the database on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/spanner_database#instance SpannerDatabase#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// A unique identifier for the database, which cannot be changed after the instance is created.
	//
	// Values are of the form '[a-z][-_a-z0-9]*[a-z0-9]'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/spanner_database#name SpannerDatabase#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The dialect of the Cloud Spanner Database. If it is not provided, "GOOGLE_STANDARD_SQL" will be used. Possible values: ["GOOGLE_STANDARD_SQL", "POSTGRESQL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/spanner_database#database_dialect SpannerDatabase#database_dialect}
	DatabaseDialect *string `field:"optional" json:"databaseDialect" yaml:"databaseDialect"`
	// An optional list of DDL statements to run inside the newly created database.
	//
	// Statements can create tables, indexes, etc. These statements
	// execute atomically with the creation of the database: if there is an
	// error in any statement, the database is not created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/spanner_database#ddl SpannerDatabase#ddl}
	Ddl *[]*string `field:"optional" json:"ddl" yaml:"ddl"`
	// Whether Terraform will be prevented from destroying the database.
	//
	// Defaults to true.
	// When a'terraform destroy' or 'terraform apply' would delete the database,
	// the command will fail if this field is not set to false in Terraform state.
	// When the field is set to true or unset in Terraform state, a 'terraform apply'
	// or 'terraform destroy' that would delete the database will fail.
	// When the field is set to false, deleting the database is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/spanner_database#deletion_protection SpannerDatabase#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Whether drop protection is enabled for this database.
	//
	// Defaults to false.
	// Drop protection is different from
	// the "deletion_protection" attribute in the following ways:
	// (1) "deletion_protection" only protects the database from deletions in Terraform.
	// whereas setting “enableDropProtection” to true protects the database from deletions in all interfaces.
	// (2) Setting "enableDropProtection" to true also prevents the deletion of the parent instance containing the database.
	// "deletion_protection" attribute does not provide protection against the deletion of the parent instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/spanner_database#enable_drop_protection SpannerDatabase#enable_drop_protection}
	EnableDropProtection interface{} `field:"optional" json:"enableDropProtection" yaml:"enableDropProtection"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/spanner_database#encryption_config SpannerDatabase#encryption_config}
	EncryptionConfig *SpannerDatabaseEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/spanner_database#id SpannerDatabase#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/spanner_database#project SpannerDatabase#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/spanner_database#timeouts SpannerDatabase#timeouts}
	Timeouts *SpannerDatabaseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The retention period for the database.
	//
	// The retention period must be between 1 hour
	// and 7 days, and can be specified in days, hours, minutes, or seconds. For example,
	// the values 1d, 24h, 1440m, and 86400s are equivalent. Default value is 1h.
	// If this property is used, you must avoid adding new DDL statements to 'ddl' that
	// update the database's version_retention_period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/spanner_database#version_retention_period SpannerDatabase#version_retention_period}
	VersionRetentionPeriod *string `field:"optional" json:"versionRetentionPeriod" yaml:"versionRetentionPeriod"`
}

