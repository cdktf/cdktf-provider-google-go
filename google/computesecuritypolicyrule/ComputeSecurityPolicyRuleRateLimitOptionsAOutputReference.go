// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computesecuritypolicyrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/computesecuritypolicyrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference interface {
	cdktf.ComplexObject
	BanDurationSec() *float64
	SetBanDurationSec(val *float64)
	BanDurationSecInput() *float64
	BanThreshold() ComputeSecurityPolicyRuleRateLimitOptionsBanThresholdAOutputReference
	BanThresholdInput() *ComputeSecurityPolicyRuleRateLimitOptionsBanThresholdA
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
	ConformAction() *string
	SetConformAction(val *string)
	ConformActionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnforceOnKey() *string
	SetEnforceOnKey(val *string)
	EnforceOnKeyConfigs() ComputeSecurityPolicyRuleRateLimitOptionsEnforceOnKeyConfigsAList
	EnforceOnKeyConfigsInput() interface{}
	EnforceOnKeyInput() *string
	EnforceOnKeyName() *string
	SetEnforceOnKeyName(val *string)
	EnforceOnKeyNameInput() *string
	ExceedAction() *string
	SetExceedAction(val *string)
	ExceedActionInput() *string
	ExceedRedirectOptions() ComputeSecurityPolicyRuleRateLimitOptionsExceedRedirectOptionsAOutputReference
	ExceedRedirectOptionsInput() *ComputeSecurityPolicyRuleRateLimitOptionsExceedRedirectOptionsA
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeSecurityPolicyRuleRateLimitOptionsA
	SetInternalValue(val *ComputeSecurityPolicyRuleRateLimitOptionsA)
	RateLimitThreshold() ComputeSecurityPolicyRuleRateLimitOptionsRateLimitThresholdAOutputReference
	RateLimitThresholdInput() *ComputeSecurityPolicyRuleRateLimitOptionsRateLimitThresholdA
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
	PutBanThreshold(value *ComputeSecurityPolicyRuleRateLimitOptionsBanThresholdA)
	PutEnforceOnKeyConfigs(value interface{})
	PutExceedRedirectOptions(value *ComputeSecurityPolicyRuleRateLimitOptionsExceedRedirectOptionsA)
	PutRateLimitThreshold(value *ComputeSecurityPolicyRuleRateLimitOptionsRateLimitThresholdA)
	ResetBanDurationSec()
	ResetBanThreshold()
	ResetConformAction()
	ResetEnforceOnKey()
	ResetEnforceOnKeyConfigs()
	ResetEnforceOnKeyName()
	ResetExceedAction()
	ResetExceedRedirectOptions()
	ResetRateLimitThreshold()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference
type jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) BanDurationSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"banDurationSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) BanDurationSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"banDurationSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) BanThreshold() ComputeSecurityPolicyRuleRateLimitOptionsBanThresholdAOutputReference {
	var returns ComputeSecurityPolicyRuleRateLimitOptionsBanThresholdAOutputReference
	_jsii_.Get(
		j,
		"banThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) BanThresholdInput() *ComputeSecurityPolicyRuleRateLimitOptionsBanThresholdA {
	var returns *ComputeSecurityPolicyRuleRateLimitOptionsBanThresholdA
	_jsii_.Get(
		j,
		"banThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) ConformAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conformAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) ConformActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conformActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) EnforceOnKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforceOnKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) EnforceOnKeyConfigs() ComputeSecurityPolicyRuleRateLimitOptionsEnforceOnKeyConfigsAList {
	var returns ComputeSecurityPolicyRuleRateLimitOptionsEnforceOnKeyConfigsAList
	_jsii_.Get(
		j,
		"enforceOnKeyConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) EnforceOnKeyConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceOnKeyConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) EnforceOnKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforceOnKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) EnforceOnKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforceOnKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) EnforceOnKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforceOnKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) ExceedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exceedAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) ExceedActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exceedActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) ExceedRedirectOptions() ComputeSecurityPolicyRuleRateLimitOptionsExceedRedirectOptionsAOutputReference {
	var returns ComputeSecurityPolicyRuleRateLimitOptionsExceedRedirectOptionsAOutputReference
	_jsii_.Get(
		j,
		"exceedRedirectOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) ExceedRedirectOptionsInput() *ComputeSecurityPolicyRuleRateLimitOptionsExceedRedirectOptionsA {
	var returns *ComputeSecurityPolicyRuleRateLimitOptionsExceedRedirectOptionsA
	_jsii_.Get(
		j,
		"exceedRedirectOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) InternalValue() *ComputeSecurityPolicyRuleRateLimitOptionsA {
	var returns *ComputeSecurityPolicyRuleRateLimitOptionsA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) RateLimitThreshold() ComputeSecurityPolicyRuleRateLimitOptionsRateLimitThresholdAOutputReference {
	var returns ComputeSecurityPolicyRuleRateLimitOptionsRateLimitThresholdAOutputReference
	_jsii_.Get(
		j,
		"rateLimitThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) RateLimitThresholdInput() *ComputeSecurityPolicyRuleRateLimitOptionsRateLimitThresholdA {
	var returns *ComputeSecurityPolicyRuleRateLimitOptionsRateLimitThresholdA
	_jsii_.Get(
		j,
		"rateLimitThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeSecurityPolicyRuleRateLimitOptionsAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference {
	_init_.Initialize()

	if err := validateNewComputeSecurityPolicyRuleRateLimitOptionsAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeSecurityPolicyRule.ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeSecurityPolicyRuleRateLimitOptionsAOutputReference_Override(c ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeSecurityPolicyRule.ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference)SetBanDurationSec(val *float64) {
	if err := j.validateSetBanDurationSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"banDurationSec",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference)SetConformAction(val *string) {
	if err := j.validateSetConformActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conformAction",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference)SetEnforceOnKey(val *string) {
	if err := j.validateSetEnforceOnKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceOnKey",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference)SetEnforceOnKeyName(val *string) {
	if err := j.validateSetEnforceOnKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceOnKeyName",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference)SetExceedAction(val *string) {
	if err := j.validateSetExceedActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exceedAction",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference)SetInternalValue(val *ComputeSecurityPolicyRuleRateLimitOptionsA) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) PutBanThreshold(value *ComputeSecurityPolicyRuleRateLimitOptionsBanThresholdA) {
	if err := c.validatePutBanThresholdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBanThreshold",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) PutEnforceOnKeyConfigs(value interface{}) {
	if err := c.validatePutEnforceOnKeyConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEnforceOnKeyConfigs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) PutExceedRedirectOptions(value *ComputeSecurityPolicyRuleRateLimitOptionsExceedRedirectOptionsA) {
	if err := c.validatePutExceedRedirectOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putExceedRedirectOptions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) PutRateLimitThreshold(value *ComputeSecurityPolicyRuleRateLimitOptionsRateLimitThresholdA) {
	if err := c.validatePutRateLimitThresholdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRateLimitThreshold",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) ResetBanDurationSec() {
	_jsii_.InvokeVoid(
		c,
		"resetBanDurationSec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) ResetBanThreshold() {
	_jsii_.InvokeVoid(
		c,
		"resetBanThreshold",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) ResetConformAction() {
	_jsii_.InvokeVoid(
		c,
		"resetConformAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) ResetEnforceOnKey() {
	_jsii_.InvokeVoid(
		c,
		"resetEnforceOnKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) ResetEnforceOnKeyConfigs() {
	_jsii_.InvokeVoid(
		c,
		"resetEnforceOnKeyConfigs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) ResetEnforceOnKeyName() {
	_jsii_.InvokeVoid(
		c,
		"resetEnforceOnKeyName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) ResetExceedAction() {
	_jsii_.InvokeVoid(
		c,
		"resetExceedAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) ResetExceedRedirectOptions() {
	_jsii_.InvokeVoid(
		c,
		"resetExceedRedirectOptions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) ResetRateLimitThreshold() {
	_jsii_.InvokeVoid(
		c,
		"resetRateLimitThreshold",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeSecurityPolicyRuleRateLimitOptionsAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

