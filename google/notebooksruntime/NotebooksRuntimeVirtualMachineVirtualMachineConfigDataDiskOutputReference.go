// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notebooksruntime

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/notebooksruntime/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference interface {
	cdktf.ComplexObject
	AutoDelete() cdktf.IResolvable
	Boot() cdktf.IResolvable
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
	DeviceName() *string
	// Experimental.
	Fqn() *string
	GuestOsFeatures() *[]*string
	Index() *float64
	InitializeParams() NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParamsOutputReference
	InitializeParamsInput() *NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParams
	Interface() *string
	SetInterface(val *string)
	InterfaceInput() *string
	InternalValue() *NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDisk
	SetInternalValue(val *NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDisk)
	Kind() *string
	Licenses() *[]*string
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	Source() *string
	SetSource(val *string)
	SourceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutInitializeParams(value *NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParams)
	ResetInitializeParams()
	ResetInterface()
	ResetMode()
	ResetSource()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference
type jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) AutoDelete() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autoDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) Boot() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"boot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) DeviceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) GuestOsFeatures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guestOsFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) Index() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"index",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) InitializeParams() NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParamsOutputReference {
	var returns NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParamsOutputReference
	_jsii_.Get(
		j,
		"initializeParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) InitializeParamsInput() *NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParams {
	var returns *NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParams
	_jsii_.Get(
		j,
		"initializeParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) Interface() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) InterfaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) InternalValue() *NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDisk {
	var returns *NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDisk
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) Licenses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"licenses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference {
	_init_.Initialize()

	if err := validateNewNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.notebooksRuntime.NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference_Override(n NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.notebooksRuntime.NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference)SetInterface(val *string) {
	if err := j.validateSetInterfaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interface",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference)SetInternalValue(val *NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDisk) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) PutInitializeParams(value *NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParams) {
	if err := n.validatePutInitializeParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putInitializeParams",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) ResetInitializeParams() {
	_jsii_.InvokeVoid(
		n,
		"resetInitializeParams",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) ResetInterface() {
	_jsii_.InvokeVoid(
		n,
		"resetInterface",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		n,
		"resetMode",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		n,
		"resetSource",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		n,
		"resetType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

