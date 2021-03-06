/*
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
 */
package com.vmware.identity.saml.idm;

import com.vmware.identity.idm.client.CasIdmClient;
import com.vmware.identity.saml.PrincipalAttributesExtractor;
import com.vmware.identity.saml.PrincipalAttributesExtractorFactory;

public class IdmPrincipalAttributesExtractorFactory implements
   PrincipalAttributesExtractorFactory {

   private String idmHostName;

   public IdmPrincipalAttributesExtractorFactory(String idmHostName) {
      this.idmHostName = idmHostName;
   }

   @Override
   public PrincipalAttributesExtractor getPrincipalAttributesExtractor(
      String tenantName) {
      return new IdmPrincipalAttributesExtractor(tenantName, new CasIdmClient(
         idmHostName));
   }

}
