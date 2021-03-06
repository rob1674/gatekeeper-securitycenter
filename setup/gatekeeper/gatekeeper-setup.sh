#!/usr/bin/env bash
#
# Copyright 2020 Google LLC
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

# This script installs OPA Gatekeeper in a GKE cluster.

set -ef -o pipefail

GATEKEEPER_VERSION=${GATEKEEPER_VERSION:-master}
SCRIPT_DIR=$(dirname "$0")

if [ ! -f "$SCRIPT_DIR/gatekeeper.yaml" ]; then
    curl -sSLo "$SCRIPT_DIR/gatekeeper.yaml" "https://raw.githubusercontent.com/open-policy-agent/gatekeeper/$GATEKEEPER_VERSION/deploy/gatekeeper.yaml"
fi

