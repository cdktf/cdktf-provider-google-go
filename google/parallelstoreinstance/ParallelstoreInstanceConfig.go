// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package parallelstoreinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ParallelstoreInstanceConfig struct {
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
	// Required. Immutable. Storage capacity of Parallelstore instance in Gibibytes (GiB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/parallelstore_instance#capacity_gib ParallelstoreInstance#capacity_gib}
	CapacityGib *string `field:"required" json:"capacityGib" yaml:"capacityGib"`
	// The logical name of the Parallelstore instance in the user project with the following restrictions:   * Must contain only lowercase letters, numbers, and hyphens.
	//
	// * Must start with a letter.
	//   * Must be between 1-63 characters.
	//   * Must end with a number or a letter.
	//   * Must be unique within the customer project/ location
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/parallelstore_instance#instance_id ParallelstoreInstance#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Part of 'parent'. See documentation of 'projectsId'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/parallelstore_instance#location ParallelstoreInstance#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Parallelstore Instance deployment type.   Possible values:   DEPLOYMENT_TYPE_UNSPECIFIED   SCRATCH   PERSISTENT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/parallelstore_instance#deployment_type ParallelstoreInstance#deployment_type}
	DeploymentType *string `field:"optional" json:"deploymentType" yaml:"deploymentType"`
	// The description of the instance. 2048 characters or less.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/parallelstore_instance#description ParallelstoreInstance#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Stripe level for directories.
	//
	// MIN when directory has a small number of files.
	// MAX when directory has a large number of files.
	//   Possible values:
	//   DIRECTORY_STRIPE_LEVEL_UNSPECIFIED
	//   DIRECTORY_STRIPE_LEVEL_MIN
	//   DIRECTORY_STRIPE_LEVEL_BALANCED
	//   DIRECTORY_STRIPE_LEVEL_MAX
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/parallelstore_instance#directory_stripe_level ParallelstoreInstance#directory_stripe_level}
	DirectoryStripeLevel *string `field:"optional" json:"directoryStripeLevel" yaml:"directoryStripeLevel"`
	// Stripe level for files.
	//
	// MIN better suited for small size files.
	// MAX higher throughput performance for larger files.
	//   Possible values:
	//   FILE_STRIPE_LEVEL_UNSPECIFIED
	//   FILE_STRIPE_LEVEL_MIN
	//   FILE_STRIPE_LEVEL_BALANCED
	//   FILE_STRIPE_LEVEL_MAX
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/parallelstore_instance#file_stripe_level ParallelstoreInstance#file_stripe_level}
	FileStripeLevel *string `field:"optional" json:"fileStripeLevel" yaml:"fileStripeLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/parallelstore_instance#id ParallelstoreInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Cloud Labels are a flexible and lightweight mechanism for organizing cloud resources into groups that reflect a customer's organizational needs and deployment strategies.
	//
	// Cloud Labels can be used to filter collections
	// of resources. They can be used to control how resource metrics are aggregated.
	// And they can be used as arguments to policy management rules (e.g. route, firewall,
	// load balancing, etc.).
	//
	// * Label keys must be between 1 and 63 characters long and must conform to
	//  the following regular expression: 'a-z{0,62}'.
	// * Label values must be between 0 and 63 characters long and must conform
	//  to the regular expression '[a-z0-9_-]{0,63}'.
	// * No more than 64 labels can be associated with a given resource.
	//
	// See https://goo.gl/xmQnxf for more information on and examples of labels.
	//
	// If you plan to use labels in your own code, please note that additional
	// characters may be allowed in the future. Therefore, you are advised to use
	// an internal label representation, such as JSON, which doesn't rely upon
	// specific characters being disallowed.  For example, representing labels
	// as the string:  'name + "_" + value' would prove problematic if we were to
	// allow '"_"' in a future release. "
	//
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/parallelstore_instance#labels ParallelstoreInstance#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Immutable. The name of the Google Compute Engine [VPC network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/parallelstore_instance#network ParallelstoreInstance#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/parallelstore_instance#project ParallelstoreInstance#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Immutable.
	//
	// Contains the id of the allocated IP address range
	// associated with the private service access connection for example, \"test-default\"
	// associated with IP range 10.0.0.0/29. If no range id is provided all ranges will
	// be considered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/parallelstore_instance#reserved_ip_range ParallelstoreInstance#reserved_ip_range}
	ReservedIpRange *string `field:"optional" json:"reservedIpRange" yaml:"reservedIpRange"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/parallelstore_instance#timeouts ParallelstoreInstance#timeouts}
	Timeouts *ParallelstoreInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

