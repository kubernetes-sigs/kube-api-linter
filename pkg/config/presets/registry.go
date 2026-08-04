/*
Copyright 2025 The Kubernetes Authors.

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
package presets

import (
	"sort"
	"sync"
)

// Preset is a list of linter names that are enabled by default when this preset is active.
type Preset []string

// Registry stores named presets.
type Registry interface {
	Register(name string, preset Preset)
	Get(name string) (Preset, bool)
	All() []string
}

//nolint:gochecknoglobals
var defaultRegistry = NewRegistry()

// DefaultRegistry returns the global preset registry.
func DefaultRegistry() Registry {
	return defaultRegistry
}

// NewRegistry returns a new empty preset registry.
func NewRegistry() Registry {
	return &presetRegistry{
		presets: make(map[string]Preset),
	}
}

type presetRegistry struct {
	lock    sync.RWMutex
	presets map[string]Preset
}

// Register adds a preset to the registry.
func (r *presetRegistry) Register(name string, preset Preset) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.presets[name] = preset
}

// Get returns the preset for the given name.
func (r *presetRegistry) Get(name string) (Preset, bool) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	p, ok := r.presets[name]

	return p, ok
}

// All returns all registered preset names in sorted order.
func (r *presetRegistry) All() []string {
	r.lock.RLock()
	defer r.lock.RUnlock()

	names := make([]string, 0, len(r.presets))
	for name := range r.presets {
		names = append(names, name)
	}

	sort.Strings(names)

	return names
}
