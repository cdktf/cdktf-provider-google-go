// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeeapiproduct

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/apigeeapiproduct/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigeeApiProductOperationGroupOperationConfigsOutputReference interface {
	cdktf.ComplexObject
	ApiSource() *string
	SetApiSource(val *string)
	ApiSourceInput() *string
	Attributes() ApigeeApiProductOperationGroupOperationConfigsAttributesList
	AttributesInput() interface{}
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
	Operations() ApigeeApiProductOperationGroupOperationConfigsOperationsList
	OperationsInput() interface{}
	Quota() ApigeeApiProductOperationGroupOperationConfigsQuotaOutputReference
	QuotaInput() *ApigeeApiProductOperationGroupOperationConfigsQuota
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
	PutAttributes(value interface{})
	PutOperations(value interface{})
	PutQuota(value *ApigeeApiProductOperationGroupOperationConfigsQuota)
	ResetApiSource()
	ResetAttributes()
	ResetOperations()
	ResetQuota()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApigeeApiProductOperationGroupOperationConfigsOutputReference
type jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) ApiSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) ApiSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) Attributes() ApigeeApiProductOperationGroupOperationConfigsAttributesList {
	var returns ApigeeApiProductOperationGroupOperationConfigsAttributesList
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) AttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) Operations() ApigeeApiProductOperationGroupOperationConfigsOperationsList {
	var returns ApigeeApiProductOperationGroupOperationConfigsOperationsList
	_jsii_.Get(
		j,
		"operations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) OperationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) Quota() ApigeeApiProductOperationGroupOperationConfigsQuotaOutputReference {
	var returns ApigeeApiProductOperationGroupOperationConfigsQuotaOutputReference
	_jsii_.Get(
		j,
		"quota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) QuotaInput() *ApigeeApiProductOperationGroupOperationConfigsQuota {
	var returns *ApigeeApiProductOperationGroupOperationConfigsQuota
	_jsii_.Get(
		j,
		"quotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApigeeApiProductOperationGroupOperationConfigsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApigeeApiProductOperationGroupOperationConfigsOutputReference {
	_init_.Initialize()

	if err := validateNewApigeeApiProductOperationGroupOperationConfigsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.apigeeApiProduct.ApigeeApiProductOperationGroupOperationConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApigeeApiProductOperationGroupOperationConfigsOutputReference_Override(a ApigeeApiProductOperationGroupOperationConfigsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.apigeeApiProduct.ApigeeApiProductOperationGroupOperationConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference)SetApiSource(val *string) {
	if err := j.validateSetApiSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiSource",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) PutAttributes(value interface{}) {
	if err := a.validatePutAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAttributes",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) PutOperations(value interface{}) {
	if err := a.validatePutOperationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOperations",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) PutQuota(value *ApigeeApiProductOperationGroupOperationConfigsQuota) {
	if err := a.validatePutQuotaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putQuota",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) ResetApiSource() {
	_jsii_.InvokeVoid(
		a,
		"resetApiSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) ResetAttributes() {
	_jsii_.InvokeVoid(
		a,
		"resetAttributes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) ResetOperations() {
	_jsii_.InvokeVoid(
		a,
		"resetOperations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) ResetQuota() {
	_jsii_.InvokeVoid(
		a,
		"resetQuota",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApigeeApiProductOperationGroupOperationConfigsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

