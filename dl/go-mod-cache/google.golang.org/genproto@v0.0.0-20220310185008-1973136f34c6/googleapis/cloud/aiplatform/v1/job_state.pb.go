// Copyright 2022 Google LLC
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/cloud/aiplatform/v1/job_state.proto

package aiplatform

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Describes the state of a job.
type JobState int32

const (
	// The job state is unspecified.
	JobState_JOB_STATE_UNSPECIFIED JobState = 0
	// The job has been just created or resumed and processing has not yet begun.
	JobState_JOB_STATE_QUEUED JobState = 1
	// The service is preparing to run the job.
	JobState_JOB_STATE_PENDING JobState = 2
	// The job is in progress.
	JobState_JOB_STATE_RUNNING JobState = 3
	// The job completed successfully.
	JobState_JOB_STATE_SUCCEEDED JobState = 4
	// The job failed.
	JobState_JOB_STATE_FAILED JobState = 5
	// The job is being cancelled. From this state the job may only go to
	// either `JOB_STATE_SUCCEEDED`, `JOB_STATE_FAILED` or `JOB_STATE_CANCELLED`.
	JobState_JOB_STATE_CANCELLING JobState = 6
	// The job has been cancelled.
	JobState_JOB_STATE_CANCELLED JobState = 7
	// The job has been stopped, and can be resumed.
	JobState_JOB_STATE_PAUSED JobState = 8
	// The job has expired.
	JobState_JOB_STATE_EXPIRED JobState = 9
)

// Enum value maps for JobState.
var (
	JobState_name = map[int32]string{
		0: "JOB_STATE_UNSPECIFIED",
		1: "JOB_STATE_QUEUED",
		2: "JOB_STATE_PENDING",
		3: "JOB_STATE_RUNNING",
		4: "JOB_STATE_SUCCEEDED",
		5: "JOB_STATE_FAILED",
		6: "JOB_STATE_CANCELLING",
		7: "JOB_STATE_CANCELLED",
		8: "JOB_STATE_PAUSED",
		9: "JOB_STATE_EXPIRED",
	}
	JobState_value = map[string]int32{
		"JOB_STATE_UNSPECIFIED": 0,
		"JOB_STATE_QUEUED":      1,
		"JOB_STATE_PENDING":     2,
		"JOB_STATE_RUNNING":     3,
		"JOB_STATE_SUCCEEDED":   4,
		"JOB_STATE_FAILED":      5,
		"JOB_STATE_CANCELLING":  6,
		"JOB_STATE_CANCELLED":   7,
		"JOB_STATE_PAUSED":      8,
		"JOB_STATE_EXPIRED":     9,
	}
)

func (x JobState) Enum() *JobState {
	p := new(JobState)
	*p = x
	return p
}

func (x JobState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobState) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_aiplatform_v1_job_state_proto_enumTypes[0].Descriptor()
}

func (JobState) Type() protoreflect.EnumType {
	return &file_google_cloud_aiplatform_v1_job_state_proto_enumTypes[0]
}

func (x JobState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobState.Descriptor instead.
func (JobState) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_job_state_proto_rawDescGZIP(), []int{0}
}

var File_google_cloud_aiplatform_v1_job_state_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1_job_state_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xf8, 0x01, 0x0a, 0x08, 0x4a, 0x6f, 0x62, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14,
	0x0a, 0x10, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x51, 0x55, 0x45, 0x55,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x4a,
	0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47,
	0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x4a,
	0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10,
	0x05, 0x12, 0x18, 0x0a, 0x14, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43,
	0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x06, 0x12, 0x17, 0x0a, 0x13, 0x4a,
	0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c,
	0x45, 0x44, 0x10, 0x07, 0x12, 0x14, 0x0a, 0x10, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x08, 0x12, 0x15, 0x0a, 0x11, 0x4a, 0x4f,
	0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10,
	0x09, 0x42, 0xd1, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31,
	0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0xaa, 0x02, 0x1a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1_job_state_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1_job_state_proto_rawDescData = file_google_cloud_aiplatform_v1_job_state_proto_rawDesc
)

func file_google_cloud_aiplatform_v1_job_state_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1_job_state_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1_job_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1_job_state_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1_job_state_proto_rawDescData
}

var file_google_cloud_aiplatform_v1_job_state_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_aiplatform_v1_job_state_proto_goTypes = []interface{}{
	(JobState)(0), // 0: google.cloud.aiplatform.v1.JobState
}
var file_google_cloud_aiplatform_v1_job_state_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1_job_state_proto_init() }
func file_google_cloud_aiplatform_v1_job_state_proto_init() {
	if File_google_cloud_aiplatform_v1_job_state_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1_job_state_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1_job_state_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1_job_state_proto_depIdxs,
		EnumInfos:         file_google_cloud_aiplatform_v1_job_state_proto_enumTypes,
	}.Build()
	File_google_cloud_aiplatform_v1_job_state_proto = out.File
	file_google_cloud_aiplatform_v1_job_state_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1_job_state_proto_goTypes = nil
	file_google_cloud_aiplatform_v1_job_state_proto_depIdxs = nil
}
