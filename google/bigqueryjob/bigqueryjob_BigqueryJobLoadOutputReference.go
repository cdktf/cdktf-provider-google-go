package bigqueryjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-google-go/google/v3/jsii"

	"github.com/hashicorp/cdktf-provider-google-go/google/v3/bigqueryjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigqueryJobLoadOutputReference interface {
	cdktf.ComplexObject
	AllowJaggedRows() interface{}
	SetAllowJaggedRows(val interface{})
	AllowJaggedRowsInput() interface{}
	AllowQuotedNewlines() interface{}
	SetAllowQuotedNewlines(val interface{})
	AllowQuotedNewlinesInput() interface{}
	Autodetect() interface{}
	SetAutodetect(val interface{})
	AutodetectInput() interface{}
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
	CreateDisposition() *string
	SetCreateDisposition(val *string)
	CreateDispositionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DestinationEncryptionConfiguration() BigqueryJobLoadDestinationEncryptionConfigurationOutputReference
	DestinationEncryptionConfigurationInput() *BigqueryJobLoadDestinationEncryptionConfiguration
	DestinationTable() BigqueryJobLoadDestinationTableOutputReference
	DestinationTableInput() *BigqueryJobLoadDestinationTable
	Encoding() *string
	SetEncoding(val *string)
	EncodingInput() *string
	FieldDelimiter() *string
	SetFieldDelimiter(val *string)
	FieldDelimiterInput() *string
	// Experimental.
	Fqn() *string
	IgnoreUnknownValues() interface{}
	SetIgnoreUnknownValues(val interface{})
	IgnoreUnknownValuesInput() interface{}
	InternalValue() *BigqueryJobLoad
	SetInternalValue(val *BigqueryJobLoad)
	JsonExtension() *string
	SetJsonExtension(val *string)
	JsonExtensionInput() *string
	MaxBadRecords() *float64
	SetMaxBadRecords(val *float64)
	MaxBadRecordsInput() *float64
	NullMarker() *string
	SetNullMarker(val *string)
	NullMarkerInput() *string
	ProjectionFields() *[]*string
	SetProjectionFields(val *[]*string)
	ProjectionFieldsInput() *[]*string
	Quote() *string
	SetQuote(val *string)
	QuoteInput() *string
	SchemaUpdateOptions() *[]*string
	SetSchemaUpdateOptions(val *[]*string)
	SchemaUpdateOptionsInput() *[]*string
	SkipLeadingRows() *float64
	SetSkipLeadingRows(val *float64)
	SkipLeadingRowsInput() *float64
	SourceFormat() *string
	SetSourceFormat(val *string)
	SourceFormatInput() *string
	SourceUris() *[]*string
	SetSourceUris(val *[]*string)
	SourceUrisInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimePartitioning() BigqueryJobLoadTimePartitioningOutputReference
	TimePartitioningInput() *BigqueryJobLoadTimePartitioning
	WriteDisposition() *string
	SetWriteDisposition(val *string)
	WriteDispositionInput() *string
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
	PutDestinationEncryptionConfiguration(value *BigqueryJobLoadDestinationEncryptionConfiguration)
	PutDestinationTable(value *BigqueryJobLoadDestinationTable)
	PutTimePartitioning(value *BigqueryJobLoadTimePartitioning)
	ResetAllowJaggedRows()
	ResetAllowQuotedNewlines()
	ResetAutodetect()
	ResetCreateDisposition()
	ResetDestinationEncryptionConfiguration()
	ResetEncoding()
	ResetFieldDelimiter()
	ResetIgnoreUnknownValues()
	ResetJsonExtension()
	ResetMaxBadRecords()
	ResetNullMarker()
	ResetProjectionFields()
	ResetQuote()
	ResetSchemaUpdateOptions()
	ResetSkipLeadingRows()
	ResetSourceFormat()
	ResetTimePartitioning()
	ResetWriteDisposition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BigqueryJobLoadOutputReference
type jsiiProxy_BigqueryJobLoadOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) AllowJaggedRows() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowJaggedRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) AllowJaggedRowsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowJaggedRowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) AllowQuotedNewlines() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowQuotedNewlines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) AllowQuotedNewlinesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowQuotedNewlinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) Autodetect() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autodetect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) AutodetectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autodetectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) CreateDisposition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDisposition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) CreateDispositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDispositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) DestinationEncryptionConfiguration() BigqueryJobLoadDestinationEncryptionConfigurationOutputReference {
	var returns BigqueryJobLoadDestinationEncryptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"destinationEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) DestinationEncryptionConfigurationInput() *BigqueryJobLoadDestinationEncryptionConfiguration {
	var returns *BigqueryJobLoadDestinationEncryptionConfiguration
	_jsii_.Get(
		j,
		"destinationEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) DestinationTable() BigqueryJobLoadDestinationTableOutputReference {
	var returns BigqueryJobLoadDestinationTableOutputReference
	_jsii_.Get(
		j,
		"destinationTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) DestinationTableInput() *BigqueryJobLoadDestinationTable {
	var returns *BigqueryJobLoadDestinationTable
	_jsii_.Get(
		j,
		"destinationTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) Encoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) EncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) FieldDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) FieldDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) IgnoreUnknownValues() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreUnknownValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) IgnoreUnknownValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreUnknownValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) InternalValue() *BigqueryJobLoad {
	var returns *BigqueryJobLoad
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) JsonExtension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonExtension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) JsonExtensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonExtensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) MaxBadRecords() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBadRecords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) MaxBadRecordsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBadRecordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) NullMarker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nullMarker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) NullMarkerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nullMarkerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) ProjectionFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projectionFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) ProjectionFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projectionFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) Quote() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) QuoteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) SchemaUpdateOptions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"schemaUpdateOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) SchemaUpdateOptionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"schemaUpdateOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) SkipLeadingRows() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"skipLeadingRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) SkipLeadingRowsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"skipLeadingRowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) SourceFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) SourceFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) SourceUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) SourceUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) TimePartitioning() BigqueryJobLoadTimePartitioningOutputReference {
	var returns BigqueryJobLoadTimePartitioningOutputReference
	_jsii_.Get(
		j,
		"timePartitioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) TimePartitioningInput() *BigqueryJobLoadTimePartitioning {
	var returns *BigqueryJobLoadTimePartitioning
	_jsii_.Get(
		j,
		"timePartitioningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) WriteDisposition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writeDisposition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference) WriteDispositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writeDispositionInput",
		&returns,
	)
	return returns
}


func NewBigqueryJobLoadOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BigqueryJobLoadOutputReference {
	_init_.Initialize()

	if err := validateNewBigqueryJobLoadOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BigqueryJobLoadOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryJob.BigqueryJobLoadOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBigqueryJobLoadOutputReference_Override(b BigqueryJobLoadOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryJob.BigqueryJobLoadOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference)SetAllowJaggedRows(val interface{}) {
	if err := j.validateSetAllowJaggedRowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowJaggedRows",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference)SetAllowQuotedNewlines(val interface{}) {
	if err := j.validateSetAllowQuotedNewlinesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowQuotedNewlines",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference)SetAutodetect(val interface{}) {
	if err := j.validateSetAutodetectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autodetect",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference)SetCreateDisposition(val *string) {
	if err := j.validateSetCreateDispositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createDisposition",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference)SetEncoding(val *string) {
	if err := j.validateSetEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encoding",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference)SetFieldDelimiter(val *string) {
	if err := j.validateSetFieldDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldDelimiter",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference)SetIgnoreUnknownValues(val interface{}) {
	if err := j.validateSetIgnoreUnknownValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreUnknownValues",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference)SetInternalValue(val *BigqueryJobLoad) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference)SetJsonExtension(val *string) {
	if err := j.validateSetJsonExtensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jsonExtension",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference)SetMaxBadRecords(val *float64) {
	if err := j.validateSetMaxBadRecordsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBadRecords",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference)SetNullMarker(val *string) {
	if err := j.validateSetNullMarkerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nullMarker",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference)SetProjectionFields(val *[]*string) {
	if err := j.validateSetProjectionFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectionFields",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference)SetQuote(val *string) {
	if err := j.validateSetQuoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quote",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference)SetSchemaUpdateOptions(val *[]*string) {
	if err := j.validateSetSchemaUpdateOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaUpdateOptions",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference)SetSkipLeadingRows(val *float64) {
	if err := j.validateSetSkipLeadingRowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipLeadingRows",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference)SetSourceFormat(val *string) {
	if err := j.validateSetSourceFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceFormat",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference)SetSourceUris(val *[]*string) {
	if err := j.validateSetSourceUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceUris",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobLoadOutputReference)SetWriteDisposition(val *string) {
	if err := j.validateSetWriteDispositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeDisposition",
		val,
	)
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) PutDestinationEncryptionConfiguration(value *BigqueryJobLoadDestinationEncryptionConfiguration) {
	if err := b.validatePutDestinationEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDestinationEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) PutDestinationTable(value *BigqueryJobLoadDestinationTable) {
	if err := b.validatePutDestinationTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDestinationTable",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) PutTimePartitioning(value *BigqueryJobLoadTimePartitioning) {
	if err := b.validatePutTimePartitioningParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTimePartitioning",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) ResetAllowJaggedRows() {
	_jsii_.InvokeVoid(
		b,
		"resetAllowJaggedRows",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) ResetAllowQuotedNewlines() {
	_jsii_.InvokeVoid(
		b,
		"resetAllowQuotedNewlines",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) ResetAutodetect() {
	_jsii_.InvokeVoid(
		b,
		"resetAutodetect",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) ResetCreateDisposition() {
	_jsii_.InvokeVoid(
		b,
		"resetCreateDisposition",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) ResetDestinationEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetDestinationEncryptionConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) ResetEncoding() {
	_jsii_.InvokeVoid(
		b,
		"resetEncoding",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) ResetFieldDelimiter() {
	_jsii_.InvokeVoid(
		b,
		"resetFieldDelimiter",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) ResetIgnoreUnknownValues() {
	_jsii_.InvokeVoid(
		b,
		"resetIgnoreUnknownValues",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) ResetJsonExtension() {
	_jsii_.InvokeVoid(
		b,
		"resetJsonExtension",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) ResetMaxBadRecords() {
	_jsii_.InvokeVoid(
		b,
		"resetMaxBadRecords",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) ResetNullMarker() {
	_jsii_.InvokeVoid(
		b,
		"resetNullMarker",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) ResetProjectionFields() {
	_jsii_.InvokeVoid(
		b,
		"resetProjectionFields",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) ResetQuote() {
	_jsii_.InvokeVoid(
		b,
		"resetQuote",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) ResetSchemaUpdateOptions() {
	_jsii_.InvokeVoid(
		b,
		"resetSchemaUpdateOptions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) ResetSkipLeadingRows() {
	_jsii_.InvokeVoid(
		b,
		"resetSkipLeadingRows",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) ResetSourceFormat() {
	_jsii_.InvokeVoid(
		b,
		"resetSourceFormat",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) ResetTimePartitioning() {
	_jsii_.InvokeVoid(
		b,
		"resetTimePartitioning",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) ResetWriteDisposition() {
	_jsii_.InvokeVoid(
		b,
		"resetWriteDisposition",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobLoadOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

