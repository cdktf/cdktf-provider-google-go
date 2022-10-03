package sqlsourcerepresentationinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SqlSourceRepresentationInstanceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The MySQL version running on your source database server. Possible values: ["MYSQL_5_5", "MYSQL_5_6", "MYSQL_5_7", "MYSQL_8_0"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_source_representation_instance#database_version SqlSourceRepresentationInstance#database_version}
	DatabaseVersion *string `field:"required" json:"databaseVersion" yaml:"databaseVersion"`
	// The externally accessible IPv4 address for the source database server.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_source_representation_instance#host SqlSourceRepresentationInstance#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// The name of the source representation instance. Use any valid Cloud SQL instance name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_source_representation_instance#name SqlSourceRepresentationInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_source_representation_instance#id SqlSourceRepresentationInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The externally accessible port for the source database server. Defaults to 3306.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_source_representation_instance#port SqlSourceRepresentationInstance#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_source_representation_instance#project SqlSourceRepresentationInstance#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The Region in which the created instance should reside. If it is not provided, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_source_representation_instance#region SqlSourceRepresentationInstance#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_source_representation_instance#timeouts SqlSourceRepresentationInstance#timeouts}
	Timeouts *SqlSourceRepresentationInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

