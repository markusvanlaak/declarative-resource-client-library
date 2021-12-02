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
// gen_go_data -package alpha -var YAML_instance_group_manager blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/alpha/instance_group_manager.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/alpha/instance_group_manager.yaml
var YAML_instance_group_manager = []byte("info:\n  title: Compute/InstanceGroupManager\n  description: The Compute InstanceGroupManager resource\n  x-dcl-struct-name: InstanceGroupManager\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a InstanceGroupManager\n    parameters:\n    - name: InstanceGroupManager\n      required: true\n      description: A full instance of a InstanceGroupManager\n  apply:\n    description: The function used to apply information about a InstanceGroupManager\n    parameters:\n    - name: InstanceGroupManager\n      required: true\n      description: A full instance of a InstanceGroupManager\n  delete:\n    description: The function used to delete a InstanceGroupManager\n    parameters:\n    - name: InstanceGroupManager\n      required: true\n      description: A full instance of a InstanceGroupManager\n  deleteAll:\n    description: The function used to delete all InstanceGroupManager\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many InstanceGroupManager\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    FixedOrPercent:\n      x-dcl-has-iam: false\n      type: object\n      x-dcl-go-name: TargetSize\n      x-dcl-go-type: InstanceGroupManagerFixedOrPercent\n      description: 'Specifies the intended number of instances to be created from\n        the `instanceTemplate`. The final number of instances created from the template\n        will be equal to: - If expressed as a fixed number, the minimum of either\n        `targetSize.fixed` or `instanceGroupManager.targetSize` is used. - if expressed\n        as a `percent`, the `targetSize` would be `(targetSize.percent/100 * InstanceGroupManager.targetSize)`\n        If there is a remainder, the number is rounded. If unset, this version will\n        update any remaining instances not updated by another `version`. Read [Starting\n        a canary update](/compute/docs/instance-groups/rolling-out-updates-to-managed-instance-groups#starting_a_canary_update)\n        for more information.'\n      properties:\n        calculated:\n          type: integer\n          format: int64\n          x-dcl-go-name: Calculated\n          readOnly: true\n          description: '[Output Only] Absolute value of VM instances calculated based\n            on the specific mode. - If the value is `fixed`, then the `calculated`\n            value is equal to the `fixed` value. - If the value is a `percent`, then\n            the `calculated` value is `percent`/100 * `targetSize`. For example, the\n            `calculated` value of a 80% of a managed instance group with 150 instances\n            would be (80/100 * 150) = 120 VM instances. If there is a remainder, the\n            number is rounded.'\n        fixed:\n          type: integer\n          format: int64\n          x-dcl-go-name: Fixed\n          description: Specifies a fixed number of VM instances. This must be a positive\n            integer.\n          x-dcl-send-empty: true\n        percent:\n          type: integer\n          format: int64\n          x-dcl-go-name: Percent\n          description: Specifies a percentage of instances between 0 to 100%, inclusive.\n            For example, specify `80` for 80%.\n          x-dcl-send-empty: true\n    InstanceGroupManager:\n      title: InstanceGroupManager\n      x-dcl-locations:\n      - zone\n      - region\n      x-dcl-parent-container: project\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - project\n      properties:\n        autoHealingPolicies:\n          type: array\n          x-dcl-go-name: AutoHealingPolicies\n          description: The autohealing policy for this managed instance group. You\n            can specify only one value.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: InstanceGroupManagerAutoHealingPolicies\n            properties:\n              healthCheck:\n                type: string\n                x-dcl-go-name: HealthCheck\n                description: The URL for the health check that signals autohealing.\n                x-dcl-references:\n                - resource: Compute/HealthCheck\n                  field: selfLink\n              initialDelaySec:\n                type: integer\n                format: int64\n                x-dcl-go-name: InitialDelaySec\n                description: The number of seconds that the managed instance group\n                  waits before it applies autohealing policies to new instances or\n                  recently recreated instances. This initial delay allows instances\n                  to initialize and run their startup scripts before the instance\n                  group determines that they are UNHEALTHY. This prevents the managed\n                  instance group from recreating its instances prematurely. This value\n                  must be from range [0, 3600].\n        baseInstanceName:\n          type: string\n          x-dcl-go-name: BaseInstanceName\n          description: The base instance name to use for instances in this group.\n            The value must be 1-58 characters long. Instances are named by appending\n            a hyphen and a random four-character string to the base instance name.\n            The base instance name must comply with [RFC1035](https://www.ietf.org/rfc/rfc1035.txt).\n        creationTimestamp:\n          type: string\n          x-dcl-go-name: CreationTimestamp\n          readOnly: true\n          description: The creation timestamp for this managed instance group in \\[RFC3339\\](https://www.ietf.org/rfc/rfc3339.txt)\n            text format.\n          x-kubernetes-immutable: true\n        currentActions:\n          type: object\n          x-dcl-go-name: CurrentActions\n          x-dcl-go-type: InstanceGroupManagerCurrentActions\n          readOnly: true\n          description: '[Output Only] The list of instance actions and the number\n            of instances in this managed instance group that are scheduled for each\n            of those actions.'\n          x-kubernetes-immutable: true\n          properties:\n            abandoning:\n              type: integer\n              format: int64\n              x-dcl-go-name: Abandoning\n              readOnly: true\n              description: '[Output Only] The total number of instances in the managed\n                instance group that are scheduled to be abandoned. Abandoning an instance\n                removes it from the managed instance group without deleting it.'\n              x-kubernetes-immutable: true\n            creating:\n              type: integer\n              format: int64\n              x-dcl-go-name: Creating\n              readOnly: true\n              description: '[Output Only] The number of instances in the managed instance\n                group that are scheduled to be created or are currently being created.\n                If the group fails to create any of these instances, it tries again\n                until it creates the instance successfully. If you have disabled creation\n                retries, this field will not be populated; instead, the `creatingWithoutRetries`\n                field will be populated.'\n              x-kubernetes-immutable: true\n            creatingWithoutRetries:\n              type: integer\n              format: int64\n              x-dcl-go-name: CreatingWithoutRetries\n              readOnly: true\n              description: '[Output Only] The number of instances that the managed\n                instance group will attempt to create. The group attempts to create\n                each instance only once. If the group fails to create any of these\n                instances, it decreases the group''s `targetSize` value accordingly.'\n              x-kubernetes-immutable: true\n            deleting:\n              type: integer\n              format: int64\n              x-dcl-go-name: Deleting\n              readOnly: true\n              description: '[Output Only] The number of instances in the managed instance\n                group that are scheduled to be deleted or are currently being deleted.'\n              x-kubernetes-immutable: true\n            none:\n              type: integer\n              format: int64\n              x-dcl-go-name: None\n              readOnly: true\n              description: '[Output Only] The number of instances in the managed instance\n                group that are running and have no scheduled actions.'\n              x-kubernetes-immutable: true\n            recreating:\n              type: integer\n              format: int64\n              x-dcl-go-name: Recreating\n              readOnly: true\n              description: '[Output Only] The number of instances in the managed instance\n                group that are scheduled to be recreated or are currently being being\n                recreated. Recreating an instance deletes the existing root persistent\n                disk and creates a new disk from the image that is defined in the\n                instance template.'\n              x-kubernetes-immutable: true\n            refreshing:\n              type: integer\n              format: int64\n              x-dcl-go-name: Refreshing\n              readOnly: true\n              description: '[Output Only] The number of instances in the managed instance\n                group that are being reconfigured with properties that do not require\n                a restart or a recreate action. For example, setting or removing target\n                pools for the instance.'\n              x-kubernetes-immutable: true\n            restarting:\n              type: integer\n              format: int64\n              x-dcl-go-name: Restarting\n              readOnly: true\n              description: '[Output Only] The number of instances in the managed instance\n                group that are scheduled to be restarted or are currently being restarted.'\n              x-kubernetes-immutable: true\n            verifying:\n              type: integer\n              format: int64\n              x-dcl-go-name: Verifying\n              readOnly: true\n              description: '[Output Only] The number of instances in the managed instance\n                group that are being verified. See the `managedInstances[].currentAction`\n                property in the `listManagedInstances` method documentation.'\n              x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: An optional description of this resource.\n          x-kubernetes-immutable: true\n        distributionPolicy:\n          type: object\n          x-dcl-go-name: DistributionPolicy\n          x-dcl-go-type: InstanceGroupManagerDistributionPolicy\n          description: Policy specifying the intended distribution of managed instances\n            across zones in a regional managed instance group.\n          properties:\n            targetShape:\n              type: string\n              x-dcl-go-name: TargetShape\n              x-dcl-go-type: InstanceGroupManagerDistributionPolicyTargetShapeEnum\n              description: 'The distribution shape to which the group converges either\n                proactively or on resize events (depending on the value set in `updatePolicy.instanceRedistributionType`).\n                Possible values: TARGET_SHAPE_UNSPECIFIED, ANY, BALANCED, ANY_SINGLE_ZONE'\n              enum:\n              - TARGET_SHAPE_UNSPECIFIED\n              - ANY\n              - BALANCED\n              - ANY_SINGLE_ZONE\n            zones:\n              type: array\n              x-dcl-go-name: Zones\n              description: Zones where the regional managed instance group will create\n                and manage its instances.\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: InstanceGroupManagerDistributionPolicyZones\n                properties:\n                  zone:\n                    type: string\n                    x-dcl-go-name: Zone\n                    description: The URL of the [zone](/compute/docs/regions-zones/#available).\n                      The zone must exist in the region where the managed instance\n                      group is located.\n                    x-kubernetes-immutable: true\n        failoverAction:\n          type: string\n          x-dcl-go-name: FailoverAction\n          x-dcl-go-type: InstanceGroupManagerFailoverActionEnum\n          description: 'The action to perform in case of zone failure. Only one value\n            is supported, `NO_FAILOVER`. The default is `NO_FAILOVER`. Possible values:\n            UNKNOWN, NO_FAILOVER'\n          enum:\n          - UNKNOWN\n          - NO_FAILOVER\n        fingerprint:\n          type: string\n          x-dcl-go-name: Fingerprint\n          readOnly: true\n          description: Fingerprint of this resource. This field may be used in optimistic\n            locking. It will be ignored when inserting an InstanceGroupManager. An\n            up-to-date fingerprint must be provided in order to update the InstanceGroupManager,\n            otherwise the request will fail with error `412 conditionNotMet`. To see\n            the latest fingerprint, make a `get()` request to retrieve an InstanceGroupManager.\n          x-kubernetes-immutable: true\n        id:\n          type: integer\n          format: int64\n          x-dcl-go-name: Id\n          readOnly: true\n          description: '[Output Only] A unique identifier for this resource type.\n            The server generates this identifier.'\n          x-kubernetes-immutable: true\n        instanceGroup:\n          type: string\n          x-dcl-go-name: InstanceGroup\n          readOnly: true\n          description: '[Output Only] The URL of the Instance Group resource.'\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Compute/InstanceGroup\n            field: selfLink\n        instanceTemplate:\n          type: string\n          x-dcl-go-name: InstanceTemplate\n          description: The URL of the instance template that is specified for this\n            managed instance group. The group uses this template to create all new\n            instances in the managed instance group. The templates for existing instances\n            in the group do not change unless you run `recreateInstances`, run `applyUpdatesToInstances`,\n            or set the group's `updatePolicy.type` to `PROACTIVE`.\n          x-dcl-forward-slash-allowed: true\n          x-dcl-references:\n          - resource: Compute/InstanceTemplate\n            field: selfLink\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location of this resource.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The name of the managed instance group. The name must be 1-63\n            characters long, and comply with [RFC1035](https://www.ietf.org/rfc/rfc1035.txt).\n          x-kubernetes-immutable: true\n        namedPorts:\n          type: array\n          x-dcl-go-name: NamedPorts\n          description: Named ports configured for the Instance Groups complementary\n            to this Instance Group Manager.\n          x-kubernetes-immutable: true\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: InstanceGroupManagerNamedPorts\n            properties:\n              name:\n                type: string\n                x-dcl-go-name: Name\n                description: The name for this named port. The name must be 1-63 characters\n                  long, and comply with [RFC1035](https://www.ietf.org/rfc/rfc1035.txt).\n                x-kubernetes-immutable: true\n              port:\n                type: integer\n                format: int64\n                x-dcl-go-name: Port\n                description: The port number, which can be a value between 1 and 65535.\n                x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n        region:\n          type: string\n          x-dcl-go-name: Region\n          readOnly: true\n          description: '[Output Only] The URL of the [region](/compute/docs/regions-zones/#available)\n            where the managed instance group resides (for regional resources).'\n          x-kubernetes-immutable: true\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          readOnly: true\n          description: '[Output Only] The URL for this managed instance group. The\n            server defines this URL.'\n          x-kubernetes-immutable: true\n        serviceAccount:\n          type: string\n          x-dcl-go-name: ServiceAccount\n          description: 'The service account to be used as credentials for all operations\n            performed by the managed instance group on instances. The service accounts\n            needs all permissions required to create and delete instances. By default,\n            the service account: {projectNumber}@cloudservices.gserviceaccount.com\n            is used.'\n          x-dcl-references:\n          - resource: Iam/ServiceAccount\n            field: email\n        statefulPolicy:\n          type: object\n          x-dcl-go-name: StatefulPolicy\n          x-dcl-go-type: InstanceGroupManagerStatefulPolicy\n          description: Stateful configuration for this Instanced Group Manager\n          properties:\n            preservedState:\n              type: object\n              x-dcl-go-name: PreservedState\n              x-dcl-go-type: InstanceGroupManagerStatefulPolicyPreservedState\n              properties:\n                disks:\n                  type: object\n                  additionalProperties:\n                    type: object\n                    x-dcl-go-type: InstanceGroupManagerStatefulPolicyPreservedStateDisks\n                    properties:\n                      autoDelete:\n                        type: string\n                        x-dcl-go-name: AutoDelete\n                        x-dcl-go-type: InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum\n                        description: 'These stateful disks will never be deleted during\n                          autohealing, update or VM instance recreate operations.\n                          This flag is used to configure if the disk should be deleted\n                          after it is no longer used by the group, e.g. when the given\n                          instance or the whole group is deleted. Note: disks attached\n                          in READ_ONLY mode cannot be auto-deleted. Possible values:\n                          NEVER, ON_PERMANENT_INSTANCE_DELETION'\n                        enum:\n                        - NEVER\n                        - ON_PERMANENT_INSTANCE_DELETION\n                  x-dcl-go-name: Disks\n                  description: Disks created on the instances that will be preserved\n                    on instance delete, update, etc. This map is keyed with the device\n                    names of the disks.\n        status:\n          type: object\n          x-dcl-go-name: Status\n          x-dcl-go-type: InstanceGroupManagerStatus\n          readOnly: true\n          description: '[Output Only] The status of this managed instance group.'\n          properties:\n            autoscaler:\n              type: string\n              x-dcl-go-name: Autoscaler\n              readOnly: true\n              description: '[Output Only] The URL of the [Autoscaler](/compute/docs/autoscaler/)\n                that targets this instance group manager.'\n              x-kubernetes-immutable: true\n            isStable:\n              type: boolean\n              x-dcl-go-name: IsStable\n              readOnly: true\n              description: '[Output Only] A bit indicating whether the managed instance\n                group is in a stable state. A stable state means that: none of the\n                instances in the managed instance group is currently undergoing any\n                type of change (for example, creation, restart, or deletion); no future\n                changes are scheduled for instances in the managed instance group;\n                and the managed instance group itself is not being modified.'\n              x-kubernetes-immutable: true\n            stateful:\n              type: object\n              x-dcl-go-name: Stateful\n              x-dcl-go-type: InstanceGroupManagerStatusStateful\n              readOnly: true\n              description: '[Output Only] Stateful status of the given Instance Group\n                Manager.'\n              properties:\n                hasStatefulConfig:\n                  type: boolean\n                  x-dcl-go-name: HasStatefulConfig\n                  readOnly: true\n                  description: '[Output Only] A bit indicating whether the managed\n                    instance group has stateful configuration, that is, if you have\n                    configured any items in a stateful policy or in per-instance configs.\n                    The group might report that it has no stateful config even when\n                    there is still some preserved state on a managed instance, for\n                    example, if you have deleted all PICs but not yet applied those\n                    deletions.'\n                  x-kubernetes-immutable: true\n                isStateful:\n                  type: boolean\n                  x-dcl-go-name: IsStateful\n                  readOnly: true\n                  description: '[Output Only] A bit indicating whether the managed\n                    instance group has stateful configuration, that is, if you have\n                    configured any items in a stateful policy or in per-instance configs.\n                    The group might report that it has no stateful config even when\n                    there is still some preserved state on a managed instance, for\n                    example, if you have deleted all PICs but not yet applied those\n                    deletions. This field is deprecated in favor of has_stateful_config.'\n                  x-kubernetes-immutable: true\n                perInstanceConfigs:\n                  type: object\n                  x-dcl-go-name: PerInstanceConfigs\n                  x-dcl-go-type: InstanceGroupManagerStatusStatefulPerInstanceConfigs\n                  readOnly: true\n                  description: '[Output Only] Status of per-instance configs on the\n                    instance.'\n                  properties:\n                    allEffective:\n                      type: boolean\n                      x-dcl-go-name: AllEffective\n                      description: A bit indicating if all of the group's per-instance\n                        configs (listed in the output of a listPerInstanceConfigs\n                        API call) have status `EFFECTIVE` or there are no per-instance-configs.\n            versionTarget:\n              type: object\n              x-dcl-go-name: VersionTarget\n              x-dcl-go-type: InstanceGroupManagerStatusVersionTarget\n              readOnly: true\n              description: '[Output Only] A status of consistency of Instances'' versions\n                with their target version specified by `version` field on Instance\n                Group Manager.'\n              x-kubernetes-immutable: true\n              properties:\n                isReached:\n                  type: boolean\n                  x-dcl-go-name: IsReached\n                  readOnly: true\n                  description: '[Output Only] A bit indicating whether version target\n                    has been reached in this managed instance group, i.e. all instances\n                    are in their target version. Instances'' target version are specified\n                    by `version` field on Instance Group Manager.'\n                  x-kubernetes-immutable: true\n        targetPools:\n          type: array\n          x-dcl-go-name: TargetPools\n          description: The URLs for all TargetPool resources to which instances in\n            the `instanceGroup` field are added. The target pools automatically apply\n            to all of the instances in the managed instance group.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n            x-dcl-references:\n            - resource: Compute/TargetPool\n              field: selfLink\n        targetSize:\n          type: integer\n          format: int64\n          x-dcl-go-name: TargetSize\n          description: The target number of running instances for this managed instance\n            group. You can reduce this number by using the instanceGroupManager deleteInstances\n            or abandonInstances methods. Resizing the group also changes this number.\n        updatePolicy:\n          type: object\n          x-dcl-go-name: UpdatePolicy\n          x-dcl-go-type: InstanceGroupManagerUpdatePolicy\n          description: The update policy for this managed instance group.\n          properties:\n            instanceRedistributionType:\n              type: string\n              x-dcl-go-name: InstanceRedistributionType\n              x-dcl-go-type: InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum\n              description: 'The [instance redistribution policy](/compute/docs/instance-groups/regional-migs#proactive_instance_redistribution)\n                for regional managed instance groups. Valid values are: - `PROACTIVE`\n                (default): The group attempts to maintain an even distribution of\n                VM instances across zones in the region. - `NONE`: For non-autoscaled\n                groups, proactive redistribution is disabled.'\n              enum:\n              - NONE\n              - PROACTIVE\n            maxSurge:\n              $ref: '#/components/schemas/FixedOrPercent'\n              x-dcl-go-name: MaxSurge\n            maxUnavailable:\n              $ref: '#/components/schemas/FixedOrPercent'\n              x-dcl-go-name: MaxUnavailable\n            minReadySec:\n              type: integer\n              format: int64\n              x-dcl-go-name: MinReadySec\n              description: Minimum number of seconds to wait for after a newly created\n                instance becomes available. This value must be from range [0, 3600].\n            minimalAction:\n              type: string\n              x-dcl-go-name: MinimalAction\n              x-dcl-go-type: InstanceGroupManagerUpdatePolicyMinimalActionEnum\n              description: Minimal action to be taken on an instance. You can specify\n                either `RESTART` to restart existing instances or `REPLACE` to delete\n                and create new instances from the target template. If you specify\n                a `RESTART`, the Updater will attempt to perform that action only.\n                However, if the Updater determines that the minimal action you specify\n                is not enough to perform the update, it might perform a more disruptive\n                action.\n              enum:\n              - REPLACE\n              - RESTART\n              - REFRESH\n              - NONE\n            mostDisruptiveAllowedAction:\n              type: string\n              x-dcl-go-name: MostDisruptiveAllowedAction\n              x-dcl-go-type: InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum\n              description: Most disruptive action that is allowed to be taken on an\n                instance. You can specify either `NONE` to forbid any actions, `REFRESH`\n                to allow actions that do not need instance restart, `RESTART` to allow\n                actions that can be applied without instance replacing or `REPLACE`\n                to allow all possible actions. If the Updater determines that the\n                minimal update action needed is more disruptive than most disruptive\n                allowed action you specify it will not perform the update at all.\n              enum:\n              - REPLACE\n              - RESTART\n              - REFRESH\n              - NONE\n            replacementMethod:\n              type: string\n              x-dcl-go-name: ReplacementMethod\n              x-dcl-go-type: InstanceGroupManagerUpdatePolicyReplacementMethodEnum\n              description: 'What action should be used to replace instances. See minimal_action.REPLACE\n                Possible values: SUBSTITUTE, RECREATE'\n              enum:\n              - SUBSTITUTE\n              - RECREATE\n            type:\n              type: string\n              x-dcl-go-name: Type\n              x-dcl-go-type: InstanceGroupManagerUpdatePolicyTypeEnum\n              description: The type of update process. You can specify either `PROACTIVE`\n                so that the instance group manager proactively executes actions in\n                order to bring instances to their target versions or `OPPORTUNISTIC`\n                so that no action is proactively executed but the update will be performed\n                as part of other actions (for example, resizes or `recreateInstances`\n                calls).\n              enum:\n              - OPPORTUNISTIC\n              - PROACTIVE\n        versions:\n          type: array\n          x-dcl-go-name: Versions\n          description: Specifies the instance templates used by this managed instance\n            group to create instances. Each version is defined by an `instanceTemplate`\n            and a `name`. Every version can appear at most once per instance group.\n            This field overrides the top-level `instanceTemplate` field. Read more\n            about the [relationships between these fields](/compute/docs/instance-groups/rolling-out-updates-to-managed-instance-groups#relationship_between_versions_and_instancetemplate_properties_for_a_managed_instance_group).\n            Exactly one `version` must leave the `targetSize` field unset. That version\n            will be applied to all remaining instances. For more information, read\n            about [canary updates](/compute/docs/instance-groups/rolling-out-updates-to-managed-instance-groups#starting_a_canary_update).\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: InstanceGroupManagerVersions\n            properties:\n              instanceTemplate:\n                type: string\n                x-dcl-go-name: InstanceTemplate\n                description: The URL of the instance template that is specified for\n                  this managed instance group. The group uses this template to create\n                  new instances in the managed instance group until the `targetSize`\n                  for this version is reached. The templates for existing instances\n                  in the group do not change unless you run `recreateInstances`, run\n                  `applyUpdatesToInstances`, or set the group's `updatePolicy.type`\n                  to `PROACTIVE`; in those cases, existing instances are updated until\n                  the `targetSize` for this version is reached.\n                x-dcl-references:\n                - resource: Compute/InstanceTemplate\n                  field: selfLink\n              name:\n                type: string\n                x-dcl-go-name: Name\n                description: Name of the version. Unique among all versions in the\n                  scope of this managed instance group.\n              targetSize:\n                $ref: '#/components/schemas/FixedOrPercent'\n                x-dcl-go-name: TargetSize\n        zone:\n          type: string\n          x-dcl-go-name: Zone\n          readOnly: true\n          description: '[Output Only] The URL of a [zone](/compute/docs/regions-zones/#available)\n            where the managed instance group is located (for zonal resources).'\n          x-kubernetes-immutable: true\n")

// 32436 bytes
// MD5: c89d44cf44efb3e22be98d17bfe162e4
