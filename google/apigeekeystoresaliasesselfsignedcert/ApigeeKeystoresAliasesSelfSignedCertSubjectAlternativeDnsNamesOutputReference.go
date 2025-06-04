// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeekeystoresaliasesselfsignedcert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/apigeekeystoresaliasesselfsignedcert/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference interface {
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
	InternalValue() *ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames
	SetInternalValue(val *ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames)
	SubjectAlternativeName() *string
	SetSubjectAlternativeName(val *string)
	SubjectAlternativeNameInput() *string
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
	ResetSubjectAlternativeName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference
type jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference) InternalValue() *ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames {
	var returns *ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference) SubjectAlternativeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectAlternativeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference) SubjectAlternativeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectAlternativeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference {
	_init_.Initialize()

	if err := validateNewApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.apigeeKeystoresAliasesSelfSignedCert.ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference_Override(a ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.apigeeKeystoresAliasesSelfSignedCert.ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference)SetInternalValue(val *ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference)SetSubjectAlternativeName(val *string) {
	if err := j.validateSetSubjectAlternativeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectAlternativeName",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference) ResetSubjectAlternativeName() {
	_jsii_.InvokeVoid(
		a,
		"resetSubjectAlternativeName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

