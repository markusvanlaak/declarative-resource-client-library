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
// gen_go_data -package alpha -var YAML_http_route blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networkservices/alpha/http_route.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networkservices/alpha/http_route.yaml
var YAML_http_route = []byte("info:\n  title: NetworkServices/HttpRoute\n  description: The NetworkServices HttpRoute resource\n  x-dcl-struct-name: HttpRoute\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a HttpRoute\n    parameters:\n    - name: HttpRoute\n      required: true\n      description: A full instance of a HttpRoute\n  apply:\n    description: The function used to apply information about a HttpRoute\n    parameters:\n    - name: HttpRoute\n      required: true\n      description: A full instance of a HttpRoute\n  delete:\n    description: The function used to delete a HttpRoute\n    parameters:\n    - name: HttpRoute\n      required: true\n      description: A full instance of a HttpRoute\n  deleteAll:\n    description: The function used to delete all HttpRoute\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many HttpRoute\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    HttpRoute:\n      title: HttpRoute\n      x-dcl-id: projects/{{project}}/locations/{{location}}/httpRoutes/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - hostnames\n      - rules\n      - project\n      - location\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The timestamp when the resource was created.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. A free-text description of the resource. Max length\n            1024 characters.\n        gateways:\n          type: array\n          x-dcl-go-name: Gateways\n          description: 'Optional. Gateways defines a list of gateways this HttpRoute\n            is attached to, as one of the routing rules to route the requests served\n            by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`'\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n            x-dcl-references:\n            - resource: Networkservices/Gateway\n              field: selfLink\n        hostnames:\n          type: array\n          x-dcl-go-name: Hostnames\n          description: Required. Hostnames define a set of hosts that should match\n            against the HTTP host header to select a HttpRoute to process the request.\n            Hostname is the fully qualified domain name of a network host, as defined\n            by RFC 1123 with the exception that ip addresses are not allowed. Wildcard\n            hosts are supported as \"*\" (no prefix or suffix allowed).\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Optional. Set of label tags associated with the HttpRoute resource.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        meshes:\n          type: array\n          x-dcl-go-name: Meshes\n          description: 'Optional. Meshes defines a list of meshes this HttpRoute is\n            attached to, as one of the routing rules to route the requests served\n            by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/`\n            The attached Mesh should be of a type SIDECAR'\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n            x-dcl-references:\n            - resource: Networkservices/Mesh\n              field: selfLink\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Required. Name of the HttpRoute resource. It matches pattern\n            `projects/*/locations/global/httpRoutes/http_route_name>`.\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        routers:\n          type: array\n          x-dcl-go-name: Routers\n          description: 'Optional. Routers define a list of routers this HttpRoute\n            should be served by. Each router reference should match the pattern: `projects/*/locations/global/routers/`\n            The attached Router should be of a type PROXY'\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n        rules:\n          type: array\n          x-dcl-go-name: Rules\n          description: Required. Rules that define how traffic is routed and handled.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: HttpRouteRules\n            properties:\n              action:\n                type: object\n                x-dcl-go-name: Action\n                x-dcl-go-type: HttpRouteRulesAction\n                description: The detailed rule defining how to route matched traffic.\n                properties:\n                  corsPolicy:\n                    type: object\n                    x-dcl-go-name: CorsPolicy\n                    x-dcl-go-type: HttpRouteRulesActionCorsPolicy\n                    description: The specification for allowing client side cross-origin\n                      requests.\n                    properties:\n                      allowCredentials:\n                        type: boolean\n                        x-dcl-go-name: AllowCredentials\n                        description: In response to a preflight request, setting this\n                          to true indicates that the actual request can include user\n                          credentials. This translates to the Access-Control-Allow-Credentials\n                          header. Default value is false.\n                      allowHeaders:\n                        type: array\n                        x-dcl-go-name: AllowHeaders\n                        description: Specifies the content for Access-Control-Allow-Headers\n                          header.\n                        x-dcl-send-empty: true\n                        x-dcl-list-type: list\n                        items:\n                          type: string\n                          x-dcl-go-type: string\n                      allowMethods:\n                        type: array\n                        x-dcl-go-name: AllowMethods\n                        description: Specifies the content for Access-Control-Allow-Methods\n                          header.\n                        x-dcl-send-empty: true\n                        x-dcl-list-type: list\n                        items:\n                          type: string\n                          x-dcl-go-type: string\n                      allowOriginRegexes:\n                        type: array\n                        x-dcl-go-name: AllowOriginRegexes\n                        description: Specifies the regular expression patterns that\n                          match allowed origins. For regular expression grammar, please\n                          see https://github.com/google/re2/wiki/Syntax.\n                        x-dcl-send-empty: true\n                        x-dcl-list-type: list\n                        items:\n                          type: string\n                          x-dcl-go-type: string\n                      allowOrigins:\n                        type: array\n                        x-dcl-go-name: AllowOrigins\n                        description: Specifies the list of origins that will be allowed\n                          to do CORS requests. An origin is allowed if it matches\n                          either an item in allow_origins or an item in allow_origin_regexes.\n                        x-dcl-send-empty: true\n                        x-dcl-list-type: list\n                        items:\n                          type: string\n                          x-dcl-go-type: string\n                      disabled:\n                        type: boolean\n                        x-dcl-go-name: Disabled\n                        description: If true, the CORS policy is disabled. The default\n                          value is false, which indicates that the CORS policy is\n                          in effect.\n                      exposeHeaders:\n                        type: array\n                        x-dcl-go-name: ExposeHeaders\n                        description: Specifies the content for Access-Control-Expose-Headers\n                          header.\n                        x-dcl-send-empty: true\n                        x-dcl-list-type: list\n                        items:\n                          type: string\n                          x-dcl-go-type: string\n                      maxAge:\n                        type: string\n                        x-dcl-go-name: MaxAge\n                        description: Specifies how long result of a preflight request\n                          can be cached in seconds. This translates to the Access-Control-Max-Age\n                          header.\n                  destinations:\n                    type: array\n                    x-dcl-go-name: Destinations\n                    description: The destination to which traffic should be forwarded.\n                    x-dcl-send-empty: true\n                    x-dcl-list-type: list\n                    items:\n                      type: object\n                      x-dcl-go-type: HttpRouteRulesActionDestinations\n                      properties:\n                        serviceName:\n                          type: string\n                          x-dcl-go-name: ServiceName\n                          description: The URL of a BackendService to route traffic\n                            to.\n                        weight:\n                          type: integer\n                          format: int64\n                          x-dcl-go-name: Weight\n                          description: 'Specifies the proportion of requests forwarded\n                            to the backend referenced by the serviceName field. This\n                            is computed as: weight/Sum(weights in this destination\n                            list). For non-zero values, there may be some epsilon\n                            from the exact proportion defined here depending on the\n                            precision an implementation supports. If only one serviceName\n                            is specified and it has a weight greater than 0, 100%\n                            of the traffic is forwarded to that backend. If weights\n                            are specified for any one service name, they need to be\n                            specified for all of them. If weights are unspecified\n                            for all services, then, traffic is distributed in equal\n                            proportions to all of them.'\n                  faultInjectionPolicy:\n                    type: object\n                    x-dcl-go-name: FaultInjectionPolicy\n                    x-dcl-go-type: HttpRouteRulesActionFaultInjectionPolicy\n                    description: The specification for fault injection introduced\n                      into traffic to test the resiliency of clients to backend service\n                      failure. As part of fault injection, when clients send requests\n                      to a backend service, delays can be introduced on a percentage\n                      of requests before sending those requests to the backend service.\n                      Similarly requests from clients can be aborted for a percentage\n                      of requests. timeout and retry_policy will be ignored by clients\n                      that are configured with a fault_injection_policy\n                    properties:\n                      abort:\n                        type: object\n                        x-dcl-go-name: Abort\n                        x-dcl-go-type: HttpRouteRulesActionFaultInjectionPolicyAbort\n                        description: The specification for aborting to client requests.\n                        properties:\n                          httpStatus:\n                            type: integer\n                            format: int64\n                            x-dcl-go-name: HttpStatus\n                            description: The HTTP status code used to abort the request.\n                              The value must be between 200 and 599 inclusive.\n                          percentage:\n                            type: integer\n                            format: int64\n                            x-dcl-go-name: Percentage\n                            description: The percentage of traffic which will be aborted.\n                              The value must be between [0, 100]\n                      delay:\n                        type: object\n                        x-dcl-go-name: Delay\n                        x-dcl-go-type: HttpRouteRulesActionFaultInjectionPolicyDelay\n                        description: The specification for injecting delay to client\n                          requests.\n                        properties:\n                          fixedDelay:\n                            type: string\n                            x-dcl-go-name: FixedDelay\n                            description: Specify a fixed delay before forwarding the\n                              request.\n                          percentage:\n                            type: integer\n                            format: int64\n                            x-dcl-go-name: Percentage\n                            description: The percentage of traffic on which delay\n                              will be injected. The value must be between [0, 100]\n                  originalDestination:\n                    type: boolean\n                    x-dcl-go-name: OriginalDestination\n                    description: If true, the matched traffic will use the destination\n                      ip and port of the original connection (as it was not processed\n                      by proxy) as the destination of the request. Only one of destinations,\n                      redirect, original_destination can be set.\n                  redirect:\n                    type: object\n                    x-dcl-go-name: Redirect\n                    x-dcl-go-type: HttpRouteRulesActionRedirect\n                    description: If set, the request is directed as configured by\n                      this field.\n                    properties:\n                      hostRedirect:\n                        type: string\n                        x-dcl-go-name: HostRedirect\n                        description: The host that will be used in the redirect response\n                          instead of the one that was supplied in the request.\n                      httpsRedirect:\n                        type: boolean\n                        x-dcl-go-name: HttpsRedirect\n                        description: If set to true, the URL scheme in the redirected\n                          request is set to https. If set to false, the URL scheme\n                          of the redirected request will remain the same as that of\n                          the request. The default is set to false.\n                      pathRedirect:\n                        type: string\n                        x-dcl-go-name: PathRedirect\n                        description: The path that will be used in the redirect response\n                          instead of the one that was supplied in the request. path_redirect\n                          can not be supplied together with prefix_redirect. Supply\n                          one alone or neither. If neither is supplied, the path of\n                          the original request will be used for the redirect.\n                      portRedirect:\n                        type: integer\n                        format: int64\n                        x-dcl-go-name: PortRedirect\n                        description: The port that will be used in the redirected\n                          request instead of the one that was supplied in the request.\n                      prefixRewrite:\n                        type: string\n                        x-dcl-go-name: PrefixRewrite\n                        description: Indicates that during redirection, the matched\n                          prefix (or path) should be swapped with this value. This\n                          option allows URLs be dynamically created based on the request.\n                      responseCode:\n                        type: string\n                        x-dcl-go-name: ResponseCode\n                        x-dcl-go-type: HttpRouteRulesActionRedirectResponseCodeEnum\n                        description: 'The HTTP Status code to use for the redirect.\n                          Possible values: MOVED_PERMANENTLY_DEFAULT, FOUND, SEE_OTHER,\n                          TEMPORARY_REDIRECT, PERMANENT_REDIRECT'\n                        enum:\n                        - MOVED_PERMANENTLY_DEFAULT\n                        - FOUND\n                        - SEE_OTHER\n                        - TEMPORARY_REDIRECT\n                        - PERMANENT_REDIRECT\n                      stripQuery:\n                        type: boolean\n                        x-dcl-go-name: StripQuery\n                        description: if set to true, any accompanying query portion\n                          of the original URL is removed prior to redirecting the\n                          request. If set to false, the query portion of the original\n                          URL is retained. The default is set to false.\n                  requestHeaderModifier:\n                    type: object\n                    x-dcl-go-name: RequestHeaderModifier\n                    x-dcl-go-type: HttpRouteRulesActionRequestHeaderModifier\n                    description: The specification for modifying the headers of a\n                      matching request prior to delivery of the request to the destination.\n                    properties:\n                      add:\n                        type: object\n                        additionalProperties:\n                          type: string\n                        x-dcl-go-name: Add\n                        description: Add the headers with given map where key is the\n                          name of the header, value is the value of the header.\n                      remove:\n                        type: array\n                        x-dcl-go-name: Remove\n                        description: Remove headers (matching by header names) specified\n                          in the list.\n                        x-dcl-send-empty: true\n                        x-dcl-list-type: list\n                        items:\n                          type: string\n                          x-dcl-go-type: string\n                      set:\n                        type: object\n                        additionalProperties:\n                          type: string\n                        x-dcl-go-name: Set\n                        description: Completely overwrite/replace the headers with\n                          given map where key is the name of the header, value is\n                          the value of the header.\n                  requestMirrorPolicy:\n                    type: object\n                    x-dcl-go-name: RequestMirrorPolicy\n                    x-dcl-go-type: HttpRouteRulesActionRequestMirrorPolicy\n                    description: Specifies the policy on how requests intended for\n                      the routes destination are shadowed to a separate mirrored destination.\n                      Proxy will not wait for the shadow destination to respond before\n                      returning the response. Prior to sending traffic to the shadow\n                      service, the host/authority header is suffixed with -shadow.\n                    properties:\n                      destination:\n                        type: object\n                        x-dcl-go-name: Destination\n                        x-dcl-go-type: HttpRouteRulesActionRequestMirrorPolicyDestination\n                        description: The destination the requests will be mirrored\n                          to. The weight of the destination will be ignored.\n                        properties:\n                          serviceName:\n                            type: string\n                            x-dcl-go-name: ServiceName\n                            description: The URL of a BackendService to route traffic\n                              to.\n                          weight:\n                            type: integer\n                            format: int64\n                            x-dcl-go-name: Weight\n                            description: 'Specifies the proportion of requests forwarded\n                              to the backend referenced by the serviceName field.\n                              This is computed as: weight/Sum(weights in this destination\n                              list). For non-zero values, there may be some epsilon\n                              from the exact proportion defined here depending on\n                              the precision an implementation supports. If only one\n                              serviceName is specified and it has a weight greater\n                              than 0, 100% of the traffic is forwarded to that backend.\n                              If weights are specified for any one service name, they\n                              need to be specified for all of them. If weights are\n                              unspecified for all services, then, traffic is distributed\n                              in equal proportions to all of them.'\n                  responseHeaderModifier:\n                    type: object\n                    x-dcl-go-name: ResponseHeaderModifier\n                    x-dcl-go-type: HttpRouteRulesActionResponseHeaderModifier\n                    description: The specification for modifying the headers of a\n                      response prior to sending the response back to the client.\n                    properties:\n                      add:\n                        type: object\n                        additionalProperties:\n                          type: string\n                        x-dcl-go-name: Add\n                        description: Add the headers with given map where key is the\n                          name of the header, value is the value of the header.\n                      remove:\n                        type: array\n                        x-dcl-go-name: Remove\n                        description: Remove headers (matching by header names) specified\n                          in the list.\n                        x-dcl-send-empty: true\n                        x-dcl-list-type: list\n                        items:\n                          type: string\n                          x-dcl-go-type: string\n                      set:\n                        type: object\n                        additionalProperties:\n                          type: string\n                        x-dcl-go-name: Set\n                        description: Completely overwrite/replace the headers with\n                          given map where key is the name of the header, value is\n                          the value of the header.\n                  retryPolicy:\n                    type: object\n                    x-dcl-go-name: RetryPolicy\n                    x-dcl-go-type: HttpRouteRulesActionRetryPolicy\n                    description: Specifies the retry policy associated with this route.\n                    properties:\n                      numRetries:\n                        type: integer\n                        format: int64\n                        x-dcl-go-name: NumRetries\n                        description: Specifies the allowed number of retries. This\n                          number must be > 0. If not specified, default to 1.\n                      perTryTimeout:\n                        type: string\n                        x-dcl-go-name: PerTryTimeout\n                        description: Specifies a non-zero timeout per retry attempt.\n                      retryConditions:\n                        type: array\n                        x-dcl-go-name: RetryConditions\n                        description: 'Specifies one or more conditions when this retry\n                          policy applies. Valid values are: 5xx: Proxy will attempt\n                          a retry if the destination service responds with any 5xx\n                          response code, of if the destination service does not respond\n                          at all, example: disconnect, reset, read timeout, connection\n                          failure and refused streams. gateway-error: Similar to 5xx,\n                          but only applies to response codes 502, 503, 504. reset:\n                          Proxy will attempt a retry if the destination service does\n                          not respond at all (disconnect/reset/read timeout) connect-failure:\n                          Proxy will retry on failures connecting to destination for\n                          example due to connection timeouts. retriable-4xx: Proxy\n                          will retry fro retriable 4xx response codes. Currently the\n                          only retriable error supported is 409. refused-stream: Proxy\n                          will retry if the destination resets the stream with a REFUSED_STREAM\n                          error code. This reset type indicates that it is safe to\n                          retry.'\n                        x-dcl-send-empty: true\n                        x-dcl-list-type: list\n                        items:\n                          type: string\n                          x-dcl-go-type: string\n                  timeout:\n                    type: string\n                    x-dcl-go-name: Timeout\n                    description: Specifies the timeout for selected route. Timeout\n                      is computed from the time the request has been fully processed\n                      (i.e. end of stream) up until the response has been completely\n                      processed. Timeout includes all retries.\n                  urlRewrite:\n                    type: object\n                    x-dcl-go-name: UrlRewrite\n                    x-dcl-go-type: HttpRouteRulesActionUrlRewrite\n                    description: The specification for rewrite URL before forwarding\n                      requests to the destination.\n                    properties:\n                      hostRewrite:\n                        type: string\n                        x-dcl-go-name: HostRewrite\n                        description: Prior to forwarding the request to the selected\n                          destination, the requests host header is replaced by this\n                          value.\n                      pathPrefixRewrite:\n                        type: string\n                        x-dcl-go-name: PathPrefixRewrite\n                        description: Prior to forwarding the request to the selected\n                          destination, the matching portion of the requests path is\n                          replaced by this value.\n              matches:\n                type: array\n                x-dcl-go-name: Matches\n                description: A list of matches define conditions used for matching\n                  the rule against incoming HTTP requests. Each match is independent,\n                  i.e. this rule will be matched if ANY one of the matches is satisfied.\n                x-dcl-send-empty: true\n                x-dcl-list-type: list\n                items:\n                  type: object\n                  x-dcl-go-type: HttpRouteRulesMatches\n                  properties:\n                    fullPathMatch:\n                      type: string\n                      x-dcl-go-name: FullPathMatch\n                      description: The HTTP request path value should exactly match\n                        this value. Only one of full_path_match, prefix_match, or\n                        regex_match should be used.\n                    headers:\n                      type: array\n                      x-dcl-go-name: Headers\n                      description: Specifies a list of HTTP request headers to match\n                        against. ALL of the supplied headers must be matched.\n                      x-dcl-send-empty: true\n                      x-dcl-list-type: list\n                      items:\n                        type: object\n                        x-dcl-go-type: HttpRouteRulesMatchesHeaders\n                        properties:\n                          exactMatch:\n                            type: string\n                            x-dcl-go-name: ExactMatch\n                            description: The value of the header should match exactly\n                              the content of exact_match.\n                          header:\n                            type: string\n                            x-dcl-go-name: Header\n                            description: The name of the HTTP header to match against.\n                          invertMatch:\n                            type: boolean\n                            x-dcl-go-name: InvertMatch\n                            description: If specified, the match result will be inverted\n                              before checking. Default value is set to false.\n                          prefixMatch:\n                            type: string\n                            x-dcl-go-name: PrefixMatch\n                            description: The value of the header must start with the\n                              contents of prefix_match.\n                          presentMatch:\n                            type: boolean\n                            x-dcl-go-name: PresentMatch\n                            description: A header with header_name must exist. The\n                              match takes place whether or not the header has a value.\n                          rangeMatch:\n                            type: object\n                            x-dcl-go-name: RangeMatch\n                            x-dcl-go-type: HttpRouteRulesMatchesHeadersRangeMatch\n                            description: If specified, the rule will match if the\n                              request header value is within the range.\n                            properties:\n                              end:\n                                type: integer\n                                format: int64\n                                x-dcl-go-name: End\n                                description: End of the range (exclusive)\n                              start:\n                                type: integer\n                                format: int64\n                                x-dcl-go-name: Start\n                                description: Start of the range (inclusive)\n                          regexMatch:\n                            type: string\n                            x-dcl-go-name: RegexMatch\n                            description: 'The value of the header must match the regular\n                              expression specified in regex_match. For regular expression\n                              grammar, please see: https://github.com/google/re2/wiki/Syntax'\n                          suffixMatch:\n                            type: string\n                            x-dcl-go-name: SuffixMatch\n                            description: The value of the header must end with the\n                              contents of suffix_match.\n                    ignoreCase:\n                      type: boolean\n                      x-dcl-go-name: IgnoreCase\n                      description: Specifies if prefix_match and full_path_match matches\n                        are case sensitive. The default value is false.\n                    prefixMatch:\n                      type: string\n                      x-dcl-go-name: PrefixMatch\n                      description: The HTTP request path value must begin with specified\n                        prefix_match. prefix_match must begin with a /. Only one of\n                        full_path_match, prefix_match, or regex_match should be used.\n                    queryParameters:\n                      type: array\n                      x-dcl-go-name: QueryParameters\n                      description: Specifies a list of query parameters to match against.\n                        ALL of the query parameters must be matched.\n                      x-dcl-send-empty: true\n                      x-dcl-list-type: list\n                      items:\n                        type: object\n                        x-dcl-go-type: HttpRouteRulesMatchesQueryParameters\n                        properties:\n                          exactMatch:\n                            type: string\n                            x-dcl-go-name: ExactMatch\n                            description: The value of the query parameter must exactly\n                              match the contents of exact_match. Only one of exact_match,\n                              regex_match, or present_match must be set.\n                          presentMatch:\n                            type: boolean\n                            x-dcl-go-name: PresentMatch\n                            description: Specifies that the QueryParameterMatcher\n                              matches if request contains query parameter, irrespective\n                              of whether the parameter has a value or not. Only one\n                              of exact_match, regex_match, or present_match must be\n                              set.\n                          queryParameter:\n                            type: string\n                            x-dcl-go-name: QueryParameter\n                            description: The name of the query parameter to match.\n                          regexMatch:\n                            type: string\n                            x-dcl-go-name: RegexMatch\n                            description: The value of the query parameter must match\n                              the regular expression specified by regex_match. For\n                              regular expression grammar, please see https://github.com/google/re2/wiki/Syntax\n                              Only one of exact_match, regex_match, or present_match\n                              must be set.\n                    regexMatch:\n                      type: string\n                      x-dcl-go-name: RegexMatch\n                      description: The HTTP request path value must satisfy the regular\n                        expression specified by regex_match after removing any query\n                        parameters and anchor supplied with the original URL. For\n                        regular expression grammar, please see https://github.com/google/re2/wiki/Syntax\n                        Only one of full_path_match, prefix_match, or regex_match\n                        should be used.\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The timestamp when the resource was updated.\n          x-kubernetes-immutable: true\n")

// 36291 bytes
// MD5: d1abb18177441b5110dd286d2003681f
