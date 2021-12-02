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
// gen_go_data -package alpha -var YAML_grpc_route blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networkservices/alpha/grpc_route.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networkservices/alpha/grpc_route.yaml
var YAML_grpc_route = []byte("info:\n  title: NetworkServices/GrpcRoute\n  description: The NetworkServices GrpcRoute resource\n  x-dcl-struct-name: GrpcRoute\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a GrpcRoute\n    parameters:\n    - name: GrpcRoute\n      required: true\n      description: A full instance of a GrpcRoute\n  apply:\n    description: The function used to apply information about a GrpcRoute\n    parameters:\n    - name: GrpcRoute\n      required: true\n      description: A full instance of a GrpcRoute\n  delete:\n    description: The function used to delete a GrpcRoute\n    parameters:\n    - name: GrpcRoute\n      required: true\n      description: A full instance of a GrpcRoute\n  deleteAll:\n    description: The function used to delete all GrpcRoute\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many GrpcRoute\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    GrpcRoute:\n      title: GrpcRoute\n      x-dcl-id: projects/{{project}}/locations/{{location}}/grpcRoutes/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - hostnames\n      - rules\n      - project\n      - location\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The timestamp when the resource was created.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. A free-text description of the resource. Max length\n            1024 characters.\n        gateways:\n          type: array\n          x-dcl-go-name: Gateways\n          description: 'Optional. Gateways defines a list of gateways this GrpcRoute\n            is attached to, as one of the routing rules to route the requests served\n            by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`'\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n            x-dcl-references:\n            - resource: Networkservices/Gateway\n              field: selfLink\n        hostnames:\n          type: array\n          x-dcl-go-name: Hostnames\n          description: 'Required. Service hostnames with an optional port for which\n            this route describes traffic. Format: [:] Hostname is the fully qualified\n            domain name of a network host. This matches the RFC 1123 definition of\n            a hostname with 2 notable exceptions: - IPs are not allowed. - A hostname\n            may be prefixed with a wildcard label (*.). The wildcard label must appear\n            by itself as the first label. Hostname can be “precise” which is a domain\n            name without the terminating dot of a network host (e.g. “foo.example.com”)\n            or “wildcard”, which is a domain name prefixed with a single wildcard\n            label (e.g. *.example.com). Note that as per RFC1035 and RFC1123, a label\n            must consist of lower case alphanumeric characters or ‘-’, and must start\n            and end with an alphanumeric character. No other punctuation is allowed.\n            The routes associated with a Router must have unique hostnames. If you\n            attempt to attach multiple routes with conflicting hostnames, the configuration\n            will be rejected. For example, while it is acceptable for routes for the\n            hostnames \"*.foo.bar.com\" and \"*.bar.com\" to be associated with the same\n            route, it is not possible to associate two routes both with \"*.bar.com\"\n            or both with \"bar.com\". In the case that multiple routes match the hostname,\n            the most specific match will be selected. For example, \"foo.bar.baz.com\"\n            will take precedence over \"*.bar.baz.com\" and \"*.bar.baz.com\" will take\n            precedence over \"*.baz.com\". If a port is specified, then gRPC clients\n            must use the channel URI with the port to match this rule (i.e. \"xds:///service:123\"),\n            otherwise they must supply the URI without a port (i.e. \"xds:///service\").'\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Optional. Set of label tags associated with the GrpcRoute resource.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        meshes:\n          type: array\n          x-dcl-go-name: Meshes\n          description: 'Optional. Meshes defines a list of meshes this GrpcRoute is\n            attached to, as one of the routing rules to route the requests served\n            by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/`'\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n            x-dcl-references:\n            - resource: Networkservices/Mesh\n              field: selfLink\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Required. Name of the GrpcRoute resource. It matches pattern\n            `projects/*/locations/global/grpcRoutes/`\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        routers:\n          type: array\n          x-dcl-go-name: Routers\n          description: 'Optional. Routers define a list of routers this GrpcRoute\n            should be served by. Each router reference should match the pattern: `projects/*/locations/global/routers/`'\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n        rules:\n          type: array\n          x-dcl-go-name: Rules\n          description: Required. A list of detailed rules defining how to route traffic.\n            Within a single GrpcRoute, the GrpcRoute.RouteAction associated with the\n            first matching GrpcRoute.RouteRule will be executed. At least one rule\n            must be supplied.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: GrpcRouteRules\n            required:\n            - action\n            properties:\n              action:\n                type: object\n                x-dcl-go-name: Action\n                x-dcl-go-type: GrpcRouteRulesAction\n                description: Required. A detailed rule defining how to route traffic.\n                  This field is required.\n                properties:\n                  destinations:\n                    type: array\n                    x-dcl-go-name: Destinations\n                    description: Optional. The destination services to which traffic\n                      should be forwarded. If multiple destinations are specified,\n                      traffic will be split between Backend Service(s) according to\n                      the weight field of these destinations.\n                    x-dcl-send-empty: true\n                    x-dcl-list-type: list\n                    items:\n                      type: object\n                      x-dcl-go-type: GrpcRouteRulesActionDestinations\n                      required:\n                      - serviceName\n                      properties:\n                        serviceName:\n                          type: string\n                          x-dcl-go-name: ServiceName\n                          description: Required. The URL of a destination service\n                            to which to route traffic. Must refer to either a BackendService\n                            or ServiceDirectoryService.\n                        weight:\n                          type: integer\n                          format: int64\n                          x-dcl-go-name: Weight\n                          description: 'Optional. Specifies the proportion of requests\n                            forwarded to the backend referenced by the serviceName\n                            field. This is computed as: weight/Sum(weights in this\n                            destination list). For non-zero values, there may be some\n                            epsilon from the exact proportion defined here depending\n                            on the precision an implementation supports. If only one\n                            serviceName is specified and it has a weight greater than\n                            0, 100% of the traffic is forwarded to that backend. If\n                            weights are specified for any one service name, they need\n                            to be specified for all of them. If weights are unspecified\n                            for all services, then, traffic is distributed in equal\n                            proportions to all of them.'\n                  faultInjectionPolicy:\n                    type: object\n                    x-dcl-go-name: FaultInjectionPolicy\n                    x-dcl-go-type: GrpcRouteRulesActionFaultInjectionPolicy\n                    description: Optional. The specification for fault injection introduced\n                      into traffic to test the resiliency of clients to destination\n                      service failure. As part of fault injection, when clients send\n                      requests to a destination, delays can be introduced on a percentage\n                      of requests before sending those requests to the destination\n                      service. Similarly requests from clients can be aborted by for\n                      a percentage of requests. timeout and retry_policy will be ignored\n                      by clients that are configured with a fault_injection_policy\n                    properties:\n                      abort:\n                        type: object\n                        x-dcl-go-name: Abort\n                        x-dcl-go-type: GrpcRouteRulesActionFaultInjectionPolicyAbort\n                        description: The specification for aborting to client requests.\n                        properties:\n                          httpStatus:\n                            type: integer\n                            format: int64\n                            x-dcl-go-name: HttpStatus\n                            description: The HTTP status code used to abort the request.\n                              The value must be between 200 and 599 inclusive.\n                          percentage:\n                            type: integer\n                            format: int64\n                            x-dcl-go-name: Percentage\n                            description: The percentage of traffic which will be aborted.\n                              The value must be between [0, 100]\n                      delay:\n                        type: object\n                        x-dcl-go-name: Delay\n                        x-dcl-go-type: GrpcRouteRulesActionFaultInjectionPolicyDelay\n                        description: The specification for injecting delay to client\n                          requests.\n                        properties:\n                          fixedDelay:\n                            type: string\n                            x-dcl-go-name: FixedDelay\n                            description: Specify a fixed delay before forwarding the\n                              request.\n                          percentage:\n                            type: integer\n                            format: int64\n                            x-dcl-go-name: Percentage\n                            description: The percentage of traffic on which delay\n                              will be injected. The value must be between [0, 100]\n                  requestHeaderModifier:\n                    type: object\n                    x-dcl-go-name: RequestHeaderModifier\n                    x-dcl-go-type: GrpcRouteRulesActionRequestHeaderModifier\n                    description: Optional. The specification for modifying the headers\n                      of a matching request prior to delivery of the request to the\n                      destination. Cannot be set if the route is attached to a Router\n                      whose type is PROXYLESS_GRPC.\n                    properties:\n                      add:\n                        type: object\n                        additionalProperties:\n                          type: string\n                        x-dcl-go-name: Add\n                        description: Add the headers with given map where key is the\n                          name of the header, value is the value of the header.\n                      remove:\n                        type: array\n                        x-dcl-go-name: Remove\n                        description: Remove headers (matching by header names) specified\n                          in the list.\n                        x-dcl-send-empty: true\n                        x-dcl-list-type: list\n                        items:\n                          type: string\n                          x-dcl-go-type: string\n                      set:\n                        type: object\n                        additionalProperties:\n                          type: string\n                        x-dcl-go-name: Set\n                        description: Completely overwrite/replace the headers with\n                          given map where key is the name of the header, value is\n                          the value of the header.\n                  requestMirrorPolicy:\n                    type: object\n                    x-dcl-go-name: RequestMirrorPolicy\n                    x-dcl-go-type: GrpcRouteRulesActionRequestMirrorPolicy\n                    description: Optional. Specifies the policy on how requests intended\n                      for the route's destination are mirrored to a separate mirrored\n                      destination. The proxy will not wait for the mirrored destination\n                      to respond before returning the response. Prior to sending traffic\n                      to the mirrored service, the host / authority header is suffixed\n                      with -shadow. Cannot be set if the route is attached to a Router\n                      whose type is PROXYLESS_GRPC.\n                    properties:\n                      destination:\n                        type: object\n                        x-dcl-go-name: Destination\n                        x-dcl-go-type: GrpcRouteRulesActionRequestMirrorPolicyDestination\n                        description: The destination the requests will be mirrored\n                          to. The weight of the destination will be ignored.\n                        required:\n                        - serviceName\n                        properties:\n                          serviceName:\n                            type: string\n                            x-dcl-go-name: ServiceName\n                            description: Required. The URL of a destination service\n                              to which to route traffic. Must refer to either a BackendService\n                              or ServiceDirectoryService.\n                          weight:\n                            type: integer\n                            format: int64\n                            x-dcl-go-name: Weight\n                            description: 'Optional. Specifies the proportion of requests\n                              forwarded to the backend referenced by the serviceName\n                              field. This is computed as: weight/Sum(weights in this\n                              destination list). For non-zero values, there may be\n                              some epsilon from the exact proportion defined here\n                              depending on the precision an implementation supports.\n                              If only one serviceName is specified and it has a weight\n                              greater than 0, 100% of the traffic is forwarded to\n                              that backend. If weights are specified for any one service\n                              name, they need to be specified for all of them. If\n                              weights are unspecified for all services, then, traffic\n                              is distributed in equal proportions to all of them.'\n                  responseHeaderModifier:\n                    type: object\n                    x-dcl-go-name: ResponseHeaderModifier\n                    x-dcl-go-type: GrpcRouteRulesActionResponseHeaderModifier\n                    description: Optional. The specification for modifying the headers\n                      of a response prior to sending the response back to the client.\n                      Cannot be set if the route is attached to a Router whose type\n                      is PROXYLESS_GRPC.\n                    properties:\n                      add:\n                        type: object\n                        additionalProperties:\n                          type: string\n                        x-dcl-go-name: Add\n                        description: Add the headers with given map where key is the\n                          name of the header, value is the value of the header.\n                      remove:\n                        type: array\n                        x-dcl-go-name: Remove\n                        description: Remove headers (matching by header names) specified\n                          in the list.\n                        x-dcl-send-empty: true\n                        x-dcl-list-type: list\n                        items:\n                          type: string\n                          x-dcl-go-type: string\n                      set:\n                        type: object\n                        additionalProperties:\n                          type: string\n                        x-dcl-go-name: Set\n                        description: Completely overwrite/replace the headers with\n                          given map where key is the name of the header, value is\n                          the value of the header.\n                  retryPolicy:\n                    type: object\n                    x-dcl-go-name: RetryPolicy\n                    x-dcl-go-type: GrpcRouteRulesActionRetryPolicy\n                    description: Optional. Specifies the retry policy associated with\n                      this route.\n                    properties:\n                      numRetries:\n                        type: integer\n                        format: int64\n                        x-dcl-go-name: NumRetries\n                        description: Specifies the allowed number of retries. This\n                          number must be > 0. If not specpfied, default to 1.\n                      perTryTimeout:\n                        type: string\n                        x-dcl-go-name: PerTryTimeout\n                        description: If not specified, will use the timeout set in\n                          the RouteAction. If timeout is not set in the RouteAction,\n                          will use the largest timeout among all Backend Services\n                          associated with the route.\n                      retryConditions:\n                        type: array\n                        x-dcl-go-name: RetryConditions\n                        description: '- connect-failure: Router will retry on failures\n                          connecting to Backend Services, for example due to connection\n                          timeouts. - refused-stream: Router will retry if the backend\n                          service resets the stream with a REFUSED_STREAM error code.\n                          This reset type indicates that it is safe to retry. - cancelled:\n                          Router will retry if the gRPC status code in the response\n                          header is set to cancelled - deadline-exceeded: Router will\n                          retry if the gRPC status code in the response header is\n                          set to deadline-exceeded - resource-exhausted: Router will\n                          retry if the gRPC status code in the response header is\n                          set to resource-exhausted - unavailable: Router will retry\n                          if the gRPC status code in the response header is set to\n                          unavailable'\n                        x-dcl-send-empty: true\n                        x-dcl-list-type: list\n                        items:\n                          type: string\n                          x-dcl-go-type: string\n                  timeout:\n                    type: string\n                    x-dcl-go-name: Timeout\n                    description: Optional. Specifies the timeout for selected route.\n                      Timeout is computed from the time the request has been fully\n                      processed (i.e. end of stream) up until the response has been\n                      completely processed. Timeout includes all retries.\n                  urlRewrite:\n                    type: object\n                    x-dcl-go-name: UrlRewrite\n                    x-dcl-go-type: GrpcRouteRulesActionUrlRewrite\n                    description: Optional. The specification for rewrite URL before\n                      forwarding requests to the destination. Cannot be set if the\n                      route is attached to a Router whose type is PROXYLESS_GRPC.\n                    properties:\n                      hostRewrite:\n                        type: string\n                        x-dcl-go-name: HostRewrite\n                        description: Prior to forwarding the request to the selected\n                          destination, the requests host header is replaced by this\n                          value.\n                      pathPrefixRewrite:\n                        type: string\n                        x-dcl-go-name: PathPrefixRewrite\n                        description: Prior to forwarding the request to the selected\n                          destination, the matching portion of the requests path is\n                          replaced by this value.\n              matches:\n                type: array\n                x-dcl-go-name: Matches\n                description: Optional. Matches define conditions used for matching\n                  the rule against incoming gRPC requests. Each match is independent,\n                  i.e. this rule will be matched if ANY one of the matches is satisfied.\n                  If no matches field is specified, this rule will unconditionally\n                  match traffic.\n                x-dcl-send-empty: true\n                x-dcl-list-type: list\n                items:\n                  type: object\n                  x-dcl-go-type: GrpcRouteRulesMatches\n                  properties:\n                    headers:\n                      type: array\n                      x-dcl-go-name: Headers\n                      description: Optional. Specifies a collection of headers to\n                        match.\n                      x-dcl-send-empty: true\n                      x-dcl-list-type: list\n                      items:\n                        type: object\n                        x-dcl-go-type: GrpcRouteRulesMatchesHeaders\n                        required:\n                        - key\n                        - value\n                        properties:\n                          key:\n                            type: string\n                            x-dcl-go-name: Key\n                            description: Required. The key of the header.\n                          type:\n                            type: string\n                            x-dcl-go-name: Type\n                            x-dcl-go-type: GrpcRouteRulesMatchesHeadersTypeEnum\n                            description: 'Optional. Specifies how to match against\n                              the value of the header. If not specified, a default\n                              value of EXACT is used. Possible values: MATCH_TYPE_UNSPECIFIED,\n                              MATCH_ANY, MATCH_ALL'\n                            enum:\n                            - MATCH_TYPE_UNSPECIFIED\n                            - MATCH_ANY\n                            - MATCH_ALL\n                          value:\n                            type: string\n                            x-dcl-go-name: Value\n                            description: Required. The value of the header.\n                    method:\n                      type: object\n                      x-dcl-go-name: Method\n                      x-dcl-go-type: GrpcRouteRulesMatchesMethod\n                      description: Optional. A gRPC method to match against. If this\n                        field is empty or omitted, will match all methods.\n                      required:\n                      - grpcService\n                      - grpcMethod\n                      properties:\n                        caseSensitive:\n                          type: boolean\n                          x-dcl-go-name: CaseSensitive\n                          description: Optional. Specifies that matches are case sensitive.\n                            The default value is true. case_sensitive must not be\n                            used with a type of REGULAR_EXPRESSION.\n                        grpcMethod:\n                          type: string\n                          x-dcl-go-name: GrpcMethod\n                          description: Required. Name of the method to match against.\n                            If unspecified, will match all methods.\n                        grpcService:\n                          type: string\n                          x-dcl-go-name: GrpcService\n                          description: Required. Name of the service to match against.\n                            If unspecified, will match all services.\n                        type:\n                          type: string\n                          x-dcl-go-name: Type\n                          x-dcl-go-type: GrpcRouteRulesMatchesMethodTypeEnum\n                          description: 'Optional. Specifies how to match against the\n                            name. If not specified, a default value of \"EXACT\" is\n                            used. Possible values: MATCH_TYPE_UNSPECIFIED, MATCH_ANY,\n                            MATCH_ALL'\n                          enum:\n                          - MATCH_TYPE_UNSPECIFIED\n                          - MATCH_ANY\n                          - MATCH_ALL\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The timestamp when the resource was updated.\n          x-kubernetes-immutable: true\n")

// 27681 bytes
// MD5: 236bda66d01fe2b2283ff0460b87f656
