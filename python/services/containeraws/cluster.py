# Copyright 2022 Google LLC. All Rights Reserved.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
from connector import channel
from google3.cloud.graphite.mmv2.services.google.container_aws import cluster_pb2
from google3.cloud.graphite.mmv2.services.google.container_aws import cluster_pb2_grpc

from typing import List


class Cluster(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        networking: dict = None,
        aws_region: str = None,
        control_plane: dict = None,
        authorization: dict = None,
        state: str = None,
        endpoint: str = None,
        uid: str = None,
        reconciling: bool = None,
        create_time: str = None,
        update_time: str = None,
        etag: str = None,
        annotations: dict = None,
        workload_identity_config: dict = None,
        project: str = None,
        location: str = None,
        fleet: dict = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.networking = networking
        self.aws_region = aws_region
        self.control_plane = control_plane
        self.authorization = authorization
        self.annotations = annotations
        self.project = project
        self.location = location
        self.fleet = fleet
        self.service_account_file = service_account_file

    def apply(self):
        stub = cluster_pb2_grpc.ContainerawsClusterServiceStub(channel.Channel())
        request = cluster_pb2.ApplyContainerawsClusterRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if ClusterNetworking.to_proto(self.networking):
            request.resource.networking.CopyFrom(
                ClusterNetworking.to_proto(self.networking)
            )
        else:
            request.resource.ClearField("networking")
        if Primitive.to_proto(self.aws_region):
            request.resource.aws_region = Primitive.to_proto(self.aws_region)

        if ClusterControlPlane.to_proto(self.control_plane):
            request.resource.control_plane.CopyFrom(
                ClusterControlPlane.to_proto(self.control_plane)
            )
        else:
            request.resource.ClearField("control_plane")
        if ClusterAuthorization.to_proto(self.authorization):
            request.resource.authorization.CopyFrom(
                ClusterAuthorization.to_proto(self.authorization)
            )
        else:
            request.resource.ClearField("authorization")
        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if ClusterFleet.to_proto(self.fleet):
            request.resource.fleet.CopyFrom(ClusterFleet.to_proto(self.fleet))
        else:
            request.resource.ClearField("fleet")
        request.service_account_file = self.service_account_file

        response = stub.ApplyContainerawsCluster(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.networking = ClusterNetworking.from_proto(response.networking)
        self.aws_region = Primitive.from_proto(response.aws_region)
        self.control_plane = ClusterControlPlane.from_proto(response.control_plane)
        self.authorization = ClusterAuthorization.from_proto(response.authorization)
        self.state = ClusterStateEnum.from_proto(response.state)
        self.endpoint = Primitive.from_proto(response.endpoint)
        self.uid = Primitive.from_proto(response.uid)
        self.reconciling = Primitive.from_proto(response.reconciling)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.etag = Primitive.from_proto(response.etag)
        self.annotations = Primitive.from_proto(response.annotations)
        self.workload_identity_config = ClusterWorkloadIdentityConfig.from_proto(
            response.workload_identity_config
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.fleet = ClusterFleet.from_proto(response.fleet)

    def delete(self):
        stub = cluster_pb2_grpc.ContainerawsClusterServiceStub(channel.Channel())
        request = cluster_pb2.DeleteContainerawsClusterRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if ClusterNetworking.to_proto(self.networking):
            request.resource.networking.CopyFrom(
                ClusterNetworking.to_proto(self.networking)
            )
        else:
            request.resource.ClearField("networking")
        if Primitive.to_proto(self.aws_region):
            request.resource.aws_region = Primitive.to_proto(self.aws_region)

        if ClusterControlPlane.to_proto(self.control_plane):
            request.resource.control_plane.CopyFrom(
                ClusterControlPlane.to_proto(self.control_plane)
            )
        else:
            request.resource.ClearField("control_plane")
        if ClusterAuthorization.to_proto(self.authorization):
            request.resource.authorization.CopyFrom(
                ClusterAuthorization.to_proto(self.authorization)
            )
        else:
            request.resource.ClearField("authorization")
        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if ClusterFleet.to_proto(self.fleet):
            request.resource.fleet.CopyFrom(ClusterFleet.to_proto(self.fleet))
        else:
            request.resource.ClearField("fleet")
        response = stub.DeleteContainerawsCluster(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = cluster_pb2_grpc.ContainerawsClusterServiceStub(channel.Channel())
        request = cluster_pb2.ListContainerawsClusterRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListContainerawsCluster(request).items

    def to_proto(self):
        resource = cluster_pb2.ContainerawsCluster()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if ClusterNetworking.to_proto(self.networking):
            resource.networking.CopyFrom(ClusterNetworking.to_proto(self.networking))
        else:
            resource.ClearField("networking")
        if Primitive.to_proto(self.aws_region):
            resource.aws_region = Primitive.to_proto(self.aws_region)
        if ClusterControlPlane.to_proto(self.control_plane):
            resource.control_plane.CopyFrom(
                ClusterControlPlane.to_proto(self.control_plane)
            )
        else:
            resource.ClearField("control_plane")
        if ClusterAuthorization.to_proto(self.authorization):
            resource.authorization.CopyFrom(
                ClusterAuthorization.to_proto(self.authorization)
            )
        else:
            resource.ClearField("authorization")
        if Primitive.to_proto(self.annotations):
            resource.annotations = Primitive.to_proto(self.annotations)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if ClusterFleet.to_proto(self.fleet):
            resource.fleet.CopyFrom(ClusterFleet.to_proto(self.fleet))
        else:
            resource.ClearField("fleet")
        return resource


class ClusterNetworking(object):
    def __init__(
        self,
        vpc_id: str = None,
        pod_address_cidr_blocks: list = None,
        service_address_cidr_blocks: list = None,
    ):
        self.vpc_id = vpc_id
        self.pod_address_cidr_blocks = pod_address_cidr_blocks
        self.service_address_cidr_blocks = service_address_cidr_blocks

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerawsClusterNetworking()
        if Primitive.to_proto(resource.vpc_id):
            res.vpc_id = Primitive.to_proto(resource.vpc_id)
        if Primitive.to_proto(resource.pod_address_cidr_blocks):
            res.pod_address_cidr_blocks.extend(
                Primitive.to_proto(resource.pod_address_cidr_blocks)
            )
        if Primitive.to_proto(resource.service_address_cidr_blocks):
            res.service_address_cidr_blocks.extend(
                Primitive.to_proto(resource.service_address_cidr_blocks)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNetworking(
            vpc_id=Primitive.from_proto(resource.vpc_id),
            pod_address_cidr_blocks=Primitive.from_proto(
                resource.pod_address_cidr_blocks
            ),
            service_address_cidr_blocks=Primitive.from_proto(
                resource.service_address_cidr_blocks
            ),
        )


class ClusterNetworkingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNetworking.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNetworking.from_proto(i) for i in resources]


class ClusterControlPlane(object):
    def __init__(
        self,
        version: str = None,
        instance_type: str = None,
        ssh_config: dict = None,
        subnet_ids: list = None,
        config_encryption: dict = None,
        security_group_ids: list = None,
        iam_instance_profile: str = None,
        root_volume: dict = None,
        main_volume: dict = None,
        database_encryption: dict = None,
        tags: dict = None,
        aws_services_authentication: dict = None,
        proxy_config: dict = None,
    ):
        self.version = version
        self.instance_type = instance_type
        self.ssh_config = ssh_config
        self.subnet_ids = subnet_ids
        self.config_encryption = config_encryption
        self.security_group_ids = security_group_ids
        self.iam_instance_profile = iam_instance_profile
        self.root_volume = root_volume
        self.main_volume = main_volume
        self.database_encryption = database_encryption
        self.tags = tags
        self.aws_services_authentication = aws_services_authentication
        self.proxy_config = proxy_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerawsClusterControlPlane()
        if Primitive.to_proto(resource.version):
            res.version = Primitive.to_proto(resource.version)
        if Primitive.to_proto(resource.instance_type):
            res.instance_type = Primitive.to_proto(resource.instance_type)
        if ClusterControlPlaneSshConfig.to_proto(resource.ssh_config):
            res.ssh_config.CopyFrom(
                ClusterControlPlaneSshConfig.to_proto(resource.ssh_config)
            )
        else:
            res.ClearField("ssh_config")
        if Primitive.to_proto(resource.subnet_ids):
            res.subnet_ids.extend(Primitive.to_proto(resource.subnet_ids))
        if ClusterControlPlaneConfigEncryption.to_proto(resource.config_encryption):
            res.config_encryption.CopyFrom(
                ClusterControlPlaneConfigEncryption.to_proto(resource.config_encryption)
            )
        else:
            res.ClearField("config_encryption")
        if Primitive.to_proto(resource.security_group_ids):
            res.security_group_ids.extend(
                Primitive.to_proto(resource.security_group_ids)
            )
        if Primitive.to_proto(resource.iam_instance_profile):
            res.iam_instance_profile = Primitive.to_proto(resource.iam_instance_profile)
        if ClusterControlPlaneRootVolume.to_proto(resource.root_volume):
            res.root_volume.CopyFrom(
                ClusterControlPlaneRootVolume.to_proto(resource.root_volume)
            )
        else:
            res.ClearField("root_volume")
        if ClusterControlPlaneMainVolume.to_proto(resource.main_volume):
            res.main_volume.CopyFrom(
                ClusterControlPlaneMainVolume.to_proto(resource.main_volume)
            )
        else:
            res.ClearField("main_volume")
        if ClusterControlPlaneDatabaseEncryption.to_proto(resource.database_encryption):
            res.database_encryption.CopyFrom(
                ClusterControlPlaneDatabaseEncryption.to_proto(
                    resource.database_encryption
                )
            )
        else:
            res.ClearField("database_encryption")
        if Primitive.to_proto(resource.tags):
            res.tags = Primitive.to_proto(resource.tags)
        if ClusterControlPlaneAwsServicesAuthentication.to_proto(
            resource.aws_services_authentication
        ):
            res.aws_services_authentication.CopyFrom(
                ClusterControlPlaneAwsServicesAuthentication.to_proto(
                    resource.aws_services_authentication
                )
            )
        else:
            res.ClearField("aws_services_authentication")
        if ClusterControlPlaneProxyConfig.to_proto(resource.proxy_config):
            res.proxy_config.CopyFrom(
                ClusterControlPlaneProxyConfig.to_proto(resource.proxy_config)
            )
        else:
            res.ClearField("proxy_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterControlPlane(
            version=Primitive.from_proto(resource.version),
            instance_type=Primitive.from_proto(resource.instance_type),
            ssh_config=ClusterControlPlaneSshConfig.from_proto(resource.ssh_config),
            subnet_ids=Primitive.from_proto(resource.subnet_ids),
            config_encryption=ClusterControlPlaneConfigEncryption.from_proto(
                resource.config_encryption
            ),
            security_group_ids=Primitive.from_proto(resource.security_group_ids),
            iam_instance_profile=Primitive.from_proto(resource.iam_instance_profile),
            root_volume=ClusterControlPlaneRootVolume.from_proto(resource.root_volume),
            main_volume=ClusterControlPlaneMainVolume.from_proto(resource.main_volume),
            database_encryption=ClusterControlPlaneDatabaseEncryption.from_proto(
                resource.database_encryption
            ),
            tags=Primitive.from_proto(resource.tags),
            aws_services_authentication=ClusterControlPlaneAwsServicesAuthentication.from_proto(
                resource.aws_services_authentication
            ),
            proxy_config=ClusterControlPlaneProxyConfig.from_proto(
                resource.proxy_config
            ),
        )


class ClusterControlPlaneArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterControlPlane.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterControlPlane.from_proto(i) for i in resources]


class ClusterControlPlaneSshConfig(object):
    def __init__(self, ec2_key_pair: str = None):
        self.ec2_key_pair = ec2_key_pair

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerawsClusterControlPlaneSshConfig()
        if Primitive.to_proto(resource.ec2_key_pair):
            res.ec2_key_pair = Primitive.to_proto(resource.ec2_key_pair)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterControlPlaneSshConfig(
            ec2_key_pair=Primitive.from_proto(resource.ec2_key_pair),
        )


class ClusterControlPlaneSshConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterControlPlaneSshConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterControlPlaneSshConfig.from_proto(i) for i in resources]


class ClusterControlPlaneConfigEncryption(object):
    def __init__(self, kms_key_arn: str = None):
        self.kms_key_arn = kms_key_arn

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerawsClusterControlPlaneConfigEncryption()
        if Primitive.to_proto(resource.kms_key_arn):
            res.kms_key_arn = Primitive.to_proto(resource.kms_key_arn)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterControlPlaneConfigEncryption(
            kms_key_arn=Primitive.from_proto(resource.kms_key_arn),
        )


class ClusterControlPlaneConfigEncryptionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterControlPlaneConfigEncryption.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterControlPlaneConfigEncryption.from_proto(i) for i in resources]


class ClusterControlPlaneRootVolume(object):
    def __init__(
        self,
        size_gib: int = None,
        volume_type: str = None,
        iops: int = None,
        kms_key_arn: str = None,
    ):
        self.size_gib = size_gib
        self.volume_type = volume_type
        self.iops = iops
        self.kms_key_arn = kms_key_arn

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerawsClusterControlPlaneRootVolume()
        if Primitive.to_proto(resource.size_gib):
            res.size_gib = Primitive.to_proto(resource.size_gib)
        if ClusterControlPlaneRootVolumeVolumeTypeEnum.to_proto(resource.volume_type):
            res.volume_type = ClusterControlPlaneRootVolumeVolumeTypeEnum.to_proto(
                resource.volume_type
            )
        if Primitive.to_proto(resource.iops):
            res.iops = Primitive.to_proto(resource.iops)
        if Primitive.to_proto(resource.kms_key_arn):
            res.kms_key_arn = Primitive.to_proto(resource.kms_key_arn)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterControlPlaneRootVolume(
            size_gib=Primitive.from_proto(resource.size_gib),
            volume_type=ClusterControlPlaneRootVolumeVolumeTypeEnum.from_proto(
                resource.volume_type
            ),
            iops=Primitive.from_proto(resource.iops),
            kms_key_arn=Primitive.from_proto(resource.kms_key_arn),
        )


class ClusterControlPlaneRootVolumeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterControlPlaneRootVolume.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterControlPlaneRootVolume.from_proto(i) for i in resources]


class ClusterControlPlaneMainVolume(object):
    def __init__(
        self,
        size_gib: int = None,
        volume_type: str = None,
        iops: int = None,
        kms_key_arn: str = None,
    ):
        self.size_gib = size_gib
        self.volume_type = volume_type
        self.iops = iops
        self.kms_key_arn = kms_key_arn

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerawsClusterControlPlaneMainVolume()
        if Primitive.to_proto(resource.size_gib):
            res.size_gib = Primitive.to_proto(resource.size_gib)
        if ClusterControlPlaneMainVolumeVolumeTypeEnum.to_proto(resource.volume_type):
            res.volume_type = ClusterControlPlaneMainVolumeVolumeTypeEnum.to_proto(
                resource.volume_type
            )
        if Primitive.to_proto(resource.iops):
            res.iops = Primitive.to_proto(resource.iops)
        if Primitive.to_proto(resource.kms_key_arn):
            res.kms_key_arn = Primitive.to_proto(resource.kms_key_arn)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterControlPlaneMainVolume(
            size_gib=Primitive.from_proto(resource.size_gib),
            volume_type=ClusterControlPlaneMainVolumeVolumeTypeEnum.from_proto(
                resource.volume_type
            ),
            iops=Primitive.from_proto(resource.iops),
            kms_key_arn=Primitive.from_proto(resource.kms_key_arn),
        )


class ClusterControlPlaneMainVolumeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterControlPlaneMainVolume.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterControlPlaneMainVolume.from_proto(i) for i in resources]


class ClusterControlPlaneDatabaseEncryption(object):
    def __init__(self, kms_key_arn: str = None):
        self.kms_key_arn = kms_key_arn

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerawsClusterControlPlaneDatabaseEncryption()
        if Primitive.to_proto(resource.kms_key_arn):
            res.kms_key_arn = Primitive.to_proto(resource.kms_key_arn)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterControlPlaneDatabaseEncryption(
            kms_key_arn=Primitive.from_proto(resource.kms_key_arn),
        )


class ClusterControlPlaneDatabaseEncryptionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterControlPlaneDatabaseEncryption.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterControlPlaneDatabaseEncryption.from_proto(i) for i in resources]


class ClusterControlPlaneAwsServicesAuthentication(object):
    def __init__(self, role_arn: str = None, role_session_name: str = None):
        self.role_arn = role_arn
        self.role_session_name = role_session_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerawsClusterControlPlaneAwsServicesAuthentication()
        if Primitive.to_proto(resource.role_arn):
            res.role_arn = Primitive.to_proto(resource.role_arn)
        if Primitive.to_proto(resource.role_session_name):
            res.role_session_name = Primitive.to_proto(resource.role_session_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterControlPlaneAwsServicesAuthentication(
            role_arn=Primitive.from_proto(resource.role_arn),
            role_session_name=Primitive.from_proto(resource.role_session_name),
        )


class ClusterControlPlaneAwsServicesAuthenticationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterControlPlaneAwsServicesAuthentication.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterControlPlaneAwsServicesAuthentication.from_proto(i)
            for i in resources
        ]


class ClusterControlPlaneProxyConfig(object):
    def __init__(self, secret_arn: str = None, secret_version: str = None):
        self.secret_arn = secret_arn
        self.secret_version = secret_version

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerawsClusterControlPlaneProxyConfig()
        if Primitive.to_proto(resource.secret_arn):
            res.secret_arn = Primitive.to_proto(resource.secret_arn)
        if Primitive.to_proto(resource.secret_version):
            res.secret_version = Primitive.to_proto(resource.secret_version)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterControlPlaneProxyConfig(
            secret_arn=Primitive.from_proto(resource.secret_arn),
            secret_version=Primitive.from_proto(resource.secret_version),
        )


class ClusterControlPlaneProxyConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterControlPlaneProxyConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterControlPlaneProxyConfig.from_proto(i) for i in resources]


class ClusterAuthorization(object):
    def __init__(self, admin_users: list = None):
        self.admin_users = admin_users

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerawsClusterAuthorization()
        if ClusterAuthorizationAdminUsersArray.to_proto(resource.admin_users):
            res.admin_users.extend(
                ClusterAuthorizationAdminUsersArray.to_proto(resource.admin_users)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAuthorization(
            admin_users=ClusterAuthorizationAdminUsersArray.from_proto(
                resource.admin_users
            ),
        )


class ClusterAuthorizationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterAuthorization.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterAuthorization.from_proto(i) for i in resources]


class ClusterAuthorizationAdminUsers(object):
    def __init__(self, username: str = None):
        self.username = username

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerawsClusterAuthorizationAdminUsers()
        if Primitive.to_proto(resource.username):
            res.username = Primitive.to_proto(resource.username)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAuthorizationAdminUsers(
            username=Primitive.from_proto(resource.username),
        )


class ClusterAuthorizationAdminUsersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterAuthorizationAdminUsers.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterAuthorizationAdminUsers.from_proto(i) for i in resources]


class ClusterWorkloadIdentityConfig(object):
    def __init__(
        self,
        issuer_uri: str = None,
        workload_pool: str = None,
        identity_provider: str = None,
    ):
        self.issuer_uri = issuer_uri
        self.workload_pool = workload_pool
        self.identity_provider = identity_provider

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerawsClusterWorkloadIdentityConfig()
        if Primitive.to_proto(resource.issuer_uri):
            res.issuer_uri = Primitive.to_proto(resource.issuer_uri)
        if Primitive.to_proto(resource.workload_pool):
            res.workload_pool = Primitive.to_proto(resource.workload_pool)
        if Primitive.to_proto(resource.identity_provider):
            res.identity_provider = Primitive.to_proto(resource.identity_provider)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterWorkloadIdentityConfig(
            issuer_uri=Primitive.from_proto(resource.issuer_uri),
            workload_pool=Primitive.from_proto(resource.workload_pool),
            identity_provider=Primitive.from_proto(resource.identity_provider),
        )


class ClusterWorkloadIdentityConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterWorkloadIdentityConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterWorkloadIdentityConfig.from_proto(i) for i in resources]


class ClusterFleet(object):
    def __init__(self, project: str = None, membership: str = None):
        self.project = project
        self.membership = membership

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerawsClusterFleet()
        if Primitive.to_proto(resource.project):
            res.project = Primitive.to_proto(resource.project)
        if Primitive.to_proto(resource.membership):
            res.membership = Primitive.to_proto(resource.membership)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterFleet(
            project=Primitive.from_proto(resource.project),
            membership=Primitive.from_proto(resource.membership),
        )


class ClusterFleetArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterFleet.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterFleet.from_proto(i) for i in resources]


class ClusterControlPlaneRootVolumeVolumeTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerawsClusterControlPlaneRootVolumeVolumeTypeEnum.Value(
            "ContainerawsClusterControlPlaneRootVolumeVolumeTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerawsClusterControlPlaneRootVolumeVolumeTypeEnum.Name(
            resource
        )[len("ContainerawsClusterControlPlaneRootVolumeVolumeTypeEnum") :]


class ClusterControlPlaneMainVolumeVolumeTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerawsClusterControlPlaneMainVolumeVolumeTypeEnum.Value(
            "ContainerawsClusterControlPlaneMainVolumeVolumeTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerawsClusterControlPlaneMainVolumeVolumeTypeEnum.Name(
            resource
        )[len("ContainerawsClusterControlPlaneMainVolumeVolumeTypeEnum") :]


class ClusterStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerawsClusterStateEnum.Value(
            "ContainerawsClusterStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerawsClusterStateEnum.Name(resource)[
            len("ContainerawsClusterStateEnum") :
        ]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s