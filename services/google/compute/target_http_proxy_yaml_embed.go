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
// gen_go_data -package compute -var YAML_target_http_proxy blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/target_http_proxy.yaml

package compute

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/target_http_proxy.yaml
var YAML_target_http_proxy = []byte("info:\n  title: Compute/TargetHttpProxy\n  description: DCL Specification for the Compute TargetHttpProxy resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a TargetHttpProxy\n    parameters:\n    - name: TargetHttpProxy\n      required: true\n      description: A full instance of a TargetHttpProxy\n  apply:\n    description: The function used to apply information about a TargetHttpProxy\n    parameters:\n    - name: TargetHttpProxy\n      required: true\n      description: A full instance of a TargetHttpProxy\n  delete:\n    description: The function used to delete a TargetHttpProxy\n    parameters:\n    - name: TargetHttpProxy\n      required: true\n      description: A full instance of a TargetHttpProxy\n  deleteAll:\n    description: The function used to delete all TargetHttpProxy\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many TargetHttpProxy\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    TargetHttpProxy:\n      title: TargetHttpProxy\n      x-dcl-id: projects/{{project}}/global/targetHttpProxies/{{name}}\n      x-dcl-locations:\n      - global\n      x-dcl-parent-container: project\n      type: object\n      required:\n      - name\n      - urlMap\n      - project\n      properties:\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: An optional description of this resource. Provide this property\n            when you create the resource.\n          x-kubernetes-immutable: true\n        id:\n          type: integer\n          format: int64\n          x-dcl-go-name: Id\n          readOnly: true\n          description: The unique identifier for the resource. This identifier is\n            defined by the server.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Name of the resource. Provided by the client when the resource\n            is created. The name must be 1-63 characters long, and comply with [RFC1035](https://www.ietf.org/rfc/rfc1035.txt).\n            The first character must be a lowercase letter, and all following characters\n            must be a dash, lowercase letter, or digit, except the last character,\n            which cannot be a dash.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          readOnly: true\n          description: Server-defined URL for the resource.\n          x-kubernetes-immutable: true\n        urlMap:\n          type: string\n          x-dcl-go-name: UrlMap\n          description: A reference to the UrlMap resource that defines the mapping\n            from URL to the BackendService.\n          x-dcl-references:\n          - resource: Compute/UrlMap\n            field: selfLink\n")

// 3250 bytes
// MD5: c525a031e7da3de610c51b73adf52a0e