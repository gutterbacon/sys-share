// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: sys_core_bootstrap_models.proto

package v2

import (
	v21 "github.com/getcouragenow/mod/mod-disco/service/go/rpc/v2"
	v2 "github.com/getcouragenow/sys-share/sys-account/service/go/rpc/v2"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type BootstrapProject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project        *v2.ProjectRequest           `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	ProjectDetails *v21.NewDiscoProjectRequest  `protobuf:"bytes,2,opt,name=project_details,json=projectDetails,proto3" json:"project_details,omitempty"`
	SurveySchema   *v21.NewSurveyProjectRequest `protobuf:"bytes,3,opt,name=survey_schema,json=surveySchema,proto3" json:"survey_schema,omitempty"`
}

func (x *BootstrapProject) Reset() {
	*x = BootstrapProject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_core_bootstrap_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BootstrapProject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BootstrapProject) ProtoMessage() {}

func (x *BootstrapProject) ProtoReflect() protoreflect.Message {
	mi := &file_sys_core_bootstrap_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BootstrapProject.ProtoReflect.Descriptor instead.
func (*BootstrapProject) Descriptor() ([]byte, []int) {
	return file_sys_core_bootstrap_models_proto_rawDescGZIP(), []int{0}
}

func (x *BootstrapProject) GetProject() *v2.ProjectRequest {
	if x != nil {
		return x.Project
	}
	return nil
}

func (x *BootstrapProject) GetProjectDetails() *v21.NewDiscoProjectRequest {
	if x != nil {
		return x.ProjectDetails
	}
	return nil
}

func (x *BootstrapProject) GetSurveySchema() *v21.NewSurveyProjectRequest {
	if x != nil {
		return x.SurveySchema
	}
	return nil
}

type BootstrapOrg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Org *v2.Org `protobuf:"bytes,1,opt,name=org,proto3" json:"org,omitempty"`
}

func (x *BootstrapOrg) Reset() {
	*x = BootstrapOrg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_core_bootstrap_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BootstrapOrg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BootstrapOrg) ProtoMessage() {}

func (x *BootstrapOrg) ProtoReflect() protoreflect.Message {
	mi := &file_sys_core_bootstrap_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BootstrapOrg.ProtoReflect.Descriptor instead.
func (*BootstrapOrg) Descriptor() ([]byte, []int) {
	return file_sys_core_bootstrap_models_proto_rawDescGZIP(), []int{1}
}

func (x *BootstrapOrg) GetOrg() *v2.Org {
	if x != nil {
		return x.Org
	}
	return nil
}

type BootstrapRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orgs              []*BootstrapOrg       `protobuf:"bytes,1,rep,name=orgs,proto3" json:"orgs,omitempty"`
	Projects          []*BootstrapProject   `protobuf:"bytes,2,rep,name=projects,proto3" json:"projects,omitempty"`
	InitialSuperusers []*v2.RegisterRequest `protobuf:"bytes,3,rep,name=initial_superusers,json=initialSuperusers,proto3" json:"initial_superusers,omitempty"`
}

func (x *BootstrapRequest) Reset() {
	*x = BootstrapRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_core_bootstrap_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BootstrapRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BootstrapRequest) ProtoMessage() {}

func (x *BootstrapRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_core_bootstrap_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BootstrapRequest.ProtoReflect.Descriptor instead.
func (*BootstrapRequest) Descriptor() ([]byte, []int) {
	return file_sys_core_bootstrap_models_proto_rawDescGZIP(), []int{2}
}

func (x *BootstrapRequest) GetOrgs() []*BootstrapOrg {
	if x != nil {
		return x.Orgs
	}
	return nil
}

func (x *BootstrapRequest) GetProjects() []*BootstrapProject {
	if x != nil {
		return x.Projects
	}
	return nil
}

func (x *BootstrapRequest) GetInitialSuperusers() []*v2.RegisterRequest {
	if x != nil {
		return x.InitialSuperusers
	}
	return nil
}

var File_sys_core_bootstrap_models_proto protoreflect.FileDescriptor

var file_sys_core_bootstrap_models_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x79, 0x73, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x62, 0x6f, 0x6f, 0x74, 0x73,
	0x74, 0x72, 0x61, 0x70, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x76, 0x32, 0x2e, 0x73, 0x79, 0x73, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x57, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x61, 0x67, 0x65, 0x6e, 0x6f, 0x77, 0x2f, 0x73, 0x79, 0x73, 0x2d, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x2f, 0x73, 0x79, 0x73, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x79, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x4d, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x61, 0x67, 0x65, 0x6e, 0x6f,
	0x77, 0x2f, 0x6d, 0x6f, 0x64, 0x2f, 0x6d, 0x6f, 0x64, 0x2d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x6f, 0x64, 0x5f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x82, 0x02, 0x0a, 0x10, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x41, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x79, 0x73, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x56, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x76, 0x32, 0x2e, 0x6d, 0x6f, 0x64, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x0e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x53, 0x0a, 0x0d, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x76, 0x32, 0x2e, 0x6d, 0x6f, 0x64, 0x5f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e,
	0x65, 0x77, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x22, 0x3e, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61,
	0x70, 0x4f, 0x72, 0x67, 0x12, 0x2e, 0x0a, 0x03, 0x6f, 0x72, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x79, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x67, 0x52,
	0x03, 0x6f, 0x72, 0x67, 0x22, 0xe7, 0x01, 0x0a, 0x10, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72,
	0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x6f, 0x72, 0x67,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x79, 0x73,
	0x5f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42,
	0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x4f, 0x72, 0x67, 0x52, 0x04, 0x6f, 0x72, 0x67,
	0x73, 0x12, 0x42, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x79, 0x73, 0x5f, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x74, 0x73,
	0x74, 0x72, 0x61, 0x70, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x57, 0x0a, 0x12, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c,
	0x5f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x79, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x11, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x70, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x73, 0x42, 0x3f,
	0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x74,
	0x63, 0x6f, 0x75, 0x72, 0x61, 0x67, 0x65, 0x6e, 0x6f, 0x77, 0x2f, 0x73, 0x79, 0x73, 0x2d, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x2f, 0x73, 0x79, 0x73, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x32, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sys_core_bootstrap_models_proto_rawDescOnce sync.Once
	file_sys_core_bootstrap_models_proto_rawDescData = file_sys_core_bootstrap_models_proto_rawDesc
)

func file_sys_core_bootstrap_models_proto_rawDescGZIP() []byte {
	file_sys_core_bootstrap_models_proto_rawDescOnce.Do(func() {
		file_sys_core_bootstrap_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_sys_core_bootstrap_models_proto_rawDescData)
	})
	return file_sys_core_bootstrap_models_proto_rawDescData
}

var file_sys_core_bootstrap_models_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sys_core_bootstrap_models_proto_goTypes = []interface{}{
	(*BootstrapProject)(nil),            // 0: v2.sys_core.services.BootstrapProject
	(*BootstrapOrg)(nil),                // 1: v2.sys_core.services.BootstrapOrg
	(*BootstrapRequest)(nil),            // 2: v2.sys_core.services.BootstrapRequest
	(*v2.ProjectRequest)(nil),           // 3: v2.sys_account.services.ProjectRequest
	(*v21.NewDiscoProjectRequest)(nil),  // 4: v2.mod_disco.services.NewDiscoProjectRequest
	(*v21.NewSurveyProjectRequest)(nil), // 5: v2.mod_disco.services.NewSurveyProjectRequest
	(*v2.Org)(nil),                      // 6: v2.sys_account.services.Org
	(*v2.RegisterRequest)(nil),          // 7: v2.sys_account.services.RegisterRequest
}
var file_sys_core_bootstrap_models_proto_depIdxs = []int32{
	3, // 0: v2.sys_core.services.BootstrapProject.project:type_name -> v2.sys_account.services.ProjectRequest
	4, // 1: v2.sys_core.services.BootstrapProject.project_details:type_name -> v2.mod_disco.services.NewDiscoProjectRequest
	5, // 2: v2.sys_core.services.BootstrapProject.survey_schema:type_name -> v2.mod_disco.services.NewSurveyProjectRequest
	6, // 3: v2.sys_core.services.BootstrapOrg.org:type_name -> v2.sys_account.services.Org
	1, // 4: v2.sys_core.services.BootstrapRequest.orgs:type_name -> v2.sys_core.services.BootstrapOrg
	0, // 5: v2.sys_core.services.BootstrapRequest.projects:type_name -> v2.sys_core.services.BootstrapProject
	7, // 6: v2.sys_core.services.BootstrapRequest.initial_superusers:type_name -> v2.sys_account.services.RegisterRequest
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_sys_core_bootstrap_models_proto_init() }
func file_sys_core_bootstrap_models_proto_init() {
	if File_sys_core_bootstrap_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sys_core_bootstrap_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BootstrapProject); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sys_core_bootstrap_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BootstrapOrg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sys_core_bootstrap_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BootstrapRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sys_core_bootstrap_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sys_core_bootstrap_models_proto_goTypes,
		DependencyIndexes: file_sys_core_bootstrap_models_proto_depIdxs,
		MessageInfos:      file_sys_core_bootstrap_models_proto_msgTypes,
	}.Build()
	File_sys_core_bootstrap_models_proto = out.File
	file_sys_core_bootstrap_models_proto_rawDesc = nil
	file_sys_core_bootstrap_models_proto_goTypes = nil
	file_sys_core_bootstrap_models_proto_depIdxs = nil
}
