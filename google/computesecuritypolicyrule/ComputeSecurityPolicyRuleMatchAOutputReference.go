// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computesecuritypolicyrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/computesecuritypolicyrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeSecurityPolicyRuleMatchAOutputReference interface {
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
	Config() ComputeSecurityPolicyRuleMatchConfigAOutputReference
	ConfigInput() *ComputeSecurityPolicyRuleMatchConfigA
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Expr() ComputeSecurityPolicyRuleMatchExprAOutputReference
	ExprInput() *ComputeSecurityPolicyRuleMatchExprA
	ExprOptions() ComputeSecurityPolicyRuleMatchExprOptionsAOutputReference
	ExprOptionsInput() *ComputeSecurityPolicyRuleMatchExprOptionsA
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeSecurityPolicyRuleMatchA
	SetInternalValue(val *ComputeSecurityPolicyRuleMatchA)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VersionedExpr() *string
	SetVersionedExpr(val *string)
	VersionedExprInput() *string
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
	PutConfig(value *ComputeSecurityPolicyRuleMatchConfigA)
	PutExpr(value *ComputeSecurityPolicyRuleMatchExprA)
	PutExprOptions(value *ComputeSecurityPolicyRuleMatchExprOptionsA)
	ResetConfig()
	ResetExpr()
	ResetExprOptions()
	ResetVersionedExpr()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeSecurityPolicyRuleMatchAOutputReference
type jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) Config() ComputeSecurityPolicyRuleMatchConfigAOutputReference {
	var returns ComputeSecurityPolicyRuleMatchConfigAOutputReference
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) ConfigInput() *ComputeSecurityPolicyRuleMatchConfigA {
	var returns *ComputeSecurityPolicyRuleMatchConfigA
	_jsii_.Get(
		j,
		"configInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) Expr() ComputeSecurityPolicyRuleMatchExprAOutputReference {
	var returns ComputeSecurityPolicyRuleMatchExprAOutputReference
	_jsii_.Get(
		j,
		"expr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) ExprInput() *ComputeSecurityPolicyRuleMatchExprA {
	var returns *ComputeSecurityPolicyRuleMatchExprA
	_jsii_.Get(
		j,
		"exprInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) ExprOptions() ComputeSecurityPolicyRuleMatchExprOptionsAOutputReference {
	var returns ComputeSecurityPolicyRuleMatchExprOptionsAOutputReference
	_jsii_.Get(
		j,
		"exprOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) ExprOptionsInput() *ComputeSecurityPolicyRuleMatchExprOptionsA {
	var returns *ComputeSecurityPolicyRuleMatchExprOptionsA
	_jsii_.Get(
		j,
		"exprOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) InternalValue() *ComputeSecurityPolicyRuleMatchA {
	var returns *ComputeSecurityPolicyRuleMatchA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) VersionedExpr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionedExpr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) VersionedExprInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionedExprInput",
		&returns,
	)
	return returns
}


func NewComputeSecurityPolicyRuleMatchAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeSecurityPolicyRuleMatchAOutputReference {
	_init_.Initialize()

	if err := validateNewComputeSecurityPolicyRuleMatchAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeSecurityPolicyRule.ComputeSecurityPolicyRuleMatchAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeSecurityPolicyRuleMatchAOutputReference_Override(c ComputeSecurityPolicyRuleMatchAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeSecurityPolicyRule.ComputeSecurityPolicyRuleMatchAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference)SetInternalValue(val *ComputeSecurityPolicyRuleMatchA) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference)SetVersionedExpr(val *string) {
	if err := j.validateSetVersionedExprParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionedExpr",
		val,
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) PutConfig(value *ComputeSecurityPolicyRuleMatchConfigA) {
	if err := c.validatePutConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) PutExpr(value *ComputeSecurityPolicyRuleMatchExprA) {
	if err := c.validatePutExprParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putExpr",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) PutExprOptions(value *ComputeSecurityPolicyRuleMatchExprOptionsA) {
	if err := c.validatePutExprOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putExprOptions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) ResetConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) ResetExpr() {
	_jsii_.InvokeVoid(
		c,
		"resetExpr",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) ResetExprOptions() {
	_jsii_.InvokeVoid(
		c,
		"resetExprOptions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) ResetVersionedExpr() {
	_jsii_.InvokeVoid(
		c,
		"resetVersionedExpr",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeSecurityPolicyRuleMatchAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

