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
// gen_go_data -package beta -var YAML_ssl_policy blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/beta/ssl_policy.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/beta/ssl_policy.yaml
var YAML_ssl_policy = []byte("info:\n  title: Compute/SslPolicy\n  description: DCL Specification for the Compute SslPolicy resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a SslPolicy\n    parameters:\n    - name: SslPolicy\n      required: true\n      description: A full instance of a SslPolicy\n  apply:\n    description: The function used to apply information about a SslPolicy\n    parameters:\n    - name: SslPolicy\n      required: true\n      description: A full instance of a SslPolicy\n  delete:\n    description: The function used to delete a SslPolicy\n    parameters:\n    - name: SslPolicy\n      required: true\n      description: A full instance of a SslPolicy\n  deleteAll:\n    description: The function used to delete all SslPolicy\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many SslPolicy\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    SslPolicy:\n      title: SslPolicy\n      x-dcl-id: projects/{{project}}/global/sslPolicies/{{name}}\n      x-dcl-locations:\n      - global\n      x-dcl-parent-container: project\n      type: object\n      required:\n      - name\n      - project\n      properties:\n        customFeature:\n          type: array\n          x-dcl-go-name: CustomFeature\n          description: |-\n            A list of features enabled when the selected profile is CUSTOM. The\n            <listAvailableFeatures> method returns the set of features that can be\n            specified in this list. This field must be empty if the profile is not\n            <code>CUSTOM</code>.\n          x-dcl-send-empty: true\n          x-dcl-list-type: set\n          items:\n            type: string\n            x-dcl-go-type: string\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: |-\n            An optional description of this resource. Provide this property when you\n            create the resource.\n          x-kubernetes-immutable: true\n        enabledFeature:\n          type: array\n          x-dcl-go-name: EnabledFeature\n          readOnly: true\n          description: '[Output Only] The list of features enabled in the SSL policy.'\n          x-dcl-send-empty: true\n          x-dcl-list-type: set\n          items:\n            type: string\n            x-dcl-go-type: string\n        id:\n          type: integer\n          format: int64\n          x-dcl-go-name: Id\n          readOnly: true\n          description: |-\n            [Output Only] The unique identifier for the resource. This identifier is\n            defined by the server.\n          x-kubernetes-immutable: true\n        minTlsVersion:\n          type: string\n          x-dcl-go-name: MinTlsVersion\n          x-dcl-go-type: SslPolicyMinTlsVersionEnum\n          description: |-\n            The minimum version of SSL protocol that can be used by the clients to\n            establish a connection with the load balancer. This can be one of\n            <code>TLS_1_0</code>, <code>TLS_1_1</code>, <code>TLS_1_2</code>.\n          default: TLS_1_0\n          enum:\n          - TLS_1_0\n          - TLS_1_1\n          - TLS_1_2\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: |-\n            Name of the resource. The name must be 1-63 characters long, and comply\n            with <a href=\"https://www.ietf.org/rfc/rfc1035.txt\"\n            target=\"_blank\">RFC1035</a>. Specifically, the name must be 1-63 characters\n            long and match the regular expression\n            `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character\n            must be a lowercase letter, and all following characters must be a dash,\n            lowercase letter, or digit, except the last character, which cannot be a\n            dash.\n          x-kubernetes-immutable: true\n        profile:\n          type: string\n          x-dcl-go-name: Profile\n          x-dcl-go-type: SslPolicyProfileEnum\n          description: |-\n            Profile specifies the set of SSL features that can be used by the load\n            balancer when negotiating SSL with clients. This can be one of\n            <code>COMPATIBLE</code>, <code>MODERN</code>, <code>RESTRICTED</code>, or\n            <code>CUSTOM</code>. If using <code>CUSTOM</code>, the set of SSL features\n            to enable must be specified in the <code>customFeatures</code> field.\n          default: COMPATIBLE\n          enum:\n          - COMPATIBLE\n          - MODERN\n          - RESTRICTED\n          - CUSTOM\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          readOnly: true\n          description: '[Output Only] Server-defined URL for the resource.'\n          x-kubernetes-immutable: true\n        warning:\n          type: array\n          x-dcl-go-name: Warning\n          readOnly: true\n          description: |-\n            [Output Only] If potential misconfigurations are detected for this\n            SSL policy, this field will be populated with warning messages.\n          x-kubernetes-immutable: true\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: SslPolicyWarning\n            properties:\n              code:\n                type: string\n                x-dcl-go-name: Code\n                x-kubernetes-immutable: true\n              data:\n                type: array\n                x-dcl-go-name: Data\n                x-kubernetes-immutable: true\n                x-dcl-send-empty: true\n                x-dcl-list-type: list\n                items:\n                  type: object\n                  x-dcl-go-type: SslPolicyWarningData\n                  properties:\n                    key:\n                      type: string\n                      x-dcl-go-name: Key\n                      x-kubernetes-immutable: true\n                    value:\n                      type: string\n                      x-dcl-go-name: Value\n                      x-kubernetes-immutable: true\n              message:\n                type: string\n                x-dcl-go-name: Message\n                x-kubernetes-immutable: true\n")

// 6529 bytes
// MD5: a5f6cd6407d3ba1279e58bb7a93ac682