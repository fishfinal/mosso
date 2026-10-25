// Copyright 2021 fishfinal <contact.fishfinal@gmail.com>;
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

package mosso

import (
	"fmt"
	"testing"
)

func TestDebugShowContentWithJSON(t *testing.T) {
	data := map[string]string{"username": "fishfinal", "email": "contact.fishfinal@gmail.com"}
	DebugShowContentWithJSON(data, WithSpecificOfVisual(true), WithSpecificOfBottomSpaceLine(1))
}

func TestDebugShowContentWithJSON_WithWriteFile(t *testing.T) {
	data := map[string]string{"username": "fishfinal", "email": "contact.fishfinal@gmail.com"}
	DebugShowContentWithJSON(data,
		WithSpecificOfVisual(true),
		WithSpecificOfBottomSpaceLine(1),
		WithSpecificOfWriteFile(true),
		WithSpecificOfFilename("hello.json"),
	)
}

func TestDebugContentWithJSON(t *testing.T) {

	data := map[string]string{"username": "fishfinal", "email": "contact.fishfinal@gmail.com"}
	json := DebugContentWithJSON(data, WithSpecificOfVisual(false), WithSpecificOfBottomSpaceLine(1))

	fmt.Println(json)
}

func TestDebugShowContentWithJSON_With(t *testing.T) {
	data := map[string]string{"username": "fishfinal", "email": "contact.fishfinal@gmail.com"}
	DebugShowContentWithJSON(data, WithInfo("用户信息"))
}
