// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computebackendservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/computebackendservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeBackendServiceSecuritySettingsOutputReference interface {
	cdktf.ComplexObject
	AwsV4Authentication() ComputeBackendServiceSecuritySettingsAwsV4AuthenticationOutputReference
	AwsV4AuthenticationInput() *ComputeBackendServiceSecuritySettingsAwsV4Authentication
	ClientTlsPolicy() *string
	SetClientTlsPolicy(val *string)
	ClientTlsPolicyInput() *string
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
	InternalValue() *ComputeBackendServiceSecuritySettings
	SetInternalValue(val *ComputeBackendServiceSecuritySettings)
	SubjectAltNames() *[]*string
	SetSubjectAltNames(val *[]*string)
	SubjectAltNamesInput() *[]*string
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
	PutAwsV4Authentication(value *ComputeBackendServiceSecuritySettingsAwsV4Authentication)
	ResetAwsV4Authentication()
	ResetClientTlsPolicy()
	ResetSubjectAltNames()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeBackendServiceSecuritySettingsOutputReference
type jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) AwsV4Authentication() ComputeBackendServiceSecuritySettingsAwsV4AuthenticationOutputReference {
	var returns ComputeBackendServiceSecuritySettingsAwsV4AuthenticationOutputReference
	_jsii_.Get(
		j,
		"awsV4Authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) AwsV4AuthenticationInput() *ComputeBackendServiceSecuritySettingsAwsV4Authentication {
	var returns *ComputeBackendServiceSecuritySettingsAwsV4Authentication
	_jsii_.Get(
		j,
		"awsV4AuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) ClientTlsPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTlsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) ClientTlsPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTlsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) InternalValue() *ComputeBackendServiceSecuritySettings {
	var returns *ComputeBackendServiceSecuritySettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) SubjectAltNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subjectAltNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) SubjectAltNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subjectAltNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeBackendServiceSecuritySettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeBackendServiceSecuritySettingsOutputReference {
	_init_.Initialize()

	if err := validateNewComputeBackendServiceSecuritySettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeBackendService.ComputeBackendServiceSecuritySettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeBackendServiceSecuritySettingsOutputReference_Override(c ComputeBackendServiceSecuritySettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeBackendService.ComputeBackendServiceSecuritySettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference)SetClientTlsPolicy(val *string) {
	if err := j.validateSetClientTlsPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientTlsPolicy",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference)SetInternalValue(val *ComputeBackendServiceSecuritySettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference)SetSubjectAltNames(val *[]*string) {
	if err := j.validateSetSubjectAltNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectAltNames",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) PutAwsV4Authentication(value *ComputeBackendServiceSecuritySettingsAwsV4Authentication) {
	if err := c.validatePutAwsV4AuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsV4Authentication",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) ResetAwsV4Authentication() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsV4Authentication",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) ResetClientTlsPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetClientTlsPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) ResetSubjectAltNames() {
	_jsii_.InvokeVoid(
		c,
		"resetSubjectAltNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeBackendServiceSecuritySettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

