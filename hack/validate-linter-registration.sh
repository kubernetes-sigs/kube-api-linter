#! /bin/bash

# Copyright 2025 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

linters=$(go list -json ./pkg/analysis/... | jq -rs '.[] | select((.GoFiles | index("initializer.go")) and (.GoFiles | index("analyzer.go"))) | .ImportPath')
registered=$(go list -json ./pkg/registration | jq -r '.Imports | .[]')

if [[ "${linters[*]}" != "${registered[*]}" ]]; then
    echo "Linter registration mismatch detected:"
    echo "  - Lines starting with '-' exist in pkg/registration but not in pkg/analysis"
    echo "  - Lines starting with '+' exist in pkg/analysis but not in pkg/registration"
    echo ""
    diff -Naup --label "pkg/registration/doc.go" --label "pkg/analysis/**/analyzer.go" <(echo "${registered}") <(echo "${linters}")
fi

