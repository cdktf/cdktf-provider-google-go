package bigqueryconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigqueryConnectionConfig struct {
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
	// aws block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_connection#aws BigqueryConnection#aws}
	Aws *BigqueryConnectionAws `field:"optional" json:"aws" yaml:"aws"`
	// azure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_connection#azure BigqueryConnection#azure}
	Azure *BigqueryConnectionAzure `field:"optional" json:"azure" yaml:"azure"`
	// cloud_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_connection#cloud_resource BigqueryConnection#cloud_resource}
	CloudResource *BigqueryConnectionCloudResource `field:"optional" json:"cloudResource" yaml:"cloudResource"`
	// cloud_spanner block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_connection#cloud_spanner BigqueryConnection#cloud_spanner}
	CloudSpanner *BigqueryConnectionCloudSpanner `field:"optional" json:"cloudSpanner" yaml:"cloudSpanner"`
	// cloud_sql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_connection#cloud_sql BigqueryConnection#cloud_sql}
	CloudSql *BigqueryConnectionCloudSql `field:"optional" json:"cloudSql" yaml:"cloudSql"`
	// Optional connection id that should be assigned to the created connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_connection#connection_id BigqueryConnection#connection_id}
	ConnectionId *string `field:"optional" json:"connectionId" yaml:"connectionId"`
	// A descriptive description for the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_connection#description BigqueryConnection#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A descriptive name for the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_connection#friendly_name BigqueryConnection#friendly_name}
	FriendlyName *string `field:"optional" json:"friendlyName" yaml:"friendlyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_connection#id BigqueryConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The geographic location where the connection should reside.
	//
	// Cloud SQL instance must be in the same location as the connection
	// with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
	// Examples: US, EU, asia-northeast1, us-central1, europe-west1.
	// Spanner Connections same as spanner region
	// AWS allowed regions are aws-us-east-1
	// Azure allowed regions are azure-eastus2
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_connection#location BigqueryConnection#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_connection#project BigqueryConnection#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_connection#timeouts BigqueryConnection#timeouts}
	Timeouts *BigqueryConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

