/*
 * SPDX-FileCopyrightText: 2019 SAP SE or an SAP affiliate company and Gardener contributors
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CertificateList is the list of Certificate items.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Certificate `json:"items"`
}

// Certificate is the certificate CR.
// +kubebuilder:storageversion
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,path=certificates,shortName=cert,singular=certificate
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name=COMMON NAME,description="Subject domain name of certificate",JSONPath=".status.commonName",type=string
// +kubebuilder:printcolumn:name=ISSUER,description="Issuer name",JSONPath=".status.issuerRef.name",type=string
// +kubebuilder:printcolumn:name=STATUS,JSONPath=".status.state",type=string,description="Status of registration"
// +kubebuilder:printcolumn:name=EXPIRATION_DATE,JSONPath=".status.expirationDate",priority=500,type=string,description="Expiration date (not valid anymore after this date)"
// +kubebuilder:printcolumn:name=DNS_NAMES,JSONPath=".status.dnsNames",priority=2000,type=string,description="Domains names in subject alternative names"
// +kubebuilder:printcolumn:name=AGE,JSONPath=".metadata.creationTimestamp",type=date,description="object creation timestamp"
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateSpec `json:"spec"`
	// +optional
	Status CertificateStatus `json:"status,omitempty"`
}

// CertificateSpec is the spec of the certificate to request.
type CertificateSpec struct {
	// CommonName is the CN for the certificate (max. 64 chars).
	// +kubebuilder:validation:MaxLength=64
	CommonName *string `json:"commonName,omitempty"`
	// DNSNames are the optional additional domain names of the certificate.
	// +optional
	DNSNames []string `json:"dnsNames,omitempty"`
	// CSR is the alternative way to provide CN,DNSNames and other information.
	// +optional
	CSR []byte `json:"csr,omitempty"`
	// IssuerRef is the reference of the issuer to use.
	IssuerRef *IssuerRef `json:"issuerRef,omitempty"`
	// SecretName is the name of the secret object to use for storing the certificate.
	SecretName *string `json:"secretName,omitempty"`
	// SecretRef is the reference of the secret object to use for storing the certificate.
	SecretRef *corev1.SecretReference `json:"secretRef,omitempty"`
	// Renew triggers a renewal if set to true
	// +optional
	Renew *bool `json:"renew,omitempty"`
	// EnsureRenewedAfter specifies a time stamp in the past. Renewing is only triggered if certificate notBefore date is before this date.
	// +optional
	EnsureRenewedAfter *metav1.Time `json:"ensureRenewedAfter,omitempty"`
}

// IssuerRef is the reference of the issuer by name.
type IssuerRef struct {
	// Name is the name of the issuer CR (in the configured issuer namespace).
	Name string `json:"name"`
}

// BackOffState stores the status for exponential back off on repeated cert request failure
type BackOffState struct {
	// ObservedGeneration is the observed generation the BackOffState is assigned to
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// RetryAfter is the timestamp this cert request is not retried before.
	RetryAfter metav1.Time `json:"recheckAfter"`
	// RetryInterval is interval to wait for retrying.
	RetryInterval metav1.Duration `json:"recheckInterval"`
}

// CertificateStatus is the status of the certificate request.
type CertificateStatus struct {
	// ObservedGeneration is the observed generation of the spec.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// State is the certificate state.
	State string `json:"state"`
	// Message is the status or error message.
	// +optional
	Message *string `json:"message,omitempty"`
	// LastPendingTimestamp contains the start timestamp of the last pending status.
	// +optional
	LastPendingTimestamp *metav1.Time `json:"lastPendingTimestamp,omitempty"`
	// CommonName is the current CN.
	// +optional
	CommonName *string `json:"commonName,omitempty"`
	// DNSNames are the current domain names.
	// +optional
	DNSNames []string `json:"dnsNames,omitempty"`
	// IssuerRef is the used issuer.
	// +optional
	IssuerRef *IssuerRefWithNamespace `json:"issuerRef,omitempty"`
	// ExpirationDate shows the notAfter validity date.
	// +optional
	ExpirationDate *string `json:"expirationDate,omitempty"`
	// BackOff contains the state to back off failed certificate requests
	// +optional
	BackOff *BackOffState `json:"backoff,omitempty"`
}

// IssuerRefWithNamespace is the full qualified issuer reference.
type IssuerRefWithNamespace struct {
	// Name is the name of the issuer CR.
	Name string `json:"name"`
	// Namespace is the namespace of the issuer CR.
	Namespace string `json:"namespace"`
}
