// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/expense_categories/service.proto

package expense_categories

import (
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

var File_proto_expense_categories_service_proto protoreflect.FileDescriptor

var file_proto_expense_categories_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x5f,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73,
	0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x25, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x6e,
	0x73, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x67, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x97, 0x02, 0x0a, 0x16, 0x45, 0x78, 0x70, 0x65,
	0x6e, 0x73, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x7c, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65,
	0x6e, 0x73, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x30, 0x2e, 0x65, 0x78,
	0x70, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e,
	0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x7f, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x2e, 0x65,
	0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x1a, 0x2e, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x72, 0x75, 0x62, 0x65, 0x6e, 0x67, 0x6f, 0x6d, 0x65, 0x73, 0x38, 0x2f, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2d, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x2d, 0x66, 0x69, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x6e,
	0x73, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_expense_categories_service_proto_goTypes = []interface{}{
	(*ExpenseCategoryCreateRequest)(nil),    // 0: expense_categories.ExpenseCategoryCreateRequest
	(*ExpenseCategoryGetRequestByName)(nil), // 1: expense_categories.ExpenseCategoryGetRequestByName
	(*ExpenseCategoryCreateResponse)(nil),   // 2: expense_categories.ExpenseCategoryCreateResponse
	(*ExpenseCategoryGetResponse)(nil),      // 3: expense_categories.ExpenseCategoryGetResponse
}
var file_proto_expense_categories_service_proto_depIdxs = []int32{
	0, // 0: expense_categories.ExpenseCategoryService.CreateExpenseCategory:input_type -> expense_categories.ExpenseCategoryCreateRequest
	1, // 1: expense_categories.ExpenseCategoryService.GetExpenseCategoryByName:input_type -> expense_categories.ExpenseCategoryGetRequestByName
	2, // 2: expense_categories.ExpenseCategoryService.CreateExpenseCategory:output_type -> expense_categories.ExpenseCategoryCreateResponse
	3, // 3: expense_categories.ExpenseCategoryService.GetExpenseCategoryByName:output_type -> expense_categories.ExpenseCategoryGetResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_expense_categories_service_proto_init() }
func file_proto_expense_categories_service_proto_init() {
	if File_proto_expense_categories_service_proto != nil {
		return
	}
	file_proto_expense_categories_create_proto_init()
	file_proto_expense_categories_get_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_expense_categories_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_expense_categories_service_proto_goTypes,
		DependencyIndexes: file_proto_expense_categories_service_proto_depIdxs,
	}.Build()
	File_proto_expense_categories_service_proto = out.File
	file_proto_expense_categories_service_proto_rawDesc = nil
	file_proto_expense_categories_service_proto_goTypes = nil
	file_proto_expense_categories_service_proto_depIdxs = nil
}