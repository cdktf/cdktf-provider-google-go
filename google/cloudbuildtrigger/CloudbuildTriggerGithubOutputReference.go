// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbuildtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/cloudbuildtrigger/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudbuildTriggerGithubOutputReference interface {
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
	EnterpriseConfigResourceName() *string
	SetEnterpriseConfigResourceName(val *string)
	EnterpriseConfigResourceNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CloudbuildTriggerGithub
	SetInternalValue(val *CloudbuildTriggerGithub)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Owner() *string
	SetOwner(val *string)
	OwnerInput() *string
	PullRequest() CloudbuildTriggerGithubPullRequestOutputReference
	PullRequestInput() *CloudbuildTriggerGithubPullRequest
	Push() CloudbuildTriggerGithubPushOutputReference
	PushInput() *CloudbuildTriggerGithubPush
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
	PutPullRequest(value *CloudbuildTriggerGithubPullRequest)
	PutPush(value *CloudbuildTriggerGithubPush)
	ResetEnterpriseConfigResourceName()
	ResetName()
	ResetOwner()
	ResetPullRequest()
	ResetPush()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudbuildTriggerGithubOutputReference
type jsiiProxy_CloudbuildTriggerGithubOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudbuildTriggerGithubOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerGithubOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerGithubOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerGithubOutputReference) EnterpriseConfigResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enterpriseConfigResourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerGithubOutputReference) EnterpriseConfigResourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enterpriseConfigResourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerGithubOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerGithubOutputReference) InternalValue() *CloudbuildTriggerGithub {
	var returns *CloudbuildTriggerGithub
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerGithubOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerGithubOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerGithubOutputReference) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerGithubOutputReference) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerGithubOutputReference) PullRequest() CloudbuildTriggerGithubPullRequestOutputReference {
	var returns CloudbuildTriggerGithubPullRequestOutputReference
	_jsii_.Get(
		j,
		"pullRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerGithubOutputReference) PullRequestInput() *CloudbuildTriggerGithubPullRequest {
	var returns *CloudbuildTriggerGithubPullRequest
	_jsii_.Get(
		j,
		"pullRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerGithubOutputReference) Push() CloudbuildTriggerGithubPushOutputReference {
	var returns CloudbuildTriggerGithubPushOutputReference
	_jsii_.Get(
		j,
		"push",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerGithubOutputReference) PushInput() *CloudbuildTriggerGithubPush {
	var returns *CloudbuildTriggerGithubPush
	_jsii_.Get(
		j,
		"pushInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerGithubOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerGithubOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudbuildTriggerGithubOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudbuildTriggerGithubOutputReference {
	_init_.Initialize()

	if err := validateNewCloudbuildTriggerGithubOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudbuildTriggerGithubOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudbuildTrigger.CloudbuildTriggerGithubOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudbuildTriggerGithubOutputReference_Override(c CloudbuildTriggerGithubOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudbuildTrigger.CloudbuildTriggerGithubOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudbuildTriggerGithubOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerGithubOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerGithubOutputReference)SetEnterpriseConfigResourceName(val *string) {
	if err := j.validateSetEnterpriseConfigResourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enterpriseConfigResourceName",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerGithubOutputReference)SetInternalValue(val *CloudbuildTriggerGithub) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerGithubOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerGithubOutputReference)SetOwner(val *string) {
	if err := j.validateSetOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerGithubOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerGithubOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudbuildTriggerGithubOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerGithubOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudbuildTriggerGithubOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudbuildTriggerGithubOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudbuildTriggerGithubOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudbuildTriggerGithubOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudbuildTriggerGithubOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudbuildTriggerGithubOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudbuildTriggerGithubOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudbuildTriggerGithubOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudbuildTriggerGithubOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerGithubOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudbuildTriggerGithubOutputReference) PutPullRequest(value *CloudbuildTriggerGithubPullRequest) {
	if err := c.validatePutPullRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPullRequest",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTriggerGithubOutputReference) PutPush(value *CloudbuildTriggerGithubPush) {
	if err := c.validatePutPushParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPush",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTriggerGithubOutputReference) ResetEnterpriseConfigResourceName() {
	_jsii_.InvokeVoid(
		c,
		"resetEnterpriseConfigResourceName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerGithubOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerGithubOutputReference) ResetOwner() {
	_jsii_.InvokeVoid(
		c,
		"resetOwner",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerGithubOutputReference) ResetPullRequest() {
	_jsii_.InvokeVoid(
		c,
		"resetPullRequest",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerGithubOutputReference) ResetPush() {
	_jsii_.InvokeVoid(
		c,
		"resetPush",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerGithubOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudbuildTriggerGithubOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

