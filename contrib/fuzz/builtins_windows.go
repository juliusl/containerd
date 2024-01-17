/*
   Copyright The containerd Authors.

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

package fuzz

import (
	// Windows specific imports
	_ "github.com/containerd/containerd/v2/plugins/diff/lcow"
	_ "github.com/containerd/containerd/v2/plugins/diff/windows"
	_ "github.com/containerd/containerd/v2/plugins/snapshots/lcow"
	_ "github.com/containerd/containerd/v2/snapshots/windows"
)
