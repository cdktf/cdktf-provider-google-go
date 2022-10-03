package storagetransferjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-google-go/google/v3/jsii"

	"github.com/hashicorp/cdktf-provider-google-go/google/v3/storagetransferjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageTransferJobTransferSpecOutputReference interface {
	cdktf.ComplexObject
	AwsS3DataSource() StorageTransferJobTransferSpecAwsS3DataSourceOutputReference
	AwsS3DataSourceInput() *StorageTransferJobTransferSpecAwsS3DataSource
	AzureBlobStorageDataSource() StorageTransferJobTransferSpecAzureBlobStorageDataSourceOutputReference
	AzureBlobStorageDataSourceInput() *StorageTransferJobTransferSpecAzureBlobStorageDataSource
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
	GcsDataSink() StorageTransferJobTransferSpecGcsDataSinkOutputReference
	GcsDataSinkInput() *StorageTransferJobTransferSpecGcsDataSink
	GcsDataSource() StorageTransferJobTransferSpecGcsDataSourceOutputReference
	GcsDataSourceInput() *StorageTransferJobTransferSpecGcsDataSource
	HttpDataSource() StorageTransferJobTransferSpecHttpDataSourceOutputReference
	HttpDataSourceInput() *StorageTransferJobTransferSpecHttpDataSource
	InternalValue() *StorageTransferJobTransferSpec
	SetInternalValue(val *StorageTransferJobTransferSpec)
	ObjectConditions() StorageTransferJobTransferSpecObjectConditionsOutputReference
	ObjectConditionsInput() *StorageTransferJobTransferSpecObjectConditions
	PosixDataSink() StorageTransferJobTransferSpecPosixDataSinkOutputReference
	PosixDataSinkInput() *StorageTransferJobTransferSpecPosixDataSink
	PosixDataSource() StorageTransferJobTransferSpecPosixDataSourceOutputReference
	PosixDataSourceInput() *StorageTransferJobTransferSpecPosixDataSource
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransferOptions() StorageTransferJobTransferSpecTransferOptionsOutputReference
	TransferOptionsInput() *StorageTransferJobTransferSpecTransferOptions
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
	PutAwsS3DataSource(value *StorageTransferJobTransferSpecAwsS3DataSource)
	PutAzureBlobStorageDataSource(value *StorageTransferJobTransferSpecAzureBlobStorageDataSource)
	PutGcsDataSink(value *StorageTransferJobTransferSpecGcsDataSink)
	PutGcsDataSource(value *StorageTransferJobTransferSpecGcsDataSource)
	PutHttpDataSource(value *StorageTransferJobTransferSpecHttpDataSource)
	PutObjectConditions(value *StorageTransferJobTransferSpecObjectConditions)
	PutPosixDataSink(value *StorageTransferJobTransferSpecPosixDataSink)
	PutPosixDataSource(value *StorageTransferJobTransferSpecPosixDataSource)
	PutTransferOptions(value *StorageTransferJobTransferSpecTransferOptions)
	ResetAwsS3DataSource()
	ResetAzureBlobStorageDataSource()
	ResetGcsDataSink()
	ResetGcsDataSource()
	ResetHttpDataSource()
	ResetObjectConditions()
	ResetPosixDataSink()
	ResetPosixDataSource()
	ResetTransferOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageTransferJobTransferSpecOutputReference
type jsiiProxy_StorageTransferJobTransferSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference) AwsS3DataSource() StorageTransferJobTransferSpecAwsS3DataSourceOutputReference {
	var returns StorageTransferJobTransferSpecAwsS3DataSourceOutputReference
	_jsii_.Get(
		j,
		"awsS3DataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference) AwsS3DataSourceInput() *StorageTransferJobTransferSpecAwsS3DataSource {
	var returns *StorageTransferJobTransferSpecAwsS3DataSource
	_jsii_.Get(
		j,
		"awsS3DataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference) AzureBlobStorageDataSource() StorageTransferJobTransferSpecAzureBlobStorageDataSourceOutputReference {
	var returns StorageTransferJobTransferSpecAzureBlobStorageDataSourceOutputReference
	_jsii_.Get(
		j,
		"azureBlobStorageDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference) AzureBlobStorageDataSourceInput() *StorageTransferJobTransferSpecAzureBlobStorageDataSource {
	var returns *StorageTransferJobTransferSpecAzureBlobStorageDataSource
	_jsii_.Get(
		j,
		"azureBlobStorageDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference) GcsDataSink() StorageTransferJobTransferSpecGcsDataSinkOutputReference {
	var returns StorageTransferJobTransferSpecGcsDataSinkOutputReference
	_jsii_.Get(
		j,
		"gcsDataSink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference) GcsDataSinkInput() *StorageTransferJobTransferSpecGcsDataSink {
	var returns *StorageTransferJobTransferSpecGcsDataSink
	_jsii_.Get(
		j,
		"gcsDataSinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference) GcsDataSource() StorageTransferJobTransferSpecGcsDataSourceOutputReference {
	var returns StorageTransferJobTransferSpecGcsDataSourceOutputReference
	_jsii_.Get(
		j,
		"gcsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference) GcsDataSourceInput() *StorageTransferJobTransferSpecGcsDataSource {
	var returns *StorageTransferJobTransferSpecGcsDataSource
	_jsii_.Get(
		j,
		"gcsDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference) HttpDataSource() StorageTransferJobTransferSpecHttpDataSourceOutputReference {
	var returns StorageTransferJobTransferSpecHttpDataSourceOutputReference
	_jsii_.Get(
		j,
		"httpDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference) HttpDataSourceInput() *StorageTransferJobTransferSpecHttpDataSource {
	var returns *StorageTransferJobTransferSpecHttpDataSource
	_jsii_.Get(
		j,
		"httpDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference) InternalValue() *StorageTransferJobTransferSpec {
	var returns *StorageTransferJobTransferSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference) ObjectConditions() StorageTransferJobTransferSpecObjectConditionsOutputReference {
	var returns StorageTransferJobTransferSpecObjectConditionsOutputReference
	_jsii_.Get(
		j,
		"objectConditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference) ObjectConditionsInput() *StorageTransferJobTransferSpecObjectConditions {
	var returns *StorageTransferJobTransferSpecObjectConditions
	_jsii_.Get(
		j,
		"objectConditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference) PosixDataSink() StorageTransferJobTransferSpecPosixDataSinkOutputReference {
	var returns StorageTransferJobTransferSpecPosixDataSinkOutputReference
	_jsii_.Get(
		j,
		"posixDataSink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference) PosixDataSinkInput() *StorageTransferJobTransferSpecPosixDataSink {
	var returns *StorageTransferJobTransferSpecPosixDataSink
	_jsii_.Get(
		j,
		"posixDataSinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference) PosixDataSource() StorageTransferJobTransferSpecPosixDataSourceOutputReference {
	var returns StorageTransferJobTransferSpecPosixDataSourceOutputReference
	_jsii_.Get(
		j,
		"posixDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference) PosixDataSourceInput() *StorageTransferJobTransferSpecPosixDataSource {
	var returns *StorageTransferJobTransferSpecPosixDataSource
	_jsii_.Get(
		j,
		"posixDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference) TransferOptions() StorageTransferJobTransferSpecTransferOptionsOutputReference {
	var returns StorageTransferJobTransferSpecTransferOptionsOutputReference
	_jsii_.Get(
		j,
		"transferOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference) TransferOptionsInput() *StorageTransferJobTransferSpecTransferOptions {
	var returns *StorageTransferJobTransferSpecTransferOptions
	_jsii_.Get(
		j,
		"transferOptionsInput",
		&returns,
	)
	return returns
}


func NewStorageTransferJobTransferSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageTransferJobTransferSpecOutputReference {
	_init_.Initialize()

	if err := validateNewStorageTransferJobTransferSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageTransferJobTransferSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.storageTransferJob.StorageTransferJobTransferSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageTransferJobTransferSpecOutputReference_Override(s StorageTransferJobTransferSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.storageTransferJob.StorageTransferJobTransferSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference)SetInternalValue(val *StorageTransferJobTransferSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) PutAwsS3DataSource(value *StorageTransferJobTransferSpecAwsS3DataSource) {
	if err := s.validatePutAwsS3DataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAwsS3DataSource",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) PutAzureBlobStorageDataSource(value *StorageTransferJobTransferSpecAzureBlobStorageDataSource) {
	if err := s.validatePutAzureBlobStorageDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAzureBlobStorageDataSource",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) PutGcsDataSink(value *StorageTransferJobTransferSpecGcsDataSink) {
	if err := s.validatePutGcsDataSinkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGcsDataSink",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) PutGcsDataSource(value *StorageTransferJobTransferSpecGcsDataSource) {
	if err := s.validatePutGcsDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGcsDataSource",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) PutHttpDataSource(value *StorageTransferJobTransferSpecHttpDataSource) {
	if err := s.validatePutHttpDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putHttpDataSource",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) PutObjectConditions(value *StorageTransferJobTransferSpecObjectConditions) {
	if err := s.validatePutObjectConditionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putObjectConditions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) PutPosixDataSink(value *StorageTransferJobTransferSpecPosixDataSink) {
	if err := s.validatePutPosixDataSinkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPosixDataSink",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) PutPosixDataSource(value *StorageTransferJobTransferSpecPosixDataSource) {
	if err := s.validatePutPosixDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPosixDataSource",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) PutTransferOptions(value *StorageTransferJobTransferSpecTransferOptions) {
	if err := s.validatePutTransferOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTransferOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) ResetAwsS3DataSource() {
	_jsii_.InvokeVoid(
		s,
		"resetAwsS3DataSource",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) ResetAzureBlobStorageDataSource() {
	_jsii_.InvokeVoid(
		s,
		"resetAzureBlobStorageDataSource",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) ResetGcsDataSink() {
	_jsii_.InvokeVoid(
		s,
		"resetGcsDataSink",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) ResetGcsDataSource() {
	_jsii_.InvokeVoid(
		s,
		"resetGcsDataSource",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) ResetHttpDataSource() {
	_jsii_.InvokeVoid(
		s,
		"resetHttpDataSource",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) ResetObjectConditions() {
	_jsii_.InvokeVoid(
		s,
		"resetObjectConditions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) ResetPosixDataSink() {
	_jsii_.InvokeVoid(
		s,
		"resetPosixDataSink",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) ResetPosixDataSource() {
	_jsii_.InvokeVoid(
		s,
		"resetPosixDataSource",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) ResetTransferOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetTransferOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

