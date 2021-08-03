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
// gen_go_data -package serviceusage -var YAML_service blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/serviceusage/service.yaml

package serviceusage

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/serviceusage/service.yaml
var YAML_service = []byte("info:\n  title: ServiceUsage/Service\n  description: DCL Specification for the ServiceUsage Service resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Service\n    parameters:\n    - name: Service\n      required: true\n      description: A full instance of a Service\n  apply:\n    description: The function used to apply information about a Service\n    parameters:\n    - name: Service\n      required: true\n      description: A full instance of a Service\n  delete:\n    description: The function used to delete a Service\n    parameters:\n    - name: Service\n      required: true\n      description: A full instance of a Service\n  deleteAll:\n    description: The function used to delete all Service\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Service\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Service:\n      title: Service\n      x-dcl-id: projects/{{project}}/services/{{name}}\n      x-dcl-parent-container: project\n      type: object\n      properties:\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The resource name of the service. A valid name would be `serviceusage.googleapis.com`\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The resource name of the consumer project.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: ServiceStateEnum\n          description: 'The state of the resource. Possible values: DISABLED, ENABLED'\n          enum:\n          - DISABLED\n          - ENABLED\n")

// 1974 bytes
// MD5: ed5cd802643680656b9bfe0cb13edfde