// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeature

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/gkehubfeature/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference interface {
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
	GcpServiceAccountEmail() *string
	SetGcpServiceAccountEmail(val *string)
	GcpServiceAccountEmailInput() *string
	HttpsProxy() *string
	SetHttpsProxy(val *string)
	HttpsProxyInput() *string
	InternalValue() *GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGit
	SetInternalValue(val *GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGit)
	PolicyDir() *string
	SetPolicyDir(val *string)
	PolicyDirInput() *string
	SecretType() *string
	SetSecretType(val *string)
	SecretTypeInput() *string
	SyncBranch() *string
	SetSyncBranch(val *string)
	SyncBranchInput() *string
	SyncRepo() *string
	SetSyncRepo(val *string)
	SyncRepoInput() *string
	SyncRev() *string
	SetSyncRev(val *string)
	SyncRevInput() *string
	SyncWaitSecs() *string
	SetSyncWaitSecs(val *string)
	SyncWaitSecsInput() *string
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
	ResetGcpServiceAccountEmail()
	ResetHttpsProxy()
	ResetPolicyDir()
	ResetSyncBranch()
	ResetSyncRepo()
	ResetSyncRev()
	ResetSyncWaitSecs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference
type jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) GcpServiceAccountEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpServiceAccountEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) GcpServiceAccountEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpServiceAccountEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) HttpsProxy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsProxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) HttpsProxyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsProxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) InternalValue() *GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGit {
	var returns *GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGit
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) PolicyDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) PolicyDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) SecretType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) SecretTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) SyncBranch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) SyncBranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncBranchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) SyncRepo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) SyncRepoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncRepoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) SyncRev() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncRev",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) SyncRevInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncRevInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) SyncWaitSecs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncWaitSecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) SyncWaitSecsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncWaitSecsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference {
	_init_.Initialize()

	if err := validateNewGkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeHubFeature.GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference_Override(g GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeHubFeature.GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference)SetGcpServiceAccountEmail(val *string) {
	if err := j.validateSetGcpServiceAccountEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcpServiceAccountEmail",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference)SetHttpsProxy(val *string) {
	if err := j.validateSetHttpsProxyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsProxy",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference)SetInternalValue(val *GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGit) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference)SetPolicyDir(val *string) {
	if err := j.validateSetPolicyDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyDir",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference)SetSecretType(val *string) {
	if err := j.validateSetSecretTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretType",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference)SetSyncBranch(val *string) {
	if err := j.validateSetSyncBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncBranch",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference)SetSyncRepo(val *string) {
	if err := j.validateSetSyncRepoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncRepo",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference)SetSyncRev(val *string) {
	if err := j.validateSetSyncRevParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncRev",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference)SetSyncWaitSecs(val *string) {
	if err := j.validateSetSyncWaitSecsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncWaitSecs",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) ResetGcpServiceAccountEmail() {
	_jsii_.InvokeVoid(
		g,
		"resetGcpServiceAccountEmail",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) ResetHttpsProxy() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpsProxy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) ResetPolicyDir() {
	_jsii_.InvokeVoid(
		g,
		"resetPolicyDir",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) ResetSyncBranch() {
	_jsii_.InvokeVoid(
		g,
		"resetSyncBranch",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) ResetSyncRepo() {
	_jsii_.InvokeVoid(
		g,
		"resetSyncRepo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) ResetSyncRev() {
	_jsii_.InvokeVoid(
		g,
		"resetSyncRev",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) ResetSyncWaitSecs() {
	_jsii_.InvokeVoid(
		g,
		"resetSyncWaitSecs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGitOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

