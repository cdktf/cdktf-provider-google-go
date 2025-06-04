// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanagerserviceperimeters

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/accesscontextmanagerserviceperimeters/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MethodSelectors() AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsMethodSelectorsList
	MethodSelectorsInput() interface{}
	ServiceName() *string
	SetServiceName(val *string)
	ServiceNameInput() *string
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
	PutMethodSelectors(value interface{})
	ResetMethodSelectors()
	ResetServiceName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference
type jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) MethodSelectors() AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsMethodSelectorsList {
	var returns AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsMethodSelectorsList
	_jsii_.Get(
		j,
		"methodSelectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) MethodSelectorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"methodSelectorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) ServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference {
	_init_.Initialize()

	if err := validateNewAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.accessContextManagerServicePerimeters.AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference_Override(a AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.accessContextManagerServicePerimeters.AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference)SetServiceName(val *string) {
	if err := j.validateSetServiceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceName",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) PutMethodSelectors(value interface{}) {
	if err := a.validatePutMethodSelectorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMethodSelectors",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) ResetMethodSelectors() {
	_jsii_.InvokeVoid(
		a,
		"resetMethodSelectors",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) ResetServiceName() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressToOperationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

