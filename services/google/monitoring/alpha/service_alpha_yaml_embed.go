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
// gen_go_data -package alpha -var YAML_service blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/monitoring/alpha/service.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/monitoring/alpha/service.yaml
var YAML_service = []byte("info:\n  title: Monitoring/Service\n  description: The Monitoring Service resource\n  x-dcl-struct-name: Service\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Service\n    parameters:\n    - name: Service\n      required: true\n      description: A full instance of a Service\n  apply:\n    description: The function used to apply information about a Service\n    parameters:\n    - name: Service\n      required: true\n      description: A full instance of a Service\n  delete:\n    description: The function used to delete a Service\n    parameters:\n    - name: Service\n      required: true\n      description: A full instance of a Service\n  deleteAll:\n    description: The function used to delete all Service\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Service\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Service:\n      title: Service\n      x-dcl-id: projects/{{project}}/services/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-labels: userLabels\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - project\n      properties:\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: Name used for UI elements listing this Service.\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'Resource name for this Service. The format is: projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]'\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        telemetry:\n          type: object\n          x-dcl-go-name: Telemetry\n          x-dcl-go-type: ServiceTelemetry\n          description: Configuration for how to query telemetry on a Service.\n          properties:\n            resourceName:\n              type: string\n              x-dcl-go-name: ResourceName\n              description: The full name of the resource that defines this service.\n                Formatted as described in https://cloud.google.com/apis/design/resource_names.\n        userLabels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: UserLabels\n          description: Labels which have been used to annotate the service. Label\n            keys must start with a letter. Label keys and values may contain lowercase\n            letters, numbers, underscores, and dashes. Label keys and values have\n            a maximum length of 63 characters, and must be less than 128 bytes in\n            size. Up to 64 label entries may be stored. For labels which do not have\n            a semantic value, the empty string may be supplied for the label value.\n")

// 3111 bytes
// MD5: 525fac7e93189d3ca8fbdeeea7e15fa1
