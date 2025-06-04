// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/containercluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerClusterUserManagedKeysConfigOutputReference interface {
	cdktf.ComplexObject
	AggregationCa() *string
	SetAggregationCa(val *string)
	AggregationCaInput() *string
	ClusterCa() *string
	SetClusterCa(val *string)
	ClusterCaInput() *string
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
	ControlPlaneDiskEncryptionKey() *string
	SetControlPlaneDiskEncryptionKey(val *string)
	ControlPlaneDiskEncryptionKeyInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EtcdApiCa() *string
	SetEtcdApiCa(val *string)
	EtcdApiCaInput() *string
	EtcdPeerCa() *string
	SetEtcdPeerCa(val *string)
	EtcdPeerCaInput() *string
	// Experimental.
	Fqn() *string
	GkeopsEtcdBackupEncryptionKey() *string
	SetGkeopsEtcdBackupEncryptionKey(val *string)
	GkeopsEtcdBackupEncryptionKeyInput() *string
	InternalValue() *ContainerClusterUserManagedKeysConfig
	SetInternalValue(val *ContainerClusterUserManagedKeysConfig)
	ServiceAccountSigningKeys() *[]*string
	SetServiceAccountSigningKeys(val *[]*string)
	ServiceAccountSigningKeysInput() *[]*string
	ServiceAccountVerificationKeys() *[]*string
	SetServiceAccountVerificationKeys(val *[]*string)
	ServiceAccountVerificationKeysInput() *[]*string
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
	ResetAggregationCa()
	ResetClusterCa()
	ResetControlPlaneDiskEncryptionKey()
	ResetEtcdApiCa()
	ResetEtcdPeerCa()
	ResetGkeopsEtcdBackupEncryptionKey()
	ResetServiceAccountSigningKeys()
	ResetServiceAccountVerificationKeys()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerClusterUserManagedKeysConfigOutputReference
type jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) AggregationCa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) AggregationCaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) ClusterCa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) ClusterCaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) ControlPlaneDiskEncryptionKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlPlaneDiskEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) ControlPlaneDiskEncryptionKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlPlaneDiskEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) EtcdApiCa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etcdApiCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) EtcdApiCaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etcdApiCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) EtcdPeerCa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etcdPeerCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) EtcdPeerCaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etcdPeerCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) GkeopsEtcdBackupEncryptionKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeopsEtcdBackupEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) GkeopsEtcdBackupEncryptionKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeopsEtcdBackupEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) InternalValue() *ContainerClusterUserManagedKeysConfig {
	var returns *ContainerClusterUserManagedKeysConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) ServiceAccountSigningKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceAccountSigningKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) ServiceAccountSigningKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceAccountSigningKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) ServiceAccountVerificationKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceAccountVerificationKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) ServiceAccountVerificationKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceAccountVerificationKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerClusterUserManagedKeysConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerClusterUserManagedKeysConfigOutputReference {
	_init_.Initialize()

	if err := validateNewContainerClusterUserManagedKeysConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterUserManagedKeysConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerClusterUserManagedKeysConfigOutputReference_Override(c ContainerClusterUserManagedKeysConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterUserManagedKeysConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference)SetAggregationCa(val *string) {
	if err := j.validateSetAggregationCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggregationCa",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference)SetClusterCa(val *string) {
	if err := j.validateSetClusterCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterCa",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference)SetControlPlaneDiskEncryptionKey(val *string) {
	if err := j.validateSetControlPlaneDiskEncryptionKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"controlPlaneDiskEncryptionKey",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference)SetEtcdApiCa(val *string) {
	if err := j.validateSetEtcdApiCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"etcdApiCa",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference)SetEtcdPeerCa(val *string) {
	if err := j.validateSetEtcdPeerCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"etcdPeerCa",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference)SetGkeopsEtcdBackupEncryptionKey(val *string) {
	if err := j.validateSetGkeopsEtcdBackupEncryptionKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gkeopsEtcdBackupEncryptionKey",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference)SetInternalValue(val *ContainerClusterUserManagedKeysConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference)SetServiceAccountSigningKeys(val *[]*string) {
	if err := j.validateSetServiceAccountSigningKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountSigningKeys",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference)SetServiceAccountVerificationKeys(val *[]*string) {
	if err := j.validateSetServiceAccountVerificationKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountVerificationKeys",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) ResetAggregationCa() {
	_jsii_.InvokeVoid(
		c,
		"resetAggregationCa",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) ResetClusterCa() {
	_jsii_.InvokeVoid(
		c,
		"resetClusterCa",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) ResetControlPlaneDiskEncryptionKey() {
	_jsii_.InvokeVoid(
		c,
		"resetControlPlaneDiskEncryptionKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) ResetEtcdApiCa() {
	_jsii_.InvokeVoid(
		c,
		"resetEtcdApiCa",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) ResetEtcdPeerCa() {
	_jsii_.InvokeVoid(
		c,
		"resetEtcdPeerCa",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) ResetGkeopsEtcdBackupEncryptionKey() {
	_jsii_.InvokeVoid(
		c,
		"resetGkeopsEtcdBackupEncryptionKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) ResetServiceAccountSigningKeys() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceAccountSigningKeys",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) ResetServiceAccountVerificationKeys() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceAccountVerificationKeys",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerClusterUserManagedKeysConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

