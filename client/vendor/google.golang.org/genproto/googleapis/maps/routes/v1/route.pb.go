// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/maps/routes/v1/route.proto

package routes

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	viewport "google.golang.org/genproto/googleapis/geo/type/viewport"
	money "google.golang.org/genproto/googleapis/type/money"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A set of values that specify the navigation action to take for the current
// step (e.g., turn left, merge, straight, etc.).
type Maneuver int32

const (
	// Not used.
	Maneuver_MANEUVER_UNSPECIFIED Maneuver = 0
	// Turn slightly to the left.
	Maneuver_TURN_SLIGHT_LEFT Maneuver = 1
	// Turn sharply to the left.
	Maneuver_TURN_SHARP_LEFT Maneuver = 2
	// Make a left u-turn.
	Maneuver_UTURN_LEFT Maneuver = 3
	// Turn left.
	Maneuver_TURN_LEFT Maneuver = 4
	// Turn slightly to the right.
	Maneuver_TURN_SLIGHT_RIGHT Maneuver = 5
	// Turn sharply to the right.
	Maneuver_TURN_SHARP_RIGHT Maneuver = 6
	// Make a right u-turn.
	Maneuver_UTURN_RIGHT Maneuver = 7
	// Turn right.
	Maneuver_TURN_RIGHT Maneuver = 8
	// Go straight.
	Maneuver_STRAIGHT Maneuver = 9
	// Take the left ramp.
	Maneuver_RAMP_LEFT Maneuver = 10
	// Take the right ramp.
	Maneuver_RAMP_RIGHT Maneuver = 11
	// Merge into traffic.
	Maneuver_MERGE Maneuver = 12
	// Take the left fork.
	Maneuver_FORK_LEFT Maneuver = 13
	// Take the right fork.
	Maneuver_FORK_RIGHT Maneuver = 14
	// Take the ferry.
	Maneuver_FERRY Maneuver = 15
	// Take the train leading onto the ferry.
	Maneuver_FERRY_TRAIN Maneuver = 16
	// Turn left at the roundabout.
	Maneuver_ROUNDABOUT_LEFT Maneuver = 17
	// Turn right at the roundabout.
	Maneuver_ROUNDABOUT_RIGHT Maneuver = 18
)

var Maneuver_name = map[int32]string{
	0:  "MANEUVER_UNSPECIFIED",
	1:  "TURN_SLIGHT_LEFT",
	2:  "TURN_SHARP_LEFT",
	3:  "UTURN_LEFT",
	4:  "TURN_LEFT",
	5:  "TURN_SLIGHT_RIGHT",
	6:  "TURN_SHARP_RIGHT",
	7:  "UTURN_RIGHT",
	8:  "TURN_RIGHT",
	9:  "STRAIGHT",
	10: "RAMP_LEFT",
	11: "RAMP_RIGHT",
	12: "MERGE",
	13: "FORK_LEFT",
	14: "FORK_RIGHT",
	15: "FERRY",
	16: "FERRY_TRAIN",
	17: "ROUNDABOUT_LEFT",
	18: "ROUNDABOUT_RIGHT",
}

var Maneuver_value = map[string]int32{
	"MANEUVER_UNSPECIFIED": 0,
	"TURN_SLIGHT_LEFT":     1,
	"TURN_SHARP_LEFT":      2,
	"UTURN_LEFT":           3,
	"TURN_LEFT":            4,
	"TURN_SLIGHT_RIGHT":    5,
	"TURN_SHARP_RIGHT":     6,
	"UTURN_RIGHT":          7,
	"TURN_RIGHT":           8,
	"STRAIGHT":             9,
	"RAMP_LEFT":            10,
	"RAMP_RIGHT":           11,
	"MERGE":                12,
	"FORK_LEFT":            13,
	"FORK_RIGHT":           14,
	"FERRY":                15,
	"FERRY_TRAIN":          16,
	"ROUNDABOUT_LEFT":      17,
	"ROUNDABOUT_RIGHT":     18,
}

func (x Maneuver) String() string {
	return proto.EnumName(Maneuver_name, int32(x))
}

func (Maneuver) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_294c3406019f26b4, []int{0}
}

// Encapsulates a route, which consists of a series of connected road segments
// that join beginning, ending, and intermediate waypoints.
type Route struct {
	// A collection of legs (path segments between waypoints) that make-up the
	// route. Each leg corresponds to the trip between two non-`via` Waypoints.
	// For example, a route with no intermediate waypoints has only one leg. A
	// route that includes one non-`via` intermediate waypoint has two legs. A
	// route that includes one `via` intermediate waypoint has one leg. The order
	// of the legs matches the order of Waypoints from `origin` to `intermediates`
	// to `destination`.
	Legs []*RouteLeg `protobuf:"bytes,1,rep,name=legs,proto3" json:"legs,omitempty"`
	// The travel distance of the route, in meters.
	DistanceMeters int32 `protobuf:"varint,2,opt,name=distance_meters,json=distanceMeters,proto3" json:"distance_meters,omitempty"`
	// The length of time needed to navigate the route. If you set the
	// `route_preference` to `TRAFFIC_UNAWARE`, then this value is the same as
	// `static_duration`. If you set the `route_preference` to either
	// `TRAFFIC_AWARE` or `TRAFFIC_AWARE_OPTIMAL`, then this value is calculated
	// taking traffic conditions into account.
	Duration *duration.Duration `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
	// The duration of traveling through the route without taking traffic
	// conditions into consideration.
	StaticDuration *duration.Duration `protobuf:"bytes,4,opt,name=static_duration,json=staticDuration,proto3" json:"static_duration,omitempty"`
	// The overall route polyline. This polyline will be the combined polyline of
	// all `legs`.
	Polyline *Polyline `protobuf:"bytes,5,opt,name=polyline,proto3" json:"polyline,omitempty"`
	// A description of the route.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// An array of warnings to show when displaying the route.
	Warnings []string `protobuf:"bytes,7,rep,name=warnings,proto3" json:"warnings,omitempty"`
	// The viewport bounding box of the polyline.
	Viewport *viewport.Viewport `protobuf:"bytes,8,opt,name=viewport,proto3" json:"viewport,omitempty"`
	// Additional information about the route.
	TravelAdvisory       *RouteTravelAdvisory `protobuf:"bytes,9,opt,name=travel_advisory,json=travelAdvisory,proto3" json:"travel_advisory,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_294c3406019f26b4, []int{0}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetLegs() []*RouteLeg {
	if m != nil {
		return m.Legs
	}
	return nil
}

func (m *Route) GetDistanceMeters() int32 {
	if m != nil {
		return m.DistanceMeters
	}
	return 0
}

func (m *Route) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *Route) GetStaticDuration() *duration.Duration {
	if m != nil {
		return m.StaticDuration
	}
	return nil
}

func (m *Route) GetPolyline() *Polyline {
	if m != nil {
		return m.Polyline
	}
	return nil
}

func (m *Route) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Route) GetWarnings() []string {
	if m != nil {
		return m.Warnings
	}
	return nil
}

func (m *Route) GetViewport() *viewport.Viewport {
	if m != nil {
		return m.Viewport
	}
	return nil
}

func (m *Route) GetTravelAdvisory() *RouteTravelAdvisory {
	if m != nil {
		return m.TravelAdvisory
	}
	return nil
}

// Encapsulates the additional information that the user should be informed
// about, such as possible traffic zone restriction etc.
type RouteTravelAdvisory struct {
	// The traffic restriction that applies to the route. A vehicle that is
	// subject to the restriction is not allowed to travel on the route. As of
	// October 2019, only Jakarta, Indonesia takes into consideration.
	TrafficRestriction *TrafficRestriction `protobuf:"bytes,1,opt,name=traffic_restriction,json=trafficRestriction,proto3" json:"traffic_restriction,omitempty"`
	// Encapsulates information about tolls on the Route.
	// This field is only populated if we expect there are tolls on the Route.
	// If this field is set but the estimated_price subfield is not populated,
	// we expect that road contains tolls but we do not know an estimated price.
	// If this field is not set, then we expect there is no toll on the Route.
	TollInfo             *TollInfo `protobuf:"bytes,2,opt,name=toll_info,json=tollInfo,proto3" json:"toll_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RouteTravelAdvisory) Reset()         { *m = RouteTravelAdvisory{} }
func (m *RouteTravelAdvisory) String() string { return proto.CompactTextString(m) }
func (*RouteTravelAdvisory) ProtoMessage()    {}
func (*RouteTravelAdvisory) Descriptor() ([]byte, []int) {
	return fileDescriptor_294c3406019f26b4, []int{1}
}

func (m *RouteTravelAdvisory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteTravelAdvisory.Unmarshal(m, b)
}
func (m *RouteTravelAdvisory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteTravelAdvisory.Marshal(b, m, deterministic)
}
func (m *RouteTravelAdvisory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteTravelAdvisory.Merge(m, src)
}
func (m *RouteTravelAdvisory) XXX_Size() int {
	return xxx_messageInfo_RouteTravelAdvisory.Size(m)
}
func (m *RouteTravelAdvisory) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteTravelAdvisory.DiscardUnknown(m)
}

var xxx_messageInfo_RouteTravelAdvisory proto.InternalMessageInfo

func (m *RouteTravelAdvisory) GetTrafficRestriction() *TrafficRestriction {
	if m != nil {
		return m.TrafficRestriction
	}
	return nil
}

func (m *RouteTravelAdvisory) GetTollInfo() *TollInfo {
	if m != nil {
		return m.TollInfo
	}
	return nil
}

// Encapsulates the additional information that the user should be informed
// about, such as possible traffic zone restriction etc. on a route leg.
type RouteLegTravelAdvisory struct {
	// Encapsulates information about tolls on the specific RouteLeg.
	// This field is only populated if we expect there are tolls on the RouteLeg.
	// If this field is set but the estimated_price subfield is not populated,
	// we expect that road contains tolls but we do not know an estimated price.
	// If this field does not exist, then there is no toll on the RouteLeg.
	TollInfo             *TollInfo `protobuf:"bytes,1,opt,name=toll_info,json=tollInfo,proto3" json:"toll_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RouteLegTravelAdvisory) Reset()         { *m = RouteLegTravelAdvisory{} }
func (m *RouteLegTravelAdvisory) String() string { return proto.CompactTextString(m) }
func (*RouteLegTravelAdvisory) ProtoMessage()    {}
func (*RouteLegTravelAdvisory) Descriptor() ([]byte, []int) {
	return fileDescriptor_294c3406019f26b4, []int{2}
}

func (m *RouteLegTravelAdvisory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteLegTravelAdvisory.Unmarshal(m, b)
}
func (m *RouteLegTravelAdvisory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteLegTravelAdvisory.Marshal(b, m, deterministic)
}
func (m *RouteLegTravelAdvisory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteLegTravelAdvisory.Merge(m, src)
}
func (m *RouteLegTravelAdvisory) XXX_Size() int {
	return xxx_messageInfo_RouteLegTravelAdvisory.Size(m)
}
func (m *RouteLegTravelAdvisory) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteLegTravelAdvisory.DiscardUnknown(m)
}

var xxx_messageInfo_RouteLegTravelAdvisory proto.InternalMessageInfo

func (m *RouteLegTravelAdvisory) GetTollInfo() *TollInfo {
	if m != nil {
		return m.TollInfo
	}
	return nil
}

// Encapsulates the traffic restriction applied to the route. As of October
// 2019, only Jakarta, Indonesia takes into consideration.
type TrafficRestriction struct {
	// The restriction based on the vehicle's license plate last character. If
	// this field does not exist, then no restriction on route.
	LicensePlateLastCharacterRestriction *LicensePlateLastCharacterRestriction `protobuf:"bytes,1,opt,name=license_plate_last_character_restriction,json=licensePlateLastCharacterRestriction,proto3" json:"license_plate_last_character_restriction,omitempty"`
	XXX_NoUnkeyedLiteral                 struct{}                              `json:"-"`
	XXX_unrecognized                     []byte                                `json:"-"`
	XXX_sizecache                        int32                                 `json:"-"`
}

func (m *TrafficRestriction) Reset()         { *m = TrafficRestriction{} }
func (m *TrafficRestriction) String() string { return proto.CompactTextString(m) }
func (*TrafficRestriction) ProtoMessage()    {}
func (*TrafficRestriction) Descriptor() ([]byte, []int) {
	return fileDescriptor_294c3406019f26b4, []int{3}
}

func (m *TrafficRestriction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficRestriction.Unmarshal(m, b)
}
func (m *TrafficRestriction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficRestriction.Marshal(b, m, deterministic)
}
func (m *TrafficRestriction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficRestriction.Merge(m, src)
}
func (m *TrafficRestriction) XXX_Size() int {
	return xxx_messageInfo_TrafficRestriction.Size(m)
}
func (m *TrafficRestriction) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficRestriction.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficRestriction proto.InternalMessageInfo

func (m *TrafficRestriction) GetLicensePlateLastCharacterRestriction() *LicensePlateLastCharacterRestriction {
	if m != nil {
		return m.LicensePlateLastCharacterRestriction
	}
	return nil
}

// Encapsulates the license plate last character restriction.
type LicensePlateLastCharacterRestriction struct {
	// The allowed last character of a license plate of a vehicle. Only vehicles
	// whose license plate's last characters match these are allowed to travel on
	// the route. If empty, no vehicle is allowed.
	AllowedLastCharacters []string `protobuf:"bytes,1,rep,name=allowed_last_characters,json=allowedLastCharacters,proto3" json:"allowed_last_characters,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *LicensePlateLastCharacterRestriction) Reset()         { *m = LicensePlateLastCharacterRestriction{} }
func (m *LicensePlateLastCharacterRestriction) String() string { return proto.CompactTextString(m) }
func (*LicensePlateLastCharacterRestriction) ProtoMessage()    {}
func (*LicensePlateLastCharacterRestriction) Descriptor() ([]byte, []int) {
	return fileDescriptor_294c3406019f26b4, []int{4}
}

func (m *LicensePlateLastCharacterRestriction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LicensePlateLastCharacterRestriction.Unmarshal(m, b)
}
func (m *LicensePlateLastCharacterRestriction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LicensePlateLastCharacterRestriction.Marshal(b, m, deterministic)
}
func (m *LicensePlateLastCharacterRestriction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LicensePlateLastCharacterRestriction.Merge(m, src)
}
func (m *LicensePlateLastCharacterRestriction) XXX_Size() int {
	return xxx_messageInfo_LicensePlateLastCharacterRestriction.Size(m)
}
func (m *LicensePlateLastCharacterRestriction) XXX_DiscardUnknown() {
	xxx_messageInfo_LicensePlateLastCharacterRestriction.DiscardUnknown(m)
}

var xxx_messageInfo_LicensePlateLastCharacterRestriction proto.InternalMessageInfo

func (m *LicensePlateLastCharacterRestriction) GetAllowedLastCharacters() []string {
	if m != nil {
		return m.AllowedLastCharacters
	}
	return nil
}

// Encapsulates a segment between non-`via` waypoints.
type RouteLeg struct {
	// The travel distance of the route leg, in meters.
	DistanceMeters int32 `protobuf:"varint,1,opt,name=distance_meters,json=distanceMeters,proto3" json:"distance_meters,omitempty"`
	// The length of time needed to navigate the leg. If the `route_preference`
	// is set to `TRAFFIC_UNAWARE`, then this value is the same as
	// `static_duration`. If the `route_preference` is either `TRAFFIC_AWARE` or
	// `TRAFFIC_AWARE_OPTIMAL`, then this value is calculated taking traffic
	// conditions into account.
	Duration *duration.Duration `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
	// The duration of traveling through the leg, calculated without taking
	// traffic conditions into consideration.
	StaticDuration *duration.Duration `protobuf:"bytes,3,opt,name=static_duration,json=staticDuration,proto3" json:"static_duration,omitempty"`
	// The overall polyline for this leg. This includes that each `step`'s
	// polyline.
	Polyline *Polyline `protobuf:"bytes,4,opt,name=polyline,proto3" json:"polyline,omitempty"`
	// The start location of this leg. This might be different from the provided
	// `origin`. For example, when the provided `origin` is not near a road, this
	// is a point on the road.
	StartLocation *Location `protobuf:"bytes,5,opt,name=start_location,json=startLocation,proto3" json:"start_location,omitempty"`
	// The end location of this leg. This might be different from the provided
	// `destination`. For example, when the provided `destination` is not near a
	// road, this is a point on the road.
	EndLocation *Location `protobuf:"bytes,6,opt,name=end_location,json=endLocation,proto3" json:"end_location,omitempty"`
	// An array of steps denoting segments within this leg. Each step represents
	// one navigation instruction.
	Steps []*RouteLegStep `protobuf:"bytes,7,rep,name=steps,proto3" json:"steps,omitempty"`
	// Encapsulates the additional information that the user should be informed
	// about, such as possible traffic zone restriction etc. on a route leg.
	TravelAdvisory       *RouteLegTravelAdvisory `protobuf:"bytes,8,opt,name=travel_advisory,json=travelAdvisory,proto3" json:"travel_advisory,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *RouteLeg) Reset()         { *m = RouteLeg{} }
func (m *RouteLeg) String() string { return proto.CompactTextString(m) }
func (*RouteLeg) ProtoMessage()    {}
func (*RouteLeg) Descriptor() ([]byte, []int) {
	return fileDescriptor_294c3406019f26b4, []int{5}
}

func (m *RouteLeg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteLeg.Unmarshal(m, b)
}
func (m *RouteLeg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteLeg.Marshal(b, m, deterministic)
}
func (m *RouteLeg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteLeg.Merge(m, src)
}
func (m *RouteLeg) XXX_Size() int {
	return xxx_messageInfo_RouteLeg.Size(m)
}
func (m *RouteLeg) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteLeg.DiscardUnknown(m)
}

var xxx_messageInfo_RouteLeg proto.InternalMessageInfo

func (m *RouteLeg) GetDistanceMeters() int32 {
	if m != nil {
		return m.DistanceMeters
	}
	return 0
}

func (m *RouteLeg) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *RouteLeg) GetStaticDuration() *duration.Duration {
	if m != nil {
		return m.StaticDuration
	}
	return nil
}

func (m *RouteLeg) GetPolyline() *Polyline {
	if m != nil {
		return m.Polyline
	}
	return nil
}

func (m *RouteLeg) GetStartLocation() *Location {
	if m != nil {
		return m.StartLocation
	}
	return nil
}

func (m *RouteLeg) GetEndLocation() *Location {
	if m != nil {
		return m.EndLocation
	}
	return nil
}

func (m *RouteLeg) GetSteps() []*RouteLegStep {
	if m != nil {
		return m.Steps
	}
	return nil
}

func (m *RouteLeg) GetTravelAdvisory() *RouteLegTravelAdvisory {
	if m != nil {
		return m.TravelAdvisory
	}
	return nil
}

// Encapsulates toll information on a `Route` or on a `RouteLeg`.
type TollInfo struct {
	// The monetary amount of tolls for the corresponding Route or RouteLeg.
	// This list contains a money amount for each currency that is expected
	// to be charged by the toll stations. Typically this list will contain only
	// one item for routes with tolls in one currency. For international trips,
	// this list may contain multiple items to reflect tolls in different
	// currencies.
	EstimatedPrice       []*money.Money `protobuf:"bytes,1,rep,name=estimated_price,json=estimatedPrice,proto3" json:"estimated_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TollInfo) Reset()         { *m = TollInfo{} }
func (m *TollInfo) String() string { return proto.CompactTextString(m) }
func (*TollInfo) ProtoMessage()    {}
func (*TollInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_294c3406019f26b4, []int{6}
}

func (m *TollInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TollInfo.Unmarshal(m, b)
}
func (m *TollInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TollInfo.Marshal(b, m, deterministic)
}
func (m *TollInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TollInfo.Merge(m, src)
}
func (m *TollInfo) XXX_Size() int {
	return xxx_messageInfo_TollInfo.Size(m)
}
func (m *TollInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TollInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TollInfo proto.InternalMessageInfo

func (m *TollInfo) GetEstimatedPrice() []*money.Money {
	if m != nil {
		return m.EstimatedPrice
	}
	return nil
}

// Encapsulates a segment of a `RouteLeg`. A step corresponds to a single
// navigation instruction. Route legs are made up of steps.
type RouteLegStep struct {
	// The travel distance of this step, in meters. In some circumstances, this
	// field might not have a value.
	DistanceMeters int32 `protobuf:"varint,1,opt,name=distance_meters,json=distanceMeters,proto3" json:"distance_meters,omitempty"`
	// The duration of travel through this step without taking traffic conditions
	// into consideration. In some circumstances, this field might not have a
	// value.
	StaticDuration *duration.Duration `protobuf:"bytes,2,opt,name=static_duration,json=staticDuration,proto3" json:"static_duration,omitempty"`
	// The polyline associated with this step.
	Polyline *Polyline `protobuf:"bytes,3,opt,name=polyline,proto3" json:"polyline,omitempty"`
	// The start location of this step.
	StartLocation *Location `protobuf:"bytes,4,opt,name=start_location,json=startLocation,proto3" json:"start_location,omitempty"`
	// The end location of this step.
	EndLocation *Location `protobuf:"bytes,5,opt,name=end_location,json=endLocation,proto3" json:"end_location,omitempty"`
	// Navigation instructions.
	NavigationInstruction *NavigationInstruction `protobuf:"bytes,6,opt,name=navigation_instruction,json=navigationInstruction,proto3" json:"navigation_instruction,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *RouteLegStep) Reset()         { *m = RouteLegStep{} }
func (m *RouteLegStep) String() string { return proto.CompactTextString(m) }
func (*RouteLegStep) ProtoMessage()    {}
func (*RouteLegStep) Descriptor() ([]byte, []int) {
	return fileDescriptor_294c3406019f26b4, []int{7}
}

func (m *RouteLegStep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteLegStep.Unmarshal(m, b)
}
func (m *RouteLegStep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteLegStep.Marshal(b, m, deterministic)
}
func (m *RouteLegStep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteLegStep.Merge(m, src)
}
func (m *RouteLegStep) XXX_Size() int {
	return xxx_messageInfo_RouteLegStep.Size(m)
}
func (m *RouteLegStep) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteLegStep.DiscardUnknown(m)
}

var xxx_messageInfo_RouteLegStep proto.InternalMessageInfo

func (m *RouteLegStep) GetDistanceMeters() int32 {
	if m != nil {
		return m.DistanceMeters
	}
	return 0
}

func (m *RouteLegStep) GetStaticDuration() *duration.Duration {
	if m != nil {
		return m.StaticDuration
	}
	return nil
}

func (m *RouteLegStep) GetPolyline() *Polyline {
	if m != nil {
		return m.Polyline
	}
	return nil
}

func (m *RouteLegStep) GetStartLocation() *Location {
	if m != nil {
		return m.StartLocation
	}
	return nil
}

func (m *RouteLegStep) GetEndLocation() *Location {
	if m != nil {
		return m.EndLocation
	}
	return nil
}

func (m *RouteLegStep) GetNavigationInstruction() *NavigationInstruction {
	if m != nil {
		return m.NavigationInstruction
	}
	return nil
}

type NavigationInstruction struct {
	// Encapsulates the navigation instructions for the current step (e.g., turn
	// left, merge, straight, etc.). This field determines which icon to display.
	Maneuver Maneuver `protobuf:"varint,1,opt,name=maneuver,proto3,enum=google.maps.routes.v1.Maneuver" json:"maneuver,omitempty"`
	// Instructions for navigating this step.
	Instructions         string   `protobuf:"bytes,2,opt,name=instructions,proto3" json:"instructions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NavigationInstruction) Reset()         { *m = NavigationInstruction{} }
func (m *NavigationInstruction) String() string { return proto.CompactTextString(m) }
func (*NavigationInstruction) ProtoMessage()    {}
func (*NavigationInstruction) Descriptor() ([]byte, []int) {
	return fileDescriptor_294c3406019f26b4, []int{8}
}

func (m *NavigationInstruction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NavigationInstruction.Unmarshal(m, b)
}
func (m *NavigationInstruction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NavigationInstruction.Marshal(b, m, deterministic)
}
func (m *NavigationInstruction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NavigationInstruction.Merge(m, src)
}
func (m *NavigationInstruction) XXX_Size() int {
	return xxx_messageInfo_NavigationInstruction.Size(m)
}
func (m *NavigationInstruction) XXX_DiscardUnknown() {
	xxx_messageInfo_NavigationInstruction.DiscardUnknown(m)
}

var xxx_messageInfo_NavigationInstruction proto.InternalMessageInfo

func (m *NavigationInstruction) GetManeuver() Maneuver {
	if m != nil {
		return m.Maneuver
	}
	return Maneuver_MANEUVER_UNSPECIFIED
}

func (m *NavigationInstruction) GetInstructions() string {
	if m != nil {
		return m.Instructions
	}
	return ""
}

func init() {
	proto.RegisterEnum("google.maps.routes.v1.Maneuver", Maneuver_name, Maneuver_value)
	proto.RegisterType((*Route)(nil), "google.maps.routes.v1.Route")
	proto.RegisterType((*RouteTravelAdvisory)(nil), "google.maps.routes.v1.RouteTravelAdvisory")
	proto.RegisterType((*RouteLegTravelAdvisory)(nil), "google.maps.routes.v1.RouteLegTravelAdvisory")
	proto.RegisterType((*TrafficRestriction)(nil), "google.maps.routes.v1.TrafficRestriction")
	proto.RegisterType((*LicensePlateLastCharacterRestriction)(nil), "google.maps.routes.v1.LicensePlateLastCharacterRestriction")
	proto.RegisterType((*RouteLeg)(nil), "google.maps.routes.v1.RouteLeg")
	proto.RegisterType((*TollInfo)(nil), "google.maps.routes.v1.TollInfo")
	proto.RegisterType((*RouteLegStep)(nil), "google.maps.routes.v1.RouteLegStep")
	proto.RegisterType((*NavigationInstruction)(nil), "google.maps.routes.v1.NavigationInstruction")
}

func init() {
	proto.RegisterFile("google/maps/routes/v1/route.proto", fileDescriptor_294c3406019f26b4)
}

var fileDescriptor_294c3406019f26b4 = []byte{
	// 1028 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0x66, 0xfd, 0x93, 0xae, 0x8f, 0x13, 0xdb, 0x9d, 0xd4, 0xad, 0x9b, 0x0b, 0x30, 0x26, 0x12,
	0xa6, 0x82, 0xb5, 0x92, 0x0a, 0x24, 0x14, 0x6e, 0x9c, 0xc6, 0x49, 0x23, 0x6c, 0xc7, 0x1a, 0xff,
	0x48, 0x54, 0x11, 0xab, 0xe9, 0x7a, 0xb2, 0xac, 0xb4, 0xde, 0x59, 0xed, 0x8e, 0x1d, 0xfc, 0x0a,
	0xdc, 0xf2, 0x02, 0x88, 0x3b, 0x7a, 0xc3, 0x7b, 0xf0, 0x08, 0xbc, 0x00, 0xaf, 0xc0, 0x25, 0x9a,
	0x99, 0x5d, 0xff, 0x24, 0x76, 0x9c, 0x92, 0xde, 0x58, 0x73, 0xce, 0x7c, 0xdf, 0x77, 0xce, 0xcc,
	0x9c, 0x73, 0xbc, 0xf0, 0xa9, 0xcd, 0x98, 0xed, 0xd2, 0xda, 0x88, 0xf8, 0x61, 0x2d, 0x60, 0x63,
	0x4e, 0xc3, 0xda, 0xe4, 0x40, 0xad, 0x0c, 0x3f, 0x60, 0x9c, 0xa1, 0xa2, 0x82, 0x18, 0x02, 0x62,
	0x28, 0x88, 0x31, 0x39, 0xd8, 0xfb, 0x38, 0x62, 0xda, 0x94, 0xd5, 0xf8, 0xd4, 0xa7, 0xb5, 0x89,
	0x43, 0xaf, 0x7d, 0x16, 0x70, 0x45, 0xdb, 0xdb, 0x5f, 0xad, 0xec, 0x33, 0x77, 0xea, 0x3a, 0x1e,
	0xbd, 0x1b, 0x75, 0x4d, 0xa6, 0x3e, 0x73, 0xbc, 0x58, 0x2b, 0x8e, 0x25, 0xad, 0xb7, 0xe3, 0xab,
	0xda, 0x70, 0x1c, 0x10, 0xee, 0x30, 0x2f, 0xda, 0x7f, 0x16, 0xed, 0xcb, 0x3c, 0x46, 0xcc, 0xa3,
	0x53, 0xb5, 0x51, 0xf9, 0x27, 0x09, 0x69, 0x2c, 0x54, 0xd1, 0x4b, 0x48, 0xb9, 0xd4, 0x0e, 0x4b,
	0x5a, 0x39, 0x59, 0xcd, 0x1e, 0x7e, 0x62, 0xac, 0x3c, 0x94, 0x21, 0xb1, 0x4d, 0x6a, 0x63, 0x09,
	0x46, 0x9f, 0x43, 0x7e, 0xe8, 0x84, 0x9c, 0x78, 0x16, 0x35, 0x47, 0x94, 0xd3, 0x20, 0x2c, 0x25,
	0xca, 0x5a, 0x35, 0x8d, 0x73, 0xb1, 0xbb, 0x25, 0xbd, 0xe8, 0x6b, 0xd0, 0xe3, 0x94, 0x4a, 0xc9,
	0xb2, 0x56, 0xcd, 0x1e, 0x3e, 0x8f, 0x23, 0xc4, 0x39, 0x1b, 0x27, 0x11, 0x00, 0xcf, 0xa0, 0xe8,
	0x18, 0xf2, 0x21, 0x27, 0xdc, 0xb1, 0xcc, 0x19, 0x3b, 0xb5, 0x89, 0x9d, 0x53, 0x8c, 0xd8, 0x46,
	0x47, 0xa0, 0xc7, 0x77, 0x5a, 0x4a, 0x4b, 0xf2, 0xba, 0xc3, 0x75, 0x22, 0x18, 0x9e, 0x11, 0x50,
	0x19, 0xb2, 0x43, 0x1a, 0x5a, 0x81, 0xe3, 0xcb, 0xe0, 0x5b, 0x65, 0xad, 0x9a, 0xc1, 0x8b, 0x2e,
	0xb4, 0x07, 0xfa, 0x35, 0x09, 0x3c, 0xc7, 0xb3, 0xc3, 0xd2, 0xa3, 0x72, 0xb2, 0x9a, 0xc1, 0x33,
	0x5b, 0x9c, 0x3a, 0x7e, 0xf4, 0x92, 0xbe, 0x9c, 0xb7, 0x4d, 0x99, 0x21, 0x5e, 0xc3, 0x18, 0x44,
	0x00, 0x3c, 0x83, 0xa2, 0x2e, 0xe4, 0x79, 0x40, 0x26, 0xd4, 0x35, 0xc9, 0x70, 0xe2, 0x84, 0x2c,
	0x98, 0x96, 0x32, 0x92, 0xfd, 0xe2, 0xae, 0x57, 0xe9, 0x49, 0x4a, 0x3d, 0x62, 0xe0, 0x1c, 0x5f,
	0xb2, 0x2b, 0x7f, 0x6a, 0xb0, 0xbb, 0x02, 0x87, 0xde, 0xc0, 0x2e, 0x0f, 0xc8, 0xd5, 0x95, 0x63,
	0x99, 0x01, 0x0d, 0x79, 0xe0, 0x58, 0xf2, 0xa4, 0x9a, 0x0c, 0xf8, 0xc5, 0x9a, 0x80, 0x3d, 0xc5,
	0xc0, 0x73, 0x02, 0x46, 0xfc, 0x96, 0x0f, 0x7d, 0x07, 0x19, 0xce, 0x5c, 0xd7, 0x74, 0xbc, 0x2b,
	0x26, 0x0b, 0x63, 0xfd, 0xdd, 0xf7, 0x98, 0xeb, 0x9e, 0x7b, 0x57, 0x0c, 0xeb, 0x3c, 0x5a, 0x55,
	0x06, 0xf0, 0x34, 0x2e, 0xb7, 0x1b, 0x39, 0x2f, 0xe9, 0x6a, 0xef, 0xab, 0xfb, 0x4e, 0x03, 0x74,
	0xfb, 0x00, 0xe8, 0x57, 0x0d, 0xaa, 0xae, 0x63, 0x51, 0x2f, 0xa4, 0xa6, 0xef, 0x12, 0x4e, 0x4d,
	0x97, 0x84, 0xdc, 0xb4, 0x7e, 0x22, 0x01, 0xb1, 0x38, 0x0d, 0x56, 0x5c, 0xcf, 0xd1, 0x9a, 0xa0,
	0x4d, 0x25, 0xd3, 0x11, 0x2a, 0x4d, 0x12, 0xf2, 0x57, 0xb1, 0xc6, 0xe2, 0x85, 0xed, 0xbb, 0xf7,
	0x40, 0x55, 0x7e, 0x84, 0xfd, 0xfb, 0xa8, 0xa1, 0x6f, 0xe0, 0x19, 0x71, 0x5d, 0x76, 0x4d, 0x87,
	0x37, 0xd2, 0x56, 0x1d, 0x9d, 0xc1, 0xc5, 0x68, 0x7b, 0x49, 0x21, 0xac, 0xfc, 0x92, 0x02, 0x3d,
	0xbe, 0xe5, 0x55, 0xed, 0xac, 0x6d, 0x6c, 0xe7, 0xc4, 0x83, 0xda, 0x39, 0xf9, 0x90, 0x76, 0x4e,
	0xbd, 0x6f, 0x3b, 0x9f, 0x82, 0x90, 0x0b, 0xb8, 0xe9, 0x32, 0x4b, 0xc5, 0xbf, 0x7b, 0x22, 0x34,
	0x23, 0x18, 0xde, 0x91, 0xb4, 0xd8, 0x44, 0xc7, 0xb0, 0x4d, 0xbd, 0xe1, 0x5c, 0x65, 0xeb, 0x7e,
	0x2a, 0x59, 0xea, 0x0d, 0x67, 0x1a, 0xdf, 0x42, 0x3a, 0xe4, 0xd4, 0x57, 0x53, 0x23, 0x7b, 0xf8,
	0xd9, 0x86, 0x89, 0xdb, 0xe5, 0xd4, 0xc7, 0x8a, 0x81, 0x06, 0xb7, 0x07, 0x84, 0x1a, 0x2f, 0x5f,
	0x6d, 0x10, 0xd9, 0x30, 0x23, 0xce, 0x40, 0x8f, 0xfb, 0x05, 0x1d, 0x41, 0x9e, 0x86, 0xdc, 0x19,
	0x11, 0x4e, 0x87, 0xa6, 0x1f, 0x38, 0x16, 0x8d, 0xfe, 0x1a, 0x50, 0x1c, 0x43, 0x8e, 0xaf, 0x96,
	0xf8, 0x33, 0xc1, 0xb9, 0x19, 0xb4, 0x23, 0x90, 0x95, 0x3f, 0x92, 0xb0, 0xbd, 0x98, 0xf8, 0xfd,
	0x2b, 0x6b, 0x45, 0x89, 0x24, 0x1e, 0x52, 0x22, 0xc9, 0x87, 0x97, 0x48, 0xea, 0x83, 0x94, 0x48,
	0xfa, 0x7f, 0x94, 0x88, 0x05, 0x4f, 0x3d, 0x32, 0x71, 0x6c, 0x69, 0x99, 0x8e, 0x17, 0xf2, 0x60,
	0x6c, 0x2d, 0x14, 0xdc, 0x97, 0x6b, 0xd4, 0xda, 0x33, 0xd2, 0xf9, 0x9c, 0x83, 0x8b, 0xde, 0x2a,
	0x77, 0xe5, 0x67, 0x28, 0xae, 0xc4, 0x8b, 0x6b, 0x1c, 0x11, 0x8f, 0x8e, 0x27, 0x34, 0x90, 0x8f,
	0x95, 0x5b, 0x9b, 0x7d, 0x2b, 0x82, 0xe1, 0x19, 0x01, 0x55, 0x60, 0x7b, 0x21, 0x5f, 0xf5, 0x59,
	0x90, 0xc1, 0x4b, 0xbe, 0x17, 0x7f, 0x27, 0x40, 0x8f, 0xa9, 0xa8, 0x04, 0x4f, 0x5a, 0xf5, 0x76,
	0xa3, 0x3f, 0x68, 0x60, 0xb3, 0xdf, 0xee, 0x76, 0x1a, 0xaf, 0xce, 0x4f, 0xcf, 0x1b, 0x27, 0x85,
	0x8f, 0xd0, 0x13, 0x28, 0xf4, 0xfa, 0xb8, 0x6d, 0x76, 0x9b, 0xe7, 0x67, 0xaf, 0x7b, 0x66, 0xb3,
	0x71, 0xda, 0x2b, 0x68, 0x68, 0x17, 0xf2, 0xca, 0xfb, 0xba, 0x8e, 0x3b, 0xca, 0x99, 0x40, 0x39,
	0x80, 0xbe, 0xf4, 0x4a, 0x3b, 0x89, 0x76, 0x20, 0x33, 0x37, 0x53, 0xa8, 0x08, 0x8f, 0x17, 0x95,
	0xb0, 0xf8, 0x2d, 0xa4, 0xe7, 0x01, 0xa4, 0x94, 0xf2, 0x6e, 0xa1, 0x3c, 0x64, 0x95, 0x96, 0x72,
	0x3c, 0x12, 0xe2, 0x0b, 0xb6, 0x8e, 0xb6, 0x41, 0xef, 0xf6, 0x70, 0x5d, 0x5a, 0x19, 0x11, 0x0a,
	0xd7, 0x5b, 0x51, 0x26, 0x20, 0xc0, 0xd2, 0x54, 0xe0, 0x2c, 0xca, 0x40, 0xba, 0xd5, 0xc0, 0x67,
	0x8d, 0xc2, 0xb6, 0x40, 0x9e, 0x5e, 0xe0, 0xef, 0x15, 0x72, 0x47, 0x20, 0xa5, 0xa9, 0x90, 0x39,
	0x81, 0x3c, 0x6d, 0x60, 0xfc, 0x43, 0x21, 0x2f, 0x52, 0x90, 0x4b, 0x53, 0x84, 0x69, 0x17, 0x0a,
	0xe2, 0xd0, 0xf8, 0xa2, 0xdf, 0x3e, 0xa9, 0x1f, 0x5f, 0xf4, 0xa3, 0x9b, 0x78, 0x2c, 0xd2, 0x5f,
	0x70, 0x2a, 0x19, 0x74, 0xfc, 0x9b, 0x06, 0xcf, 0x2d, 0x36, 0x5a, 0xfd, 0x62, 0xc7, 0x20, 0xbb,
	0xb3, 0x23, 0x5a, 0xa9, 0xa3, 0xbd, 0x39, 0x8a, 0x3f, 0x4a, 0x98, 0x4b, 0x3c, 0xdb, 0x60, 0x81,
	0x5d, 0xb3, 0xa9, 0x27, 0x1b, 0xad, 0xa6, 0xb6, 0x88, 0xef, 0x84, 0x37, 0xbe, 0x41, 0x8f, 0xd4,
	0xea, 0x5f, 0x4d, 0xfb, 0x3d, 0x91, 0x3a, 0x6b, 0xe1, 0xee, 0xbb, 0x44, 0xf1, 0x4c, 0xe9, 0xb4,
	0x44, 0x30, 0xac, 0x82, 0x0d, 0x0e, 0xfe, 0x8a, 0xfd, 0x97, 0xc2, 0x7f, 0xa9, 0xfc, 0x97, 0x83,
	0x83, 0xb7, 0x5b, 0x32, 0xc2, 0xcb, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x63, 0x2f, 0xb2, 0xd5,
	0x64, 0x0b, 0x00, 0x00,
}
