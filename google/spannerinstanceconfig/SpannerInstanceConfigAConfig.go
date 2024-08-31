// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spannerinstanceconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpannerInstanceConfigAConfig struct {
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
	// The name of this instance configuration as it appears in UIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/spanner_instance_config#display_name SpannerInstanceConfigA#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// replicas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/spanner_instance_config#replicas SpannerInstanceConfigA#replicas}
	Replicas interface{} `field:"required" json:"replicas" yaml:"replicas"`
	// Base configuration name, e.g. nam3, based on which this configuration is created. Only set for user managed configurations. baseConfig must refer to a configuration of type GOOGLE_MANAGED in the same project as this configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/spanner_instance_config#base_config SpannerInstanceConfigA#base_config}
	BaseConfig *string `field:"optional" json:"baseConfig" yaml:"baseConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/spanner_instance_config#id SpannerInstanceConfigA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/spanner_instance_config#labels SpannerInstanceConfigA#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// A unique identifier for the instance configuration. Values are of the form projects/<project>/instanceConfigs/[a-z][-a-z0-9]*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/spanner_instance_config#name SpannerInstanceConfigA#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/spanner_instance_config#project SpannerInstanceConfigA#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/spanner_instance_config#timeouts SpannerInstanceConfigA#timeouts}
	Timeouts *SpannerInstanceConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

