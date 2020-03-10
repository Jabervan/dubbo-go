/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package extension

import (
	"github.com/apache/dubbo-go/protocol/rest/rest_interface"
)

var (
	restClients = make(map[string]func(restOptions *rest_interface.RestOptions) rest_interface.RestClient, 8)
)

func SetRestClient(name string, fun func(restOptions *rest_interface.RestOptions) rest_interface.RestClient) {
	restClients[name] = fun
}

func GetNewRestClient(name string, restOptions *rest_interface.RestOptions) rest_interface.RestClient {
	if restClients[name] == nil {
		panic("restClient for " + name + " is not existing, make sure you have import the package.")
	}
	return restClients[name](restOptions)
}