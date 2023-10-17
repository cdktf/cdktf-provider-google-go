// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanageraccesslevels

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v12/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v12/accesscontextmanageraccesslevels/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DevicePolicy() AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference
	DevicePolicyInput() *AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicy
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpSubnetworks() *[]*string
	SetIpSubnetworks(val *[]*string)
	IpSubnetworksInput() *[]*string
	Members() *[]*string
	SetMembers(val *[]*string)
	MembersInput() *[]*string
	Negate() interface{}
	SetNegate(val interface{})
	NegateInput() interface{}
	Regions() *[]*string
	SetRegions(val *[]*string)
	RegionsInput() *[]*string
	RequiredAccessLevels() *[]*string
	SetRequiredAccessLevels(val *[]*string)
	RequiredAccessLevelsInput() *[]*string
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
	PutDevicePolicy(value *AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicy)
	ResetDevicePolicy()
	ResetIpSubnetworks()
	ResetMembers()
	ResetNegate()
	ResetRegions()
	ResetRequiredAccessLevels()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference
type jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) DevicePolicy() AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference {
	var returns AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference
	_jsii_.Get(
		j,
		"devicePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) DevicePolicyInput() *AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicy {
	var returns *AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicy
	_jsii_.Get(
		j,
		"devicePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) IpSubnetworks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipSubnetworks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) IpSubnetworksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipSubnetworksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) Members() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"members",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) MembersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"membersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) Negate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) NegateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) Regions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) RegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) RequiredAccessLevels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requiredAccessLevels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) RequiredAccessLevelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requiredAccessLevelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewAccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.accessContextManagerAccessLevels.AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference_Override(a AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.accessContextManagerAccessLevels.AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference)SetIpSubnetworks(val *[]*string) {
	if err := j.validateSetIpSubnetworksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipSubnetworks",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference)SetMembers(val *[]*string) {
	if err := j.validateSetMembersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"members",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference)SetNegate(val interface{}) {
	if err := j.validateSetNegateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"negate",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference)SetRegions(val *[]*string) {
	if err := j.validateSetRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regions",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference)SetRequiredAccessLevels(val *[]*string) {
	if err := j.validateSetRequiredAccessLevelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredAccessLevels",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) PutDevicePolicy(value *AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicy) {
	if err := a.validatePutDevicePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDevicePolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) ResetDevicePolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetDevicePolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) ResetIpSubnetworks() {
	_jsii_.InvokeVoid(
		a,
		"resetIpSubnetworks",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) ResetMembers() {
	_jsii_.InvokeVoid(
		a,
		"resetMembers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) ResetNegate() {
	_jsii_.InvokeVoid(
		a,
		"resetNegate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) ResetRegions() {
	_jsii_.InvokeVoid(
		a,
		"resetRegions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) ResetRequiredAccessLevels() {
	_jsii_.InvokeVoid(
		a,
		"resetRequiredAccessLevels",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

