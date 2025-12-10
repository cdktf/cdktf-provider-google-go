// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package firebaseapphostingbackend

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FirebaseAppHostingBackendManagedResourcesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FirebaseAppHostingBackendManagedResourcesList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FirebaseAppHostingBackendManagedResourcesList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FirebaseAppHostingBackendManagedResourcesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FirebaseAppHostingBackendManagedResourcesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FirebaseAppHostingBackendManagedResourcesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFirebaseAppHostingBackendManagedResourcesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

