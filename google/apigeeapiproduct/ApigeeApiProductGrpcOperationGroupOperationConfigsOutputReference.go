// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeeapiproduct

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/apigeeapiproduct/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference interface {
	cdktf.ComplexObject
	ApiSource() *string
	SetApiSource(val *string)
	ApiSourceInput() *string
	Attributes() ApigeeApiProductGrpcOperationGroupOperationConfigsAttributesList
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
	Methods() *[]*string
	SetMethods(val *[]*string)
	MethodsInput() *[]*string
	Quota() ApigeeApiProductGrpcOperationGroupOperationConfigsQuotaOutputReference
	QuotaInput() *ApigeeApiProductGrpcOperationGroupOperationConfigsQuota
	Service() *string
	SetService(val *string)
	ServiceInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutAttributes(value interface{})
	PutQuota(value *ApigeeApiProductGrpcOperationGroupOperationConfigsQuota)
	ResetApiSource()
	ResetAttributes()
	ResetMethods()
	ResetQuota()
	ResetService()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference
type jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) ApiSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) ApiSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) Attributes() ApigeeApiProductGrpcOperationGroupOperationConfigsAttributesList {
	var returns ApigeeApiProductGrpcOperationGroupOperationConfigsAttributesList
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) AttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) Methods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"methods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) MethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"methodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) Quota() ApigeeApiProductGrpcOperationGroupOperationConfigsQuotaOutputReference {
	var returns ApigeeApiProductGrpcOperationGroupOperationConfigsQuotaOutputReference
	_jsii_.Get(
		j,
		"quota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) QuotaInput() *ApigeeApiProductGrpcOperationGroupOperationConfigsQuota {
	var returns *ApigeeApiProductGrpcOperationGroupOperationConfigsQuota
	_jsii_.Get(
		j,
		"quotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference {
	_init_.Initialize()

	if err := validateNewApigeeApiProductGrpcOperationGroupOperationConfigsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.apigeeApiProduct.ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference_Override(a ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.apigeeApiProduct.ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference)SetApiSource(val *string) {
	if err := j.validateSetApiSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiSource",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference)SetMethods(val *[]*string) {
	if err := j.validateSetMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"methods",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) PutAttributes(value interface{}) {
	if err := a.validatePutAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAttributes",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) PutQuota(value *ApigeeApiProductGrpcOperationGroupOperationConfigsQuota) {
	if err := a.validatePutQuotaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putQuota",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) ResetApiSource() {
	_jsii_.InvokeVoid(
		a,
		"resetApiSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) ResetAttributes() {
	_jsii_.InvokeVoid(
		a,
		"resetAttributes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) ResetMethods() {
	_jsii_.InvokeVoid(
		a,
		"resetMethods",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) ResetQuota() {
	_jsii_.InvokeVoid(
		a,
		"resetQuota",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) ResetService() {
	_jsii_.InvokeVoid(
		a,
		"resetService",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeApiProductGrpcOperationGroupOperationConfigsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

