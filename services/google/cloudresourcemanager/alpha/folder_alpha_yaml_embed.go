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
// gen_go_data -package alpha -var YAML_folder blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudresourcemanager/alpha/folder.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudresourcemanager/alpha/folder.yaml
var YAML_folder = []byte("info:\n  title: CloudResourceManager/Folder\n  description: The CloudResourceManager Folder resource\n  x-dcl-struct-name: Folder\n  x-dcl-has-iam: true\npaths:\n  get:\n    description: The function used to get information about a Folder\n    parameters:\n    - name: Folder\n      required: true\n      description: A full instance of a Folder\n  apply:\n    description: The function used to apply information about a Folder\n    parameters:\n    - name: Folder\n      required: true\n      description: A full instance of a Folder\n  delete:\n    description: The function used to delete a Folder\n    parameters:\n    - name: Folder\n      required: true\n      description: A full instance of a Folder\n  deleteAll:\n    description: The function used to delete all Folder\n    parameters:\n    - name: parent\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Folder\n    parameters:\n    - name: parent\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Folder:\n      title: Folder\n      x-dcl-id: folders/{{name}}\n      x-dcl-has-iam: true\n      type: object\n      required:\n      - parent\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. Timestamp when the Folder was created. Assigned\n            by the server.\n          x-kubernetes-immutable: true\n        deleteTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: DeleteTime\n          readOnly: true\n          description: Output only. Timestamp when the Folder was requested to be\n            deleted.\n          x-kubernetes-immutable: true\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: 'The folder''s display name. A folder''s display name must\n            be unique amongst its siblings, e.g. no two folders with the same parent\n            can share the same display name. The display name must start and end with\n            a letter or digit, may contain letters, digits, spaces, hyphens and underscores\n            and can be no longer than 30 characters. This is captured by the regular\n            expression: `[p{L}p{N}]([p{L}p{N}_- ]{0,28}[p{L}p{N}])?`.'\n        etag:\n          type: string\n          x-dcl-go-name: Etag\n          readOnly: true\n          description: Output only. A checksum computed by the server based on the\n            current value of the Folder resource. This may be sent on update and delete\n            requests to ensure the client has an up-to-date value before proceeding.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Output only. The resource name of the Folder.\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n        parent:\n          type: string\n          x-dcl-go-name: Parent\n          description: Required. The Folder's parent's resource name. Updates to the\n            folder's parent must be performed via MoveFolder.\n          x-dcl-forward-slash-allowed: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: FolderStateEnum\n          readOnly: true\n          description: 'Output only. The lifecycle state of the folder. Possible values:\n            LIFECYCLE_STATE_UNSPECIFIED, ACTIVE, DELETE_REQUESTED'\n          x-kubernetes-immutable: true\n          enum:\n          - LIFECYCLE_STATE_UNSPECIFIED\n          - ACTIVE\n          - DELETE_REQUESTED\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. Timestamp when the Folder was last modified.\n          x-kubernetes-immutable: true\n")

// 3896 bytes
// MD5: a204e2566be9c3ea2227c0e40b87a34c
