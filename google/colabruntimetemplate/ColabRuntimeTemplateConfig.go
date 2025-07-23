// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package colabruntimetemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ColabRuntimeTemplateConfig struct {
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
	// Required. The display name of the Runtime Template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/colab_runtime_template#display_name ColabRuntimeTemplate#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The location for the resource: https://cloud.google.com/colab/docs/locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/colab_runtime_template#location ColabRuntimeTemplate#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// data_persistent_disk_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/colab_runtime_template#data_persistent_disk_spec ColabRuntimeTemplate#data_persistent_disk_spec}
	DataPersistentDiskSpec *ColabRuntimeTemplateDataPersistentDiskSpec `field:"optional" json:"dataPersistentDiskSpec" yaml:"dataPersistentDiskSpec"`
	// The description of the Runtime Template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/colab_runtime_template#description ColabRuntimeTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// encryption_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/colab_runtime_template#encryption_spec ColabRuntimeTemplate#encryption_spec}
	EncryptionSpec *ColabRuntimeTemplateEncryptionSpec `field:"optional" json:"encryptionSpec" yaml:"encryptionSpec"`
	// euc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/colab_runtime_template#euc_config ColabRuntimeTemplate#euc_config}
	EucConfig *ColabRuntimeTemplateEucConfig `field:"optional" json:"eucConfig" yaml:"eucConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/colab_runtime_template#id ColabRuntimeTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// idle_shutdown_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/colab_runtime_template#idle_shutdown_config ColabRuntimeTemplate#idle_shutdown_config}
	IdleShutdownConfig *ColabRuntimeTemplateIdleShutdownConfig `field:"optional" json:"idleShutdownConfig" yaml:"idleShutdownConfig"`
	// Labels to identify and group the runtime template.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/colab_runtime_template#labels ColabRuntimeTemplate#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// machine_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/colab_runtime_template#machine_spec ColabRuntimeTemplate#machine_spec}
	MachineSpec *ColabRuntimeTemplateMachineSpec `field:"optional" json:"machineSpec" yaml:"machineSpec"`
	// The resource name of the Runtime Template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/colab_runtime_template#name ColabRuntimeTemplate#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// network_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/colab_runtime_template#network_spec ColabRuntimeTemplate#network_spec}
	NetworkSpec *ColabRuntimeTemplateNetworkSpec `field:"optional" json:"networkSpec" yaml:"networkSpec"`
	// Applies the given Compute Engine tags to the runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/colab_runtime_template#network_tags ColabRuntimeTemplate#network_tags}
	NetworkTags *[]*string `field:"optional" json:"networkTags" yaml:"networkTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/colab_runtime_template#project ColabRuntimeTemplate#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// shielded_vm_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/colab_runtime_template#shielded_vm_config ColabRuntimeTemplate#shielded_vm_config}
	ShieldedVmConfig *ColabRuntimeTemplateShieldedVmConfig `field:"optional" json:"shieldedVmConfig" yaml:"shieldedVmConfig"`
	// software_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/colab_runtime_template#software_config ColabRuntimeTemplate#software_config}
	SoftwareConfig *ColabRuntimeTemplateSoftwareConfig `field:"optional" json:"softwareConfig" yaml:"softwareConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/colab_runtime_template#timeouts ColabRuntimeTemplate#timeouts}
	Timeouts *ColabRuntimeTemplateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

