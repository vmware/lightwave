/*
 *
 *  Copyright (c) 2012-2015 VMware, Inc.  All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License"); you may not
 *  use this file except in compliance with the License.  You may obtain a copy
 *  of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS, without
 *  warranties or conditions of any kind, EITHER EXPRESS OR IMPLIED.  See the
 *  License for the specific language governing permissions and limitations
 *  under the License.
 *
 */

package com.vmware.identity.idm.server;

import java.io.File;
import java.util.concurrent.ConcurrentHashMap;

import org.apache.commons.lang.Validate;

import com.rsa.authagent.authapi.AuthAgentException;
import com.rsa.authagent.authapi.AuthSessionFactory;
import com.vmware.identity.diagnostics.DiagnosticsLoggerFactory;
import com.vmware.identity.diagnostics.IDiagnosticsLogger;
import com.vmware.identity.idm.IDMException;
import com.vmware.identity.interop.IdmUtils;

public class AuthSessionFactoryCache {

    private ConcurrentHashMap<String, AuthSessionFactory> _factoryLookup; // lookup table for AuthSessionFactory keyed to tenant name.
    private static final IDiagnosticsLogger logger = DiagnosticsLoggerFactory.getLogger(AuthSessionFactoryCache.class);

    static String _configDir;

    AuthSessionFactoryCache() {
        _configDir = IdmUtils.getIdentityServicesConfigDir();
        _factoryLookup = new ConcurrentHashMap<String, AuthSessionFactory>();
    }

    /**
     * Return AuthSessionFactory for the tenant. Use cache if it exist.
     * Create one if none was found.
     * @param info
     * @return
     * @throws IDMException
     */
    public AuthSessionFactory getAuthnFactory(TenantInformation info) throws IDMException {

        Validate.notNull(info, "TenantInformation");
        String tenantName = info.getTenant().getName();
        Validate.notEmpty(tenantName);

        AuthSessionFactory factory = _factoryLookup.get(tenantName);

        if (factory == null) {
            logger.debug("Creating RSA AuthSessionFactory. tenant: "+tenantName);
            File configFolder = new File(_configDir, tenantName);
            File apiFile = new File(configFolder, "rsa_api.properties");
            String apiPropsFile = apiFile.getAbsolutePath();
            Validate.notEmpty(apiPropsFile, "rsa_api.properties");

            try {
                info.get_rsaConfigFilesLock().readLock().lock();
                factory = AuthSessionFactory.getInstance(apiPropsFile);
            } catch (AuthAgentException ex) {
                throw new IDMException(String.format(
                        "Failed to create AuthSessionFactory instance with %s", apiPropsFile));
            } finally {
                info.get_rsaConfigFilesLock().readLock().unlock();
            }

            _factoryLookup.putIfAbsent(tenantName, factory);  //this avoids synchronization
            logger.debug("Created RSA AuthSessionFactory for tenant: "+tenantName);
        }

        return factory;
    }

    public void removeFactory(String tenantName) throws AuthAgentException {
        Validate.notEmpty(tenantName);
        AuthSessionFactory api = _factoryLookup.remove(tenantName);  //remove is thread safe
        if (api != null) {
            api.shutdown();
        }

        logger.debug("Removed RSA AuthSessionFactory for tenant: "+tenantName);
    }
}
