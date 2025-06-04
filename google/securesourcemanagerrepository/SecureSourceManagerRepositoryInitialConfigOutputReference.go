// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securesourcemanagerrepository

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/securesourcemanagerrepository/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecureSourceManagerRepositoryInitialConfigOutputReference interface {
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
	DefaultBranch() *string
	SetDefaultBranch(val *string)
	DefaultBranchInput() *string
	// Experimental.
	Fqn() *string
	Gitignores() *[]*string
	SetGitignores(val *[]*string)
	GitignoresInput() *[]*string
	InternalValue() *SecureSourceManagerRepositoryInitialConfig
	SetInternalValue(val *SecureSourceManagerRepositoryInitialConfig)
	License() *string
	SetLicense(val *string)
	LicenseInput() *string
	Readme() *string
	SetReadme(val *string)
	ReadmeInput() *string
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
	ResetDefaultBranch()
	ResetGitignores()
	ResetLicense()
	ResetReadme()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecureSourceManagerRepositoryInitialConfigOutputReference
type jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) DefaultBranch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) DefaultBranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBranchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) Gitignores() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gitignores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) GitignoresInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gitignoresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) InternalValue() *SecureSourceManagerRepositoryInitialConfig {
	var returns *SecureSourceManagerRepositoryInitialConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) License() *string {
	var returns *string
	_jsii_.Get(
		j,
		"license",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) LicenseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) Readme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) ReadmeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readmeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSecureSourceManagerRepositoryInitialConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SecureSourceManagerRepositoryInitialConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSecureSourceManagerRepositoryInitialConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.secureSourceManagerRepository.SecureSourceManagerRepositoryInitialConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSecureSourceManagerRepositoryInitialConfigOutputReference_Override(s SecureSourceManagerRepositoryInitialConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.secureSourceManagerRepository.SecureSourceManagerRepositoryInitialConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference)SetDefaultBranch(val *string) {
	if err := j.validateSetDefaultBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultBranch",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference)SetGitignores(val *[]*string) {
	if err := j.validateSetGitignoresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitignores",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference)SetInternalValue(val *SecureSourceManagerRepositoryInitialConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference)SetLicense(val *string) {
	if err := j.validateSetLicenseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"license",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference)SetReadme(val *string) {
	if err := j.validateSetReadmeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readme",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) ResetDefaultBranch() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultBranch",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) ResetGitignores() {
	_jsii_.InvokeVoid(
		s,
		"resetGitignores",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) ResetLicense() {
	_jsii_.InvokeVoid(
		s,
		"resetLicense",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) ResetReadme() {
	_jsii_.InvokeVoid(
		s,
		"resetReadme",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerRepositoryInitialConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

