// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwarecluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/gkeonpremvmwarecluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeonpremVmwareClusterVcenterOutputReference interface {
	cdktf.ComplexObject
	Address() *string
	CaCertData() *string
	SetCaCertData(val *string)
	CaCertDataInput() *string
	Cluster() *string
	SetCluster(val *string)
	ClusterInput() *string
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
	Datacenter() *string
	SetDatacenter(val *string)
	DatacenterInput() *string
	Datastore() *string
	SetDatastore(val *string)
	DatastoreInput() *string
	Folder() *string
	SetFolder(val *string)
	FolderInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GkeonpremVmwareClusterVcenter
	SetInternalValue(val *GkeonpremVmwareClusterVcenter)
	ResourcePool() *string
	SetResourcePool(val *string)
	ResourcePoolInput() *string
	StoragePolicyName() *string
	SetStoragePolicyName(val *string)
	StoragePolicyNameInput() *string
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
	ResetCaCertData()
	ResetCluster()
	ResetDatacenter()
	ResetDatastore()
	ResetFolder()
	ResetResourcePool()
	ResetStoragePolicyName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeonpremVmwareClusterVcenterOutputReference
type jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) CaCertData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) CaCertDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) Cluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) ClusterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) Datacenter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) DatacenterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) Datastore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) DatastoreInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) Folder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) FolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) InternalValue() *GkeonpremVmwareClusterVcenter {
	var returns *GkeonpremVmwareClusterVcenter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) ResourcePool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) ResourcePoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) StoragePolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storagePolicyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) StoragePolicyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storagePolicyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGkeonpremVmwareClusterVcenterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GkeonpremVmwareClusterVcenterOutputReference {
	_init_.Initialize()

	if err := validateNewGkeonpremVmwareClusterVcenterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremVmwareCluster.GkeonpremVmwareClusterVcenterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGkeonpremVmwareClusterVcenterOutputReference_Override(g GkeonpremVmwareClusterVcenterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremVmwareCluster.GkeonpremVmwareClusterVcenterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference)SetCaCertData(val *string) {
	if err := j.validateSetCaCertDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertData",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference)SetCluster(val *string) {
	if err := j.validateSetClusterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cluster",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference)SetDatacenter(val *string) {
	if err := j.validateSetDatacenterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datacenter",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference)SetDatastore(val *string) {
	if err := j.validateSetDatastoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datastore",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference)SetFolder(val *string) {
	if err := j.validateSetFolderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"folder",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference)SetInternalValue(val *GkeonpremVmwareClusterVcenter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference)SetResourcePool(val *string) {
	if err := j.validateSetResourcePoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcePool",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference)SetStoragePolicyName(val *string) {
	if err := j.validateSetStoragePolicyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storagePolicyName",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) ResetCaCertData() {
	_jsii_.InvokeVoid(
		g,
		"resetCaCertData",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) ResetCluster() {
	_jsii_.InvokeVoid(
		g,
		"resetCluster",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) ResetDatacenter() {
	_jsii_.InvokeVoid(
		g,
		"resetDatacenter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) ResetDatastore() {
	_jsii_.InvokeVoid(
		g,
		"resetDatastore",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) ResetFolder() {
	_jsii_.InvokeVoid(
		g,
		"resetFolder",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) ResetResourcePool() {
	_jsii_.InvokeVoid(
		g,
		"resetResourcePool",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) ResetStoragePolicyName() {
	_jsii_.InvokeVoid(
		g,
		"resetStoragePolicyName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GkeonpremVmwareClusterVcenterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

