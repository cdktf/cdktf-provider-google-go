package logginglinkeddataset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoggingLinkedDatasetConfig struct {
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
	// The bucket to which the linked dataset is attached.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/logging_linked_dataset#bucket LoggingLinkedDataset#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The id of the linked dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/logging_linked_dataset#link_id LoggingLinkedDataset#link_id}
	LinkId *string `field:"required" json:"linkId" yaml:"linkId"`
	// bigquery_dataset block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/logging_linked_dataset#bigquery_dataset LoggingLinkedDataset#bigquery_dataset}
	BigqueryDataset interface{} `field:"optional" json:"bigqueryDataset" yaml:"bigqueryDataset"`
	// Describes this link. The maximum length of the description is 8000 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/logging_linked_dataset#description LoggingLinkedDataset#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/logging_linked_dataset#id LoggingLinkedDataset#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The location of the linked dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/logging_linked_dataset#location LoggingLinkedDataset#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The parent of the linked dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/logging_linked_dataset#parent LoggingLinkedDataset#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/logging_linked_dataset#timeouts LoggingLinkedDataset#timeouts}
	Timeouts *LoggingLinkedDatasetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

