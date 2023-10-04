// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbuildtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v10/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v10/cloudbuildtrigger/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudbuildTriggerBuildOutputReference interface {
	cdktf.ComplexObject
	Artifacts() CloudbuildTriggerBuildArtifactsOutputReference
	ArtifactsInput() *CloudbuildTriggerBuildArtifacts
	AvailableSecrets() CloudbuildTriggerBuildAvailableSecretsOutputReference
	AvailableSecretsInput() *CloudbuildTriggerBuildAvailableSecrets
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
	Images() *[]*string
	SetImages(val *[]*string)
	ImagesInput() *[]*string
	InternalValue() *CloudbuildTriggerBuild
	SetInternalValue(val *CloudbuildTriggerBuild)
	LogsBucket() *string
	SetLogsBucket(val *string)
	LogsBucketInput() *string
	Options() CloudbuildTriggerBuildOptionsOutputReference
	OptionsInput() *CloudbuildTriggerBuildOptions
	QueueTtl() *string
	SetQueueTtl(val *string)
	QueueTtlInput() *string
	Secret() CloudbuildTriggerBuildSecretList
	SecretInput() interface{}
	Source() CloudbuildTriggerBuildSourceOutputReference
	SourceInput() *CloudbuildTriggerBuildSource
	Step() CloudbuildTriggerBuildStepList
	StepInput() interface{}
	Substitutions() *map[string]*string
	SetSubstitutions(val *map[string]*string)
	SubstitutionsInput() *map[string]*string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *string
	SetTimeout(val *string)
	TimeoutInput() *string
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
	PutArtifacts(value *CloudbuildTriggerBuildArtifacts)
	PutAvailableSecrets(value *CloudbuildTriggerBuildAvailableSecrets)
	PutOptions(value *CloudbuildTriggerBuildOptions)
	PutSecret(value interface{})
	PutSource(value *CloudbuildTriggerBuildSource)
	PutStep(value interface{})
	ResetArtifacts()
	ResetAvailableSecrets()
	ResetImages()
	ResetLogsBucket()
	ResetOptions()
	ResetQueueTtl()
	ResetSecret()
	ResetSource()
	ResetSubstitutions()
	ResetTags()
	ResetTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudbuildTriggerBuildOutputReference
type jsiiProxy_CloudbuildTriggerBuildOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) Artifacts() CloudbuildTriggerBuildArtifactsOutputReference {
	var returns CloudbuildTriggerBuildArtifactsOutputReference
	_jsii_.Get(
		j,
		"artifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) ArtifactsInput() *CloudbuildTriggerBuildArtifacts {
	var returns *CloudbuildTriggerBuildArtifacts
	_jsii_.Get(
		j,
		"artifactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) AvailableSecrets() CloudbuildTriggerBuildAvailableSecretsOutputReference {
	var returns CloudbuildTriggerBuildAvailableSecretsOutputReference
	_jsii_.Get(
		j,
		"availableSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) AvailableSecretsInput() *CloudbuildTriggerBuildAvailableSecrets {
	var returns *CloudbuildTriggerBuildAvailableSecrets
	_jsii_.Get(
		j,
		"availableSecretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) Images() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"images",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) ImagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"imagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) InternalValue() *CloudbuildTriggerBuild {
	var returns *CloudbuildTriggerBuild
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) LogsBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logsBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) LogsBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logsBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) Options() CloudbuildTriggerBuildOptionsOutputReference {
	var returns CloudbuildTriggerBuildOptionsOutputReference
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) OptionsInput() *CloudbuildTriggerBuildOptions {
	var returns *CloudbuildTriggerBuildOptions
	_jsii_.Get(
		j,
		"optionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) QueueTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) QueueTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) Secret() CloudbuildTriggerBuildSecretList {
	var returns CloudbuildTriggerBuildSecretList
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) Source() CloudbuildTriggerBuildSourceOutputReference {
	var returns CloudbuildTriggerBuildSourceOutputReference
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) SourceInput() *CloudbuildTriggerBuildSource {
	var returns *CloudbuildTriggerBuildSource
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) Step() CloudbuildTriggerBuildStepList {
	var returns CloudbuildTriggerBuildStepList
	_jsii_.Get(
		j,
		"step",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) StepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) Substitutions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"substitutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) SubstitutionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"substitutionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference) TimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}


func NewCloudbuildTriggerBuildOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudbuildTriggerBuildOutputReference {
	_init_.Initialize()

	if err := validateNewCloudbuildTriggerBuildOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudbuildTriggerBuildOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudbuildTrigger.CloudbuildTriggerBuildOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudbuildTriggerBuildOutputReference_Override(c CloudbuildTriggerBuildOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudbuildTrigger.CloudbuildTriggerBuildOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference)SetImages(val *[]*string) {
	if err := j.validateSetImagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"images",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference)SetInternalValue(val *CloudbuildTriggerBuild) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference)SetLogsBucket(val *string) {
	if err := j.validateSetLogsBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logsBucket",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference)SetQueueTtl(val *string) {
	if err := j.validateSetQueueTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queueTtl",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference)SetSubstitutions(val *map[string]*string) {
	if err := j.validateSetSubstitutionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"substitutions",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildOutputReference)SetTimeout(val *string) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) PutArtifacts(value *CloudbuildTriggerBuildArtifacts) {
	if err := c.validatePutArtifactsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putArtifacts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) PutAvailableSecrets(value *CloudbuildTriggerBuildAvailableSecrets) {
	if err := c.validatePutAvailableSecretsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAvailableSecrets",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) PutOptions(value *CloudbuildTriggerBuildOptions) {
	if err := c.validatePutOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOptions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) PutSecret(value interface{}) {
	if err := c.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSecret",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) PutSource(value *CloudbuildTriggerBuildSource) {
	if err := c.validatePutSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSource",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) PutStep(value interface{}) {
	if err := c.validatePutStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStep",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) ResetArtifacts() {
	_jsii_.InvokeVoid(
		c,
		"resetArtifacts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) ResetAvailableSecrets() {
	_jsii_.InvokeVoid(
		c,
		"resetAvailableSecrets",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) ResetImages() {
	_jsii_.InvokeVoid(
		c,
		"resetImages",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) ResetLogsBucket() {
	_jsii_.InvokeVoid(
		c,
		"resetLogsBucket",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) ResetOptions() {
	_jsii_.InvokeVoid(
		c,
		"resetOptions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) ResetQueueTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetQueueTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		c,
		"resetSecret",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		c,
		"resetSource",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) ResetSubstitutions() {
	_jsii_.InvokeVoid(
		c,
		"resetSubstitutions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudbuildTriggerBuildOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

