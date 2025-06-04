// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeployautomation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/clouddeployautomation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference interface {
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
	Retry() ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference
	RetryInput() *ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetry
	Rollback() ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRollbackOutputReference
	RollbackInput() *ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRollback
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
	PutRetry(value *ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetry)
	PutRollback(value *ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRollback)
	ResetRetry()
	ResetRollback()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference
type jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) Retry() ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference {
	var returns ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference
	_jsii_.Get(
		j,
		"retry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) RetryInput() *ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetry {
	var returns *ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetry
	_jsii_.Get(
		j,
		"retryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) Rollback() ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRollbackOutputReference {
	var returns ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRollbackOutputReference
	_jsii_.Get(
		j,
		"rollback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) RollbackInput() *ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRollback {
	var returns *ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRollback
	_jsii_.Get(
		j,
		"rollbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference {
	_init_.Initialize()

	if err := validateNewClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.clouddeployAutomation.ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference_Override(c ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.clouddeployAutomation.ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) PutRetry(value *ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetry) {
	if err := c.validatePutRetryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRetry",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) PutRollback(value *ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRollback) {
	if err := c.validatePutRollbackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRollback",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) ResetRetry() {
	_jsii_.InvokeVoid(
		c,
		"resetRetry",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) ResetRollback() {
	_jsii_.InvokeVoid(
		c,
		"resetRollback",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

