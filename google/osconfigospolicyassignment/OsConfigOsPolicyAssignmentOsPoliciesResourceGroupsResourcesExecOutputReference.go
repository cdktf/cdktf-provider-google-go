// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigospolicyassignment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/osconfigospolicyassignment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference interface {
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
	Enforce() OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecEnforceOutputReference
	EnforceInput() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecEnforce
	// Experimental.
	Fqn() *string
	InternalValue() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec
	SetInternalValue(val *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Validate() OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecValidateOutputReference
	ValidateInput() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecValidate
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
	PutEnforce(value *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecEnforce)
	PutValidate(value *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecValidate)
	ResetEnforce()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference
type jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) Enforce() OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecEnforceOutputReference {
	var returns OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecEnforceOutputReference
	_jsii_.Get(
		j,
		"enforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) EnforceInput() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecEnforce {
	var returns *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecEnforce
	_jsii_.Get(
		j,
		"enforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) InternalValue() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec {
	var returns *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) Validate() OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecValidateOutputReference {
	var returns OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecValidateOutputReference
	_jsii_.Get(
		j,
		"validate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) ValidateInput() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecValidate {
	var returns *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecValidate
	_jsii_.Get(
		j,
		"validateInput",
		&returns,
	)
	return returns
}


func NewOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference {
	_init_.Initialize()

	if err := validateNewOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.osConfigOsPolicyAssignment.OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference_Override(o OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.osConfigOsPolicyAssignment.OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference)SetInternalValue(val *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) PutEnforce(value *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecEnforce) {
	if err := o.validatePutEnforceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putEnforce",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) PutValidate(value *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecValidate) {
	if err := o.validatePutValidateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putValidate",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) ResetEnforce() {
	_jsii_.InvokeVoid(
		o,
		"resetEnforce",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

