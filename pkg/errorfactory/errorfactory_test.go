// Copyright 2020 Orange SA
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
// limitations under the License.package apis

package errorfactory

import (
	"errors"
	"reflect"
	"testing"

	emperrors "emperror.dev/errors"
)

var errorTypes = []error{
	ResourceNotReady{},
	APIFailure{},
	VaultAPIFailure{},
	StatusUpdateError{},
	NodesUnreachable{},
	NodesNotReady{},
	NodesRequestError{},
	GracefulUpscaleFailed{},
	TooManyResources{},
	InternalError{},
	FatalReconcileError{},
	NifiClusterNotReady{},
	NifiClusterTaskRunning{},
}

func TestNew(t *testing.T) {
	for _, errType := range errorTypes {
		err := New(errType, errors.New("test-error"), "test-message")
		expected := "test-message: test-error"
		got := err.Error()
		if got != expected {
			t.Error("Expected:", expected, "got:", got)
		}
		if !emperrors.As(err, &errType) {
			t.Error("Expected:", reflect.TypeOf(errType), "got:", reflect.TypeOf(err))
		}
	}
}
