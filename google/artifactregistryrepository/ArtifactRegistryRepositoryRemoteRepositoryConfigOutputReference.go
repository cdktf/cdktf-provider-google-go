package artifactregistryrepository

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v6/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v6/artifactregistryrepository/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference interface {
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DockerRepository() ArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference
	DockerRepositoryInput() *ArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepository
	// Experimental.
	Fqn() *string
	InternalValue() *ArtifactRegistryRepositoryRemoteRepositoryConfig
	SetInternalValue(val *ArtifactRegistryRepositoryRemoteRepositoryConfig)
	MavenRepository() ArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepositoryOutputReference
	MavenRepositoryInput() *ArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepository
	NpmRepository() ArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepositoryOutputReference
	NpmRepositoryInput() *ArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepository
	PythonRepository() ArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference
	PythonRepositoryInput() *ArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepository
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
	PutDockerRepository(value *ArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepository)
	PutMavenRepository(value *ArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepository)
	PutNpmRepository(value *ArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepository)
	PutPythonRepository(value *ArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepository)
	ResetDescription()
	ResetDockerRepository()
	ResetMavenRepository()
	ResetNpmRepository()
	ResetPythonRepository()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference
type jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) DockerRepository() ArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference {
	var returns ArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryOutputReference
	_jsii_.Get(
		j,
		"dockerRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) DockerRepositoryInput() *ArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepository {
	var returns *ArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepository
	_jsii_.Get(
		j,
		"dockerRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) InternalValue() *ArtifactRegistryRepositoryRemoteRepositoryConfig {
	var returns *ArtifactRegistryRepositoryRemoteRepositoryConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) MavenRepository() ArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepositoryOutputReference {
	var returns ArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepositoryOutputReference
	_jsii_.Get(
		j,
		"mavenRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) MavenRepositoryInput() *ArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepository {
	var returns *ArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepository
	_jsii_.Get(
		j,
		"mavenRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) NpmRepository() ArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepositoryOutputReference {
	var returns ArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepositoryOutputReference
	_jsii_.Get(
		j,
		"npmRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) NpmRepositoryInput() *ArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepository {
	var returns *ArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepository
	_jsii_.Get(
		j,
		"npmRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) PythonRepository() ArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference {
	var returns ArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference
	_jsii_.Get(
		j,
		"pythonRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) PythonRepositoryInput() *ArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepository {
	var returns *ArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepository
	_jsii_.Get(
		j,
		"pythonRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference {
	_init_.Initialize()

	if err := validateNewArtifactRegistryRepositoryRemoteRepositoryConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.artifactRegistryRepository.ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference_Override(a ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.artifactRegistryRepository.ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference)SetInternalValue(val *ArtifactRegistryRepositoryRemoteRepositoryConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) PutDockerRepository(value *ArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepository) {
	if err := a.validatePutDockerRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDockerRepository",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) PutMavenRepository(value *ArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepository) {
	if err := a.validatePutMavenRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMavenRepository",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) PutNpmRepository(value *ArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepository) {
	if err := a.validatePutNpmRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNpmRepository",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) PutPythonRepository(value *ArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepository) {
	if err := a.validatePutPythonRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPythonRepository",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ResetDockerRepository() {
	_jsii_.InvokeVoid(
		a,
		"resetDockerRepository",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ResetMavenRepository() {
	_jsii_.InvokeVoid(
		a,
		"resetMavenRepository",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ResetNpmRepository() {
	_jsii_.InvokeVoid(
		a,
		"resetNpmRepository",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ResetPythonRepository() {
	_jsii_.InvokeVoid(
		a,
		"resetPythonRepository",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

