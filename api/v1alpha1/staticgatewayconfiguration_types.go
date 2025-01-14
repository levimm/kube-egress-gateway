// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GatewayVmssProfile finds an existing gateway VMSS (virtual machine scale set).
type GatewayVmssProfile struct {
	// Resource group of the VMSS. Must be in the same subscription.
	VmssResourceGroup string `json:"vmssResourceGroup,omitempty"`

	// Name of the VMSS
	VmssName string `json:"vmssName,omitempty"`

	// Public IP prefix size to be applied to this VMSS.
	//+kubebuilder:validation:Minimum=0
	//+kubebuilder:validation:Maximum=31
	PublicIpPrefixSize int32 `json:"publicIpPrefixSize,omitempty"`
}

// RouteType defines the type of defaultRoute.
// +kubebuilder:validation:Enum=azureNetworking;staticEgressGateway
type RouteType string

const (
	// RouteStaticEgressGateway defines static egress gateway as the default route.
	RouteStaticEgressGateway RouteType = "staticEgressGateway"

	// RouteAzureNetworking defines azure networking as the default route.
	RouteAzureNetworking RouteType = "azureNetworking"
)

// StaticGatewayConfigurationSpec defines the desired state of StaticGatewayConfiguration
type StaticGatewayConfigurationSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Name of the gateway nodepool to apply the gateway configuration.
	// +optional
	GatewayNodepoolName string `json:"gatewayNodepoolName,omitempty"`

	// Profile of the gateway VMSS to apply the gateway configuration.
	// +optional
	GatewayVmssProfile `json:"gatewayVmssProfile,omitempty"`

	// Pod default route, should be either azureNetworking (pod's eth0) or staticEgressGateway (default).
	//+kubebuilder:default=staticEgressGateway
	DefaultRoute RouteType `json:"defaultRoute,omitempty"`

	// Whether to provision public IP prefixes for outbound.
	//+kubebuilder:default=true
	ProvisionPublicIps bool `json:"provisionPublicIps"`

	// BYO Resource ID of public IP prefix to be used as outbound. This can only be specified when provisionPublicIps is true.
	// +optional
	PublicIpPrefixId string `json:"publicIpPrefixId,omitempty"`

	// CIDRs to be excluded from the default route.
	ExcludeCidrs []string `json:"excludeCidrs,omitempty"`
}

// GatewayWireguardProfile provides details about gateway side wireguard configuration.
type GatewayWireguardProfile struct {
	// Gateway IP for wireguard connection.
	WireguardServerIp string `json:"wireguardServerIp,omitempty"`

	// Listening port of the gateway side wireguard daemon.
	WireguardServerPort int32 `json:"wireguardServerPort,omitempty"`

	// Gateway side wireguard public key.
	WireguardPublicKey string `json:"wireguardPublicKey,omitempty"`

	// Reference of the secret that holds gateway side wireguard private key.
	WireguardPrivateKeySecretRef *corev1.ObjectReference `json:"wireguardPrivateKeySecretRef,omitempty"`
}

// StaticGatewayConfigurationStatus defines the observed state of StaticGatewayConfiguration
type StaticGatewayConfigurationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Egress IP Prefix CIDR used for this gateway configuration.
	EgressIpPrefix string `json:"egressIpPrefix,omitempty"`

	// Gateway side wireguard profile.
	GatewayWireguardProfile `json:"gatewayWireguardProfile,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// StaticGatewayConfiguration is the Schema for the staticgatewayconfigurations API
type StaticGatewayConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StaticGatewayConfigurationSpec   `json:"spec,omitempty"`
	Status StaticGatewayConfigurationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// StaticGatewayConfigurationList contains a list of StaticGatewayConfiguration
type StaticGatewayConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StaticGatewayConfiguration `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StaticGatewayConfiguration{}, &StaticGatewayConfigurationList{})
}
