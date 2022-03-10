/*
Copyright 2022 The Kubernetes Authors.

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

package util

const (
	InstanceHandleSplitLen = 3
	InstanceURISplitLen    = 6
	ShareHandleSplitLen    = 4
	ShareURISplitLen       = 8

	MinMultishareInstanceSizeBytes    int64 = 1 * Tb
	MaxShareSizeBytes                 int64 = 1 * Tb
	MinShareSizeBytes                 int64 = 100 * Gb
	MaxSharesPerInstance                    = 10
	NewMultishareInstancePrefix             = "fs-"
	ParamMultishareInstanceScLabel          = "instanceStorageClassLabel"
	ParamMultishareInstanceScLabelKey       = "storage_gke_io_storage-class-id"
)

type OperationType int

const (
	InstanceCreate OperationType = iota
	InstanceDelete
	InstanceExpand
	InstanceShrink
	ShareCreate
	ShareDelete
	ShareExpand
	UnknownOp
)

type OperationStatus int

const (
	StatusSuccess OperationStatus = iota
	StatusRunning
	StatusFailed
	StatusUnknown
)