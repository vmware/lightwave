MKDIR=/bin/mkdir
RM=/bin/rm
RMDIR=/bin/rm -rf
CP=/bin/cp
MV=/bin/mv
TAR=/bin/tar
RPM=/bin/rpm
SED=/bin/sed

SRCROOT := $(realpath $(SRCROOT))

LIGHTWAVE_STAGE_DIR=$(SRCROOT)/stage
LIGHTWAVE_LOGDIR=/var/log/lightwave/

ARCH=x86_64

LW_SERVER_SRCROOT=$(SRCROOT)/lw-server
LW_SERVER_PKGDIR=$(LW_SERVER_SRCROOT)/rpmbuild/RPMS/$(ARCH)

LW_SERVER_MAJOR_VER=1
LW_SERVER_MINOR_VER=3
LW_SERVER_RELEASE_VER=1
LW_SERVER_PATCH_VER=2
LW_SERVER_VERSION=$(LW_SERVER_MAJOR_VER).$(LW_SERVER_MINOR_VER).$(LW_SERVER_RELEASE_VER)-$(LW_SERVER_PATCH_VER)
LW_SERVER_PKG_NAME=vmware-lightwave-server
LW_SERVER_RPM=$(LW_SERVER_PKG_NAME)-$(LW_SERVER_VERSION).$(ARCH).rpm

LW_CLIENTS_SRCROOT=$(SRCROOT)/lw-clients
LW_CLIENTS_PKGDIR=$(LW_CLIENTS_SRCROOT)/rpmbuild/RPMS/$(ARCH)

LW_CLIENTS_MAJOR_VER=1
LW_CLIENTS_MINOR_VER=3
LW_CLIENTS_RELEASE_VER=0
LW_CLIENTS_PATCH_VER=1
LW_CLIENTS_VERSION=$(LW_CLIENTS_MAJOR_VER).$(LW_CLIENTS_MINOR_VER).$(LW_CLIENTS_RELEASE_VER)-$(LW_CLIENTS_PATCH_VER)
LW_CLIENTS_PKG_NAME=vmware-lightwave-clients
LW_CLIENTS_RPM=$(LW_CLIENTS_PKG_NAME)-$(LW_CLIENTS_VERSION).$(ARCH).rpm

LW_RAFT_SRCROOT=$(SRCROOT)/lw-raft
LW_RAFT_PKGDIR=$(LW_RAFT_SRCROOT)/rpmbuild/RPMS/$(ARCH)

LW_RAFT_MAJOR_VER=1
LW_RAFT_MINOR_VER=3
LW_RAFT_RELEASE_VER=0
LW_RAFT_PATCH_VER=5
LW_RAFT_VERSION=$(LW_RAFT_MAJOR_VER).$(LW_RAFT_MINOR_VER).$(LW_RAFT_RELEASE_VER)-$(LW_RAFT_PATCH_VER)
LW_RAFT_PKG_NAME=vmware-objectstore
LW_RAFT_RPM=$(LW_RAFT_PKG_NAME)-$(LW_RAFT_VERSION).$(ARCH).rpm


LW_BUILD_SRCROOT=$(SRCROOT)/build

VMEVENT_SRCROOT=$(SRCROOT)/vmevent
VMEVENT_PKGDIR=$(VMEVENT_SRCROOT)/build/rpmbuild/RPMS/$(ARCH)

VMEVENT_MAJOR_VER=1
VMEVENT_MINOR_VER=3
VMEVENT_RELEASE_VER=0
VMEVENT_PATCH_VER=3
VMEVENT_VERSION=$(VMEVENT_MAJOR_VER).$(VMEVENT_MINOR_VER).$(VMEVENT_RELEASE_VER)-$(VMEVENT_PATCH_VER)
VMEVENT_CLIENT_DEVEL_PKG_NAME=vmware-event-devel
VMEVENT_CLIENT_DEVEL_RPM=$(VMEVENT_CLIENT_DEVEL_PKG_NAME)-$(VMEVENT_VERSION).$(ARCH).rpm
VMEVENT_RPMS=$(VMEVENT_CLIENT_DEVEL_RPM)

VMDIR_SRCROOT=$(SRCROOT)/vmdir
VMDIR_PKGDIR=$(VMDIR_SRCROOT)/build/rpmbuild/RPMS/$(ARCH)

VMDIR_MAJOR_VER=1
VMDIR_MINOR_VER=3
VMDIR_RELEASE_VER=0
VMDIR_PATCH_VER=5
VMDIR_VERSION=$(VMDIR_MAJOR_VER).$(VMDIR_MINOR_VER).$(VMDIR_RELEASE_VER)-$(VMDIR_PATCH_VER)
VMDIR_SERVER_PKG_NAME=vmware-directory
VMDIR_CLIENT_PKG_NAME=vmware-directory-client
VMDIR_CLIENT_DEVEL_PKG_NAME=vmware-directory-client-devel
VMDIR_DBGINFO_PKG_NAME=vmware-directory-debuginfo
VMDIR_SERVER_RPM=$(VMDIR_SERVER_PKG_NAME)-$(VMDIR_VERSION).$(ARCH).rpm
VMDIR_CLIENT_RPM=$(VMDIR_CLIENT_PKG_NAME)-$(VMDIR_VERSION).$(ARCH).rpm
VMDIR_CLIENT_DEVEL_RPM=$(VMDIR_CLIENT_DEVEL_PKG_NAME)-$(VMDIR_VERSION).$(ARCH).rpm
VMDIR_DBGINFO_RPM=$(VMDIR_DBGINFO_PKG_NAME)-$(VMDIR_VERSION).$(ARCH).rpm
VMDIR_RPMS=$(VMDIR_SERVER_RPM) $(VMDIR_CLIENT_RPM) $(VMDIR_CLIENT_DEVEL_RPM) $(VMDIR_DBGINFO_RPM)

LWRAFT_SRCROOT=$(SRCROOT)/lwraft
LWRAFT_PKGDIR=$(LWRAFT_SRCROOT)/build/rpmbuild/RPMS/$(ARCH)

LWRAFT_MAJOR_VER=1
LWRAFT_MINOR_VER=3
LWRAFT_RELEASE_VER=0
LWRAFT_PATCH_VER=5
LWRAFT_VERSION=$(LWRAFT_MAJOR_VER).$(LWRAFT_MINOR_VER).$(LWRAFT_RELEASE_VER)-$(LWRAFT_PATCH_VER)
LWRAFT_SERVER_PKG_NAME=vmware-post
LWRAFT_CLIENT_PKG_NAME=vmware-post-client
LWRAFT_CLIENT_DEVEL_PKG_NAME=vmware-post-client-devel
LWRAFT_DEBUGINFO_PKG_NAME=vmware-post-debuginfo
LWRAFT_SERVER_RPM=$(LWRAFT_SERVER_PKG_NAME)-$(LWRAFT_VERSION).$(ARCH).rpm
LWRAFT_CLIENT_RPM=$(LWRAFT_CLIENT_PKG_NAME)-$(LWRAFT_VERSION).$(ARCH).rpm
LWRAFT_CLIENT_DEVEL_RPM=$(LWRAFT_CLIENT_DEVEL_PKG_NAME)-$(LWRAFT_VERSION).$(ARCH).rpm
LWRAFT_DEBUGINFO_RPM=$(LWRAFT_DEBUGINFO_PKG_NAME)-$(LWRAFT_VERSION).$(ARCH).rpm
LWRAFT_RPMS=$(LWRAFT_SERVER_RPM) $(LWRAFT_CLIENT_RPM) $(LWRAFT_CLIENT_DEVEL_RPM) $(LWRAFT_DEBUGINFO_RPM)

VMDNS_SRCROOT=$(SRCROOT)/vmdns
VMDNS_PKGDIR=$(VMDNS_SRCROOT)/build/rpmbuild/RPMS/$(ARCH)

VMDNS_MAJOR_VER=1
VMDNS_MINOR_VER=3
VMDNS_RELEASE_VER=0
VMDNS_PATCH_VER=3
VMDNS_VERSION=$(VMDNS_MAJOR_VER).$(VMDNS_MINOR_VER).$(VMDNS_RELEASE_VER)-$(VMDNS_PATCH_VER)
VMDNS_SERVER_PKG_NAME=vmware-dns
VMDNS_CLIENT_PKG_NAME=vmware-dns-client
VMDNS_CLIENT_DEVEL_PKG_NAME=vmware-dns-client-devel
VMDNS_DEBUGINFO_PKG_NAME=vmware-dns-debuginfo
VMDNS_SERVER_RPM=$(VMDNS_SERVER_PKG_NAME)-$(VMDNS_VERSION).$(ARCH).rpm
VMDNS_CLIENT_RPM=$(VMDNS_CLIENT_PKG_NAME)-$(VMDNS_VERSION).$(ARCH).rpm
VMDNS_CLIENT_DEVEL_RPM=$(VMDNS_CLIENT_DEVEL_PKG_NAME)-$(VMDNS_VERSION).$(ARCH).rpm
VMDNS_DEBUGINFO_RPM=$(VMDNS_DEBUGINFO_PKG_NAME)-$(VMDNS_VERSION).$(ARCH).rpm
VMDNS_RPMS=$(VMDNS_SERVER_RPM) $(VMDNS_CLIENT_RPM) $(VMDNS_CLIENT_DEVEL_RPM) $(VMDNS_DEBUGINFO_RPM)

VMAFD_SRCROOT=$(SRCROOT)/vmafd
VMAFD_PKGDIR=$(VMAFD_SRCROOT)/build/rpmbuild/RPMS/$(ARCH)

VMAFD_MAJOR_VER=1
VMAFD_MINOR_VER=3
VMAFD_RELEASE_VER=0
VMAFD_PATCH_VER=3
VMAFD_VERSION=$(VMAFD_MAJOR_VER).$(VMAFD_MINOR_VER).$(VMAFD_RELEASE_VER)-$(VMAFD_PATCH_VER)
VMAFD_SERVER_PKG_NAME=vmware-afd
VMAFD_CLIENT_PKG_NAME=vmware-afd-client
VMAFD_CLIENT_DEVEL_PKG_NAME=vmware-afd-client-devel
VMAFD_CLIENT_PYTHON_PKG_NAME=vmware-afd-client-python
VMAFD_DEBUGINFO_PKG_NAME=vmware-afd-debuginfo
VMAFD_SERVER_RPM=$(VMAFD_SERVER_PKG_NAME)-$(VMAFD_VERSION).$(ARCH).rpm
VMAFD_CLIENT_RPM=$(VMAFD_CLIENT_PKG_NAME)-$(VMAFD_VERSION).$(ARCH).rpm
VMAFD_CLIENT_PYTHON_RPM=$(VMAFD_CLIENT_PYTHON_PKG_NAME)-$(VMAFD_VERSION).$(ARCH).rpm
VMAFD_CLIENT_DEVEL_RPM=$(VMAFD_CLIENT_DEVEL_PKG_NAME)-$(VMAFD_VERSION).$(ARCH).rpm
VMAFD_DEBUGINFO_RPM=$(VMAFD_DEBUGINFO_PKG_NAME)-$(VMAFD_VERSION).$(ARCH).rpm
VMAFD_RPMS=$(VMAFD_SERVER_RPM) $(VMAFD_CLIENT_RPM) $(VMAFD_CLIENT_DEVEL_RPM) $(VMAFD_DEBUGINFO_RPM)
VMCA_SRCROOT=$(SRCROOT)/vmca
VMCA_PKGDIR=$(VMCA_SRCROOT)/build/rpmbuild/RPMS/$(ARCH)

VMCA_MAJOR_VER=1
VMCA_MINOR_VER=3
VMCA_RELEASE_VER=0
VMCA_PATCH_VER=3
VMCA_VERSION=$(VMCA_MAJOR_VER).$(VMCA_MINOR_VER).$(VMCA_RELEASE_VER)-$(VMCA_PATCH_VER)
VMCA_SERVER_PKG_NAME=vmware-ca
VMCA_CLIENT_PKG_NAME=vmware-ca-client
VMCA_CLIENT_DEVEL_PKG_NAME=vmware-ca-client-devel
VMCA_DEBUGINFO_PKG_NAME=vmware-ca-debuginfo
VMCA_SERVER_RPM=$(VMCA_SERVER_PKG_NAME)-$(VMCA_VERSION).$(ARCH).rpm
VMCA_CLIENT_RPM=$(VMCA_CLIENT_PKG_NAME)-$(VMCA_VERSION).$(ARCH).rpm
VMCA_CLIENT_DEVEL_RPM=$(VMCA_CLIENT_DEVEL_PKG_NAME)-$(VMCA_VERSION).$(ARCH).rpm
VMCA_DEBUGINFO_RPM=$(VMCA_DEBUGINFO_PKG_NAME)-$(VMCA_VERSION).$(ARCH).rpm
VMCA_RPMS=$(VMCA_SERVER_RPM) $(VMCA_CLIENT_RPM) $(VMCA_CLIENT_DEVEL_RPM) $(VMCA_DEBUGINFO_RPM)

VMSTS_SRCROOT=$(SRCROOT)/vmidentity
VMSTS_PKGDIR=$(VMSTS_SRCROOT)/build/rpmbuild/RPMS/$(ARCH)
VMSTS_REST=$(VMSTS_SRCROOT)/rest
VMSTS_REST_IDM_TARGET=$(VMSTS_REST)/idm/samples/target
VMSTS_REST_VMDIR_CLIENT_TARGET=$(VMSTS_REST)/vmdir/client/target
VMSTS_REST_VMDIR_COMMON_TARGET=$(VMSTS_REST)/vmdir/common/target
VMSTS_REST_VMDIR_SERVER_TARGET=$(VMSTS_REST)/vmdir/server/target
VMSTS_PROPERTY_FILES=$(VMSTS_SRCROOT)/websso/src/main/resources
VMSTS_DIAGNOSTICS_LIB=$(VMSTS_SRCROOT)/diagnostics/lib
VMSTS_LWUI_SRC_MAIN_RESOURCES=$(VMSTS_SRCROOT)/lightwaveui/src/main/resources

VMSTS_MAJOR_VER=1
VMSTS_MINOR_VER=3
VMSTS_RELEASE_VER=0
VMSTS_PATCH_VER=3
VMSTS_VERSION=$(VMSTS_MAJOR_VER).$(VMSTS_MINOR_VER).$(VMSTS_RELEASE_VER)-$(VMSTS_PATCH_VER)
VMSTS_SERVER_PKG_NAME=vmware-sts
VMSTS_CLIENT_PKG_NAME=vmware-sts-client
VMSTS_C_CLIENT_PKG_NAME=vmware-sts-c-client
VMSTS_DEBUGINFO_PKG_NAME=vmware-sts-debuginfo
VMSTS_SERVER_RPM=$(VMSTS_SERVER_PKG_NAME)-$(VMSTS_VERSION).$(ARCH).rpm
VMSTS_CLIENT_RPM=$(VMSTS_CLIENT_PKG_NAME)-$(VMSTS_VERSION).$(ARCH).rpm
VMSTS_C_CLIENT_RPM=$(VMSTS_C_CLIENT_PKG_NAME)-$(VMSTS_VERSION).$(ARCH).rpm
VMSTS_DEBUGINFO_RPM=$(VMSTS_DEBUGINFO_PKG_NAME)-$(VMSTS_VERSION).$(ARCH).rpm
VMSTS_RPMS=$(VMSTS_SERVER_RPM) $(VMSTS_CLIENT_RPM) $(VMSTS_DEBUGINFO_RPM)

CFG_SRCROOT=$(SRCROOT)/config
CFG_PKGDIR=$(CFG_SRCROOT)/build/rpmbuild/RPMS/$(ARCH)

CFG_MAJOR_VER=1
CFG_MINOR_VER=3
CFG_RELEASE_VER=0
CFG_PATCH_VER=3
CFG_VERSION=$(CFG_MAJOR_VER).$(CFG_MINOR_VER).$(CFG_RELEASE_VER)-$(CFG_PATCH_VER)
CFG_PKG_NAME=vmware-ic-config
CFG_DEBUGINFO_PKG_NAME=vmware-ic-config-debuginfo
CFG_RPM=$(CFG_PKG_NAME)-$(CFG_VERSION).$(ARCH).rpm
CFG_DEBUGINFO_RPM=$(CFG_DEBUGINFO_PKG_NAME)-$(CFG_VERSION).$(ARCH).rpm
CFG_RPMS=$(CFG_RPM) $(CFG_DEBUGINFO_RPM)

LIGHTWAVE_OVA_STAGE=$(LIGHTWAVE_STAGE_DIR)/lw-appliance
LIGHTWAVE_OVA=$(LIGHTWAVE_OVA_STAGE)/lightwave-$(LW_SERVER_VERSION).ova
LIGHTWAVE_SB_OVA=$(LIGHTWAVE_OVA_STAGE)/lightwave-sb.ova
PHOTON_OVA=$(LIGHTWAVE_STAGE_DIR)/lw-appliance/photon-ova-build/photon-ova-virtualbox.ova
APPLIANCE_BUILDER=$(SRCROOT)/appliance/build-lightwave-ova.sh

DOCKER_SRCROOT=$(SRCROOT)/support/docker
DOCKER_BUILDER=$(DOCKER_SRCROOT)/build-lightwave-container.sh
DOCKER_IMAGE_VERSION=1.3.0
DOCKER_IMAGE_TAG=vmware/lightwave-sts:$(DOCKER_IMAGE_VERSION)
DOCKER_IMAGE=$(LIGHTWAVE_STAGE_DIR)/vmware-lightwave-sts-$(DOCKER_IMAGE_VERSION).tar

DOCKER_CLIENT_SRCROOT=$(SRCROOT)/support/docker
DOCKER_CLIENT_BUILDER=$(DOCKER_SRCROOT)/build-lightwave-client-container.sh
DOCKER_CLIENT_IMAGE=$(LIGHTWAVE_STAGE_DIR)/vmware-lightwave-client.tar

JAVA_HOME?=/etc/alternatives/jre/../
COMMONS_DAEMON?=/usr/share/java/
ANT_HOME?=/var/opt/apache-ant
TOMCAT_HOME?=/var/opt/apache-tomcat
MAVEN_HOME?=/var/opt/apache-maven
