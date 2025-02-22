// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	internalupstream "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/internal_upstream/v3"
	rawbuffer "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/raw_buffer/v3"
	metadata "github.com/envoyproxy/go-control-plane/envoy/type/metadata/v3"

	"istio.io/istio/pilot/pkg/util/protoconv"
)

var TunnelHostMetadata = &internalupstream.InternalUpstreamTransport_MetadataValueSource{
	Kind: &metadata.MetadataKind{Kind: &metadata.MetadataKind_Host_{Host: &metadata.MetadataKind_Host{}}},
	Name: "tunnel",
}

func InternalUpstreamTransportSocket(passthroughMetadata ...*internalupstream.InternalUpstreamTransport_MetadataValueSource) *core.TransportSocket {
	return &core.TransportSocket{
		Name: "internal_upstream",
		ConfigType: &core.TransportSocket_TypedConfig{TypedConfig: protoconv.MessageToAny(&internalupstream.InternalUpstreamTransport{
			PassthroughMetadata: passthroughMetadata,
			TransportSocket: &core.TransportSocket{
				Name:       "raw_buffer",
				ConfigType: &core.TransportSocket_TypedConfig{TypedConfig: protoconv.MessageToAny(&rawbuffer.RawBuffer{})},
			},
		})},
	}
}
