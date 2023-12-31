// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: proto/booking.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// adjust later after implement update booking status function
type BookingStatusEnum int32

const (
	BookingStatusEnum_UNKNOWN    BookingStatusEnum = 0
	BookingStatusEnum_BOOKED     BookingStatusEnum = 1
	BookingStatusEnum_CHECKED_IN BookingStatusEnum = 2
	BookingStatusEnum_COMPLETED  BookingStatusEnum = 3
	BookingStatusEnum_MISSED     BookingStatusEnum = 4
)

// Enum value maps for BookingStatusEnum.
var (
	BookingStatusEnum_name = map[int32]string{
		0: "UNKNOWN",
		1: "BOOKED",
		2: "CHECKED_IN",
		3: "COMPLETED",
		4: "MISSED",
	}
	BookingStatusEnum_value = map[string]int32{
		"UNKNOWN":    0,
		"BOOKED":     1,
		"CHECKED_IN": 2,
		"COMPLETED":  3,
		"MISSED":     4,
	}
)

func (x BookingStatusEnum) Enum() *BookingStatusEnum {
	p := new(BookingStatusEnum)
	*p = x
	return p
}

func (x BookingStatusEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BookingStatusEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_booking_proto_enumTypes[0].Descriptor()
}

func (BookingStatusEnum) Type() protoreflect.EnumType {
	return &file_proto_booking_proto_enumTypes[0]
}

func (x BookingStatusEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BookingStatusEnum.Descriptor instead.
func (BookingStatusEnum) EnumDescriptor() ([]byte, []int) {
	return file_proto_booking_proto_rawDescGZIP(), []int{0}
}

type BookingId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BookingId) Reset() {
	*x = BookingId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingId) ProtoMessage() {}

func (x *BookingId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingId.ProtoReflect.Descriptor instead.
func (*BookingId) Descriptor() ([]byte, []int) {
	return file_proto_booking_proto_rawDescGZIP(), []int{0}
}

func (x *BookingId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UserId) Reset() {
	*x = UserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserId) ProtoMessage() {}

func (x *UserId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserId.ProtoReflect.Descriptor instead.
func (*UserId) Descriptor() ([]byte, []int) {
	return file_proto_booking_proto_rawDescGZIP(), []int{1}
}

func (x *UserId) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Seat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeatId int32 `protobuf:"varint,1,opt,name=seatId,proto3" json:"seatId,omitempty"`
}

func (x *Seat) Reset() {
	*x = Seat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Seat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Seat) ProtoMessage() {}

func (x *Seat) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Seat.ProtoReflect.Descriptor instead.
func (*Seat) Descriptor() ([]byte, []int) {
	return file_proto_booking_proto_rawDescGZIP(), []int{2}
}

func (x *Seat) GetSeatId() int32 {
	if x != nil {
		return x.SeatId
	}
	return 0
}

type SeatList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seats []*Seat `protobuf:"bytes,1,rep,name=seats,proto3" json:"seats,omitempty"`
}

func (x *SeatList) Reset() {
	*x = SeatList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeatList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeatList) ProtoMessage() {}

func (x *SeatList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeatList.ProtoReflect.Descriptor instead.
func (*SeatList) Descriptor() ([]byte, []int) {
	return file_proto_booking_proto_rawDescGZIP(), []int{3}
}

func (x *SeatList) GetSeats() []*Seat {
	if x != nil {
		return x.Seats
	}
	return nil
}

type BookingTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime int64 `protobuf:"varint,1,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime   int64 `protobuf:"varint,2,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (x *BookingTime) Reset() {
	*x = BookingTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingTime) ProtoMessage() {}

func (x *BookingTime) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingTime.ProtoReflect.Descriptor instead.
func (*BookingTime) Descriptor() ([]byte, []int) {
	return file_proto_booking_proto_rawDescGZIP(), []int{4}
}

func (x *BookingTime) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *BookingTime) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

type BookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User        *UserId           `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	BookingTime *BookingTime      `protobuf:"bytes,2,opt,name=bookingTime,proto3" json:"bookingTime,omitempty"`
	Seat        *Seat             `protobuf:"bytes,3,opt,name=seat,proto3" json:"seat,omitempty"`
	Status      BookingStatusEnum `protobuf:"varint,4,opt,name=status,proto3,enum=booking.BookingStatusEnum" json:"status,omitempty"`
}

func (x *BookingRequest) Reset() {
	*x = BookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingRequest) ProtoMessage() {}

func (x *BookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingRequest.ProtoReflect.Descriptor instead.
func (*BookingRequest) Descriptor() ([]byte, []int) {
	return file_proto_booking_proto_rawDescGZIP(), []int{5}
}

func (x *BookingRequest) GetUser() *UserId {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *BookingRequest) GetBookingTime() *BookingTime {
	if x != nil {
		return x.BookingTime
	}
	return nil
}

func (x *BookingRequest) GetSeat() *Seat {
	if x != nil {
		return x.Seat
	}
	return nil
}

func (x *BookingRequest) GetStatus() BookingStatusEnum {
	if x != nil {
		return x.Status
	}
	return BookingStatusEnum_UNKNOWN
}

type BookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           *BookingId      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BookingData  *BookingRequest `protobuf:"bytes,2,opt,name=bookingData,proto3" json:"bookingData,omitempty"`
	CheckinTime  int64           `protobuf:"varint,3,opt,name=checkinTime,proto3" json:"checkinTime,omitempty"`
	CheckoutTime int64           `protobuf:"varint,4,opt,name=checkoutTime,proto3" json:"checkoutTime,omitempty"`
	IsActive     bool            `protobuf:"varint,5,opt,name=isActive,proto3" json:"isActive,omitempty"`
}

func (x *BookingResponse) Reset() {
	*x = BookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingResponse) ProtoMessage() {}

func (x *BookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingResponse.ProtoReflect.Descriptor instead.
func (*BookingResponse) Descriptor() ([]byte, []int) {
	return file_proto_booking_proto_rawDescGZIP(), []int{6}
}

func (x *BookingResponse) GetId() *BookingId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *BookingResponse) GetBookingData() *BookingRequest {
	if x != nil {
		return x.BookingData
	}
	return nil
}

func (x *BookingResponse) GetCheckinTime() int64 {
	if x != nil {
		return x.CheckinTime
	}
	return 0
}

func (x *BookingResponse) GetCheckoutTime() int64 {
	if x != nil {
		return x.CheckoutTime
	}
	return 0
}

func (x *BookingResponse) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type BookingList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bookings []*BookingResponse `protobuf:"bytes,1,rep,name=bookings,proto3" json:"bookings,omitempty"`
}

func (x *BookingList) Reset() {
	*x = BookingList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingList) ProtoMessage() {}

func (x *BookingList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingList.ProtoReflect.Descriptor instead.
func (*BookingList) Descriptor() ([]byte, []int) {
	return file_proto_booking_proto_rawDescGZIP(), []int{7}
}

func (x *BookingList) GetBookings() []*BookingResponse {
	if x != nil {
		return x.Bookings
	}
	return nil
}

type BookingUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          *BookingId      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BookingData *BookingRequest `protobuf:"bytes,2,opt,name=bookingData,proto3" json:"bookingData,omitempty"`
}

func (x *BookingUpdateRequest) Reset() {
	*x = BookingUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingUpdateRequest) ProtoMessage() {}

func (x *BookingUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingUpdateRequest.ProtoReflect.Descriptor instead.
func (*BookingUpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_booking_proto_rawDescGZIP(), []int{8}
}

func (x *BookingUpdateRequest) GetId() *BookingId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *BookingUpdateRequest) GetBookingData() *BookingRequest {
	if x != nil {
		return x.BookingData
	}
	return nil
}

type UpdateBookingStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     *BookingId        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status BookingStatusEnum `protobuf:"varint,2,opt,name=status,proto3,enum=booking.BookingStatusEnum" json:"status,omitempty"`
}

func (x *UpdateBookingStatusRequest) Reset() {
	*x = UpdateBookingStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookingStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookingStatusRequest) ProtoMessage() {}

func (x *UpdateBookingStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookingStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateBookingStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_booking_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateBookingStatusRequest) GetId() *BookingId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *UpdateBookingStatusRequest) GetStatus() BookingStatusEnum {
	if x != nil {
		return x.Status
	}
	return BookingStatusEnum_UNKNOWN
}

var File_proto_booking_proto protoreflect.FileDescriptor

var file_proto_booking_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x09, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x1e, 0x0a, 0x04, 0x53, 0x65,
	0x61, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x65, 0x61, 0x74, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x08, 0x53, 0x65,
	0x61, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x53, 0x65, 0x61, 0x74, 0x52, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x22, 0x45, 0x0a, 0x0b, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0xc4, 0x01, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x0b, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x73, 0x65, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x52,
	0x04, 0x73, 0x65, 0x61, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75,
	0x6d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xd2, 0x01, 0x0a, 0x0f, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x39, 0x0a, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x43,
	0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x34, 0x0a,
	0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x73, 0x22, 0x75, 0x0a, 0x14, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x39, 0x0a, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x22, 0x74, 0x0a, 0x1a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2a, 0x57, 0x0a, 0x11, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x4f, 0x4f, 0x4b, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0a, 0x0a,
	0x06, 0x4d, 0x49, 0x53, 0x53, 0x45, 0x44, 0x10, 0x04, 0x32, 0xe5, 0x03, 0x0a, 0x0e, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0f,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a,
	0x14, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x12, 0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x1a, 0x18, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x53, 0x65, 0x61, 0x74, 0x12, 0x14, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0x11, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x42, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x12, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54,
	0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_booking_proto_rawDescOnce sync.Once
	file_proto_booking_proto_rawDescData = file_proto_booking_proto_rawDesc
)

func file_proto_booking_proto_rawDescGZIP() []byte {
	file_proto_booking_proto_rawDescOnce.Do(func() {
		file_proto_booking_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_booking_proto_rawDescData)
	})
	return file_proto_booking_proto_rawDescData
}

var file_proto_booking_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_booking_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_booking_proto_goTypes = []interface{}{
	(BookingStatusEnum)(0),             // 0: booking.BookingStatusEnum
	(*BookingId)(nil),                  // 1: booking.BookingId
	(*UserId)(nil),                     // 2: booking.UserId
	(*Seat)(nil),                       // 3: booking.Seat
	(*SeatList)(nil),                   // 4: booking.SeatList
	(*BookingTime)(nil),                // 5: booking.BookingTime
	(*BookingRequest)(nil),             // 6: booking.BookingRequest
	(*BookingResponse)(nil),            // 7: booking.BookingResponse
	(*BookingList)(nil),                // 8: booking.BookingList
	(*BookingUpdateRequest)(nil),       // 9: booking.BookingUpdateRequest
	(*UpdateBookingStatusRequest)(nil), // 10: booking.UpdateBookingStatusRequest
	(*emptypb.Empty)(nil),              // 11: google.protobuf.Empty
}
var file_proto_booking_proto_depIdxs = []int32{
	3,  // 0: booking.SeatList.seats:type_name -> booking.Seat
	2,  // 1: booking.BookingRequest.user:type_name -> booking.UserId
	5,  // 2: booking.BookingRequest.bookingTime:type_name -> booking.BookingTime
	3,  // 3: booking.BookingRequest.seat:type_name -> booking.Seat
	0,  // 4: booking.BookingRequest.status:type_name -> booking.BookingStatusEnum
	1,  // 5: booking.BookingResponse.id:type_name -> booking.BookingId
	6,  // 6: booking.BookingResponse.bookingData:type_name -> booking.BookingRequest
	7,  // 7: booking.BookingList.bookings:type_name -> booking.BookingResponse
	1,  // 8: booking.BookingUpdateRequest.id:type_name -> booking.BookingId
	6,  // 9: booking.BookingUpdateRequest.bookingData:type_name -> booking.BookingRequest
	1,  // 10: booking.UpdateBookingStatusRequest.id:type_name -> booking.BookingId
	0,  // 11: booking.UpdateBookingStatusRequest.status:type_name -> booking.BookingStatusEnum
	2,  // 12: booking.BookingService.GetUserHistory:input_type -> booking.UserId
	1,  // 13: booking.BookingService.GetBooking:input_type -> booking.BookingId
	5,  // 14: booking.BookingService.GetUnavailableSeat:input_type -> booking.BookingTime
	6,  // 15: booking.BookingService.CreateBooking:input_type -> booking.BookingRequest
	9,  // 16: booking.BookingService.UpdateBooking:input_type -> booking.BookingUpdateRequest
	10, // 17: booking.BookingService.UpdateBookingStatus:input_type -> booking.UpdateBookingStatusRequest
	1,  // 18: booking.BookingService.DeleteBooking:input_type -> booking.BookingId
	8,  // 19: booking.BookingService.GetUserHistory:output_type -> booking.BookingList
	7,  // 20: booking.BookingService.GetBooking:output_type -> booking.BookingResponse
	4,  // 21: booking.BookingService.GetUnavailableSeat:output_type -> booking.SeatList
	7,  // 22: booking.BookingService.CreateBooking:output_type -> booking.BookingResponse
	7,  // 23: booking.BookingService.UpdateBooking:output_type -> booking.BookingResponse
	7,  // 24: booking.BookingService.UpdateBookingStatus:output_type -> booking.BookingResponse
	11, // 25: booking.BookingService.DeleteBooking:output_type -> google.protobuf.Empty
	19, // [19:26] is the sub-list for method output_type
	12, // [12:19] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_proto_booking_proto_init() }
func file_proto_booking_proto_init() {
	if File_proto_booking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_booking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingId); i {
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
		file_proto_booking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserId); i {
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
		file_proto_booking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Seat); i {
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
		file_proto_booking_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeatList); i {
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
		file_proto_booking_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingTime); i {
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
		file_proto_booking_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingRequest); i {
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
		file_proto_booking_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingResponse); i {
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
		file_proto_booking_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingList); i {
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
		file_proto_booking_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingUpdateRequest); i {
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
		file_proto_booking_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBookingStatusRequest); i {
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
			RawDescriptor: file_proto_booking_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_booking_proto_goTypes,
		DependencyIndexes: file_proto_booking_proto_depIdxs,
		EnumInfos:         file_proto_booking_proto_enumTypes,
		MessageInfos:      file_proto_booking_proto_msgTypes,
	}.Build()
	File_proto_booking_proto = out.File
	file_proto_booking_proto_rawDesc = nil
	file_proto_booking_proto_goTypes = nil
	file_proto_booking_proto_depIdxs = nil
}
