// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagoogleartifactregistryrepository

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/datagoogleartifactregistryrepository/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference interface {
	cdktf.ComplexObject
	AptRepository() DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigAptRepositoryList
	CommonRepository() DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigCommonRepositoryList
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
	Description() *string
	DisableUpstreamValidation() cdktf.IResolvable
	DockerRepository() DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryList
	// Experimental.
	Fqn() *string
	InternalValue() *DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfig
	SetInternalValue(val *DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfig)
	MavenRepository() DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepositoryList
	NpmRepository() DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepositoryList
	PythonRepository() DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpstreamCredentials() DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentialsList
	YumRepository() DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigYumRepositoryList
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference
type jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) AptRepository() DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigAptRepositoryList {
	var returns DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigAptRepositoryList
	_jsii_.Get(
		j,
		"aptRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) CommonRepository() DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigCommonRepositoryList {
	var returns DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigCommonRepositoryList
	_jsii_.Get(
		j,
		"commonRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) DisableUpstreamValidation() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"disableUpstreamValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) DockerRepository() DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryList {
	var returns DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryList
	_jsii_.Get(
		j,
		"dockerRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) InternalValue() *DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfig {
	var returns *DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) MavenRepository() DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepositoryList {
	var returns DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepositoryList
	_jsii_.Get(
		j,
		"mavenRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) NpmRepository() DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepositoryList {
	var returns DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepositoryList
	_jsii_.Get(
		j,
		"npmRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) PythonRepository() DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryList {
	var returns DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryList
	_jsii_.Get(
		j,
		"pythonRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) UpstreamCredentials() DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentialsList {
	var returns DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentialsList
	_jsii_.Get(
		j,
		"upstreamCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) YumRepository() DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigYumRepositoryList {
	var returns DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigYumRepositoryList
	_jsii_.Get(
		j,
		"yumRepository",
		&returns,
	)
	return returns
}


func NewDataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleArtifactRegistryRepository.DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference_Override(d DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleArtifactRegistryRepository.DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference)SetInternalValue(val *DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataGoogleArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

