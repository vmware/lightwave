/*
 * Copyright © 2018 VMware, Inc.  All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the “License”); you may not
 * use this file except in compliance with the License.  You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an “AS IS” BASIS, without
 * warranties or conditions of any kind, EITHER EXPRESS OR IMPLIED.  See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

DWORD
LwCAPluginLoad(
   PLWCA_DB_FUNCTION_TABLE pFt
   );

VOID
LwCAPluginUnload(
   VOID
   );

DWORD
LwCADbTestPluginInitialize(
    PCSTR            pcszPluginConfigPath,
    PLWCA_DB_HANDLE  *ppHandle
    );

DWORD
LwCADbTestPluginAddCA(
    PLWCA_DB_HANDLE         pHandle,
    PCSTR                   pcszCAId,
    PLWCA_DB_CA_DATA        pCAData,
    PCSTR                   pcszParentCAId
    );

DWORD
LwCADbTestPluginAddCertData(
    PLWCA_DB_HANDLE         pHandle,
    PCSTR                   pcszCAId,
    PLWCA_DB_CERT_DATA      pCertData
    );

DWORD
LwCADbTestPluginCheckCA(
    PLWCA_DB_HANDLE        pHandle,
    PCSTR                  pcszCAId,
    PBOOLEAN               pbExists
    );

DWORD
LwCADbTestPluginCheckCertData(
    PLWCA_DB_HANDLE         pHandle,
    PCSTR                   pcszCAId,
    PCSTR                   pcszSerialNumber,
    PBOOLEAN                pbExists
    );

DWORD
LwCADbTestPluginGetCA(
    PLWCA_DB_HANDLE         pHandle,
    PCSTR                   pcszCAId,
    PLWCA_DB_CA_DATA        *ppCAData
    );

DWORD
LwCADbTestPluginGetCACertificates(
    PLWCA_DB_HANDLE              pHandle,
    PCSTR                        pcszCAId,
    PLWCA_CERTIFICATE_ARRAY      *ppCertArray
    );

DWORD
LwCADbTestPluginGetCertData(
    PLWCA_DB_HANDLE             pHandle,
    PCSTR                       pcszCAId,
    PLWCA_DB_CERT_DATA_ARRAY    *ppCertDataArray
    );

DWORD
LwCADbTestPluginGetCACRLNumber(
    PLWCA_DB_HANDLE             pHandle,
    PCSTR                       pcszCAId,
    PSTR                        *ppszCRLNumber
    );

DWORD
LwCADbTestPluginGetParentCAId(
    PLWCA_DB_HANDLE             pHandle,
    PCSTR                       pcszCAId,
    PSTR                        *ppszParentCAId
    );

DWORD
LwCADbTestPluginGetCAStatus(
    PLWCA_DB_HANDLE             pHandle,
    PCSTR                       pcszCAId,
    PLWCA_CA_STATUS             pStatus
    );

DWORD
LwCADbTestPluginGetCAAuthBlob(
    PLWCA_DB_HANDLE             pHandle,
    PCSTR                       pcszCAId,
    PSTR                        *ppszAuthBlob
    );

DWORD
LwCADbTestPluginUpdateCA(
    PLWCA_DB_HANDLE         pHandle,
    PCSTR                   pcszCAId,
    PLWCA_DB_CA_DATA        pCAData
    );


DWORD
LwCADbTestPluginUpdateCAStatus(
    PLWCA_DB_HANDLE         pHandle,
    PCSTR                   pcszCAId,
    LWCA_CA_STATUS          status
    );

DWORD
LwCADbTestPluginUpdateCertData(
    PLWCA_DB_HANDLE         pHandle,
    PCSTR                   pcszCAId,
    PLWCA_DB_CERT_DATA      pCertData
    );

DWORD
LwCADbTestPluginUpdateCACRLNumber(
    PLWCA_DB_HANDLE             pHandle,
    PCSTR                       pcszCAId,
    PCSTR                       pcszCRLNumber
    );

VOID
LwCADbTestPluginFreeCAData(
    PLWCA_DB_CA_DATA pCAData
    );

VOID
LwCADbTestPluginFreeCertDataArray(
    PLWCA_DB_CERT_DATA_ARRAY pCertDataArray
    );

VOID
LwCADbTestPluginFreeCertificates(
    PLWCA_CERTIFICATE_ARRAY pCertArray
    );

VOID
LwCADbTestPluginFreeString(
    PSTR  pszString
    );

VOID
LwCADbTestPluginFreeHandle(
    PLWCA_DB_HANDLE pDbHandle
    );

DWORD
LwCADbTestPluginLockCA(
    PLWCA_DB_HANDLE pHandle,
    PCSTR           pcszCAId,
    PSTR            *ppszUuid
    );

DWORD
LwCADbTestPluginUnlockCA(
    PLWCA_DB_HANDLE pHandle,
    PCSTR           pcszCAId,
    PCSTR           pcszUuid
    );

DWORD
LwCADbTestPluginLockCert(
    PLWCA_DB_HANDLE pHandle,
    PCSTR           pcszCAId,
    PCSTR           pcszSerialNumber,
    PSTR            *ppszUuid
    );

DWORD
LwCADbTestPluginUnlockCert(
    PLWCA_DB_HANDLE pHandle,
    PCSTR           pcszCAId,
    PCSTR           pcszSerialNumber,
    PCSTR           pcszUuid
    );
