// Copyright 2019 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import "fmt"

// Resource of a resource.
type Resource interface {

	// Debugging ResourceID for the resource instance.
	ID() ResourceID
}

// ResourceID for the resource instance. This is allocated by the framework and passed here.
type ResourceID interface {
	fmt.Stringer
}
