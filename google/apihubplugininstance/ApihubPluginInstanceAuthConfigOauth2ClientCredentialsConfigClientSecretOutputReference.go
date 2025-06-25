// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apihubplugininstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/apihubplugininstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference interface {
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
	InternalValue() *ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecret
	SetInternalValue(val *ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecret)
	SecretVersion() *string
	SetSecretVersion(val *string)
	SecretVersionInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference
type jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference) InternalValue() *ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecret {
	var returns *ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecret
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference) SecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference) SecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference {
	_init_.Initialize()

	if err := validateNewApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.apihubPluginInstance.ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference_Override(a ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.apihubPluginInstance.ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference)SetInternalValue(val *ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecret) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference)SetSecretVersion(val *string) {
	if err := j.validateSetSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretVersion",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigClientSecretOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

