// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/dialogflowcxagent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference interface {
	cdktf.ComplexObject
	AccessToken() *string
	SetAccessToken(val *string)
	AccessTokenInput() *string
	Branches() *[]*string
	SetBranches(val *[]*string)
	BranchesInput() *[]*string
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
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DialogflowCxAgentGitIntegrationSettingsGithubSettings
	SetInternalValue(val *DialogflowCxAgentGitIntegrationSettingsGithubSettings)
	RepositoryUri() *string
	SetRepositoryUri(val *string)
	RepositoryUriInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrackingBranch() *string
	SetTrackingBranch(val *string)
	TrackingBranchInput() *string
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
	ResetAccessToken()
	ResetBranches()
	ResetDisplayName()
	ResetRepositoryUri()
	ResetTrackingBranch()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference
type jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) AccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) AccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) Branches() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) BranchesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"branchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) InternalValue() *DialogflowCxAgentGitIntegrationSettingsGithubSettings {
	var returns *DialogflowCxAgentGitIntegrationSettingsGithubSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) RepositoryUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) RepositoryUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) TrackingBranch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trackingBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) TrackingBranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trackingBranchInput",
		&returns,
	)
	return returns
}


func NewDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxAgent.DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference_Override(d DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxAgent.DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference)SetAccessToken(val *string) {
	if err := j.validateSetAccessTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessToken",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference)SetBranches(val *[]*string) {
	if err := j.validateSetBranchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branches",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference)SetInternalValue(val *DialogflowCxAgentGitIntegrationSettingsGithubSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference)SetRepositoryUri(val *string) {
	if err := j.validateSetRepositoryUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repositoryUri",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference)SetTrackingBranch(val *string) {
	if err := j.validateSetTrackingBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trackingBranch",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) ResetAccessToken() {
	_jsii_.InvokeVoid(
		d,
		"resetAccessToken",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) ResetBranches() {
	_jsii_.InvokeVoid(
		d,
		"resetBranches",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) ResetRepositoryUri() {
	_jsii_.InvokeVoid(
		d,
		"resetRepositoryUri",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) ResetTrackingBranch() {
	_jsii_.InvokeVoid(
		d,
		"resetTrackingBranch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgentGitIntegrationSettingsGithubSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

