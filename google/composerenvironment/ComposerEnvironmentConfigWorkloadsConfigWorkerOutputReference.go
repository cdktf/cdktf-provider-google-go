// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package composerenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/composerenvironment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference interface {
	cdktf.ComplexObject
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
	Cpu() *float64
	SetCpu(val *float64)
	CpuInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ComposerEnvironmentConfigWorkloadsConfigWorker
	SetInternalValue(val *ComposerEnvironmentConfigWorkloadsConfigWorker)
	MaxCount() *float64
	SetMaxCount(val *float64)
	MaxCountInput() *float64
	MemoryGb() *float64
	SetMemoryGb(val *float64)
	MemoryGbInput() *float64
	MinCount() *float64
	SetMinCount(val *float64)
	MinCountInput() *float64
	StorageGb() *float64
	SetStorageGb(val *float64)
	StorageGbInput() *float64
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
	ResetCpu()
	ResetMaxCount()
	ResetMemoryGb()
	ResetMinCount()
	ResetStorageGb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference
type jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) Cpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) CpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) InternalValue() *ComposerEnvironmentConfigWorkloadsConfigWorker {
	var returns *ComposerEnvironmentConfigWorkloadsConfigWorker
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) MaxCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) MaxCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) MemoryGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) MemoryGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) MinCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) MinCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) StorageGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) StorageGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference {
	_init_.Initialize()

	if err := validateNewComposerEnvironmentConfigWorkloadsConfigWorkerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.composerEnvironment.ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference_Override(c ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.composerEnvironment.ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference)SetCpu(val *float64) {
	if err := j.validateSetCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpu",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference)SetInternalValue(val *ComposerEnvironmentConfigWorkloadsConfigWorker) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference)SetMaxCount(val *float64) {
	if err := j.validateSetMaxCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCount",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference)SetMemoryGb(val *float64) {
	if err := j.validateSetMemoryGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryGb",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference)SetMinCount(val *float64) {
	if err := j.validateSetMinCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCount",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference)SetStorageGb(val *float64) {
	if err := j.validateSetStorageGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageGb",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) ResetCpu() {
	_jsii_.InvokeVoid(
		c,
		"resetCpu",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) ResetMaxCount() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) ResetMemoryGb() {
	_jsii_.InvokeVoid(
		c,
		"resetMemoryGb",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) ResetMinCount() {
	_jsii_.InvokeVoid(
		c,
		"resetMinCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) ResetStorageGb() {
	_jsii_.InvokeVoid(
		c,
		"resetStorageGb",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

