// Copyright 2021 Google LLC. All Rights Reserved.
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
// GENERATED BY gen_go_data.go
// gen_go_data -package beta -var YAML_packet_mirroring blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/beta/packet_mirroring.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/beta/packet_mirroring.yaml
var YAML_packet_mirroring = []byte("info:\n  title: Compute/PacketMirroring\n  description: The Compute PacketMirroring resource\n  x-dcl-struct-name: PacketMirroring\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a PacketMirroring\n    parameters:\n    - name: PacketMirroring\n      required: true\n      description: A full instance of a PacketMirroring\n  apply:\n    description: The function used to apply information about a PacketMirroring\n    parameters:\n    - name: PacketMirroring\n      required: true\n      description: A full instance of a PacketMirroring\n  delete:\n    description: The function used to delete a PacketMirroring\n    parameters:\n    - name: PacketMirroring\n      required: true\n      description: A full instance of a PacketMirroring\n  deleteAll:\n    description: The function used to delete all PacketMirroring\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many PacketMirroring\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    PacketMirroring:\n      title: PacketMirroring\n      x-dcl-id: projects/{{project}}/regions/{{location}}/packetMirrorings/{{name}}\n      x-dcl-locations:\n      - region\n      x-dcl-parent-container: project\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - network\n      - collectorIlb\n      - mirroredResources\n      - project\n      - location\n      properties:\n        collectorIlb:\n          type: object\n          x-dcl-go-name: CollectorIlb\n          x-dcl-go-type: PacketMirroringCollectorIlb\n          description: The Forwarding Rule resource of type `loadBalancingScheme=INTERNAL`\n            that will be used as collector for mirrored traffic. The specified forwarding\n            rule must have `isMirroringCollector` set to true.\n          required:\n          - url\n          properties:\n            canonicalUrl:\n              type: string\n              x-dcl-go-name: CanonicalUrl\n              readOnly: true\n              description: Unique identifier for the forwarding rule; defined by the\n                server.\n            url:\n              type: string\n              x-dcl-go-name: Url\n              description: Resource URL to the forwarding rule representing the ILB\n                configured as destination of the mirrored traffic.\n              x-dcl-references:\n              - resource: Compute/ForwardingRule\n                field: selfLink\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: An optional description of this resource. Provide this property\n            when you create the resource.\n        enable:\n          type: string\n          x-dcl-go-name: Enable\n          x-dcl-go-type: PacketMirroringEnableEnum\n          description: Indicates whether or not this packet mirroring takes effect.\n            If set to FALSE, this packet mirroring policy will not be enforced on\n            the network. The default is TRUE.\n          enum:\n          - \"TRUE\"\n          - \"FALSE\"\n        filter:\n          type: object\n          x-dcl-go-name: Filter\n          x-dcl-go-type: PacketMirroringFilter\n          description: Filter for mirrored traffic. If unspecified, all traffic is\n            mirrored.\n          properties:\n            cidrRanges:\n              type: array\n              x-dcl-go-name: CidrRanges\n              description: IP CIDR ranges that apply as filter on the source (ingress)\n                or destination (egress) IP in the IP header. Only IPv4 is supported.\n                If no ranges are specified, all traffic that matches the specified\n                IPProtocols is mirrored. If neither cidrRanges nor IPProtocols is\n                specified, all traffic is mirrored.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            direction:\n              type: string\n              x-dcl-go-name: Direction\n              x-dcl-go-type: PacketMirroringFilterDirectionEnum\n              description: Direction of traffic to mirror, either INGRESS, EGRESS,\n                or BOTH. The default is BOTH.\n              enum:\n              - INGRESS\n              - EGRESS\n            ipProtocols:\n              type: array\n              x-dcl-go-name: IPProtocols\n              description: Protocols that apply as filter on mirrored traffic. If\n                no protocols are specified, all traffic that matches the specified\n                CIDR ranges is mirrored. If neither cidrRanges nor IPProtocols is\n                specified, all traffic is mirrored.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n        id:\n          type: integer\n          format: int64\n          x-dcl-go-name: Id\n          readOnly: true\n          description: The unique identifier for the resource. This identifier is\n            defined by the server.\n          x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        mirroredResources:\n          type: object\n          x-dcl-go-name: MirroredResources\n          x-dcl-go-type: PacketMirroringMirroredResources\n          description: PacketMirroring mirroredResourceInfos. MirroredResourceInfo\n            specifies a set of mirrored VM instances, subnetworks and/or tags for\n            which traffic from/to all VM instances will be mirrored.\n          properties:\n            instances:\n              type: array\n              x-dcl-go-name: Instances\n              description: A set of virtual machine instances that are being mirrored.\n                They must live in zones contained in the same region as this packetMirroring.\n                Note that this config will apply only to those network interfaces\n                of the Instances that belong to the network specified in this packetMirroring.\n                You may specify a maximum of 50 Instances.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: PacketMirroringMirroredResourcesInstances\n                properties:\n                  canonicalUrl:\n                    type: string\n                    x-dcl-go-name: CanonicalUrl\n                    readOnly: true\n                    description: Unique identifier for the instance; defined by the\n                      server.\n                    x-kubernetes-immutable: true\n                  url:\n                    type: string\n                    x-dcl-go-name: Url\n                    description: Resource URL to the virtual machine instance which\n                      is being mirrored.\n                    x-dcl-references:\n                    - resource: Compute/Instance\n                      field: selfLink\n            subnetworks:\n              type: array\n              x-dcl-go-name: Subnetworks\n              description: A set of subnetworks for which traffic from/to all VM instances\n                will be mirrored. They must live in the same region as this packetMirroring.\n                You may specify a maximum of 5 subnetworks.\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: PacketMirroringMirroredResourcesSubnetworks\n                properties:\n                  canonicalUrl:\n                    type: string\n                    x-dcl-go-name: CanonicalUrl\n                    readOnly: true\n                    description: Unique identifier for the subnetwork; defined by\n                      the server.\n                    x-kubernetes-immutable: true\n                  url:\n                    type: string\n                    x-dcl-go-name: Url\n                    description: Resource URL to the subnetwork for which traffic\n                      from/to all VM instances will be mirrored.\n                    x-kubernetes-immutable: true\n                    x-dcl-references:\n                    - resource: Compute/Subnetwork\n                      field: selfLink\n            tags:\n              type: array\n              x-dcl-go-name: Tags\n              description: A set of mirrored tags. Traffic from/to all VM instances\n                that have one or more of these tags will be mirrored.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Name of the resource; provided by the client when the resource\n            is created. The name must be 1-63 characters long, and comply with [RFC1035](https://www.ietf.org/rfc/rfc1035.txt).\n            Specifically, the name must be 1-63 characters long and match the regular\n            expression `)?` which means the first character must be a lowercase letter,\n            and all following characters must be a dash, lowercase letter, or digit,\n            except the last character, which cannot be a dash.\n          x-kubernetes-immutable: true\n        network:\n          type: object\n          x-dcl-go-name: Network\n          x-dcl-go-type: PacketMirroringNetwork\n          description: Specifies the mirrored VPC network. Only packets in this network\n            will be mirrored. All mirrored VMs should have a NIC in the given network.\n            All mirrored subnetworks should belong to the given network.\n          x-kubernetes-immutable: true\n          required:\n          - url\n          properties:\n            canonicalUrl:\n              type: string\n              x-dcl-go-name: CanonicalUrl\n              readOnly: true\n              description: Unique identifier for the network; defined by the server.\n              x-kubernetes-immutable: true\n            url:\n              type: string\n              x-dcl-go-name: Url\n              description: URL of the network resource.\n              x-kubernetes-immutable: true\n              x-dcl-references:\n              - resource: Compute/Network\n                field: selfLink\n        priority:\n          type: integer\n          format: int64\n          x-dcl-go-name: Priority\n          description: The priority of applying this configuration. Priority is used\n            to break ties in cases where there is more than one matching rule. In\n            the case of two rules that apply for a given Instance, the one with the\n            lowest-numbered priority value wins. Default value is 1000. Valid range\n            is 0 through 65535.\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        region:\n          type: string\n          x-dcl-go-name: Region\n          readOnly: true\n          description: URI of the region where the packetMirroring resides.\n          x-kubernetes-immutable: true\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          readOnly: true\n          description: Server-defined URL for the resource.\n          x-kubernetes-immutable: true\n")

// 11861 bytes
// MD5: 23287c212714d3f3bb9ac19748b53242
