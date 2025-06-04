// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notebooksruntime

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/notebooksruntime/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference interface {
	cdktf.ComplexObject
	AcceleratorConfig() NotebooksRuntimeVirtualMachineVirtualMachineConfigAcceleratorConfigOutputReference
	AcceleratorConfigInput() *NotebooksRuntimeVirtualMachineVirtualMachineConfigAcceleratorConfig
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
	ContainerImages() NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList
	ContainerImagesInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataDisk() NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference
	DataDiskInput() *NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDisk
	EncryptionConfig() NotebooksRuntimeVirtualMachineVirtualMachineConfigEncryptionConfigOutputReference
	EncryptionConfigInput() *NotebooksRuntimeVirtualMachineVirtualMachineConfigEncryptionConfig
	// Experimental.
	Fqn() *string
	GuestAttributes() cdktf.StringMap
	InternalIpOnly() interface{}
	SetInternalIpOnly(val interface{})
	InternalIpOnlyInput() interface{}
	InternalValue() *NotebooksRuntimeVirtualMachineVirtualMachineConfig
	SetInternalValue(val *NotebooksRuntimeVirtualMachineVirtualMachineConfig)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataInput() *map[string]*string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	NicType() *string
	SetNicType(val *string)
	NicTypeInput() *string
	ReservedIpRange() *string
	SetReservedIpRange(val *string)
	ReservedIpRangeInput() *string
	ShieldedInstanceConfig() NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference
	ShieldedInstanceConfigInput() *NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfig
	Subnet() *string
	SetSubnet(val *string)
	SubnetInput() *string
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
	Zone() *string
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
	PutAcceleratorConfig(value *NotebooksRuntimeVirtualMachineVirtualMachineConfigAcceleratorConfig)
	PutContainerImages(value interface{})
	PutDataDisk(value *NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDisk)
	PutEncryptionConfig(value *NotebooksRuntimeVirtualMachineVirtualMachineConfigEncryptionConfig)
	PutShieldedInstanceConfig(value *NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfig)
	ResetAcceleratorConfig()
	ResetContainerImages()
	ResetEncryptionConfig()
	ResetInternalIpOnly()
	ResetLabels()
	ResetMetadata()
	ResetNetwork()
	ResetNicType()
	ResetReservedIpRange()
	ResetShieldedInstanceConfig()
	ResetSubnet()
	ResetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference
type jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) AcceleratorConfig() NotebooksRuntimeVirtualMachineVirtualMachineConfigAcceleratorConfigOutputReference {
	var returns NotebooksRuntimeVirtualMachineVirtualMachineConfigAcceleratorConfigOutputReference
	_jsii_.Get(
		j,
		"acceleratorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) AcceleratorConfigInput() *NotebooksRuntimeVirtualMachineVirtualMachineConfigAcceleratorConfig {
	var returns *NotebooksRuntimeVirtualMachineVirtualMachineConfigAcceleratorConfig
	_jsii_.Get(
		j,
		"acceleratorConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ContainerImages() NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList {
	var returns NotebooksRuntimeVirtualMachineVirtualMachineConfigContainerImagesList
	_jsii_.Get(
		j,
		"containerImages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ContainerImagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerImagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) DataDisk() NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference {
	var returns NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference
	_jsii_.Get(
		j,
		"dataDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) DataDiskInput() *NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDisk {
	var returns *NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDisk
	_jsii_.Get(
		j,
		"dataDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) EncryptionConfig() NotebooksRuntimeVirtualMachineVirtualMachineConfigEncryptionConfigOutputReference {
	var returns NotebooksRuntimeVirtualMachineVirtualMachineConfigEncryptionConfigOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) EncryptionConfigInput() *NotebooksRuntimeVirtualMachineVirtualMachineConfigEncryptionConfig {
	var returns *NotebooksRuntimeVirtualMachineVirtualMachineConfigEncryptionConfig
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) GuestAttributes() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"guestAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) InternalIpOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalIpOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) InternalIpOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalIpOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) InternalValue() *NotebooksRuntimeVirtualMachineVirtualMachineConfig {
	var returns *NotebooksRuntimeVirtualMachineVirtualMachineConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) NicType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nicType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) NicTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nicTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ReservedIpRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reservedIpRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ReservedIpRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reservedIpRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ShieldedInstanceConfig() NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference {
	var returns NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ShieldedInstanceConfigInput() *NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfig {
	var returns *NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfig
	_jsii_.Get(
		j,
		"shieldedInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) Subnet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) SubnetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}


func NewNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference {
	_init_.Initialize()

	if err := validateNewNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.notebooksRuntime.NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference_Override(n NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.notebooksRuntime.NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetInternalIpOnly(val interface{}) {
	if err := j.validateSetInternalIpOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalIpOnly",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetInternalValue(val *NotebooksRuntimeVirtualMachineVirtualMachineConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetNicType(val *string) {
	if err := j.validateSetNicTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nicType",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetReservedIpRange(val *string) {
	if err := j.validateSetReservedIpRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservedIpRange",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetSubnet(val *string) {
	if err := j.validateSetSubnetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnet",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) PutAcceleratorConfig(value *NotebooksRuntimeVirtualMachineVirtualMachineConfigAcceleratorConfig) {
	if err := n.validatePutAcceleratorConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putAcceleratorConfig",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) PutContainerImages(value interface{}) {
	if err := n.validatePutContainerImagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putContainerImages",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) PutDataDisk(value *NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDisk) {
	if err := n.validatePutDataDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putDataDisk",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) PutEncryptionConfig(value *NotebooksRuntimeVirtualMachineVirtualMachineConfigEncryptionConfig) {
	if err := n.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) PutShieldedInstanceConfig(value *NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfig) {
	if err := n.validatePutShieldedInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putShieldedInstanceConfig",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ResetAcceleratorConfig() {
	_jsii_.InvokeVoid(
		n,
		"resetAcceleratorConfig",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ResetContainerImages() {
	_jsii_.InvokeVoid(
		n,
		"resetContainerImages",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		n,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ResetInternalIpOnly() {
	_jsii_.InvokeVoid(
		n,
		"resetInternalIpOnly",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		n,
		"resetLabels",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		n,
		"resetMetadata",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		n,
		"resetNetwork",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ResetNicType() {
	_jsii_.InvokeVoid(
		n,
		"resetNicType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ResetReservedIpRange() {
	_jsii_.InvokeVoid(
		n,
		"resetReservedIpRange",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ResetShieldedInstanceConfig() {
	_jsii_.InvokeVoid(
		n,
		"resetShieldedInstanceConfig",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ResetSubnet() {
	_jsii_.InvokeVoid(
		n,
		"resetSubnet",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		n,
		"resetTags",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksRuntimeVirtualMachineVirtualMachineConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

