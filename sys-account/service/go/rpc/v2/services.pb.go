// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: services.proto

package v2

import (
	pkg "github.com/getcouragenow/sys-share/pkg"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var File_services_proto protoreflect.FileDescriptor

var file_services_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x0c, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe5, 0x03, 0x0a, 0x0e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30,
	0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x0f, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00,
	0x12, 0x3a, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x1b, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x13,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x22, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x41, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x0f, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12,
	0x42, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x00, 0x32, 0x88, 0x03, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0e,
	0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1d,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4e, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x1c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5d, 0x0a, 0x12, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x42,
	0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x74,
	0x63, 0x6f, 0x75, 0x72, 0x61, 0x67, 0x65, 0x6e, 0x6f, 0x77, 0x2f, 0x73, 0x79, 0x73, 0x2d, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x2f, 0x73, 0x79, 0x73, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_services_proto_goTypes = []interface{}{
	(*pkg.Account)(nil),                    // 0: models.Account
	(*pkg.GetAccountRequest)(nil),          // 1: models.GetAccountRequest
	(*pkg.ListAccountsRequest)(nil),        // 2: models.ListAccountsRequest
	(*pkg.SearchAccountsRequest)(nil),      // 3: models.SearchAccountsRequest
	(*pkg.AssignAccountToRoleRequest)(nil), // 4: models.AssignAccountToRoleRequest
	(*pkg.DisableAccountRequest)(nil),      // 5: models.DisableAccountRequest
	(*pkg.RegisterRequest)(nil),            // 6: models.RegisterRequest
	(*pkg.LoginRequest)(nil),               // 7: models.LoginRequest
	(*pkg.ForgotPasswordRequest)(nil),      // 8: models.ForgotPasswordRequest
	(*pkg.ResetPasswordRequest)(nil),       // 9: models.ResetPasswordRequest
	(*pkg.RefreshAccessTokenRequest)(nil),  // 10: models.RefreshAccessTokenRequest
	(*pkg.ListAccountsResponse)(nil),       // 11: models.ListAccountsResponse
	(*pkg.SearchAccountsResponse)(nil),     // 12: models.SearchAccountsResponse
	(*pkg.RegisterResponse)(nil),           // 13: models.RegisterResponse
	(*pkg.LoginResponse)(nil),              // 14: models.LoginResponse
	(*pkg.ForgotPasswordResponse)(nil),     // 15: models.ForgotPasswordResponse
	(*pkg.ResetPasswordResponse)(nil),      // 16: models.ResetPasswordResponse
	(*pkg.RefreshAccessTokenResponse)(nil), // 17: models.RefreshAccessTokenResponse
}
var file_services_proto_depIdxs = []int32{
	0,  // 0: v2.services.AccountService.NewAccount:input_type -> models.Account
	1,  // 1: v2.services.AccountService.GetAccount:input_type -> models.GetAccountRequest
	2,  // 2: v2.services.AccountService.ListAccounts:input_type -> models.ListAccountsRequest
	3,  // 3: v2.services.AccountService.SearchAccounts:input_type -> models.SearchAccountsRequest
	4,  // 4: v2.services.AccountService.AssignAccountToRole:input_type -> models.AssignAccountToRoleRequest
	0,  // 5: v2.services.AccountService.UpdateAccount:input_type -> models.Account
	5,  // 6: v2.services.AccountService.DisableAccount:input_type -> models.DisableAccountRequest
	6,  // 7: v2.services.AuthService.Register:input_type -> models.RegisterRequest
	7,  // 8: v2.services.AuthService.Login:input_type -> models.LoginRequest
	8,  // 9: v2.services.AuthService.ForgotPassword:input_type -> models.ForgotPasswordRequest
	9,  // 10: v2.services.AuthService.ResetPassword:input_type -> models.ResetPasswordRequest
	10, // 11: v2.services.AuthService.RefreshAccessToken:input_type -> models.RefreshAccessTokenRequest
	0,  // 12: v2.services.AccountService.NewAccount:output_type -> models.Account
	0,  // 13: v2.services.AccountService.GetAccount:output_type -> models.Account
	11, // 14: v2.services.AccountService.ListAccounts:output_type -> models.ListAccountsResponse
	12, // 15: v2.services.AccountService.SearchAccounts:output_type -> models.SearchAccountsResponse
	0,  // 16: v2.services.AccountService.AssignAccountToRole:output_type -> models.Account
	0,  // 17: v2.services.AccountService.UpdateAccount:output_type -> models.Account
	0,  // 18: v2.services.AccountService.DisableAccount:output_type -> models.Account
	13, // 19: v2.services.AuthService.Register:output_type -> models.RegisterResponse
	14, // 20: v2.services.AuthService.Login:output_type -> models.LoginResponse
	15, // 21: v2.services.AuthService.ForgotPassword:output_type -> models.ForgotPasswordResponse
	16, // 22: v2.services.AuthService.ResetPassword:output_type -> models.ResetPasswordResponse
	17, // 23: v2.services.AuthService.RefreshAccessToken:output_type -> models.RefreshAccessTokenResponse
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_services_proto_init() }
func file_services_proto_init() {
	if File_services_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_services_proto_goTypes,
		DependencyIndexes: file_services_proto_depIdxs,
	}.Build()
	File_services_proto = out.File
	file_services_proto_rawDesc = nil
	file_services_proto_goTypes = nil
	file_services_proto_depIdxs = nil
}
