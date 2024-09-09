// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package artifactregistryrepository

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/artifactregistryrepository/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference interface {
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
	InternalValue() *ArtifactRegistryRepositoryCleanupPoliciesCondition
	SetInternalValue(val *ArtifactRegistryRepositoryCleanupPoliciesCondition)
	NewerThan() *string
	SetNewerThan(val *string)
	NewerThanInput() *string
	OlderThan() *string
	SetOlderThan(val *string)
	OlderThanInput() *string
	PackageNamePrefixes() *[]*string
	SetPackageNamePrefixes(val *[]*string)
	PackageNamePrefixesInput() *[]*string
	TagPrefixes() *[]*string
	SetTagPrefixes(val *[]*string)
	TagPrefixesInput() *[]*string
	TagState() *string
	SetTagState(val *string)
	TagStateInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VersionNamePrefixes() *[]*string
	SetVersionNamePrefixes(val *[]*string)
	VersionNamePrefixesInput() *[]*string
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
	ResetNewerThan()
	ResetOlderThan()
	ResetPackageNamePrefixes()
	ResetTagPrefixes()
	ResetTagState()
	ResetVersionNamePrefixes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference
type jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) InternalValue() *ArtifactRegistryRepositoryCleanupPoliciesCondition {
	var returns *ArtifactRegistryRepositoryCleanupPoliciesCondition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) NewerThan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newerThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) NewerThanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newerThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) OlderThan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"olderThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) OlderThanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"olderThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) PackageNamePrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packageNamePrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) PackageNamePrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packageNamePrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) TagPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) TagPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) TagState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) TagStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) VersionNamePrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"versionNamePrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) VersionNamePrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"versionNamePrefixesInput",
		&returns,
	)
	return returns
}


func NewArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference {
	_init_.Initialize()

	if err := validateNewArtifactRegistryRepositoryCleanupPoliciesConditionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.artifactRegistryRepository.ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference_Override(a ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.artifactRegistryRepository.ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference)SetInternalValue(val *ArtifactRegistryRepositoryCleanupPoliciesCondition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference)SetNewerThan(val *string) {
	if err := j.validateSetNewerThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newerThan",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference)SetOlderThan(val *string) {
	if err := j.validateSetOlderThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"olderThan",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference)SetPackageNamePrefixes(val *[]*string) {
	if err := j.validateSetPackageNamePrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packageNamePrefixes",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference)SetTagPrefixes(val *[]*string) {
	if err := j.validateSetTagPrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagPrefixes",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference)SetTagState(val *string) {
	if err := j.validateSetTagStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagState",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference)SetVersionNamePrefixes(val *[]*string) {
	if err := j.validateSetVersionNamePrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionNamePrefixes",
		val,
	)
}

func (a *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) ResetNewerThan() {
	_jsii_.InvokeVoid(
		a,
		"resetNewerThan",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) ResetOlderThan() {
	_jsii_.InvokeVoid(
		a,
		"resetOlderThan",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) ResetPackageNamePrefixes() {
	_jsii_.InvokeVoid(
		a,
		"resetPackageNamePrefixes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) ResetTagPrefixes() {
	_jsii_.InvokeVoid(
		a,
		"resetTagPrefixes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) ResetTagState() {
	_jsii_.InvokeVoid(
		a,
		"resetTagState",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) ResetVersionNamePrefixes() {
	_jsii_.InvokeVoid(
		a,
		"resetVersionNamePrefixes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepositoryCleanupPoliciesConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

