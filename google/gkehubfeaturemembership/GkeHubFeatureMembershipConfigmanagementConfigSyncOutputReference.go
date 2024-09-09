// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeaturemembership

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/gkehubfeaturemembership/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	Git() GkeHubFeatureMembershipConfigmanagementConfigSyncGitOutputReference
	GitInput() *GkeHubFeatureMembershipConfigmanagementConfigSyncGit
	InternalValue() *GkeHubFeatureMembershipConfigmanagementConfigSync
	SetInternalValue(val *GkeHubFeatureMembershipConfigmanagementConfigSync)
	MetricsGcpServiceAccountEmail() *string
	SetMetricsGcpServiceAccountEmail(val *string)
	MetricsGcpServiceAccountEmailInput() *string
	Oci() GkeHubFeatureMembershipConfigmanagementConfigSyncOciOutputReference
	OciInput() *GkeHubFeatureMembershipConfigmanagementConfigSyncOci
	PreventDrift() interface{}
	SetPreventDrift(val interface{})
	PreventDriftInput() interface{}
	SourceFormat() *string
	SetSourceFormat(val *string)
	SourceFormatInput() *string
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
	PutGit(value *GkeHubFeatureMembershipConfigmanagementConfigSyncGit)
	PutOci(value *GkeHubFeatureMembershipConfigmanagementConfigSyncOci)
	ResetEnabled()
	ResetGit()
	ResetMetricsGcpServiceAccountEmail()
	ResetOci()
	ResetPreventDrift()
	ResetSourceFormat()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference
type jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) Git() GkeHubFeatureMembershipConfigmanagementConfigSyncGitOutputReference {
	var returns GkeHubFeatureMembershipConfigmanagementConfigSyncGitOutputReference
	_jsii_.Get(
		j,
		"git",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) GitInput() *GkeHubFeatureMembershipConfigmanagementConfigSyncGit {
	var returns *GkeHubFeatureMembershipConfigmanagementConfigSyncGit
	_jsii_.Get(
		j,
		"gitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) InternalValue() *GkeHubFeatureMembershipConfigmanagementConfigSync {
	var returns *GkeHubFeatureMembershipConfigmanagementConfigSync
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) MetricsGcpServiceAccountEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsGcpServiceAccountEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) MetricsGcpServiceAccountEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsGcpServiceAccountEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) Oci() GkeHubFeatureMembershipConfigmanagementConfigSyncOciOutputReference {
	var returns GkeHubFeatureMembershipConfigmanagementConfigSyncOciOutputReference
	_jsii_.Get(
		j,
		"oci",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) OciInput() *GkeHubFeatureMembershipConfigmanagementConfigSyncOci {
	var returns *GkeHubFeatureMembershipConfigmanagementConfigSyncOci
	_jsii_.Get(
		j,
		"ociInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) PreventDrift() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventDrift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) PreventDriftInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preventDriftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) SourceFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) SourceFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference {
	_init_.Initialize()

	if err := validateNewGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeHubFeatureMembership.GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference_Override(g GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeHubFeatureMembership.GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference)SetInternalValue(val *GkeHubFeatureMembershipConfigmanagementConfigSync) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference)SetMetricsGcpServiceAccountEmail(val *string) {
	if err := j.validateSetMetricsGcpServiceAccountEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsGcpServiceAccountEmail",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference)SetPreventDrift(val interface{}) {
	if err := j.validateSetPreventDriftParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preventDrift",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference)SetSourceFormat(val *string) {
	if err := j.validateSetSourceFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceFormat",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) PutGit(value *GkeHubFeatureMembershipConfigmanagementConfigSyncGit) {
	if err := g.validatePutGitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGit",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) PutOci(value *GkeHubFeatureMembershipConfigmanagementConfigSyncOci) {
	if err := g.validatePutOciParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOci",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) ResetGit() {
	_jsii_.InvokeVoid(
		g,
		"resetGit",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) ResetMetricsGcpServiceAccountEmail() {
	_jsii_.InvokeVoid(
		g,
		"resetMetricsGcpServiceAccountEmail",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) ResetOci() {
	_jsii_.InvokeVoid(
		g,
		"resetOci",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) ResetPreventDrift() {
	_jsii_.InvokeVoid(
		g,
		"resetPreventDrift",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) ResetSourceFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

