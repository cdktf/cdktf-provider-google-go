package sqlsourcerepresentationinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SqlSourceRepresentationInstanceConfig struct {
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
	// The MySQL version running on your source database server.
	//
	// Possible values: ["MYSQL_5_6", "MYSQL_5_7", "MYSQL_8_0", "POSTGRES_9_6", "POSTGRES_10", "POSTGRES_11", "POSTGRES_12", "POSTGRES_13", "POSTGRES_14"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/sql_source_representation_instance#database_version SqlSourceRepresentationInstance#database_version}
	DatabaseVersion *string `field:"required" json:"databaseVersion" yaml:"databaseVersion"`
	// The externally accessible IPv4 address for the source database server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/sql_source_representation_instance#host SqlSourceRepresentationInstance#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// The name of the source representation instance. Use any valid Cloud SQL instance name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/sql_source_representation_instance#name SqlSourceRepresentationInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The CA certificate on the external server. Include only if SSL/TLS is used on the external server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/sql_source_representation_instance#ca_certificate SqlSourceRepresentationInstance#ca_certificate}
	CaCertificate *string `field:"optional" json:"caCertificate" yaml:"caCertificate"`
	// The client certificate on the external server.
	//
	// Required only for server-client authentication. Include only if SSL/TLS is used on the external server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/sql_source_representation_instance#client_certificate SqlSourceRepresentationInstance#client_certificate}
	ClientCertificate *string `field:"optional" json:"clientCertificate" yaml:"clientCertificate"`
	// The private key file for the client certificate on the external server.
	//
	// Required only for server-client authentication. Include only if SSL/TLS is used on the external server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/sql_source_representation_instance#client_key SqlSourceRepresentationInstance#client_key}
	ClientKey *string `field:"optional" json:"clientKey" yaml:"clientKey"`
	// A file in the bucket that contains the data from the external server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/sql_source_representation_instance#dump_file_path SqlSourceRepresentationInstance#dump_file_path}
	DumpFilePath *string `field:"optional" json:"dumpFilePath" yaml:"dumpFilePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/sql_source_representation_instance#id SqlSourceRepresentationInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The password for the replication user account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/sql_source_representation_instance#password SqlSourceRepresentationInstance#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The externally accessible port for the source database server. Defaults to 3306.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/sql_source_representation_instance#port SqlSourceRepresentationInstance#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/sql_source_representation_instance#project SqlSourceRepresentationInstance#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The Region in which the created instance should reside. If it is not provided, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/sql_source_representation_instance#region SqlSourceRepresentationInstance#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/sql_source_representation_instance#timeouts SqlSourceRepresentationInstance#timeouts}
	Timeouts *SqlSourceRepresentationInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The replication user account on the external server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/sql_source_representation_instance#username SqlSourceRepresentationInstance#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

