// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocworkflowtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v12/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v12/dataprocworkflowtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference interface {
	cdktf.ComplexObject
	Accelerators() DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAcceleratorsList
	AcceleratorsInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DiskConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfigOutputReference
	DiskConfigInput() *DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig
	// Experimental.
	Fqn() *string
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	InstanceNames() *[]*string
	InternalValue() *DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfig
	SetInternalValue(val *DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfig)
	IsPreemptible() cdktf.IResolvable
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	ManagedGroupConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList
	MinCpuPlatform() *string
	SetMinCpuPlatform(val *string)
	MinCpuPlatformInput() *string
	NumInstances() *float64
	SetNumInstances(val *float64)
	NumInstancesInput() *float64
	Preemptibility() *string
	SetPreemptibility(val *string)
	PreemptibilityInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutAccelerators(value interface{})
	PutDiskConfig(value *DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig)
	ResetAccelerators()
	ResetDiskConfig()
	ResetImage()
	ResetMachineType()
	ResetMinCpuPlatform()
	ResetNumInstances()
	ResetPreemptibility()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference
type jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) Accelerators() DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAcceleratorsList {
	var returns DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAcceleratorsList
	_jsii_.Get(
		j,
		"accelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) AcceleratorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceleratorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) DiskConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfigOutputReference {
	var returns DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfigOutputReference
	_jsii_.Get(
		j,
		"diskConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) DiskConfigInput() *DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig {
	var returns *DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig
	_jsii_.Get(
		j,
		"diskConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) InstanceNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) InternalValue() *DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfig {
	var returns *DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) IsPreemptible() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isPreemptible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) ManagedGroupConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList {
	var returns DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList
	_jsii_.Get(
		j,
		"managedGroupConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) MinCpuPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) MinCpuPlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) NumInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) NumInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) Preemptibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preemptibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) PreemptibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preemptibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocWorkflowTemplate.DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference_Override(d DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocWorkflowTemplate.DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference)SetInternalValue(val *DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference)SetMinCpuPlatform(val *string) {
	if err := j.validateSetMinCpuPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCpuPlatform",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference)SetNumInstances(val *float64) {
	if err := j.validateSetNumInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numInstances",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference)SetPreemptibility(val *string) {
	if err := j.validateSetPreemptibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preemptibility",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) PutAccelerators(value interface{}) {
	if err := d.validatePutAcceleratorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAccelerators",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) PutDiskConfig(value *DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig) {
	if err := d.validatePutDiskConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDiskConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) ResetAccelerators() {
	_jsii_.InvokeVoid(
		d,
		"resetAccelerators",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) ResetDiskConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetDiskConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		d,
		"resetImage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) ResetMachineType() {
	_jsii_.InvokeVoid(
		d,
		"resetMachineType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) ResetMinCpuPlatform() {
	_jsii_.InvokeVoid(
		d,
		"resetMinCpuPlatform",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) ResetNumInstances() {
	_jsii_.InvokeVoid(
		d,
		"resetNumInstances",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) ResetPreemptibility() {
	_jsii_.InvokeVoid(
		d,
		"resetPreemptibility",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

