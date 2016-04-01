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

package com.vmware.identity.openidconnect.common;

import java.util.HashMap;
import java.util.Map;

import org.apache.commons.lang3.Validate;

/**
 * @author Yehia Zayour
 */
public enum TokenClass {
    // tokens generated by the server
    ID_TOKEN("id_token"),
    ACCESS_TOKEN("access_token"),
    REFRESH_TOKEN("refresh_token"),

    // tokens generated by clients
    SOLUTION_USER_ASSERTION("solution_user_assertion"),
    CLIENT_ASSERTION("client_assertion");

    private static final Map<String, TokenClass> stringToEnumMap = new HashMap<String, TokenClass>();
    static {
        for (TokenClass v : TokenClass.values()) {
            stringToEnumMap.put(v.getValue(), v);
        }
    }

    private final String value;

    private TokenClass(String value) {
        this.value = value;
    }

    public String getValue() {
        return this.value;
    }

    public static TokenClass parse(String value) throws ParseException {
        Validate.notEmpty(value, "value");
        TokenClass result = stringToEnumMap.get(value);
        if (result == null) {
            throw new ParseException("invalid token_class value");
        }
        return result;
    }
}