// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apphubservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/apphubservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApphubServiceAttributesOutputReference interface {
	cdktf.ComplexObject
	BusinessOwners() ApphubServiceAttributesBusinessOwnersList
	BusinessOwnersInput() interface{}
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
	Criticality() ApphubServiceAttributesCriticalityOutputReference
	CriticalityInput() *ApphubServiceAttributesCriticality
	DeveloperOwners() ApphubServiceAttributesDeveloperOwnersList
	DeveloperOwnersInput() interface{}
	Environment() ApphubServiceAttributesEnvironmentOutputReference
	EnvironmentInput() *ApphubServiceAttributesEnvironment
	// Experimental.
	Fqn() *string
	InternalValue() *ApphubServiceAttributes
	SetInternalValue(val *ApphubServiceAttributes)
	OperatorOwners() ApphubServiceAttributesOperatorOwnersList
	OperatorOwnersInput() interface{}
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
	PutBusinessOwners(value interface{})
	PutCriticality(value *ApphubServiceAttributesCriticality)
	PutDeveloperOwners(value interface{})
	PutEnvironment(value *ApphubServiceAttributesEnvironment)
	PutOperatorOwners(value interface{})
	ResetBusinessOwners()
	ResetCriticality()
	ResetDeveloperOwners()
	ResetEnvironment()
	ResetOperatorOwners()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApphubServiceAttributesOutputReference
type jsiiProxy_ApphubServiceAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApphubServiceAttributesOutputReference) BusinessOwners() ApphubServiceAttributesBusinessOwnersList {
	var returns ApphubServiceAttributesBusinessOwnersList
	_jsii_.Get(
		j,
		"businessOwners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubServiceAttributesOutputReference) BusinessOwnersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"businessOwnersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubServiceAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubServiceAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubServiceAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubServiceAttributesOutputReference) Criticality() ApphubServiceAttributesCriticalityOutputReference {
	var returns ApphubServiceAttributesCriticalityOutputReference
	_jsii_.Get(
		j,
		"criticality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubServiceAttributesOutputReference) CriticalityInput() *ApphubServiceAttributesCriticality {
	var returns *ApphubServiceAttributesCriticality
	_jsii_.Get(
		j,
		"criticalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubServiceAttributesOutputReference) DeveloperOwners() ApphubServiceAttributesDeveloperOwnersList {
	var returns ApphubServiceAttributesDeveloperOwnersList
	_jsii_.Get(
		j,
		"developerOwners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubServiceAttributesOutputReference) DeveloperOwnersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"developerOwnersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubServiceAttributesOutputReference) Environment() ApphubServiceAttributesEnvironmentOutputReference {
	var returns ApphubServiceAttributesEnvironmentOutputReference
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubServiceAttributesOutputReference) EnvironmentInput() *ApphubServiceAttributesEnvironment {
	var returns *ApphubServiceAttributesEnvironment
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubServiceAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubServiceAttributesOutputReference) InternalValue() *ApphubServiceAttributes {
	var returns *ApphubServiceAttributes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubServiceAttributesOutputReference) OperatorOwners() ApphubServiceAttributesOperatorOwnersList {
	var returns ApphubServiceAttributesOperatorOwnersList
	_jsii_.Get(
		j,
		"operatorOwners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubServiceAttributesOutputReference) OperatorOwnersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operatorOwnersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubServiceAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApphubServiceAttributesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApphubServiceAttributesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApphubServiceAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewApphubServiceAttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApphubServiceAttributesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.apphubService.ApphubServiceAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApphubServiceAttributesOutputReference_Override(a ApphubServiceAttributesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.apphubService.ApphubServiceAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApphubServiceAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApphubServiceAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApphubServiceAttributesOutputReference)SetInternalValue(val *ApphubServiceAttributes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApphubServiceAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApphubServiceAttributesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApphubServiceAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApphubServiceAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApphubServiceAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApphubServiceAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApphubServiceAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApphubServiceAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApphubServiceAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApphubServiceAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApphubServiceAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApphubServiceAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApphubServiceAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApphubServiceAttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApphubServiceAttributesOutputReference) PutBusinessOwners(value interface{}) {
	if err := a.validatePutBusinessOwnersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBusinessOwners",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApphubServiceAttributesOutputReference) PutCriticality(value *ApphubServiceAttributesCriticality) {
	if err := a.validatePutCriticalityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCriticality",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApphubServiceAttributesOutputReference) PutDeveloperOwners(value interface{}) {
	if err := a.validatePutDeveloperOwnersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeveloperOwners",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApphubServiceAttributesOutputReference) PutEnvironment(value *ApphubServiceAttributesEnvironment) {
	if err := a.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApphubServiceAttributesOutputReference) PutOperatorOwners(value interface{}) {
	if err := a.validatePutOperatorOwnersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOperatorOwners",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApphubServiceAttributesOutputReference) ResetBusinessOwners() {
	_jsii_.InvokeVoid(
		a,
		"resetBusinessOwners",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApphubServiceAttributesOutputReference) ResetCriticality() {
	_jsii_.InvokeVoid(
		a,
		"resetCriticality",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApphubServiceAttributesOutputReference) ResetDeveloperOwners() {
	_jsii_.InvokeVoid(
		a,
		"resetDeveloperOwners",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApphubServiceAttributesOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApphubServiceAttributesOutputReference) ResetOperatorOwners() {
	_jsii_.InvokeVoid(
		a,
		"resetOperatorOwners",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApphubServiceAttributesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApphubServiceAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

