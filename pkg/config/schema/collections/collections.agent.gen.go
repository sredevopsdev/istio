//go:build agent
// +build agent

// GENERATED FILE -- DO NOT EDIT
//

package collections

import (
	"reflect"

	istioioapiextensionsv1alpha1 "istio.io/api/extensions/v1alpha1"
	istioioapimeshv1alpha1 "istio.io/api/mesh/v1alpha1"
	istioioapimetav1alpha1 "istio.io/api/meta/v1alpha1"
	istioioapinetworkingv1alpha3 "istio.io/api/networking/v1alpha3"
	istioioapinetworkingv1beta1 "istio.io/api/networking/v1beta1"
	istioioapisecurityv1beta1 "istio.io/api/security/v1beta1"
	istioioapitelemetryv1alpha1 "istio.io/api/telemetry/v1alpha1"
	"istio.io/istio/pkg/config/schema/collection"
	"istio.io/istio/pkg/config/schema/resource"
	"istio.io/istio/pkg/config/validation"
)

var (
	AuthorizationPolicy = resource.Builder{
		Identifier: "AuthorizationPolicy",
		Group:      "security.istio.io",
		Kind:       "AuthorizationPolicy",
		Plural:     "authorizationpolicies",
		Version:    "v1beta1",
		VersionAliases: []string{
			"v1",
		},
		Proto: "istio.security.v1beta1.AuthorizationPolicy", StatusProto: "istio.meta.v1alpha1.IstioStatus",
		ReflectType: reflect.TypeOf(&istioioapisecurityv1beta1.AuthorizationPolicy{}).Elem(), StatusType: reflect.TypeOf(&istioioapimetav1alpha1.IstioStatus{}).Elem(),
		ProtoPackage: "istio.io/api/security/v1beta1", StatusPackage: "istio.io/api/meta/v1alpha1",
		ClusterScoped: false,
		Builtin:       false,
		ValidateProto: validation.ValidateAuthorizationPolicy,
	}.MustBuild()

	DestinationRule = resource.Builder{
		Identifier: "DestinationRule",
		Group:      "networking.istio.io",
		Kind:       "DestinationRule",
		Plural:     "destinationrules",
		Version:    "v1alpha3",
		VersionAliases: []string{
			"v1beta1",
		},
		Proto: "istio.networking.v1alpha3.DestinationRule", StatusProto: "istio.meta.v1alpha1.IstioStatus",
		ReflectType: reflect.TypeOf(&istioioapinetworkingv1alpha3.DestinationRule{}).Elem(), StatusType: reflect.TypeOf(&istioioapimetav1alpha1.IstioStatus{}).Elem(),
		ProtoPackage: "istio.io/api/networking/v1alpha3", StatusPackage: "istio.io/api/meta/v1alpha1",
		ClusterScoped: false,
		Builtin:       false,
		ValidateProto: validation.ValidateDestinationRule,
	}.MustBuild()

	EnvoyFilter = resource.Builder{
		Identifier: "EnvoyFilter",
		Group:      "networking.istio.io",
		Kind:       "EnvoyFilter",
		Plural:     "envoyfilters",
		Version:    "v1alpha3",
		Proto:      "istio.networking.v1alpha3.EnvoyFilter", StatusProto: "istio.meta.v1alpha1.IstioStatus",
		ReflectType: reflect.TypeOf(&istioioapinetworkingv1alpha3.EnvoyFilter{}).Elem(), StatusType: reflect.TypeOf(&istioioapimetav1alpha1.IstioStatus{}).Elem(),
		ProtoPackage: "istio.io/api/networking/v1alpha3", StatusPackage: "istio.io/api/meta/v1alpha1",
		ClusterScoped: false,
		Builtin:       false,
		ValidateProto: validation.ValidateEnvoyFilter,
	}.MustBuild()

	Gateway = resource.Builder{
		Identifier: "Gateway",
		Group:      "networking.istio.io",
		Kind:       "Gateway",
		Plural:     "gateways",
		Version:    "v1alpha3",
		VersionAliases: []string{
			"v1beta1",
		},
		Proto: "istio.networking.v1alpha3.Gateway", StatusProto: "istio.meta.v1alpha1.IstioStatus",
		ReflectType: reflect.TypeOf(&istioioapinetworkingv1alpha3.Gateway{}).Elem(), StatusType: reflect.TypeOf(&istioioapimetav1alpha1.IstioStatus{}).Elem(),
		ProtoPackage: "istio.io/api/networking/v1alpha3", StatusPackage: "istio.io/api/meta/v1alpha1",
		ClusterScoped: false,
		Builtin:       false,
		ValidateProto: validation.ValidateGateway,
	}.MustBuild()

	MeshConfig = resource.Builder{
		Identifier:    "MeshConfig",
		Group:         "",
		Kind:          "MeshConfig",
		Plural:        "meshconfigs",
		Version:       "v1alpha1",
		Proto:         "istio.mesh.v1alpha1.MeshConfig",
		ReflectType:   reflect.TypeOf(&istioioapimeshv1alpha1.MeshConfig{}).Elem(),
		ProtoPackage:  "istio.io/api/mesh/v1alpha1",
		ClusterScoped: false,
		Builtin:       false,
		ValidateProto: validation.EmptyValidate,
	}.MustBuild()

	MeshNetworks = resource.Builder{
		Identifier:    "MeshNetworks",
		Group:         "",
		Kind:          "MeshNetworks",
		Plural:        "meshnetworks",
		Version:       "v1alpha1",
		Proto:         "istio.mesh.v1alpha1.MeshNetworks",
		ReflectType:   reflect.TypeOf(&istioioapimeshv1alpha1.MeshNetworks{}).Elem(),
		ProtoPackage:  "istio.io/api/mesh/v1alpha1",
		ClusterScoped: false,
		Builtin:       false,
		ValidateProto: validation.EmptyValidate,
	}.MustBuild()

	PeerAuthentication = resource.Builder{
		Identifier: "PeerAuthentication",
		Group:      "security.istio.io",
		Kind:       "PeerAuthentication",
		Plural:     "peerauthentications",
		Version:    "v1beta1",
		Proto:      "istio.security.v1beta1.PeerAuthentication", StatusProto: "istio.meta.v1alpha1.IstioStatus",
		ReflectType: reflect.TypeOf(&istioioapisecurityv1beta1.PeerAuthentication{}).Elem(), StatusType: reflect.TypeOf(&istioioapimetav1alpha1.IstioStatus{}).Elem(),
		ProtoPackage: "istio.io/api/security/v1beta1", StatusPackage: "istio.io/api/meta/v1alpha1",
		ClusterScoped: false,
		Builtin:       false,
		ValidateProto: validation.ValidatePeerAuthentication,
	}.MustBuild()

	ProxyConfig = resource.Builder{
		Identifier: "ProxyConfig",
		Group:      "networking.istio.io",
		Kind:       "ProxyConfig",
		Plural:     "proxyconfigs",
		Version:    "v1beta1",
		Proto:      "istio.networking.v1beta1.ProxyConfig", StatusProto: "istio.meta.v1alpha1.IstioStatus",
		ReflectType: reflect.TypeOf(&istioioapinetworkingv1beta1.ProxyConfig{}).Elem(), StatusType: reflect.TypeOf(&istioioapimetav1alpha1.IstioStatus{}).Elem(),
		ProtoPackage: "istio.io/api/networking/v1beta1", StatusPackage: "istio.io/api/meta/v1alpha1",
		ClusterScoped: false,
		Builtin:       false,
		ValidateProto: validation.ValidateProxyConfig,
	}.MustBuild()

	RequestAuthentication = resource.Builder{
		Identifier: "RequestAuthentication",
		Group:      "security.istio.io",
		Kind:       "RequestAuthentication",
		Plural:     "requestauthentications",
		Version:    "v1beta1",
		VersionAliases: []string{
			"v1",
		},
		Proto: "istio.security.v1beta1.RequestAuthentication", StatusProto: "istio.meta.v1alpha1.IstioStatus",
		ReflectType: reflect.TypeOf(&istioioapisecurityv1beta1.RequestAuthentication{}).Elem(), StatusType: reflect.TypeOf(&istioioapimetav1alpha1.IstioStatus{}).Elem(),
		ProtoPackage: "istio.io/api/security/v1beta1", StatusPackage: "istio.io/api/meta/v1alpha1",
		ClusterScoped: false,
		Builtin:       false,
		ValidateProto: validation.ValidateRequestAuthentication,
	}.MustBuild()

	ServiceEntry = resource.Builder{
		Identifier: "ServiceEntry",
		Group:      "networking.istio.io",
		Kind:       "ServiceEntry",
		Plural:     "serviceentries",
		Version:    "v1alpha3",
		VersionAliases: []string{
			"v1beta1",
		},
		Proto: "istio.networking.v1alpha3.ServiceEntry", StatusProto: "istio.meta.v1alpha1.IstioStatus",
		ReflectType: reflect.TypeOf(&istioioapinetworkingv1alpha3.ServiceEntry{}).Elem(), StatusType: reflect.TypeOf(&istioioapimetav1alpha1.IstioStatus{}).Elem(),
		ProtoPackage: "istio.io/api/networking/v1alpha3", StatusPackage: "istio.io/api/meta/v1alpha1",
		ClusterScoped: false,
		Builtin:       false,
		ValidateProto: validation.ValidateServiceEntry,
	}.MustBuild()

	Sidecar = resource.Builder{
		Identifier: "Sidecar",
		Group:      "networking.istio.io",
		Kind:       "Sidecar",
		Plural:     "sidecars",
		Version:    "v1alpha3",
		VersionAliases: []string{
			"v1beta1",
		},
		Proto: "istio.networking.v1alpha3.Sidecar", StatusProto: "istio.meta.v1alpha1.IstioStatus",
		ReflectType: reflect.TypeOf(&istioioapinetworkingv1alpha3.Sidecar{}).Elem(), StatusType: reflect.TypeOf(&istioioapimetav1alpha1.IstioStatus{}).Elem(),
		ProtoPackage: "istio.io/api/networking/v1alpha3", StatusPackage: "istio.io/api/meta/v1alpha1",
		ClusterScoped: false,
		Builtin:       false,
		ValidateProto: validation.ValidateSidecar,
	}.MustBuild()

	Telemetry = resource.Builder{
		Identifier: "Telemetry",
		Group:      "telemetry.istio.io",
		Kind:       "Telemetry",
		Plural:     "telemetries",
		Version:    "v1alpha1",
		Proto:      "istio.telemetry.v1alpha1.Telemetry", StatusProto: "istio.meta.v1alpha1.IstioStatus",
		ReflectType: reflect.TypeOf(&istioioapitelemetryv1alpha1.Telemetry{}).Elem(), StatusType: reflect.TypeOf(&istioioapimetav1alpha1.IstioStatus{}).Elem(),
		ProtoPackage: "istio.io/api/telemetry/v1alpha1", StatusPackage: "istio.io/api/meta/v1alpha1",
		ClusterScoped: false,
		Builtin:       false,
		ValidateProto: validation.ValidateTelemetry,
	}.MustBuild()

	VirtualService = resource.Builder{
		Identifier: "VirtualService",
		Group:      "networking.istio.io",
		Kind:       "VirtualService",
		Plural:     "virtualservices",
		Version:    "v1alpha3",
		VersionAliases: []string{
			"v1beta1",
		},
		Proto: "istio.networking.v1alpha3.VirtualService", StatusProto: "istio.meta.v1alpha1.IstioStatus",
		ReflectType: reflect.TypeOf(&istioioapinetworkingv1alpha3.VirtualService{}).Elem(), StatusType: reflect.TypeOf(&istioioapimetav1alpha1.IstioStatus{}).Elem(),
		ProtoPackage: "istio.io/api/networking/v1alpha3", StatusPackage: "istio.io/api/meta/v1alpha1",
		ClusterScoped: false,
		Builtin:       false,
		ValidateProto: validation.ValidateVirtualService,
	}.MustBuild()

	WasmPlugin = resource.Builder{
		Identifier: "WasmPlugin",
		Group:      "extensions.istio.io",
		Kind:       "WasmPlugin",
		Plural:     "wasmplugins",
		Version:    "v1alpha1",
		Proto:      "istio.extensions.v1alpha1.WasmPlugin", StatusProto: "istio.meta.v1alpha1.IstioStatus",
		ReflectType: reflect.TypeOf(&istioioapiextensionsv1alpha1.WasmPlugin{}).Elem(), StatusType: reflect.TypeOf(&istioioapimetav1alpha1.IstioStatus{}).Elem(),
		ProtoPackage: "istio.io/api/extensions/v1alpha1", StatusPackage: "istio.io/api/meta/v1alpha1",
		ClusterScoped: false,
		Builtin:       false,
		ValidateProto: validation.ValidateWasmPlugin,
	}.MustBuild()

	WorkloadEntry = resource.Builder{
		Identifier: "WorkloadEntry",
		Group:      "networking.istio.io",
		Kind:       "WorkloadEntry",
		Plural:     "workloadentries",
		Version:    "v1alpha3",
		VersionAliases: []string{
			"v1beta1",
		},
		Proto: "istio.networking.v1alpha3.WorkloadEntry", StatusProto: "istio.meta.v1alpha1.IstioStatus",
		ReflectType: reflect.TypeOf(&istioioapinetworkingv1alpha3.WorkloadEntry{}).Elem(), StatusType: reflect.TypeOf(&istioioapimetav1alpha1.IstioStatus{}).Elem(),
		ProtoPackage: "istio.io/api/networking/v1alpha3", StatusPackage: "istio.io/api/meta/v1alpha1",
		ClusterScoped: false,
		Builtin:       false,
		ValidateProto: validation.ValidateWorkloadEntry,
	}.MustBuild()

	WorkloadGroup = resource.Builder{
		Identifier: "WorkloadGroup",
		Group:      "networking.istio.io",
		Kind:       "WorkloadGroup",
		Plural:     "workloadgroups",
		Version:    "v1alpha3",
		VersionAliases: []string{
			"v1beta1",
		},
		Proto: "istio.networking.v1alpha3.WorkloadGroup", StatusProto: "istio.meta.v1alpha1.IstioStatus",
		ReflectType: reflect.TypeOf(&istioioapinetworkingv1alpha3.WorkloadGroup{}).Elem(), StatusType: reflect.TypeOf(&istioioapimetav1alpha1.IstioStatus{}).Elem(),
		ProtoPackage: "istio.io/api/networking/v1alpha3", StatusPackage: "istio.io/api/meta/v1alpha1",
		ClusterScoped: false,
		Builtin:       false,
		ValidateProto: validation.ValidateWorkloadGroup,
	}.MustBuild()

	// All contains all collections in the system.
	All = collection.NewSchemasBuilder().
		MustAdd(AuthorizationPolicy).
		MustAdd(DestinationRule).
		MustAdd(EnvoyFilter).
		MustAdd(Gateway).
		MustAdd(MeshConfig).
		MustAdd(MeshNetworks).
		MustAdd(PeerAuthentication).
		MustAdd(ProxyConfig).
		MustAdd(RequestAuthentication).
		MustAdd(ServiceEntry).
		MustAdd(Sidecar).
		MustAdd(Telemetry).
		MustAdd(VirtualService).
		MustAdd(WasmPlugin).
		MustAdd(WorkloadEntry).
		MustAdd(WorkloadGroup).
		Build()

	// Kube contains only kubernetes collections.
	Kube = collection.NewSchemasBuilder().
		Build()

	// Pilot contains only collections used by Pilot.
	Pilot = collection.NewSchemasBuilder().
		MustAdd(AuthorizationPolicy).
		MustAdd(DestinationRule).
		MustAdd(EnvoyFilter).
		MustAdd(Gateway).
		MustAdd(PeerAuthentication).
		MustAdd(ProxyConfig).
		MustAdd(RequestAuthentication).
		MustAdd(ServiceEntry).
		MustAdd(Sidecar).
		MustAdd(Telemetry).
		MustAdd(VirtualService).
		MustAdd(WasmPlugin).
		MustAdd(WorkloadEntry).
		MustAdd(WorkloadGroup).
		Build()

	// PilotGatewayAPI contains only collections used by Pilot, including experimental Service Api.
	PilotGatewayAPI = collection.NewSchemasBuilder().
			MustAdd(AuthorizationPolicy).
			MustAdd(DestinationRule).
			MustAdd(EnvoyFilter).
			MustAdd(Gateway).
			MustAdd(PeerAuthentication).
			MustAdd(ProxyConfig).
			MustAdd(RequestAuthentication).
			MustAdd(ServiceEntry).
			MustAdd(Sidecar).
			MustAdd(Telemetry).
			MustAdd(VirtualService).
			MustAdd(WasmPlugin).
			MustAdd(WorkloadEntry).
			MustAdd(WorkloadGroup).
			Build()
)
