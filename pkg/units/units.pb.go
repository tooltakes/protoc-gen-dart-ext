// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/units/units.proto

package units // import "github.com/empirefox/protoc-gen-dart-ext/pkg/units"

/*
https://physics.nist.gov/cuu/Units/units.html
*/

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AtomType int32

const (
	AtomType_symbol   AtomType = 0
	AtomType_singular AtomType = 1
	AtomType_plural   AtomType = 2
	AtomType_parse    AtomType = 3
)

var AtomType_name = map[int32]string{
	0: "symbol",
	1: "singular",
	2: "plural",
	3: "parse",
}
var AtomType_value = map[string]int32{
	"symbol":   0,
	"singular": 1,
	"plural":   2,
	"parse":    3,
}

func (x AtomType) String() string {
	return proto.EnumName(AtomType_name, int32(x))
}
func (AtomType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_units_6b7eddbd503602da, []int{0}
}

type Unit_PrefixType int32

const (
	Unit_symbol Unit_PrefixType = 0
	Unit_name   Unit_PrefixType = 1
)

var Unit_PrefixType_name = map[int32]string{
	0: "symbol",
	1: "name",
}
var Unit_PrefixType_value = map[string]int32{
	"symbol": 0,
	"name":   1,
}

func (x Unit_PrefixType) String() string {
	return proto.EnumName(Unit_PrefixType_name, int32(x))
}
func (Unit_PrefixType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_units_6b7eddbd503602da, []int{1, 0}
}

type CurrencyUnit struct {
	// Types that are valid to be assigned to Type:
	//	*CurrencyUnit_Code
	//	*CurrencyUnit_Symbol
	//	*CurrencyUnit_Name
	Type isCurrencyUnit_Type `protobuf_oneof:"type"`
	// divided before format
	DividedBy int32 `protobuf:"varint,4,opt,name=dividedBy,proto3" json:"dividedBy,omitempty"`
	// custom format name, which must be provided by user.
	Format               string   `protobuf:"bytes,5,opt,name=format,proto3" json:"format,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CurrencyUnit) Reset()         { *m = CurrencyUnit{} }
func (m *CurrencyUnit) String() string { return proto.CompactTextString(m) }
func (*CurrencyUnit) ProtoMessage()    {}
func (*CurrencyUnit) Descriptor() ([]byte, []int) {
	return fileDescriptor_units_6b7eddbd503602da, []int{0}
}
func (m *CurrencyUnit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrencyUnit.Unmarshal(m, b)
}
func (m *CurrencyUnit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrencyUnit.Marshal(b, m, deterministic)
}
func (dst *CurrencyUnit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrencyUnit.Merge(dst, src)
}
func (m *CurrencyUnit) XXX_Size() int {
	return xxx_messageInfo_CurrencyUnit.Size(m)
}
func (m *CurrencyUnit) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrencyUnit.DiscardUnknown(m)
}

var xxx_messageInfo_CurrencyUnit proto.InternalMessageInfo

type isCurrencyUnit_Type interface {
	isCurrencyUnit_Type()
}

type CurrencyUnit_Code struct {
	Code CurrencyV1 `protobuf:"varint,1,opt,name=code,proto3,enum=units.CurrencyV1,oneof"`
}
type CurrencyUnit_Symbol struct {
	Symbol CurrencyV1 `protobuf:"varint,2,opt,name=symbol,proto3,enum=units.CurrencyV1,oneof"`
}
type CurrencyUnit_Name struct {
	Name CurrencyV1 `protobuf:"varint,3,opt,name=name,proto3,enum=units.CurrencyV1,oneof"`
}

func (*CurrencyUnit_Code) isCurrencyUnit_Type()   {}
func (*CurrencyUnit_Symbol) isCurrencyUnit_Type() {}
func (*CurrencyUnit_Name) isCurrencyUnit_Type()   {}

func (m *CurrencyUnit) GetType() isCurrencyUnit_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *CurrencyUnit) GetCode() CurrencyV1 {
	if x, ok := m.GetType().(*CurrencyUnit_Code); ok {
		return x.Code
	}
	return CurrencyV1_XXX
}

func (m *CurrencyUnit) GetSymbol() CurrencyV1 {
	if x, ok := m.GetType().(*CurrencyUnit_Symbol); ok {
		return x.Symbol
	}
	return CurrencyV1_XXX
}

func (m *CurrencyUnit) GetName() CurrencyV1 {
	if x, ok := m.GetType().(*CurrencyUnit_Name); ok {
		return x.Name
	}
	return CurrencyV1_XXX
}

func (m *CurrencyUnit) GetDividedBy() int32 {
	if m != nil {
		return m.DividedBy
	}
	return 0
}

func (m *CurrencyUnit) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CurrencyUnit) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CurrencyUnit_OneofMarshaler, _CurrencyUnit_OneofUnmarshaler, _CurrencyUnit_OneofSizer, []interface{}{
		(*CurrencyUnit_Code)(nil),
		(*CurrencyUnit_Symbol)(nil),
		(*CurrencyUnit_Name)(nil),
	}
}

func _CurrencyUnit_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CurrencyUnit)
	// type
	switch x := m.Type.(type) {
	case *CurrencyUnit_Code:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Code))
	case *CurrencyUnit_Symbol:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Symbol))
	case *CurrencyUnit_Name:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Name))
	case nil:
	default:
		return fmt.Errorf("CurrencyUnit.Type has unexpected type %T", x)
	}
	return nil
}

func _CurrencyUnit_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CurrencyUnit)
	switch tag {
	case 1: // type.code
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Type = &CurrencyUnit_Code{CurrencyV1(x)}
		return true, err
	case 2: // type.symbol
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Type = &CurrencyUnit_Symbol{CurrencyV1(x)}
		return true, err
	case 3: // type.name
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Type = &CurrencyUnit_Name{CurrencyV1(x)}
		return true, err
	default:
		return false, nil
	}
}

func _CurrencyUnit_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CurrencyUnit)
	// type
	switch x := m.Type.(type) {
	case *CurrencyUnit_Code:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Code))
	case *CurrencyUnit_Symbol:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Symbol))
	case *CurrencyUnit_Name:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Name))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Unit struct {
	PrefixType           Unit_PrefixType `protobuf:"varint,1,opt,name=prefixType,proto3,enum=units.Unit_PrefixType" json:"prefixType,omitempty"`
	AtomType             AtomType        `protobuf:"varint,2,opt,name=atomType,proto3,enum=units.AtomType" json:"atomType,omitempty"`
	Dots                 []*Cell         `protobuf:"bytes,3,rep,name=dots,proto3" json:"dots,omitempty"`
	Per                  []*Cell         `protobuf:"bytes,4,rep,name=per,proto3" json:"per,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Unit) Reset()         { *m = Unit{} }
func (m *Unit) String() string { return proto.CompactTextString(m) }
func (*Unit) ProtoMessage()    {}
func (*Unit) Descriptor() ([]byte, []int) {
	return fileDescriptor_units_6b7eddbd503602da, []int{1}
}
func (m *Unit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Unit.Unmarshal(m, b)
}
func (m *Unit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Unit.Marshal(b, m, deterministic)
}
func (dst *Unit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Unit.Merge(dst, src)
}
func (m *Unit) XXX_Size() int {
	return xxx_messageInfo_Unit.Size(m)
}
func (m *Unit) XXX_DiscardUnknown() {
	xxx_messageInfo_Unit.DiscardUnknown(m)
}

var xxx_messageInfo_Unit proto.InternalMessageInfo

func (m *Unit) GetPrefixType() Unit_PrefixType {
	if m != nil {
		return m.PrefixType
	}
	return Unit_symbol
}

func (m *Unit) GetAtomType() AtomType {
	if m != nil {
		return m.AtomType
	}
	return AtomType_symbol
}

func (m *Unit) GetDots() []*Cell {
	if m != nil {
		return m.Dots
	}
	return nil
}

func (m *Unit) GetPer() []*Cell {
	if m != nil {
		return m.Per
	}
	return nil
}

type Cell struct {
	Exponent int32    `protobuf:"varint,1,opt,name=exponent,proto3" json:"exponent,omitempty"`
	Prefix   PrefixV1 `protobuf:"varint,2,opt,name=prefix,proto3,enum=units.PrefixV1" json:"prefix,omitempty"`
	// Types that are valid to be assigned to Unit:
	//	*Cell_Atom
	//	*Cell_Symbol
	Unit                 isCell_Unit `protobuf_oneof:"unit"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Cell) Reset()         { *m = Cell{} }
func (m *Cell) String() string { return proto.CompactTextString(m) }
func (*Cell) ProtoMessage()    {}
func (*Cell) Descriptor() ([]byte, []int) {
	return fileDescriptor_units_6b7eddbd503602da, []int{2}
}
func (m *Cell) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cell.Unmarshal(m, b)
}
func (m *Cell) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cell.Marshal(b, m, deterministic)
}
func (dst *Cell) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cell.Merge(dst, src)
}
func (m *Cell) XXX_Size() int {
	return xxx_messageInfo_Cell.Size(m)
}
func (m *Cell) XXX_DiscardUnknown() {
	xxx_messageInfo_Cell.DiscardUnknown(m)
}

var xxx_messageInfo_Cell proto.InternalMessageInfo

type isCell_Unit interface {
	isCell_Unit()
}

type Cell_Atom struct {
	Atom AtomV1 `protobuf:"varint,3,opt,name=atom,proto3,enum=units.AtomV1,oneof"`
}
type Cell_Symbol struct {
	Symbol string `protobuf:"bytes,4,opt,name=symbol,proto3,oneof"`
}

func (*Cell_Atom) isCell_Unit()   {}
func (*Cell_Symbol) isCell_Unit() {}

func (m *Cell) GetUnit() isCell_Unit {
	if m != nil {
		return m.Unit
	}
	return nil
}

func (m *Cell) GetExponent() int32 {
	if m != nil {
		return m.Exponent
	}
	return 0
}

func (m *Cell) GetPrefix() PrefixV1 {
	if m != nil {
		return m.Prefix
	}
	return PrefixV1_noPrefix
}

func (m *Cell) GetAtom() AtomV1 {
	if x, ok := m.GetUnit().(*Cell_Atom); ok {
		return x.Atom
	}
	return AtomV1_noAtom
}

func (m *Cell) GetSymbol() string {
	if x, ok := m.GetUnit().(*Cell_Symbol); ok {
		return x.Symbol
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Cell) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Cell_OneofMarshaler, _Cell_OneofUnmarshaler, _Cell_OneofSizer, []interface{}{
		(*Cell_Atom)(nil),
		(*Cell_Symbol)(nil),
	}
}

func _Cell_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Cell)
	// unit
	switch x := m.Unit.(type) {
	case *Cell_Atom:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Atom))
	case *Cell_Symbol:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Symbol)
	case nil:
	default:
		return fmt.Errorf("Cell.Unit has unexpected type %T", x)
	}
	return nil
}

func _Cell_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Cell)
	switch tag {
	case 3: // unit.atom
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Unit = &Cell_Atom{AtomV1(x)}
		return true, err
	case 4: // unit.symbol
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Unit = &Cell_Symbol{x}
		return true, err
	default:
		return false, nil
	}
}

func _Cell_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Cell)
	// unit
	switch x := m.Unit.(type) {
	case *Cell_Atom:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Atom))
	case *Cell_Symbol:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Symbol)))
		n += len(x.Symbol)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

var E_Currency = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*CurrencyUnit)(nil),
	Field:         919111,
	Name:          "units.currency",
	Tag:           "bytes,919111,opt,name=currency",
	Filename:      "protos/units/units.proto",
}

var E_Unit = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*Unit)(nil),
	Field:         919112,
	Name:          "units.unit",
	Tag:           "bytes,919112,opt,name=unit",
	Filename:      "protos/units/units.proto",
}

func init() {
	proto.RegisterType((*CurrencyUnit)(nil), "units.CurrencyUnit")
	proto.RegisterType((*Unit)(nil), "units.Unit")
	proto.RegisterType((*Cell)(nil), "units.Cell")
	proto.RegisterEnum("units.AtomType", AtomType_name, AtomType_value)
	proto.RegisterEnum("units.Unit_PrefixType", Unit_PrefixType_name, Unit_PrefixType_value)
	proto.RegisterExtension(E_Currency)
	proto.RegisterExtension(E_Unit)
}

func init() { proto.RegisterFile("protos/units/units.proto", fileDescriptor_units_6b7eddbd503602da) }

var fileDescriptor_units_6b7eddbd503602da = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0x1b, 0x27, 0x72, 0x26, 0x05, 0xc2, 0x22, 0x15, 0x53, 0xa8, 0xa8, 0xc2, 0x81, 0x0a,
	0x14, 0x5b, 0x0d, 0x08, 0xa1, 0x72, 0x81, 0x20, 0x21, 0x6e, 0x54, 0x2b, 0xe8, 0x81, 0x9b, 0x63,
	0x6f, 0xcc, 0x0a, 0xdb, 0xbb, 0x5a, 0xaf, 0x51, 0xf2, 0x1f, 0x38, 0xf1, 0xa7, 0x40, 0xe2, 0xc0,
	0x5f, 0x62, 0xbf, 0xec, 0x24, 0x45, 0xc0, 0x25, 0xd2, 0xce, 0x7b, 0xf3, 0x32, 0xef, 0xcd, 0x18,
	0x42, 0x2e, 0x98, 0x64, 0x75, 0xdc, 0x54, 0x54, 0xba, 0xdf, 0xc8, 0x94, 0x50, 0xdf, 0x3c, 0x8e,
	0x4e, 0x72, 0xc6, 0xf2, 0x82, 0xc4, 0xa6, 0xb8, 0x68, 0x96, 0x71, 0x46, 0xea, 0x54, 0x50, 0x2e,
	0x99, 0xb0, 0xc4, 0xa3, 0xbb, 0x3b, 0x12, 0x69, 0x23, 0x04, 0xa9, 0xd2, 0xb5, 0x03, 0xef, 0xec,
	0x80, 0x5c, 0x90, 0x25, 0x5d, 0x39, 0xe8, 0xf6, 0x0e, 0x94, 0x48, 0x56, 0x5a, 0x60, 0xf2, 0xd3,
	0x83, 0x83, 0xd7, 0x4e, 0xe6, 0x83, 0x02, 0xd1, 0x43, 0xf0, 0x53, 0x96, 0x91, 0xd0, 0x3b, 0xf1,
	0x4e, 0xaf, 0xcf, 0x6e, 0x46, 0x76, 0xcc, 0x96, 0x72, 0x79, 0xf6, 0x76, 0x0f, 0x1b, 0x02, 0x7a,
	0x0c, 0x83, 0x7a, 0x5d, 0x2e, 0x58, 0x11, 0xee, 0xff, 0x9d, 0xea, 0x28, 0x5a, 0xb5, 0x4a, 0x4a,
	0x12, 0xf6, 0xfe, 0xa1, 0xaa, 0x09, 0xe8, 0x1e, 0x0c, 0x33, 0xfa, 0x85, 0x66, 0x24, 0x9b, 0xaf,
	0x43, 0x5f, 0xb1, 0xfb, 0x78, 0x53, 0x40, 0x87, 0x30, 0x58, 0x32, 0x51, 0x26, 0x32, 0xec, 0x2b,
	0x68, 0x88, 0xdd, 0x6b, 0x3e, 0x00, 0x5f, 0xae, 0x39, 0x99, 0xfc, 0xf2, 0xc0, 0x37, 0x2e, 0x9e,
	0x01, 0x58, 0xff, 0xef, 0x55, 0xd9, 0x79, 0x39, 0x74, 0xff, 0xaa, 0x09, 0xd1, 0x45, 0x87, 0xe2,
	0x2d, 0xa6, 0x32, 0x15, 0xe8, 0x70, 0x4c, 0x97, 0xb5, 0x75, 0xc3, 0x75, 0xbd, 0x72, 0x65, 0xdc,
	0x11, 0xd0, 0x7d, 0xf0, 0x33, 0x26, 0x6b, 0x65, 0xaa, 0x77, 0x3a, 0x9a, 0x8d, 0x5a, 0x53, 0xa4,
	0x28, 0xb0, 0x01, 0xd0, 0x31, 0xf4, 0x38, 0x11, 0xca, 0xc6, 0x1f, 0xb8, 0xae, 0x4f, 0x26, 0x00,
	0x9b, 0x31, 0x10, 0xb4, 0x79, 0x8e, 0xf7, 0x50, 0x60, 0xe3, 0x1a, 0x7b, 0x93, 0x6f, 0xca, 0x91,
	0xee, 0x40, 0x47, 0x10, 0x90, 0x15, 0x67, 0x15, 0xa9, 0xa4, 0xf1, 0xd3, 0xc7, 0xdd, 0x5b, 0xa5,
	0x3b, 0xb0, 0x1e, 0xae, 0xcc, 0x6c, 0xd5, 0x2f, 0xcf, 0xb0, 0x83, 0xd1, 0x03, 0xf0, 0xf5, 0xf4,
	0x6e, 0x0d, 0xd7, 0xb6, 0xac, 0xd9, 0x15, 0x68, 0x10, 0x85, 0xdd, 0x62, 0x75, 0xfe, 0xc3, 0xcd,
	0x16, 0x75, 0xcc, 0xba, 0xe3, 0xd1, 0x0b, 0x08, 0xda, 0x38, 0x76, 0xc6, 0x3e, 0x80, 0xa0, 0xa6,
	0x55, 0xde, 0x14, 0x89, 0x18, 0x7b, 0x1a, 0xe1, 0x45, 0x23, 0x92, 0x62, 0xbc, 0x8f, 0x86, 0xd0,
	0xe7, 0x89, 0xa8, 0xc9, 0xb8, 0x77, 0x7e, 0x01, 0x41, 0x7b, 0xb7, 0xe8, 0x38, 0xb2, 0x17, 0x1f,
	0xb5, 0x17, 0x1f, 0xbd, 0xa1, 0xa4, 0xc8, 0xde, 0x71, 0x49, 0x59, 0x55, 0x87, 0xdf, 0xbf, 0x3e,
	0x57, 0x13, 0x8c, 0x66, 0xb7, 0xae, 0xdc, 0x8b, 0xde, 0x20, 0xee, 0x54, 0xce, 0x5f, 0xda, 0xb1,
	0xfe, 0xa7, 0xf6, 0xc3, 0xa9, 0x8d, 0xb6, 0xee, 0x00, 0x9b, 0xce, 0xf9, 0xd3, 0x8f, 0xb3, 0x9c,
	0xca, 0x4f, 0xcd, 0x22, 0x4a, 0x59, 0x19, 0x93, 0x92, 0x53, 0x95, 0x17, 0x5b, 0xd9, 0x0f, 0x31,
	0x9d, 0xe6, 0xa4, 0x9a, 0x66, 0x89, 0x90, 0x53, 0xb2, 0x92, 0x31, 0xff, 0x9c, 0xdb, 0xcf, 0x68,
	0x31, 0x30, 0xf0, 0x93, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1a, 0xb9, 0xfe, 0xb7, 0xd8, 0x03,
	0x00, 0x00,
}
