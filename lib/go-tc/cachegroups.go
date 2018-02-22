package tc

/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// CacheGroupResponse ...
type CacheGroupsResponse struct {
	Response []CacheGroup `json:"response"`
}

// CacheGroup contains information about a given Cachegroup in Traffic Ops.
type CacheGroup struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	ShortName   string  `json:"shortName"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	ParentName  string  `json:"parentCachegroupName,omitempty"`
	Type        string  `json:"typeName,omitempty"`
	LastUpdated string  `json:"lastUpdated,omitempty"`
}
