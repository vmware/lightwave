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

/**
 * @author Yehia Zayour
 */
public enum StatusCode {
    // these names and numeric values are copied from HttpServletResponse
    CONTINUE(100),
    SWITCHING_PROTOCOLS(101),
    OK(200),
    CREATED(201),
    ACCEPTED(202),
    NON_AUTHORITATIVE_INFORMATION(203),
    NO_CONTENT(204),
    RESET_CONTENT(205),
    PARTIAL_CONTENT(206),
    MULTIPLE_CHOICES(300),
    MOVED_PERMANENTLY(301),
    //MOVED_TEMPORARILY(302), we don't want two enum values with the same underlying int value
    FOUND(302),
    SEE_OTHER(303),
    NOT_MODIFIED(304),
    USE_PROXY(305),
    TEMPORARY_REDIRECT(307),
    BAD_REQUEST(400),
    UNAUTHORIZED(401),
    PAYMENT_REQUIRED(402),
    FORBIDDEN(403),
    NOT_FOUND(404),
    METHOD_NOT_ALLOWED(405),
    NOT_ACCEPTABLE(406),
    PROXY_AUTHENTICATION_REQUIRED(407),
    REQUEST_TIMEOUT(408),
    CONFLICT(409),
    GONE(410),
    LENGTH_REQUIRED(411),
    PRECONDITION_FAILED(412),
    REQUEST_ENTITY_TOO_LARGE(413),
    REQUEST_URI_TOO_LONG(414),
    UNSUPPORTED_MEDIA_TYPE(415),
    REQUESTED_RANGE_NOT_SATISFIABLE(416),
    EXPECTATION_FAILED(417),
    INTERNAL_SERVER_ERROR(500),
    NOT_IMPLEMENTED(501),
    BAD_GATEWAY(502),
    SERVICE_UNAVAILABLE(503),
    GATEWAY_TIMEOUT(504),
    HTTP_VERSION_NOT_SUPPORTED(505);

    private static final Map<Integer, StatusCode> intToEnumMap = new HashMap<Integer, StatusCode>();
    static {
        for (StatusCode v : StatusCode.values()) {
            intToEnumMap.put(v.getValue(), v);
        }
    }

    private final int value;

    private StatusCode(int value) {
        this.value = value;
    }

    public int getValue() {
        return this.value;
    }

    public static StatusCode parse(int value) throws ParseException {
        StatusCode result = intToEnumMap.get(value);
        if (result == null) {
            throw new ParseException("unrecognized status code");
        }
        return result;
    }
}