// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationsauthconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/integrationsauthconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference interface {
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
	Key() IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesKeyOutputReference
	KeyInput() *IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesKey
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Value() IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesValueOutputReference
	ValueInput() *IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesValue
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
	PutKey(value *IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesKey)
	PutValue(value *IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesValue)
	ResetKey()
	ResetValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference
type jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) Key() IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesKeyOutputReference {
	var returns IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesKeyOutputReference
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) KeyInput() *IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesKey {
	var returns *IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesKey
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) Value() IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesValueOutputReference {
	var returns IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesValueOutputReference
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) ValueInput() *IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesValue {
	var returns *IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesValue
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewIntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference {
	_init_.Initialize()

	if err := validateNewIntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.integrationsAuthConfig.IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference_Override(i IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.integrationsAuthConfig.IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) PutKey(value *IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesKey) {
	if err := i.validatePutKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putKey",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) PutValue(value *IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesValue) {
	if err := i.validatePutValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putValue",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) ResetKey() {
	_jsii_.InvokeVoid(
		i,
		"resetKey",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		i,
		"resetValue",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParamsEntriesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

