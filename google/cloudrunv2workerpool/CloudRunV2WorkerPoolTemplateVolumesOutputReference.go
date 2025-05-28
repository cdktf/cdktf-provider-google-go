// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2workerpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/cloudrunv2workerpool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudRunV2WorkerPoolTemplateVolumesOutputReference interface {
	cdktf.ComplexObject
	CloudSqlInstance() CloudRunV2WorkerPoolTemplateVolumesCloudSqlInstanceOutputReference
	CloudSqlInstanceInput() *CloudRunV2WorkerPoolTemplateVolumesCloudSqlInstance
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
	EmptyDir() CloudRunV2WorkerPoolTemplateVolumesEmptyDirOutputReference
	EmptyDirInput() *CloudRunV2WorkerPoolTemplateVolumesEmptyDir
	// Experimental.
	Fqn() *string
	Gcs() CloudRunV2WorkerPoolTemplateVolumesGcsOutputReference
	GcsInput() *CloudRunV2WorkerPoolTemplateVolumesGcs
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nfs() CloudRunV2WorkerPoolTemplateVolumesNfsOutputReference
	NfsInput() *CloudRunV2WorkerPoolTemplateVolumesNfs
	Secret() CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference
	SecretInput() *CloudRunV2WorkerPoolTemplateVolumesSecret
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
	PutCloudSqlInstance(value *CloudRunV2WorkerPoolTemplateVolumesCloudSqlInstance)
	PutEmptyDir(value *CloudRunV2WorkerPoolTemplateVolumesEmptyDir)
	PutGcs(value *CloudRunV2WorkerPoolTemplateVolumesGcs)
	PutNfs(value *CloudRunV2WorkerPoolTemplateVolumesNfs)
	PutSecret(value *CloudRunV2WorkerPoolTemplateVolumesSecret)
	ResetCloudSqlInstance()
	ResetEmptyDir()
	ResetGcs()
	ResetNfs()
	ResetSecret()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudRunV2WorkerPoolTemplateVolumesOutputReference
type jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) CloudSqlInstance() CloudRunV2WorkerPoolTemplateVolumesCloudSqlInstanceOutputReference {
	var returns CloudRunV2WorkerPoolTemplateVolumesCloudSqlInstanceOutputReference
	_jsii_.Get(
		j,
		"cloudSqlInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) CloudSqlInstanceInput() *CloudRunV2WorkerPoolTemplateVolumesCloudSqlInstance {
	var returns *CloudRunV2WorkerPoolTemplateVolumesCloudSqlInstance
	_jsii_.Get(
		j,
		"cloudSqlInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) EmptyDir() CloudRunV2WorkerPoolTemplateVolumesEmptyDirOutputReference {
	var returns CloudRunV2WorkerPoolTemplateVolumesEmptyDirOutputReference
	_jsii_.Get(
		j,
		"emptyDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) EmptyDirInput() *CloudRunV2WorkerPoolTemplateVolumesEmptyDir {
	var returns *CloudRunV2WorkerPoolTemplateVolumesEmptyDir
	_jsii_.Get(
		j,
		"emptyDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) Gcs() CloudRunV2WorkerPoolTemplateVolumesGcsOutputReference {
	var returns CloudRunV2WorkerPoolTemplateVolumesGcsOutputReference
	_jsii_.Get(
		j,
		"gcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) GcsInput() *CloudRunV2WorkerPoolTemplateVolumesGcs {
	var returns *CloudRunV2WorkerPoolTemplateVolumesGcs
	_jsii_.Get(
		j,
		"gcsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) Nfs() CloudRunV2WorkerPoolTemplateVolumesNfsOutputReference {
	var returns CloudRunV2WorkerPoolTemplateVolumesNfsOutputReference
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) NfsInput() *CloudRunV2WorkerPoolTemplateVolumesNfs {
	var returns *CloudRunV2WorkerPoolTemplateVolumesNfs
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) Secret() CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference {
	var returns CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) SecretInput() *CloudRunV2WorkerPoolTemplateVolumesSecret {
	var returns *CloudRunV2WorkerPoolTemplateVolumesSecret
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudRunV2WorkerPoolTemplateVolumesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CloudRunV2WorkerPoolTemplateVolumesOutputReference {
	_init_.Initialize()

	if err := validateNewCloudRunV2WorkerPoolTemplateVolumesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudRunV2WorkerPool.CloudRunV2WorkerPoolTemplateVolumesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCloudRunV2WorkerPoolTemplateVolumesOutputReference_Override(c CloudRunV2WorkerPoolTemplateVolumesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudRunV2WorkerPool.CloudRunV2WorkerPoolTemplateVolumesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) PutCloudSqlInstance(value *CloudRunV2WorkerPoolTemplateVolumesCloudSqlInstance) {
	if err := c.validatePutCloudSqlInstanceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCloudSqlInstance",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) PutEmptyDir(value *CloudRunV2WorkerPoolTemplateVolumesEmptyDir) {
	if err := c.validatePutEmptyDirParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEmptyDir",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) PutGcs(value *CloudRunV2WorkerPoolTemplateVolumesGcs) {
	if err := c.validatePutGcsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGcs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) PutNfs(value *CloudRunV2WorkerPoolTemplateVolumesNfs) {
	if err := c.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNfs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) PutSecret(value *CloudRunV2WorkerPoolTemplateVolumesSecret) {
	if err := c.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSecret",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) ResetCloudSqlInstance() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudSqlInstance",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) ResetEmptyDir() {
	_jsii_.InvokeVoid(
		c,
		"resetEmptyDir",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) ResetGcs() {
	_jsii_.InvokeVoid(
		c,
		"resetGcs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		c,
		"resetNfs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		c,
		"resetSecret",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

