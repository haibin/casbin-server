// Copyright 2017 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import "github.com/casbin/casbin"

var base_dir string = "J:/github_repos/patron_rest/etc/patron/custom_policy/"
var policy_global_enable string = base_dir + "../enable.json"
var policy_global_restrict string = base_dir + "../policy.json"
var policy_tenant1_custom string = base_dir + "tenant1/custom-policy.json"
var policy_tenant2_custom string = base_dir + "tenant2/default-policy.json"

func enforceForFile(path string, sc SecurityContext) bool {
	e := casbin.NewEnforcer("authz_model.conf", path)
	return e.Enforce(sc.Tenant, sc.Sub, sc.Obj, sc.Act)
}

func enforce(sc SecurityContext) bool {
	if sc.Tenant == "admin" {
		return true
	}

	if sc.Tenant == "tenant1" {
		if !enforceForFile(policy_global_restrict, sc) {
			return false
		}

		if enforceForFile(policy_global_enable, sc) {
			return true
		}

		return enforceForFile(policy_tenant1_custom, sc)
	}

	if sc.Tenant == "tenant2" {
		if !enforceForFile(policy_global_restrict, sc) {
			return false
		}

		if enforceForFile(policy_global_enable, sc) {
			return true
		}

		return enforceForFile(policy_tenant2_custom, sc)
	}

	if sc.Tenant == "tenant3" {
		if !enforceForFile(policy_global_restrict, sc) {
			return false
		}

		if enforceForFile(policy_global_enable, sc) {
			return true
		}

		return false
	}

	return false
}