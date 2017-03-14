/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dnsprovider

import "testing"

func TestStore(t *testing.T) {
	t.Run("RegisterAndLookup", testRegisterAndLookup)
	t.Run("LookupMultiple", testLookupMultiple)
}

// testRegisterAndLookup tests that a DNSProvider can be registered and looked up by name.
func testRegisterAndLookup(t *testing.T) {
	for _, tc := range []struct {
		title            string
		givenAndExpected map[string]DNSProvider
	}{
		{
			"registered provider is found by name",
			map[string]DNSProvider{
				"foo": NewInMemoryProvider(),
			},
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			for k, v := range tc.givenAndExpected {
				Register(k, v)
			}

			for k, v := range tc.givenAndExpected {
				if Lookup(k) != v {
					t.Errorf("expected %#v, got %#v", v, Lookup(k))
				}
			}
		})
	}
}

// testLookupMultiple tests that DNSProviders can be looked up by providing multiple names.
func testLookupMultiple(t *testing.T) {
	for _, tc := range []struct {
		title            string
		givenAndExpected map[string]DNSProvider
	}{
		{
			"multiple registered providers are found by names",
			map[string]DNSProvider{
				"foo": NewInMemoryProvider(),
				"bar": NewInMemoryProvider(),
			},
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			for k, v := range tc.givenAndExpected {
				Register(k, v)
			}

			names, sources := []string{}, []DNSProvider{}
			for k, v := range tc.givenAndExpected {
				names = append(names, k)
				sources = append(sources, v)
			}

			lookup := LookupMultiple(names...)

			for i := range names {
				if lookup[i] != sources[i] {
					t.Errorf("expected %#v, got %#v", sources[i], lookup[i])
				}
			}
		})
	}
}
