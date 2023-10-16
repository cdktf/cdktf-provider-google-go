// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeekeystoresaliaseskeycertfile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v11/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v11/apigeekeystoresaliaseskeycertfile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference interface {
	cdktf.ComplexObject
	BasicConstraints() *string
	SetBasicConstraints(val *string)
	BasicConstraintsInput() *string
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
	ExpiryDate() *string
	SetExpiryDate(val *string)
	ExpiryDateInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Issuer() *string
	SetIssuer(val *string)
	IssuerInput() *string
	IsValid() *string
	SetIsValid(val *string)
	IsValidInput() *string
	PublicKey() *string
	SetPublicKey(val *string)
	PublicKeyInput() *string
	SerialNumber() *string
	SetSerialNumber(val *string)
	SerialNumberInput() *string
	SigAlgName() *string
	SetSigAlgName(val *string)
	SigAlgNameInput() *string
	Subject() *string
	SetSubject(val *string)
	SubjectAlternativeNames() *[]*string
	SetSubjectAlternativeNames(val *[]*string)
	SubjectAlternativeNamesInput() *[]*string
	SubjectInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ValidFrom() *string
	SetValidFrom(val *string)
	ValidFromInput() *string
	Version() *float64
	SetVersion(val *float64)
	VersionInput() *float64
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
	ResetBasicConstraints()
	ResetExpiryDate()
	ResetIssuer()
	ResetIsValid()
	ResetPublicKey()
	ResetSerialNumber()
	ResetSigAlgName()
	ResetSubject()
	ResetSubjectAlternativeNames()
	ResetValidFrom()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference
type jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) BasicConstraints() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basicConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) BasicConstraintsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basicConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) ExpiryDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiryDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) ExpiryDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiryDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) IsValid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isValid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) IsValidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isValidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) PublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) PublicKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) SerialNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serialNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) SerialNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serialNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) SigAlgName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sigAlgName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) SigAlgNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sigAlgNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) Subject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) SubjectAlternativeNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subjectAlternativeNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) SubjectAlternativeNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subjectAlternativeNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) SubjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) ValidFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) ValidFromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) Version() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) VersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference {
	_init_.Initialize()

	if err := validateNewApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.apigeeKeystoresAliasesKeyCertFile.ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference_Override(a ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.apigeeKeystoresAliasesKeyCertFile.ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference)SetBasicConstraints(val *string) {
	if err := j.validateSetBasicConstraintsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"basicConstraints",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference)SetExpiryDate(val *string) {
	if err := j.validateSetExpiryDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expiryDate",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference)SetIsValid(val *string) {
	if err := j.validateSetIsValidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isValid",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference)SetPublicKey(val *string) {
	if err := j.validateSetPublicKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicKey",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference)SetSerialNumber(val *string) {
	if err := j.validateSetSerialNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serialNumber",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference)SetSigAlgName(val *string) {
	if err := j.validateSetSigAlgNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sigAlgName",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference)SetSubject(val *string) {
	if err := j.validateSetSubjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subject",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference)SetSubjectAlternativeNames(val *[]*string) {
	if err := j.validateSetSubjectAlternativeNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectAlternativeNames",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference)SetValidFrom(val *string) {
	if err := j.validateSetValidFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validFrom",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference)SetVersion(val *float64) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) ResetBasicConstraints() {
	_jsii_.InvokeVoid(
		a,
		"resetBasicConstraints",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) ResetExpiryDate() {
	_jsii_.InvokeVoid(
		a,
		"resetExpiryDate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) ResetIssuer() {
	_jsii_.InvokeVoid(
		a,
		"resetIssuer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) ResetIsValid() {
	_jsii_.InvokeVoid(
		a,
		"resetIsValid",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) ResetPublicKey() {
	_jsii_.InvokeVoid(
		a,
		"resetPublicKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) ResetSerialNumber() {
	_jsii_.InvokeVoid(
		a,
		"resetSerialNumber",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) ResetSigAlgName() {
	_jsii_.InvokeVoid(
		a,
		"resetSigAlgName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) ResetSubject() {
	_jsii_.InvokeVoid(
		a,
		"resetSubject",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) ResetSubjectAlternativeNames() {
	_jsii_.InvokeVoid(
		a,
		"resetSubjectAlternativeNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) ResetValidFrom() {
	_jsii_.InvokeVoid(
		a,
		"resetValidFrom",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

