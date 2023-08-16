package datagooglednsrecordset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleDnsRecordSetConfig struct {
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
	// The Name of the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/data-sources/dns_record_set#managed_zone DataGoogleDnsRecordSet#managed_zone}
	ManagedZone *string `field:"required" json:"managedZone" yaml:"managedZone"`
	// The DNS name for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/data-sources/dns_record_set#name DataGoogleDnsRecordSet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The identifier of a supported record type. See the list of Supported DNS record types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/data-sources/dns_record_set#type DataGoogleDnsRecordSet#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The ID of the project for the Google Cloud.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/data-sources/dns_record_set#project DataGoogleDnsRecordSet#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

