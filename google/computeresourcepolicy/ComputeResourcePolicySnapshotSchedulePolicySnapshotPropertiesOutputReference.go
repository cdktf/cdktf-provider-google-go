// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeresourcepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/computeresourcepolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference interface {
	cdktf.ComplexObject
	ChainName() *string
	SetChainName(val *string)
	ChainNameInput() *string
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
	// Experimental.
	Fqn() *string
	GuestFlush() interface{}
	SetGuestFlush(val interface{})
	GuestFlushInput() interface{}
	InternalValue() *ComputeResourcePolicySnapshotSchedulePolicySnapshotProperties
	SetInternalValue(val *ComputeResourcePolicySnapshotSchedulePolicySnapshotProperties)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	StorageLocations() *[]*string
	SetStorageLocations(val *[]*string)
	StorageLocationsInput() *[]*string
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
	ResetChainName()
	ResetGuestFlush()
	ResetLabels()
	ResetStorageLocations()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference
type jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) ChainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) ChainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) GuestFlush() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guestFlush",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) GuestFlushInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guestFlushInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) InternalValue() *ComputeResourcePolicySnapshotSchedulePolicySnapshotProperties {
	var returns *ComputeResourcePolicySnapshotSchedulePolicySnapshotProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) StorageLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"storageLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) StorageLocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"storageLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeResourcePolicy.ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference_Override(c ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeResourcePolicy.ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference)SetChainName(val *string) {
	if err := j.validateSetChainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"chainName",
		val,
	)
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference)SetGuestFlush(val interface{}) {
	if err := j.validateSetGuestFlushParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guestFlush",
		val,
	)
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference)SetInternalValue(val *ComputeResourcePolicySnapshotSchedulePolicySnapshotProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference)SetStorageLocations(val *[]*string) {
	if err := j.validateSetStorageLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageLocations",
		val,
	)
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) ResetChainName() {
	_jsii_.InvokeVoid(
		c,
		"resetChainName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) ResetGuestFlush() {
	_jsii_.InvokeVoid(
		c,
		"resetGuestFlush",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		c,
		"resetLabels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) ResetStorageLocations() {
	_jsii_.InvokeVoid(
		c,
		"resetStorageLocations",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

