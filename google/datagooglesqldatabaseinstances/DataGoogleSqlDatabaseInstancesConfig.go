package datagooglesqldatabaseinstances

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleSqlDatabaseInstancesConfig struct {
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
	// To filter out the database instances which are of the specified database version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/data-sources/sql_database_instances#database_version DataGoogleSqlDatabaseInstances#database_version}
	DatabaseVersion *string `field:"optional" json:"databaseVersion" yaml:"databaseVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/data-sources/sql_database_instances#id DataGoogleSqlDatabaseInstances#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Project ID of the project that contains the instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/data-sources/sql_database_instances#project DataGoogleSqlDatabaseInstances#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// To filter out the database instances which are located in this specified region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/data-sources/sql_database_instances#region DataGoogleSqlDatabaseInstances#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// To filter out the database instances based on the current state of the database instance, valid values include : "SQL_INSTANCE_STATE_UNSPECIFIED", "RUNNABLE", "SUSPENDED", "PENDING_DELETE", "PENDING_CREATE", "MAINTENANCE" and "FAILED".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/data-sources/sql_database_instances#state DataGoogleSqlDatabaseInstances#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// To filter out the database instances based on the machine type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/data-sources/sql_database_instances#tier DataGoogleSqlDatabaseInstances#tier}
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
	// To filter out the database instances which are located in this specified zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/data-sources/sql_database_instances#zone DataGoogleSqlDatabaseInstances#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

