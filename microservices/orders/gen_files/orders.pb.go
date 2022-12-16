// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: microservices/orders/gen_files/orders.proto

package orders

import (
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

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	City     string `protobuf:"bytes,2,opt,name=City,proto3" json:"City,omitempty"`
	Street   string `protobuf:"bytes,3,opt,name=Street,proto3" json:"Street,omitempty"`
	House    string `protobuf:"bytes,4,opt,name=House,proto3" json:"House,omitempty"`
	Flat     string `protobuf:"bytes,5,opt,name=Flat,proto3" json:"Flat,omitempty"`
	Priority bool   `protobuf:"varint,6,opt,name=Priority,proto3" json:"Priority,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_orders_gen_files_orders_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_orders_gen_files_orders_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_microservices_orders_gen_files_orders_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Address) GetHouse() string {
	if x != nil {
		return x.House
	}
	return ""
}

func (x *Address) GetFlat() string {
	if x != nil {
		return x.Flat
	}
	return ""
}

func (x *Address) GetPriority() bool {
	if x != nil {
		return x.Priority
	}
	return false
}

type PaymentMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PaymentType string `protobuf:"bytes,2,opt,name=PaymentType,proto3" json:"PaymentType,omitempty"`
	Number      string `protobuf:"bytes,3,opt,name=Number,proto3" json:"Number,omitempty"`
	ExpiryDate  int64  `protobuf:"varint,4,opt,name=ExpiryDate,proto3" json:"ExpiryDate,omitempty"`
	Priority    bool   `protobuf:"varint,5,opt,name=Priority,proto3" json:"Priority,omitempty"`
}

func (x *PaymentMethod) Reset() {
	*x = PaymentMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_orders_gen_files_orders_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentMethod) ProtoMessage() {}

func (x *PaymentMethod) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_orders_gen_files_orders_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentMethod.ProtoReflect.Descriptor instead.
func (*PaymentMethod) Descriptor() ([]byte, []int) {
	return file_microservices_orders_gen_files_orders_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentMethod) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *PaymentMethod) GetPaymentType() string {
	if x != nil {
		return x.PaymentType
	}
	return ""
}

func (x *PaymentMethod) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *PaymentMethod) GetExpiryDate() int64 {
	if x != nil {
		return x.ExpiryDate
	}
	return 0
}

func (x *PaymentMethod) GetPriority() bool {
	if x != nil {
		return x.Priority
	}
	return false
}

type MakeOrderType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID        int32   `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Items         []int32 `protobuf:"varint,2,rep,packed,name=Items,proto3" json:"Items,omitempty"`
	AddressID     int32   `protobuf:"varint,3,opt,name=AddressID,proto3" json:"AddressID,omitempty"`
	PaymentcardID int32   `protobuf:"varint,4,opt,name=PaymentcardID,proto3" json:"PaymentcardID,omitempty"`
	DeliveryDate  int64   `protobuf:"varint,5,opt,name=DeliveryDate,proto3" json:"DeliveryDate,omitempty"`
}

func (x *MakeOrderType) Reset() {
	*x = MakeOrderType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_orders_gen_files_orders_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeOrderType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeOrderType) ProtoMessage() {}

func (x *MakeOrderType) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_orders_gen_files_orders_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeOrderType.ProtoReflect.Descriptor instead.
func (*MakeOrderType) Descriptor() ([]byte, []int) {
	return file_microservices_orders_gen_files_orders_proto_rawDescGZIP(), []int{2}
}

func (x *MakeOrderType) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *MakeOrderType) GetItems() []int32 {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *MakeOrderType) GetAddressID() int32 {
	if x != nil {
		return x.AddressID
	}
	return 0
}

func (x *MakeOrderType) GetPaymentcardID() int32 {
	if x != nil {
		return x.PaymentcardID
	}
	return 0
}

func (x *MakeOrderType) GetDeliveryDate() int64 {
	if x != nil {
		return x.DeliveryDate
	}
	return 0
}

type CartProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           int32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name         string  `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Count        int32   `protobuf:"varint,3,opt,name=Count,proto3" json:"Count,omitempty"`
	Price        float64 `protobuf:"fixed64,4,opt,name=Price,proto3" json:"Price,omitempty"`
	NominalPrice float64 `protobuf:"fixed64,5,opt,name=NominalPrice,proto3" json:"NominalPrice,omitempty"`
	Imgsrc       *string `protobuf:"bytes,6,opt,name=Imgsrc,proto3,oneof" json:"Imgsrc,omitempty"`
}

func (x *CartProduct) Reset() {
	*x = CartProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_orders_gen_files_orders_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartProduct) ProtoMessage() {}

func (x *CartProduct) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_orders_gen_files_orders_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartProduct.ProtoReflect.Descriptor instead.
func (*CartProduct) Descriptor() ([]byte, []int) {
	return file_microservices_orders_gen_files_orders_proto_rawDescGZIP(), []int{3}
}

func (x *CartProduct) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *CartProduct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CartProduct) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CartProduct) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CartProduct) GetNominalPrice() float64 {
	if x != nil {
		return x.NominalPrice
	}
	return 0
}

func (x *CartProduct) GetImgsrc() string {
	if x != nil && x.Imgsrc != nil {
		return *x.Imgsrc
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID            int32          `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID        int32          `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Items         []*CartProduct `protobuf:"bytes,3,rep,name=Items,proto3" json:"Items,omitempty"`
	OrderStatus   string         `protobuf:"bytes,4,opt,name=OrderStatus,proto3" json:"OrderStatus,omitempty"`
	PaymentStatus string         `protobuf:"bytes,5,opt,name=PaymentStatus,proto3" json:"PaymentStatus,omitempty"`
	Address       *Address       `protobuf:"bytes,6,opt,name=Address,proto3" json:"Address,omitempty"`
	PaymentMethod *PaymentMethod `protobuf:"bytes,7,opt,name=PaymentMethod,proto3" json:"PaymentMethod,omitempty"`
	CreationDate  int64          `protobuf:"varint,8,opt,name=CreationDate,proto3" json:"CreationDate,omitempty"`
	DeliveryDate  int64          `protobuf:"varint,9,opt,name=DeliveryDate,proto3" json:"DeliveryDate,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_orders_gen_files_orders_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_orders_gen_files_orders_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_microservices_orders_gen_files_orders_proto_rawDescGZIP(), []int{4}
}

func (x *Order) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Order) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Order) GetItems() []*CartProduct {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Order) GetOrderStatus() string {
	if x != nil {
		return x.OrderStatus
	}
	return ""
}

func (x *Order) GetPaymentStatus() string {
	if x != nil {
		return x.PaymentStatus
	}
	return ""
}

func (x *Order) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Order) GetPaymentMethod() *PaymentMethod {
	if x != nil {
		return x.PaymentMethod
	}
	return nil
}

func (x *Order) GetCreationDate() int64 {
	if x != nil {
		return x.CreationDate
	}
	return 0
}

func (x *Order) GetDeliveryDate() int64 {
	if x != nil {
		return x.DeliveryDate
	}
	return 0
}

type OrdersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *OrdersResponse) Reset() {
	*x = OrdersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_orders_gen_files_orders_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrdersResponse) ProtoMessage() {}

func (x *OrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_orders_gen_files_orders_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrdersResponse.ProtoReflect.Descriptor instead.
func (*OrdersResponse) Descriptor() ([]byte, []int) {
	return file_microservices_orders_gen_files_orders_proto_rawDescGZIP(), []int{5}
}

func (x *OrdersResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type Nothing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccessful bool `protobuf:"varint,1,opt,name=isSuccessful,proto3" json:"isSuccessful,omitempty"`
}

func (x *Nothing) Reset() {
	*x = Nothing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_orders_gen_files_orders_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nothing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nothing) ProtoMessage() {}

func (x *Nothing) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_orders_gen_files_orders_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nothing.ProtoReflect.Descriptor instead.
func (*Nothing) Descriptor() ([]byte, []int) {
	return file_microservices_orders_gen_files_orders_proto_rawDescGZIP(), []int{6}
}

func (x *Nothing) GetIsSuccessful() bool {
	if x != nil {
		return x.IsSuccessful
	}
	return false
}

type UserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int32 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *UserID) Reset() {
	*x = UserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_orders_gen_files_orders_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserID) ProtoMessage() {}

func (x *UserID) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_orders_gen_files_orders_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserID.ProtoReflect.Descriptor instead.
func (*UserID) Descriptor() ([]byte, []int) {
	return file_microservices_orders_gen_files_orders_proto_rawDescGZIP(), []int{7}
}

func (x *UserID) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

var File_microservices_orders_gen_files_orders_proto protoreflect.FileDescriptor

var file_microservices_orders_gen_files_orders_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x8b, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x48, 0x6f,
	0x75, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x6c, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x46, 0x6c, 0x61, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x22, 0x95, 0x01, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1e, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0xa9, 0x01, 0x0a, 0x0d,
	0x4d, 0x61, 0x6b, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x05, 0x42, 0x02, 0x10, 0x01, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x44, 0x12, 0x24, 0x0a,
	0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x63, 0x61, 0x72, 0x64, 0x49, 0x44, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x63, 0x61, 0x72,
	0x64, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x22, 0xa9, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x4e, 0x6f, 0x6d, 0x69, 0x6e,
	0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x4e,
	0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x49,
	0x6d, 0x67, 0x73, 0x72, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x49,
	0x6d, 0x67, 0x73, 0x72, 0x63, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x49, 0x6d, 0x67,
	0x73, 0x72, 0x63, 0x22, 0xd2, 0x02, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x29, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x43, 0x61,
	0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x29, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x3b, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x52, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x12, 0x22, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x22, 0x37, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x22, 0x2d, 0x0a, 0x07, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x22, 0x0a, 0x0c,
	0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c,
	0x22, 0x20, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x32, 0x7c, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x12, 0x35, 0x0a, 0x09, 0x4d, 0x61, 0x6b, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x0f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e,
	0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x0e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x16, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_microservices_orders_gen_files_orders_proto_rawDescOnce sync.Once
	file_microservices_orders_gen_files_orders_proto_rawDescData = file_microservices_orders_gen_files_orders_proto_rawDesc
)

func file_microservices_orders_gen_files_orders_proto_rawDescGZIP() []byte {
	file_microservices_orders_gen_files_orders_proto_rawDescOnce.Do(func() {
		file_microservices_orders_gen_files_orders_proto_rawDescData = protoimpl.X.CompressGZIP(file_microservices_orders_gen_files_orders_proto_rawDescData)
	})
	return file_microservices_orders_gen_files_orders_proto_rawDescData
}

var file_microservices_orders_gen_files_orders_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_microservices_orders_gen_files_orders_proto_goTypes = []interface{}{
	(*Address)(nil),        // 0: orders.Address
	(*PaymentMethod)(nil),  // 1: orders.PaymentMethod
	(*MakeOrderType)(nil),  // 2: orders.MakeOrderType
	(*CartProduct)(nil),    // 3: orders.CartProduct
	(*Order)(nil),          // 4: orders.Order
	(*OrdersResponse)(nil), // 5: orders.OrdersResponse
	(*Nothing)(nil),        // 6: orders.Nothing
	(*UserID)(nil),         // 7: orders.UserID
}
var file_microservices_orders_gen_files_orders_proto_depIdxs = []int32{
	3, // 0: orders.Order.Items:type_name -> orders.CartProduct
	0, // 1: orders.Order.Address:type_name -> orders.Address
	1, // 2: orders.Order.PaymentMethod:type_name -> orders.PaymentMethod
	4, // 3: orders.OrdersResponse.orders:type_name -> orders.Order
	2, // 4: orders.OrdersWorker.MakeOrder:input_type -> orders.MakeOrderType
	7, // 5: orders.OrdersWorker.GetOrders:input_type -> orders.UserID
	6, // 6: orders.OrdersWorker.MakeOrder:output_type -> orders.Nothing
	5, // 7: orders.OrdersWorker.GetOrders:output_type -> orders.OrdersResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_microservices_orders_gen_files_orders_proto_init() }
func file_microservices_orders_gen_files_orders_proto_init() {
	if File_microservices_orders_gen_files_orders_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_microservices_orders_gen_files_orders_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_microservices_orders_gen_files_orders_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentMethod); i {
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
		file_microservices_orders_gen_files_orders_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakeOrderType); i {
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
		file_microservices_orders_gen_files_orders_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartProduct); i {
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
		file_microservices_orders_gen_files_orders_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_microservices_orders_gen_files_orders_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrdersResponse); i {
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
		file_microservices_orders_gen_files_orders_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nothing); i {
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
		file_microservices_orders_gen_files_orders_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserID); i {
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
	file_microservices_orders_gen_files_orders_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_microservices_orders_gen_files_orders_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_microservices_orders_gen_files_orders_proto_goTypes,
		DependencyIndexes: file_microservices_orders_gen_files_orders_proto_depIdxs,
		MessageInfos:      file_microservices_orders_gen_files_orders_proto_msgTypes,
	}.Build()
	File_microservices_orders_gen_files_orders_proto = out.File
	file_microservices_orders_gen_files_orders_proto_rawDesc = nil
	file_microservices_orders_gen_files_orders_proto_goTypes = nil
	file_microservices_orders_gen_files_orders_proto_depIdxs = nil
}
