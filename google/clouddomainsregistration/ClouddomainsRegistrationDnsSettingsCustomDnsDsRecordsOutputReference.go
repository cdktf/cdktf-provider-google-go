// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddomainsregistration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/clouddomainsregistration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference interface {
	cdktf.ComplexObject
	Algorithm() *string
	SetAlgorithm(val *string)
	AlgorithmInput() *string
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
	Digest() *string
	SetDigest(val *string)
	DigestInput() *string
	DigestType() *string
	SetDigestType(val *string)
	DigestTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyTag() *float64
	SetKeyTag(val *float64)
	KeyTagInput() *float64
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
	ResetAlgorithm()
	ResetDigest()
	ResetDigestType()
	ResetKeyTag()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference
type jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) Algorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) AlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) Digest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"digest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) DigestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"digestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) DigestType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"digestType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) DigestTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"digestTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) KeyTag() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) KeyTagInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference {
	_init_.Initialize()

	if err := validateNewClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.clouddomainsRegistration.ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference_Override(c ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.clouddomainsRegistration.ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference)SetAlgorithm(val *string) {
	if err := j.validateSetAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"algorithm",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference)SetDigest(val *string) {
	if err := j.validateSetDigestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"digest",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference)SetDigestType(val *string) {
	if err := j.validateSetDigestTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"digestType",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference)SetKeyTag(val *float64) {
	if err := j.validateSetKeyTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyTag",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) ResetAlgorithm() {
	_jsii_.InvokeVoid(
		c,
		"resetAlgorithm",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) ResetDigest() {
	_jsii_.InvokeVoid(
		c,
		"resetDigest",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) ResetDigestType() {
	_jsii_.InvokeVoid(
		c,
		"resetDigestType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) ResetKeyTag() {
	_jsii_.InvokeVoid(
		c,
		"resetKeyTag",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ClouddomainsRegistrationDnsSettingsCustomDnsDsRecordsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

