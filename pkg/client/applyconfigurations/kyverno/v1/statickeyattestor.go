/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// StaticKeyAttestorApplyConfiguration represents an declarative configuration of the StaticKeyAttestor type for use
// with apply.
type StaticKeyAttestorApplyConfiguration struct {
	PublicKeys         *string                            `json:"publicKeys,omitempty"`
	SignatureAlgorithm *string                            `json:"signatureAlgorithm,omitempty"`
	KMS                *string                            `json:"kms,omitempty"`
	Secret             *SecretReferenceApplyConfiguration `json:"secret,omitempty"`
	Rekor              *CTLogApplyConfiguration           `json:"rekor,omitempty"`
}

// StaticKeyAttestorApplyConfiguration constructs an declarative configuration of the StaticKeyAttestor type for use with
// apply.
func StaticKeyAttestor() *StaticKeyAttestorApplyConfiguration {
	return &StaticKeyAttestorApplyConfiguration{}
}

// WithPublicKeys sets the PublicKeys field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PublicKeys field is set to the value of the last call.
func (b *StaticKeyAttestorApplyConfiguration) WithPublicKeys(value string) *StaticKeyAttestorApplyConfiguration {
	b.PublicKeys = &value
	return b
}

// WithSignatureAlgorithm sets the SignatureAlgorithm field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SignatureAlgorithm field is set to the value of the last call.
func (b *StaticKeyAttestorApplyConfiguration) WithSignatureAlgorithm(value string) *StaticKeyAttestorApplyConfiguration {
	b.SignatureAlgorithm = &value
	return b
}

// WithKMS sets the KMS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KMS field is set to the value of the last call.
func (b *StaticKeyAttestorApplyConfiguration) WithKMS(value string) *StaticKeyAttestorApplyConfiguration {
	b.KMS = &value
	return b
}

// WithSecret sets the Secret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Secret field is set to the value of the last call.
func (b *StaticKeyAttestorApplyConfiguration) WithSecret(value *SecretReferenceApplyConfiguration) *StaticKeyAttestorApplyConfiguration {
	b.Secret = value
	return b
}

// WithRekor sets the Rekor field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Rekor field is set to the value of the last call.
func (b *StaticKeyAttestorApplyConfiguration) WithRekor(value *CTLogApplyConfiguration) *StaticKeyAttestorApplyConfiguration {
	b.Rekor = value
	return b
}
