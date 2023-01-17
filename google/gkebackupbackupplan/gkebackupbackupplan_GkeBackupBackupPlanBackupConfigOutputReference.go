package gkebackupbackupplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v5/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v5/gkebackupbackupplan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeBackupBackupPlanBackupConfigOutputReference interface {
	cdktf.ComplexObject
	AllNamespaces() interface{}
	SetAllNamespaces(val interface{})
	AllNamespacesInput() interface{}
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
	EncryptionKey() GkeBackupBackupPlanBackupConfigEncryptionKeyOutputReference
	EncryptionKeyInput() *GkeBackupBackupPlanBackupConfigEncryptionKey
	// Experimental.
	Fqn() *string
	IncludeSecrets() interface{}
	SetIncludeSecrets(val interface{})
	IncludeSecretsInput() interface{}
	IncludeVolumeData() interface{}
	SetIncludeVolumeData(val interface{})
	IncludeVolumeDataInput() interface{}
	InternalValue() *GkeBackupBackupPlanBackupConfig
	SetInternalValue(val *GkeBackupBackupPlanBackupConfig)
	SelectedApplications() GkeBackupBackupPlanBackupConfigSelectedApplicationsOutputReference
	SelectedApplicationsInput() *GkeBackupBackupPlanBackupConfigSelectedApplications
	SelectedNamespaces() GkeBackupBackupPlanBackupConfigSelectedNamespacesOutputReference
	SelectedNamespacesInput() *GkeBackupBackupPlanBackupConfigSelectedNamespaces
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
	PutEncryptionKey(value *GkeBackupBackupPlanBackupConfigEncryptionKey)
	PutSelectedApplications(value *GkeBackupBackupPlanBackupConfigSelectedApplications)
	PutSelectedNamespaces(value *GkeBackupBackupPlanBackupConfigSelectedNamespaces)
	ResetAllNamespaces()
	ResetEncryptionKey()
	ResetIncludeSecrets()
	ResetIncludeVolumeData()
	ResetSelectedApplications()
	ResetSelectedNamespaces()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeBackupBackupPlanBackupConfigOutputReference
type jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) AllNamespaces() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) AllNamespacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) EncryptionKey() GkeBackupBackupPlanBackupConfigEncryptionKeyOutputReference {
	var returns GkeBackupBackupPlanBackupConfigEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) EncryptionKeyInput() *GkeBackupBackupPlanBackupConfigEncryptionKey {
	var returns *GkeBackupBackupPlanBackupConfigEncryptionKey
	_jsii_.Get(
		j,
		"encryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) IncludeSecrets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) IncludeSecretsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSecretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) IncludeVolumeData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeVolumeData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) IncludeVolumeDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeVolumeDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) InternalValue() *GkeBackupBackupPlanBackupConfig {
	var returns *GkeBackupBackupPlanBackupConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) SelectedApplications() GkeBackupBackupPlanBackupConfigSelectedApplicationsOutputReference {
	var returns GkeBackupBackupPlanBackupConfigSelectedApplicationsOutputReference
	_jsii_.Get(
		j,
		"selectedApplications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) SelectedApplicationsInput() *GkeBackupBackupPlanBackupConfigSelectedApplications {
	var returns *GkeBackupBackupPlanBackupConfigSelectedApplications
	_jsii_.Get(
		j,
		"selectedApplicationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) SelectedNamespaces() GkeBackupBackupPlanBackupConfigSelectedNamespacesOutputReference {
	var returns GkeBackupBackupPlanBackupConfigSelectedNamespacesOutputReference
	_jsii_.Get(
		j,
		"selectedNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) SelectedNamespacesInput() *GkeBackupBackupPlanBackupConfigSelectedNamespaces {
	var returns *GkeBackupBackupPlanBackupConfigSelectedNamespaces
	_jsii_.Get(
		j,
		"selectedNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGkeBackupBackupPlanBackupConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GkeBackupBackupPlanBackupConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGkeBackupBackupPlanBackupConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeBackupBackupPlan.GkeBackupBackupPlanBackupConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGkeBackupBackupPlanBackupConfigOutputReference_Override(g GkeBackupBackupPlanBackupConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeBackupBackupPlan.GkeBackupBackupPlanBackupConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference)SetAllNamespaces(val interface{}) {
	if err := j.validateSetAllNamespacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allNamespaces",
		val,
	)
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference)SetIncludeSecrets(val interface{}) {
	if err := j.validateSetIncludeSecretsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSecrets",
		val,
	)
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference)SetIncludeVolumeData(val interface{}) {
	if err := j.validateSetIncludeVolumeDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeVolumeData",
		val,
	)
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference)SetInternalValue(val *GkeBackupBackupPlanBackupConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) PutEncryptionKey(value *GkeBackupBackupPlanBackupConfigEncryptionKey) {
	if err := g.validatePutEncryptionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEncryptionKey",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) PutSelectedApplications(value *GkeBackupBackupPlanBackupConfigSelectedApplications) {
	if err := g.validatePutSelectedApplicationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSelectedApplications",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) PutSelectedNamespaces(value *GkeBackupBackupPlanBackupConfigSelectedNamespaces) {
	if err := g.validatePutSelectedNamespacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSelectedNamespaces",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) ResetAllNamespaces() {
	_jsii_.InvokeVoid(
		g,
		"resetAllNamespaces",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) ResetEncryptionKey() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptionKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) ResetIncludeSecrets() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeSecrets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) ResetIncludeVolumeData() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeVolumeData",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) ResetSelectedApplications() {
	_jsii_.InvokeVoid(
		g,
		"resetSelectedApplications",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) ResetSelectedNamespaces() {
	_jsii_.InvokeVoid(
		g,
		"resetSelectedNamespaces",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GkeBackupBackupPlanBackupConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

