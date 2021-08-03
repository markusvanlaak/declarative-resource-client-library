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
// gen_go_data -package servicenetworking -var YAML_connection blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/servicenetworking/connection.yaml

package servicenetworking

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/servicenetworking/connection.yaml
var YAML_connection = []byte("info:\n  title: ServiceNetworking/Connection\n  description: DCL Specification for the ServiceNetworking Connection resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Connection\n    parameters:\n    - name: Connection\n      required: true\n      description: A full instance of a Connection\n  apply:\n    description: The function used to apply information about a Connection\n    parameters:\n    - name: Connection\n      required: true\n      description: A full instance of a Connection\n  delete:\n    description: The function used to delete a Connection\n    parameters:\n    - name: Connection\n      required: true\n      description: A full instance of a Connection\n  deleteAll:\n    description: The function used to delete all Connection\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: network\n      required: true\n      schema:\n        type: string\n    - name: service\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Connection\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: network\n      required: true\n      schema:\n        type: string\n    - name: service\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Connection:\n      title: Connection\n      x-dcl-id: names/{{name}}\n      x-dcl-locations:\n      - global\n      x-dcl-parent-container: project\n      type: object\n      required:\n      - network\n      - project\n      - reservedPeeringRanges\n      properties:\n        name:\n          type: string\n          x-dcl-go-name: Name\n          readOnly: true\n          description: Output only. The name of the VPC Network Peering connection\n            that was created by the service producer.\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n        network:\n          type: string\n          x-dcl-go-name: Network\n          description: 'The name of service consumer''s VPC network that''s connected\n            with service producer network, in the following format: `projects/{project}/global/networks/{network}`.\n            `{project}` is a project number, such as in `12345` that includes the\n            VPC service consumer''s VPC network. `{network}` is the name of the service\n            consumer''s VPC network.'\n          x-dcl-references:\n          - resource: Compute/Network\n            field: selfLink\n            parent: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project in use for this network.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        reservedPeeringRanges:\n          type: array\n          x-dcl-go-name: ReservedPeeringRanges\n          description: The name of one or more allocated IP address ranges for this\n            service producer of type `PEERING`. Note that invoking CreateConnection\n            method with a different range when connection is already established will\n            not modify already provisioned service producer subnetworks. If CreateConnection\n            method is invoked repeatedly to reconnect when peering connection had\n            been disconnected on the consumer side, leaving this field empty will\n            restore previously allocated IP ranges.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n            x-dcl-references:\n            - resource: Compute/Address\n              field: name\n        service:\n          type: string\n          x-dcl-go-name: Service\n          description: The name of the peering service that's associated with this\n            connection.\n          default: services/servicenetworking.googleapis.com\n")

// 4009 bytes
// MD5: 9bcf501d5095af82c4ad0ce24c9505b0