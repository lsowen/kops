/*
Copyright 2016 The Kubernetes Authors.

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

// Code generated by ""fitask" -type=ManagedFile"; DO NOT EDIT

package fitasks

import (
	"encoding/json"

	"k8s.io/kops/upup/pkg/fi"
)

// ManagedFile

// JSON marshalling boilerplate
type realManagedFile ManagedFile

func (o *ManagedFile) UnmarshalJSON(data []byte) error {
	var jsonName string
	if err := json.Unmarshal(data, &jsonName); err == nil {
		o.Name = &jsonName
		return nil
	}

	var r realManagedFile
	if err := json.Unmarshal(data, &r); err != nil {
		return err
	}
	*o = ManagedFile(r)
	return nil
}

var _ fi.HasName = &ManagedFile{}

func (e *ManagedFile) GetName() *string {
	return e.Name
}

func (e *ManagedFile) SetName(name string) {
	e.Name = &name
}

func (e *ManagedFile) String() string {
	return fi.TaskAsString(e)
}