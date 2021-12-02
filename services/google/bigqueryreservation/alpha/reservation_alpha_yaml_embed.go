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
// gen_go_data -package alpha -var YAML_reservation blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/bigqueryreservation/alpha/reservation.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/bigqueryreservation/alpha/reservation.yaml
var YAML_reservation = []byte("info:\n  title: BigqueryReservation/Reservation\n  description: The BigqueryReservation Reservation resource\n  x-dcl-struct-name: Reservation\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Reservation\n    parameters:\n    - name: Reservation\n      required: true\n      description: A full instance of a Reservation\n  apply:\n    description: The function used to apply information about a Reservation\n    parameters:\n    - name: Reservation\n      required: true\n      description: A full instance of a Reservation\n  delete:\n    description: The function used to delete a Reservation\n    parameters:\n    - name: Reservation\n      required: true\n      description: A full instance of a Reservation\n  deleteAll:\n    description: The function used to delete all Reservation\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Reservation\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Reservation:\n      title: Reservation\n      x-dcl-id: projects/{{project}}/locations/{{location}}/reservations/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - project\n      - location\n      properties:\n        creationTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreationTime\n          readOnly: true\n          description: Output only. Creation time of the reservation.\n          x-kubernetes-immutable: true\n        ignoreIdleSlots:\n          type: boolean\n          x-dcl-go-name: IgnoreIdleSlots\n          description: If false, any query using this reservation will use idle slots\n            from other reservations within the same admin project. If true, a query\n            using this reservation will execute with the slot capacity specified above\n            at most.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        maxConcurrency:\n          type: integer\n          format: int64\n          x-dcl-go-name: MaxConcurrency\n          description: Maximum number of queries that are allowed to run concurrently\n            in this reservation. Default value is 0 which means that maximum concurrency\n            will be automatically set based on the reservation size.\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The resource name of the reservation.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        slotCapacity:\n          type: integer\n          format: int64\n          x-dcl-go-name: SlotCapacity\n          description: Minimum slots available to this reservation. A slot is a unit\n            of computational power in BigQuery, and serves as the unit of parallelism.\n            Queries using this reservation might use more slots during runtime if\n            ignore_idle_slots is set to false. If the new reservation's slot capacity\n            exceed the parent's slot capacity or if total slot capacity of the new\n            reservation and its siblings exceeds the parent's slot capacity, the request\n            will fail with `google.rpc.Code.RESOURCE_EXHAUSTED`.\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. Last update time of the reservation.\n          x-kubernetes-immutable: true\n")

// 4077 bytes
// MD5: ce3c9d09de3116d066d3c0aa7b1fa36d
