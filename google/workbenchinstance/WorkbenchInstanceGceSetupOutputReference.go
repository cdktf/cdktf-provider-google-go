// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workbenchinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/workbenchinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkbenchInstanceGceSetupOutputReference interface {
	cdktf.ComplexObject
	AcceleratorConfigs() WorkbenchInstanceGceSetupAcceleratorConfigsList
	AcceleratorConfigsInput() interface{}
	BootDisk() WorkbenchInstanceGceSetupBootDiskOutputReference
	BootDiskInput() *WorkbenchInstanceGceSetupBootDisk
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
	ContainerImage() WorkbenchInstanceGceSetupContainerImageOutputReference
	ContainerImageInput() *WorkbenchInstanceGceSetupContainerImage
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataDisks() WorkbenchInstanceGceSetupDataDisksOutputReference
	DataDisksInput() *WorkbenchInstanceGceSetupDataDisks
	DisablePublicIp() interface{}
	SetDisablePublicIp(val interface{})
	DisablePublicIpInput() interface{}
	EnableIpForwarding() interface{}
	SetEnableIpForwarding(val interface{})
	EnableIpForwardingInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *WorkbenchInstanceGceSetup
	SetInternalValue(val *WorkbenchInstanceGceSetup)
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataInput() *map[string]*string
	NetworkInterfaces() WorkbenchInstanceGceSetupNetworkInterfacesList
	NetworkInterfacesInput() interface{}
	ServiceAccounts() WorkbenchInstanceGceSetupServiceAccountsList
	ServiceAccountsInput() interface{}
	ShieldedInstanceConfig() WorkbenchInstanceGceSetupShieldedInstanceConfigOutputReference
	ShieldedInstanceConfigInput() *WorkbenchInstanceGceSetupShieldedInstanceConfig
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
	VmImage() WorkbenchInstanceGceSetupVmImageOutputReference
	VmImageInput() *WorkbenchInstanceGceSetupVmImage
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
	PutAcceleratorConfigs(value interface{})
	PutBootDisk(value *WorkbenchInstanceGceSetupBootDisk)
	PutContainerImage(value *WorkbenchInstanceGceSetupContainerImage)
	PutDataDisks(value *WorkbenchInstanceGceSetupDataDisks)
	PutNetworkInterfaces(value interface{})
	PutServiceAccounts(value interface{})
	PutShieldedInstanceConfig(value *WorkbenchInstanceGceSetupShieldedInstanceConfig)
	PutVmImage(value *WorkbenchInstanceGceSetupVmImage)
	ResetAcceleratorConfigs()
	ResetBootDisk()
	ResetContainerImage()
	ResetDataDisks()
	ResetDisablePublicIp()
	ResetEnableIpForwarding()
	ResetMachineType()
	ResetMetadata()
	ResetNetworkInterfaces()
	ResetServiceAccounts()
	ResetShieldedInstanceConfig()
	ResetTags()
	ResetVmImage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkbenchInstanceGceSetupOutputReference
type jsiiProxy_WorkbenchInstanceGceSetupOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) AcceleratorConfigs() WorkbenchInstanceGceSetupAcceleratorConfigsList {
	var returns WorkbenchInstanceGceSetupAcceleratorConfigsList
	_jsii_.Get(
		j,
		"acceleratorConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) AcceleratorConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceleratorConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) BootDisk() WorkbenchInstanceGceSetupBootDiskOutputReference {
	var returns WorkbenchInstanceGceSetupBootDiskOutputReference
	_jsii_.Get(
		j,
		"bootDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) BootDiskInput() *WorkbenchInstanceGceSetupBootDisk {
	var returns *WorkbenchInstanceGceSetupBootDisk
	_jsii_.Get(
		j,
		"bootDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) ContainerImage() WorkbenchInstanceGceSetupContainerImageOutputReference {
	var returns WorkbenchInstanceGceSetupContainerImageOutputReference
	_jsii_.Get(
		j,
		"containerImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) ContainerImageInput() *WorkbenchInstanceGceSetupContainerImage {
	var returns *WorkbenchInstanceGceSetupContainerImage
	_jsii_.Get(
		j,
		"containerImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) DataDisks() WorkbenchInstanceGceSetupDataDisksOutputReference {
	var returns WorkbenchInstanceGceSetupDataDisksOutputReference
	_jsii_.Get(
		j,
		"dataDisks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) DataDisksInput() *WorkbenchInstanceGceSetupDataDisks {
	var returns *WorkbenchInstanceGceSetupDataDisks
	_jsii_.Get(
		j,
		"dataDisksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) DisablePublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) DisablePublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) EnableIpForwarding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIpForwarding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) EnableIpForwardingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIpForwardingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) InternalValue() *WorkbenchInstanceGceSetup {
	var returns *WorkbenchInstanceGceSetup
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) NetworkInterfaces() WorkbenchInstanceGceSetupNetworkInterfacesList {
	var returns WorkbenchInstanceGceSetupNetworkInterfacesList
	_jsii_.Get(
		j,
		"networkInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) NetworkInterfacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) ServiceAccounts() WorkbenchInstanceGceSetupServiceAccountsList {
	var returns WorkbenchInstanceGceSetupServiceAccountsList
	_jsii_.Get(
		j,
		"serviceAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) ServiceAccountsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) ShieldedInstanceConfig() WorkbenchInstanceGceSetupShieldedInstanceConfigOutputReference {
	var returns WorkbenchInstanceGceSetupShieldedInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) ShieldedInstanceConfigInput() *WorkbenchInstanceGceSetupShieldedInstanceConfig {
	var returns *WorkbenchInstanceGceSetupShieldedInstanceConfig
	_jsii_.Get(
		j,
		"shieldedInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) VmImage() WorkbenchInstanceGceSetupVmImageOutputReference {
	var returns WorkbenchInstanceGceSetupVmImageOutputReference
	_jsii_.Get(
		j,
		"vmImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) VmImageInput() *WorkbenchInstanceGceSetupVmImage {
	var returns *WorkbenchInstanceGceSetupVmImage
	_jsii_.Get(
		j,
		"vmImageInput",
		&returns,
	)
	return returns
}


func NewWorkbenchInstanceGceSetupOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WorkbenchInstanceGceSetupOutputReference {
	_init_.Initialize()

	if err := validateNewWorkbenchInstanceGceSetupOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkbenchInstanceGceSetupOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.workbenchInstance.WorkbenchInstanceGceSetupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkbenchInstanceGceSetupOutputReference_Override(w WorkbenchInstanceGceSetupOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.workbenchInstance.WorkbenchInstanceGceSetupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference)SetDisablePublicIp(val interface{}) {
	if err := j.validateSetDisablePublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disablePublicIp",
		val,
	)
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference)SetEnableIpForwarding(val interface{}) {
	if err := j.validateSetEnableIpForwardingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableIpForwarding",
		val,
	)
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference)SetInternalValue(val *WorkbenchInstanceGceSetup) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) PutAcceleratorConfigs(value interface{}) {
	if err := w.validatePutAcceleratorConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAcceleratorConfigs",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) PutBootDisk(value *WorkbenchInstanceGceSetupBootDisk) {
	if err := w.validatePutBootDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBootDisk",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) PutContainerImage(value *WorkbenchInstanceGceSetupContainerImage) {
	if err := w.validatePutContainerImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putContainerImage",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) PutDataDisks(value *WorkbenchInstanceGceSetupDataDisks) {
	if err := w.validatePutDataDisksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putDataDisks",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) PutNetworkInterfaces(value interface{}) {
	if err := w.validatePutNetworkInterfacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putNetworkInterfaces",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) PutServiceAccounts(value interface{}) {
	if err := w.validatePutServiceAccountsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putServiceAccounts",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) PutShieldedInstanceConfig(value *WorkbenchInstanceGceSetupShieldedInstanceConfig) {
	if err := w.validatePutShieldedInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putShieldedInstanceConfig",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) PutVmImage(value *WorkbenchInstanceGceSetupVmImage) {
	if err := w.validatePutVmImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putVmImage",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) ResetAcceleratorConfigs() {
	_jsii_.InvokeVoid(
		w,
		"resetAcceleratorConfigs",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) ResetBootDisk() {
	_jsii_.InvokeVoid(
		w,
		"resetBootDisk",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) ResetContainerImage() {
	_jsii_.InvokeVoid(
		w,
		"resetContainerImage",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) ResetDataDisks() {
	_jsii_.InvokeVoid(
		w,
		"resetDataDisks",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) ResetDisablePublicIp() {
	_jsii_.InvokeVoid(
		w,
		"resetDisablePublicIp",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) ResetEnableIpForwarding() {
	_jsii_.InvokeVoid(
		w,
		"resetEnableIpForwarding",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) ResetMachineType() {
	_jsii_.InvokeVoid(
		w,
		"resetMachineType",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		w,
		"resetMetadata",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) ResetNetworkInterfaces() {
	_jsii_.InvokeVoid(
		w,
		"resetNetworkInterfaces",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) ResetServiceAccounts() {
	_jsii_.InvokeVoid(
		w,
		"resetServiceAccounts",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) ResetShieldedInstanceConfig() {
	_jsii_.InvokeVoid(
		w,
		"resetShieldedInstanceConfig",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		w,
		"resetTags",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) ResetVmImage() {
	_jsii_.InvokeVoid(
		w,
		"resetVmImage",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

