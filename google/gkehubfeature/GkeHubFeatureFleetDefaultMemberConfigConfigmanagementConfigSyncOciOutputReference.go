// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeature

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/gkehubfeature/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference interface {
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
	InternalValue() *GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOci
	SetInternalValue(val *GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOci)
	PolicyDir() *string
	SetPolicyDir(val *string)
	PolicyDirInput() *string
	SecretType() *string
	SetSecretType(val *string)
	SecretTypeInput() *string
	SyncRepo() *string
	SetSyncRepo(val *string)
	SyncRepoInput() *string
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
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	ResetPolicyDir()
	ResetSyncRepo()
	ResetSyncWaitSecs()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference
type jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) GcpServiceAccountEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpServiceAccountEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) GcpServiceAccountEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpServiceAccountEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) InternalValue() *GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOci {
	var returns *GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOci
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) PolicyDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) PolicyDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) SecretType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) SecretTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) SyncRepo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) SyncRepoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncRepoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) SyncWaitSecs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncWaitSecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) SyncWaitSecsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncWaitSecsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewGkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference {
	_init_.Initialize()

	if err := validateNewGkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeHubFeature.GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference_Override(g GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeHubFeature.GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference)SetGcpServiceAccountEmail(val *string) {
	if err := j.validateSetGcpServiceAccountEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcpServiceAccountEmail",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference)SetInternalValue(val *GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOci) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference)SetPolicyDir(val *string) {
	if err := j.validateSetPolicyDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyDir",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference)SetSecretType(val *string) {
	if err := j.validateSetSecretTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretType",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference)SetSyncRepo(val *string) {
	if err := j.validateSetSyncRepoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncRepo",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference)SetSyncWaitSecs(val *string) {
	if err := j.validateSetSyncWaitSecsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncWaitSecs",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) ResetGcpServiceAccountEmail() {
	_jsii_.InvokeVoid(
		g,
		"resetGcpServiceAccountEmail",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) ResetPolicyDir() {
	_jsii_.InvokeVoid(
		g,
		"resetPolicyDir",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) ResetSyncRepo() {
	_jsii_.InvokeVoid(
		g,
		"resetSyncRepo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) ResetSyncWaitSecs() {
	_jsii_.InvokeVoid(
		g,
		"resetSyncWaitSecs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GkeHubFeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOciOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

