// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v3.12.4
// source: cart.proto

package cart

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddItemRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductId     string                 `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity      int32                  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddItemRequest) Reset() {
	*x = AddItemRequest{}
	mi := &file_cart_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemRequest) ProtoMessage() {}

func (x *AddItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemRequest.ProtoReflect.Descriptor instead.
func (*AddItemRequest) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{0}
}

func (x *AddItemRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddItemRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *AddItemRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type AddItemResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddItemResponse) Reset() {
	*x = AddItemResponse{}
	mi := &file_cart_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemResponse) ProtoMessage() {}

func (x *AddItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemResponse.ProtoReflect.Descriptor instead.
func (*AddItemResponse) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{1}
}

func (x *AddItemResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AddItemResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RemoveItemRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductId     string                 `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveItemRequest) Reset() {
	*x = RemoveItemRequest{}
	mi := &file_cart_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveItemRequest) ProtoMessage() {}

func (x *RemoveItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveItemRequest.ProtoReflect.Descriptor instead.
func (*RemoveItemRequest) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveItemRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RemoveItemRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type RemoveItemResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveItemResponse) Reset() {
	*x = RemoveItemResponse{}
	mi := &file_cart_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveItemResponse) ProtoMessage() {}

func (x *RemoveItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveItemResponse.ProtoReflect.Descriptor instead.
func (*RemoveItemResponse) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveItemResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RemoveItemResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetCartRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCartRequest) Reset() {
	*x = GetCartRequest{}
	mi := &file_cart_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartRequest) ProtoMessage() {}

func (x *GetCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartRequest.ProtoReflect.Descriptor instead.
func (*GetCartRequest) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{4}
}

func (x *GetCartRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetCartResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items         []*CartItem            `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	TotalPrice    float64                `protobuf:"fixed64,3,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"` // Optional: Useful for UI display
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCartResponse) Reset() {
	*x = GetCartResponse{}
	mi := &file_cart_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartResponse) ProtoMessage() {}

func (x *GetCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartResponse.ProtoReflect.Descriptor instead.
func (*GetCartResponse) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{5}
}

func (x *GetCartResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetCartResponse) GetItems() []*CartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *GetCartResponse) GetTotalPrice() float64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

type CartItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity      int32                  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price         float64                `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"` // Optional: Store product price at the time of adding to cart
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	mi := &file_cart_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{6}
}

func (x *CartItem) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CartItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *CartItem) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type ClearCartRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClearCartRequest) Reset() {
	*x = ClearCartRequest{}
	mi := &file_cart_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClearCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearCartRequest) ProtoMessage() {}

func (x *ClearCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearCartRequest.ProtoReflect.Descriptor instead.
func (*ClearCartRequest) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{7}
}

func (x *ClearCartRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ClearCartResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClearCartResponse) Reset() {
	*x = ClearCartResponse{}
	mi := &file_cart_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClearCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearCartResponse) ProtoMessage() {}

func (x *ClearCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearCartResponse.ProtoReflect.Descriptor instead.
func (*ClearCartResponse) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{8}
}

func (x *ClearCartResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ClearCartResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_cart_proto protoreflect.FileDescriptor

var file_cart_proto_rawDesc = string([]byte{
	0x0a, 0x0a, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x61,
	0x72, 0x74, 0x22, 0x64, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x45, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x4b, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x12,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x71, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63,
	0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x22, 0x5b, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x22, 0x2b, 0x0a, 0x10, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x47,
	0x0a, 0x11, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xfc, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x14, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e,
	0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3f, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x2e,
	0x63, 0x61, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x36, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x12, 0x14, 0x2e, 0x63, 0x61,
	0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x43, 0x6c, 0x65, 0x61,
	0x72, 0x43, 0x61, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x6c, 0x65,
	0x61, 0x72, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x55, 0x6b, 0x6b, 0x65, 0x6e, 0x6a, 0x69, 0x6a, 0x6f, 0x2f, 0x54,
	0x72, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x65, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_cart_proto_rawDescOnce sync.Once
	file_cart_proto_rawDescData []byte
)

func file_cart_proto_rawDescGZIP() []byte {
	file_cart_proto_rawDescOnce.Do(func() {
		file_cart_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_cart_proto_rawDesc), len(file_cart_proto_rawDesc)))
	})
	return file_cart_proto_rawDescData
}

var file_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_cart_proto_goTypes = []any{
	(*AddItemRequest)(nil),     // 0: cart.AddItemRequest
	(*AddItemResponse)(nil),    // 1: cart.AddItemResponse
	(*RemoveItemRequest)(nil),  // 2: cart.RemoveItemRequest
	(*RemoveItemResponse)(nil), // 3: cart.RemoveItemResponse
	(*GetCartRequest)(nil),     // 4: cart.GetCartRequest
	(*GetCartResponse)(nil),    // 5: cart.GetCartResponse
	(*CartItem)(nil),           // 6: cart.CartItem
	(*ClearCartRequest)(nil),   // 7: cart.ClearCartRequest
	(*ClearCartResponse)(nil),  // 8: cart.ClearCartResponse
}
var file_cart_proto_depIdxs = []int32{
	6, // 0: cart.GetCartResponse.items:type_name -> cart.CartItem
	0, // 1: cart.CartService.AddItem:input_type -> cart.AddItemRequest
	2, // 2: cart.CartService.RemoveItem:input_type -> cart.RemoveItemRequest
	4, // 3: cart.CartService.GetCart:input_type -> cart.GetCartRequest
	7, // 4: cart.CartService.ClearCart:input_type -> cart.ClearCartRequest
	1, // 5: cart.CartService.AddItem:output_type -> cart.AddItemResponse
	3, // 6: cart.CartService.RemoveItem:output_type -> cart.RemoveItemResponse
	5, // 7: cart.CartService.GetCart:output_type -> cart.GetCartResponse
	8, // 8: cart.CartService.ClearCart:output_type -> cart.ClearCartResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cart_proto_init() }
func file_cart_proto_init() {
	if File_cart_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_cart_proto_rawDesc), len(file_cart_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cart_proto_goTypes,
		DependencyIndexes: file_cart_proto_depIdxs,
		MessageInfos:      file_cart_proto_msgTypes,
	}.Build()
	File_cart_proto = out.File
	file_cart_proto_goTypes = nil
	file_cart_proto_depIdxs = nil
}
