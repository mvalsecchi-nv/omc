/*
Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.

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
package get

// Options holds the per-invocation configuration for a get call. Lifting
// these out of vars makes the get pipeline reentrant and is a precondition
// for any future concurrent producer/consumer in this command. Fields are
// populated either from cobra-bound flag values or from upstream defaults
// before Run is invoked.
type Options struct {
	RootPath          string
	Namespace         string
	Output            string
	LabelSelector     string
	NoHeaders         bool
	AllNamespaces     bool
	ShowLabels        bool
	Wide              bool
	ShowKind          bool
	ShowNamespace     bool
	SortBy            string
	SingleResource    bool
	ShowManagedFields bool
	GetArgs           map[string]map[string]struct{}
}

func newOptions() Options {
	return Options{
		GetArgs: make(map[string]map[string]struct{}),
	}
}
