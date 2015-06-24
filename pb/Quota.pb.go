// Code generated by protoc-gen-go.
// source: Quota.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type QuotaScope int32

const (
	QuotaScope_CLUSTER QuotaScope = 1
	QuotaScope_MACHINE QuotaScope = 2
)

var QuotaScope_name = map[int32]string{
	1: "CLUSTER",
	2: "MACHINE",
}
var QuotaScope_value = map[string]int32{
	"CLUSTER": 1,
	"MACHINE": 2,
}

func (x QuotaScope) Enum() *QuotaScope {
	p := new(QuotaScope)
	*p = x
	return p
}
func (x QuotaScope) String() string {
	return proto.EnumName(QuotaScope_name, int32(x))
}
func (x *QuotaScope) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(QuotaScope_value, data, "QuotaScope")
	if err != nil {
		return err
	}
	*x = QuotaScope(value)
	return nil
}

type ThrottleType int32

const (
	ThrottleType_REQUEST_NUMBER ThrottleType = 1
	ThrottleType_REQUEST_SIZE   ThrottleType = 2
	ThrottleType_WRITE_NUMBER   ThrottleType = 3
	ThrottleType_WRITE_SIZE     ThrottleType = 4
	ThrottleType_READ_NUMBER    ThrottleType = 5
	ThrottleType_READ_SIZE      ThrottleType = 6
)

var ThrottleType_name = map[int32]string{
	1: "REQUEST_NUMBER",
	2: "REQUEST_SIZE",
	3: "WRITE_NUMBER",
	4: "WRITE_SIZE",
	5: "READ_NUMBER",
	6: "READ_SIZE",
}
var ThrottleType_value = map[string]int32{
	"REQUEST_NUMBER": 1,
	"REQUEST_SIZE":   2,
	"WRITE_NUMBER":   3,
	"WRITE_SIZE":     4,
	"READ_NUMBER":    5,
	"READ_SIZE":      6,
}

func (x ThrottleType) Enum() *ThrottleType {
	p := new(ThrottleType)
	*p = x
	return p
}
func (x ThrottleType) String() string {
	return proto.EnumName(ThrottleType_name, int32(x))
}
func (x *ThrottleType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ThrottleType_value, data, "ThrottleType")
	if err != nil {
		return err
	}
	*x = ThrottleType(value)
	return nil
}

type QuotaType int32

const (
	QuotaType_THROTTLE QuotaType = 1
)

var QuotaType_name = map[int32]string{
	1: "THROTTLE",
}
var QuotaType_value = map[string]int32{
	"THROTTLE": 1,
}

func (x QuotaType) Enum() *QuotaType {
	p := new(QuotaType)
	*p = x
	return p
}
func (x QuotaType) String() string {
	return proto.EnumName(QuotaType_name, int32(x))
}
func (x *QuotaType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(QuotaType_value, data, "QuotaType")
	if err != nil {
		return err
	}
	*x = QuotaType(value)
	return nil
}

type TimedQuota struct {
	TimeUnit         *TimeUnit   `protobuf:"varint,1,req,name=time_unit,enum=pb.TimeUnit" json:"time_unit,omitempty"`
	SoftLimit        *uint64     `protobuf:"varint,2,opt,name=soft_limit" json:"soft_limit,omitempty"`
	Share            *float32    `protobuf:"fixed32,3,opt,name=share" json:"share,omitempty"`
	Scope            *QuotaScope `protobuf:"varint,4,opt,name=scope,enum=pb.QuotaScope,def=2" json:"scope,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *TimedQuota) Reset()         { *m = TimedQuota{} }
func (m *TimedQuota) String() string { return proto.CompactTextString(m) }
func (*TimedQuota) ProtoMessage()    {}

const Default_TimedQuota_Scope QuotaScope = QuotaScope_MACHINE

func (m *TimedQuota) GetTimeUnit() TimeUnit {
	if m != nil && m.TimeUnit != nil {
		return *m.TimeUnit
	}
	return TimeUnit_NANOSECONDS
}

func (m *TimedQuota) GetSoftLimit() uint64 {
	if m != nil && m.SoftLimit != nil {
		return *m.SoftLimit
	}
	return 0
}

func (m *TimedQuota) GetShare() float32 {
	if m != nil && m.Share != nil {
		return *m.Share
	}
	return 0
}

func (m *TimedQuota) GetScope() QuotaScope {
	if m != nil && m.Scope != nil {
		return *m.Scope
	}
	return Default_TimedQuota_Scope
}

type Throttle struct {
	ReqNum           *TimedQuota `protobuf:"bytes,1,opt,name=req_num" json:"req_num,omitempty"`
	ReqSize          *TimedQuota `protobuf:"bytes,2,opt,name=req_size" json:"req_size,omitempty"`
	WriteNum         *TimedQuota `protobuf:"bytes,3,opt,name=write_num" json:"write_num,omitempty"`
	WriteSize        *TimedQuota `protobuf:"bytes,4,opt,name=write_size" json:"write_size,omitempty"`
	ReadNum          *TimedQuota `protobuf:"bytes,5,opt,name=read_num" json:"read_num,omitempty"`
	ReadSize         *TimedQuota `protobuf:"bytes,6,opt,name=read_size" json:"read_size,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Throttle) Reset()         { *m = Throttle{} }
func (m *Throttle) String() string { return proto.CompactTextString(m) }
func (*Throttle) ProtoMessage()    {}

func (m *Throttle) GetReqNum() *TimedQuota {
	if m != nil {
		return m.ReqNum
	}
	return nil
}

func (m *Throttle) GetReqSize() *TimedQuota {
	if m != nil {
		return m.ReqSize
	}
	return nil
}

func (m *Throttle) GetWriteNum() *TimedQuota {
	if m != nil {
		return m.WriteNum
	}
	return nil
}

func (m *Throttle) GetWriteSize() *TimedQuota {
	if m != nil {
		return m.WriteSize
	}
	return nil
}

func (m *Throttle) GetReadNum() *TimedQuota {
	if m != nil {
		return m.ReadNum
	}
	return nil
}

func (m *Throttle) GetReadSize() *TimedQuota {
	if m != nil {
		return m.ReadSize
	}
	return nil
}

type ThrottleRequest struct {
	Type             *ThrottleType `protobuf:"varint,1,opt,name=type,enum=pb.ThrottleType" json:"type,omitempty"`
	TimedQuota       *TimedQuota   `protobuf:"bytes,2,opt,name=timed_quota" json:"timed_quota,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *ThrottleRequest) Reset()         { *m = ThrottleRequest{} }
func (m *ThrottleRequest) String() string { return proto.CompactTextString(m) }
func (*ThrottleRequest) ProtoMessage()    {}

func (m *ThrottleRequest) GetType() ThrottleType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ThrottleType_REQUEST_NUMBER
}

func (m *ThrottleRequest) GetTimedQuota() *TimedQuota {
	if m != nil {
		return m.TimedQuota
	}
	return nil
}

type Quotas struct {
	BypassGlobals    *bool     `protobuf:"varint,1,opt,name=bypass_globals,def=0" json:"bypass_globals,omitempty"`
	Throttle         *Throttle `protobuf:"bytes,2,opt,name=throttle" json:"throttle,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Quotas) Reset()         { *m = Quotas{} }
func (m *Quotas) String() string { return proto.CompactTextString(m) }
func (*Quotas) ProtoMessage()    {}

const Default_Quotas_BypassGlobals bool = false

func (m *Quotas) GetBypassGlobals() bool {
	if m != nil && m.BypassGlobals != nil {
		return *m.BypassGlobals
	}
	return Default_Quotas_BypassGlobals
}

func (m *Quotas) GetThrottle() *Throttle {
	if m != nil {
		return m.Throttle
	}
	return nil
}

type QuotaUsage struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *QuotaUsage) Reset()         { *m = QuotaUsage{} }
func (m *QuotaUsage) String() string { return proto.CompactTextString(m) }
func (*QuotaUsage) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("pb.QuotaScope", QuotaScope_name, QuotaScope_value)
	proto.RegisterEnum("pb.ThrottleType", ThrottleType_name, ThrottleType_value)
	proto.RegisterEnum("pb.QuotaType", QuotaType_name, QuotaType_value)
}