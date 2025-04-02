// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GoogleProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (g *jsiiProxy_GoogleProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateGoogleProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateGoogleProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGoogleProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateGoogleProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_GoogleProvider) validateSetAddTerraformAttributionLabelParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GoogleProvider) validateSetBatchingParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GoogleProvider) validateSetExternalCredentialsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GoogleProvider) validateSetUserProjectOverrideParameters(val interface{}) error {
	return nil
}

func validateNewGoogleProviderParameters(scope constructs.Construct, id *string, config *GoogleProviderConfig) error {
	return nil
}

