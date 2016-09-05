// Code generated by protoc-gen-go.
// source: coprocess_session_state.proto
// DO NOT EDIT!

package coprocess

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AccessSpec struct {
	Url     string   `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Methods []string `protobuf:"bytes,2,rep,name=methods" json:"methods,omitempty"`
}

func (m *AccessSpec) Reset()                    { *m = AccessSpec{} }
func (m *AccessSpec) String() string            { return proto.CompactTextString(m) }
func (*AccessSpec) ProtoMessage()               {}
func (*AccessSpec) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type AccessDefinition struct {
	ApiName     string        `protobuf:"bytes,1,opt,name=api_name,json=apiName" json:"api_name,omitempty"`
	ApiId       string        `protobuf:"bytes,2,opt,name=api_id,json=apiId" json:"api_id,omitempty"`
	Versions    []string      `protobuf:"bytes,3,rep,name=versions" json:"versions,omitempty"`
	AllowedUrls []*AccessSpec `protobuf:"bytes,4,rep,name=allowed_urls,json=allowedUrls" json:"allowed_urls,omitempty"`
}

func (m *AccessDefinition) Reset()                    { *m = AccessDefinition{} }
func (m *AccessDefinition) String() string            { return proto.CompactTextString(m) }
func (*AccessDefinition) ProtoMessage()               {}
func (*AccessDefinition) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *AccessDefinition) GetAllowedUrls() []*AccessSpec {
	if m != nil {
		return m.AllowedUrls
	}
	return nil
}

type BasicAuthData struct {
	Password string `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
	Hash     string `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
}

func (m *BasicAuthData) Reset()                    { *m = BasicAuthData{} }
func (m *BasicAuthData) String() string            { return proto.CompactTextString(m) }
func (*BasicAuthData) ProtoMessage()               {}
func (*BasicAuthData) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

type JWTData struct {
	Secret string `protobuf:"bytes,1,opt,name=secret" json:"secret,omitempty"`
}

func (m *JWTData) Reset()                    { *m = JWTData{} }
func (m *JWTData) String() string            { return proto.CompactTextString(m) }
func (*JWTData) ProtoMessage()               {}
func (*JWTData) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

type Monitor struct {
	TriggerLimits []float64 `protobuf:"fixed64,1,rep,packed,name=trigger_limits,json=triggerLimits" json:"trigger_limits,omitempty"`
}

func (m *Monitor) Reset()                    { *m = Monitor{} }
func (m *Monitor) String() string            { return proto.CompactTextString(m) }
func (*Monitor) ProtoMessage()               {}
func (*Monitor) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

type SessionState struct {
	LastCheck               int64                        `protobuf:"varint,1,opt,name=last_check,json=lastCheck" json:"last_check,omitempty"`
	Allowance               float64                      `protobuf:"fixed64,2,opt,name=allowance" json:"allowance,omitempty"`
	Rate                    float64                      `protobuf:"fixed64,3,opt,name=rate" json:"rate,omitempty"`
	Per                     float64                      `protobuf:"fixed64,4,opt,name=per" json:"per,omitempty"`
	Expires                 int64                        `protobuf:"varint,5,opt,name=expires" json:"expires,omitempty"`
	QuotaMax                int64                        `protobuf:"varint,6,opt,name=quota_max,json=quotaMax" json:"quota_max,omitempty"`
	QuotaRenews             int64                        `protobuf:"varint,7,opt,name=quota_renews,json=quotaRenews" json:"quota_renews,omitempty"`
	QuotaRemaining          int64                        `protobuf:"varint,8,opt,name=quota_remaining,json=quotaRemaining" json:"quota_remaining,omitempty"`
	QuotaRenewalRate        int64                        `protobuf:"varint,9,opt,name=quota_renewal_rate,json=quotaRenewalRate" json:"quota_renewal_rate,omitempty"`
	AccessRights            map[string]*AccessDefinition `protobuf:"bytes,10,rep,name=access_rights,json=accessRights" json:"access_rights,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	OrgId                   string                       `protobuf:"bytes,11,opt,name=org_id,json=orgId" json:"org_id,omitempty"`
	OauthClientId           string                       `protobuf:"bytes,12,opt,name=oauth_client_id,json=oauthClientId" json:"oauth_client_id,omitempty"`
	OauthKeys               map[string]string            `protobuf:"bytes,13,rep,name=oauth_keys,json=oauthKeys" json:"oauth_keys,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	BasicAuthData           *BasicAuthData               `protobuf:"bytes,14,opt,name=basic_auth_data,json=basicAuthData" json:"basic_auth_data,omitempty"`
	JwtData                 *JWTData                     `protobuf:"bytes,15,opt,name=jwt_data,json=jwtData" json:"jwt_data,omitempty"`
	HmacEnabled             bool                         `protobuf:"varint,16,opt,name=hmac_enabled,json=hmacEnabled" json:"hmac_enabled,omitempty"`
	HmacSecret              string                       `protobuf:"bytes,17,opt,name=hmac_secret,json=hmacSecret" json:"hmac_secret,omitempty"`
	IsInactive              bool                         `protobuf:"varint,18,opt,name=is_inactive,json=isInactive" json:"is_inactive,omitempty"`
	ApplyPolicyId           string                       `protobuf:"bytes,19,opt,name=apply_policy_id,json=applyPolicyId" json:"apply_policy_id,omitempty"`
	DataExpires             int64                        `protobuf:"varint,20,opt,name=data_expires,json=dataExpires" json:"data_expires,omitempty"`
	Monitor                 *Monitor                     `protobuf:"bytes,21,opt,name=monitor" json:"monitor,omitempty"`
	EnableDetailedRecording bool                         `protobuf:"varint,22,opt,name=enable_detailed_recording,json=enableDetailedRecording" json:"enable_detailed_recording,omitempty"`
	Metadata                string                       `protobuf:"bytes,23,opt,name=metadata" json:"metadata,omitempty"`
	Tags                    []string                     `protobuf:"bytes,24,rep,name=tags" json:"tags,omitempty"`
	Alias                   string                       `protobuf:"bytes,25,opt,name=alias" json:"alias,omitempty"`
	LastUpdated             string                       `protobuf:"bytes,26,opt,name=last_updated,json=lastUpdated" json:"last_updated,omitempty"`
}

func (m *SessionState) Reset()                    { *m = SessionState{} }
func (m *SessionState) String() string            { return proto.CompactTextString(m) }
func (*SessionState) ProtoMessage()               {}
func (*SessionState) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *SessionState) GetAccessRights() map[string]*AccessDefinition {
	if m != nil {
		return m.AccessRights
	}
	return nil
}

func (m *SessionState) GetOauthKeys() map[string]string {
	if m != nil {
		return m.OauthKeys
	}
	return nil
}

func (m *SessionState) GetBasicAuthData() *BasicAuthData {
	if m != nil {
		return m.BasicAuthData
	}
	return nil
}

func (m *SessionState) GetJwtData() *JWTData {
	if m != nil {
		return m.JwtData
	}
	return nil
}

func (m *SessionState) GetMonitor() *Monitor {
	if m != nil {
		return m.Monitor
	}
	return nil
}

func init() {
	proto.RegisterType((*AccessSpec)(nil), "coprocess.AccessSpec")
	proto.RegisterType((*AccessDefinition)(nil), "coprocess.AccessDefinition")
	proto.RegisterType((*BasicAuthData)(nil), "coprocess.BasicAuthData")
	proto.RegisterType((*JWTData)(nil), "coprocess.JWTData")
	proto.RegisterType((*Monitor)(nil), "coprocess.Monitor")
	proto.RegisterType((*SessionState)(nil), "coprocess.SessionState")
}

func init() { proto.RegisterFile("coprocess_session_state.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 836 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x55, 0x51, 0x6f, 0x23, 0x35,
	0x10, 0xd6, 0x36, 0x6d, 0x93, 0x9d, 0x24, 0x6d, 0xcf, 0xb4, 0x77, 0x6e, 0x8f, 0x13, 0x69, 0x24,
	0x8e, 0x20, 0x1d, 0x15, 0x94, 0x97, 0xea, 0x84, 0x04, 0xc7, 0xb5, 0x0f, 0x01, 0xee, 0x40, 0x5b,
	0x4e, 0xbc, 0x20, 0x59, 0xd3, 0x5d, 0x93, 0xb8, 0xdd, 0x5d, 0x2f, 0xb6, 0xd3, 0x34, 0xff, 0x80,
	0xdf, 0xc0, 0xaf, 0x45, 0x9e, 0xf5, 0x36, 0x5b, 0x95, 0x7b, 0xf3, 0x7c, 0xf3, 0x79, 0xf6, 0xb3,
	0xe7, 0xf3, 0x2c, 0xbc, 0x48, 0x75, 0x65, 0x74, 0x2a, 0xad, 0x15, 0x56, 0x5a, 0xab, 0x74, 0x29,
	0xac, 0x43, 0x27, 0x4f, 0x2a, 0xa3, 0x9d, 0x66, 0xf1, 0x7d, 0x7a, 0x7c, 0x06, 0xf0, 0x26, 0xf5,
	0xab, 0xcb, 0x4a, 0xa6, 0x6c, 0x0f, 0x3a, 0x0b, 0x93, 0xf3, 0x68, 0x14, 0x4d, 0xe2, 0xc4, 0x2f,
	0x19, 0x87, 0x6e, 0x21, 0xdd, 0x5c, 0x67, 0x96, 0x6f, 0x8c, 0x3a, 0x93, 0x38, 0x69, 0xc2, 0xf1,
	0xbf, 0x11, 0xec, 0xd5, 0x5b, 0xcf, 0xe5, 0x5f, 0xaa, 0x54, 0x4e, 0xe9, 0x92, 0x1d, 0x42, 0x0f,
	0x2b, 0x25, 0x4a, 0x2c, 0x64, 0xa8, 0xd2, 0xc5, 0x4a, 0xbd, 0xc7, 0x42, 0xb2, 0x03, 0xd8, 0xf6,
	0x29, 0x95, 0xf1, 0x0d, 0x4a, 0x6c, 0x61, 0xa5, 0xa6, 0x19, 0x3b, 0x82, 0xde, 0xad, 0x34, 0x5e,
	0xa2, 0xe5, 0x1d, 0xfa, 0xc2, 0x7d, 0xcc, 0xce, 0x60, 0x80, 0x79, 0xae, 0x97, 0x32, 0x13, 0x0b,
	0x93, 0x5b, 0xbe, 0x39, 0xea, 0x4c, 0xfa, 0xa7, 0x07, 0x27, 0xf7, 0xf2, 0x4f, 0xd6, 0xda, 0x93,
	0x7e, 0xa0, 0x7e, 0x30, 0xb9, 0x1d, 0x7f, 0x0f, 0xc3, 0x1f, 0xd1, 0xaa, 0xf4, 0xcd, 0xc2, 0xcd,
	0xcf, 0xd1, 0xa1, 0xff, 0x4c, 0x85, 0xd6, 0x2e, 0xb5, 0xc9, 0x82, 0xb0, 0xfb, 0x98, 0x31, 0xd8,
	0x9c, 0xa3, 0x9d, 0x07, 0x5d, 0xb4, 0x1e, 0x1f, 0x43, 0xf7, 0xa7, 0x3f, 0x7e, 0xa7, 0xad, 0x4f,
	0x61, 0xdb, 0xca, 0xd4, 0x48, 0x17, 0x36, 0x86, 0x68, 0xfc, 0x35, 0x74, 0xdf, 0xe9, 0x52, 0x39,
	0x6d, 0xd8, 0xe7, 0xb0, 0xe3, 0x8c, 0x9a, 0xcd, 0xa4, 0x11, 0xb9, 0x2a, 0x94, 0xb3, 0x3c, 0x1a,
	0x75, 0x26, 0x51, 0x32, 0x0c, 0xe8, 0x2f, 0x04, 0x8e, 0xff, 0x89, 0x61, 0x70, 0x59, 0xf7, 0xe3,
	0xd2, 0xb7, 0x83, 0xbd, 0x00, 0xc8, 0xd1, 0x3a, 0x91, 0xce, 0x65, 0x7a, 0x43, 0xe5, 0x3b, 0x49,
	0xec, 0x91, 0xb7, 0x1e, 0x60, 0x9f, 0x42, 0x4c, 0x87, 0xc2, 0x32, 0x95, 0xa4, 0x2e, 0x4a, 0xd6,
	0x80, 0x97, 0x6d, 0xd0, 0x49, 0xde, 0xa1, 0x04, 0xad, 0x7d, 0x03, 0x2b, 0x69, 0xf8, 0x26, 0x41,
	0x7e, 0xe9, 0x1b, 0x28, 0xef, 0x2a, 0x65, 0xa4, 0xe5, 0x5b, 0x54, 0xbf, 0x09, 0xd9, 0x73, 0x88,
	0xff, 0x5e, 0x68, 0x87, 0xa2, 0xc0, 0x3b, 0xbe, 0x4d, 0xb9, 0x1e, 0x01, 0xef, 0xf0, 0x8e, 0x1d,
	0xc3, 0xa0, 0x4e, 0x1a, 0x59, 0xca, 0xa5, 0xe5, 0x5d, 0xca, 0xf7, 0x09, 0x4b, 0x08, 0x62, 0x5f,
	0xc0, 0x6e, 0x43, 0x29, 0x50, 0x95, 0xaa, 0x9c, 0xf1, 0x1e, 0xb1, 0x76, 0x02, 0x2b, 0xa0, 0xec,
	0x15, 0xb0, 0x56, 0x2d, 0xcc, 0x05, 0xc9, 0x8e, 0x89, 0xbb, 0xb7, 0xae, 0x88, 0x79, 0xe2, 0x8f,
	0xf0, 0x1e, 0x86, 0x48, 0x5d, 0x15, 0x46, 0xcd, 0xe6, 0xce, 0x72, 0xa0, 0xae, 0x7f, 0xd9, 0xea,
	0x7a, 0xfb, 0x0e, 0x83, 0x05, 0x12, 0xe2, 0x5e, 0x94, 0xce, 0xac, 0x92, 0x01, 0xb6, 0x20, 0xef,
	0x3b, 0x6d, 0x66, 0xde, 0x77, 0xfd, 0xda, 0x77, 0xda, 0xcc, 0xa6, 0x19, 0x7b, 0x09, 0xbb, 0x1a,
	0x17, 0x6e, 0x2e, 0xd2, 0x5c, 0xc9, 0xd2, 0xf9, 0xfc, 0x80, 0xf2, 0x43, 0x82, 0xdf, 0x12, 0x3a,
	0xcd, 0xd8, 0x05, 0x40, 0xcd, 0xbb, 0x91, 0x2b, 0xcb, 0x87, 0xa4, 0xe5, 0xe5, 0xc7, 0xb4, 0xfc,
	0xea, 0x99, 0x3f, 0xcb, 0x55, 0x10, 0x12, 0xeb, 0x26, 0x66, 0x3f, 0xc0, 0xee, 0x95, 0x37, 0xa4,
	0xa0, 0x5a, 0x19, 0x3a, 0xe4, 0x3b, 0xa3, 0x68, 0xd2, 0x3f, 0xe5, 0xad, 0x5a, 0x0f, 0x2c, 0x9b,
	0x0c, 0xaf, 0x1e, 0x38, 0xf8, 0x2b, 0xe8, 0x5d, 0x2f, 0x5d, 0xbd, 0x75, 0x97, 0xb6, 0xb2, 0xd6,
	0xd6, 0x60, 0xd6, 0xa4, 0x7b, 0xbd, 0x74, 0x44, 0x3f, 0x86, 0xc1, 0xbc, 0xc0, 0x54, 0xc8, 0x12,
	0xaf, 0x72, 0x99, 0xf1, 0xbd, 0x51, 0x34, 0xe9, 0x25, 0x7d, 0x8f, 0x5d, 0xd4, 0x10, 0xfb, 0x0c,
	0x28, 0x14, 0xc1, 0xdd, 0x4f, 0xe8, 0xf8, 0xe0, 0xa1, 0x4b, 0x42, 0x3c, 0x41, 0x59, 0xa1, 0x4a,
	0x4c, 0x9d, 0xba, 0x95, 0x9c, 0x51, 0x09, 0x50, 0x76, 0x1a, 0x10, 0x7f, 0x89, 0x58, 0x55, 0xf9,
	0x4a, 0x54, 0x3a, 0x57, 0xe9, 0xca, 0x5f, 0xe2, 0x27, 0xf5, 0x25, 0x12, 0xfc, 0x1b, 0xa1, 0xd3,
	0xcc, 0x8b, 0xf1, 0xba, 0x45, 0xe3, 0xc4, 0xfd, 0xda, 0x4d, 0x1e, 0xbb, 0x08, 0x6e, 0x7c, 0x05,
	0xdd, 0xa2, 0x7e, 0x4d, 0xfc, 0xe0, 0xd1, 0xe9, 0xc2, 0x3b, 0x4b, 0x1a, 0x0a, 0x7b, 0x0d, 0x87,
	0xf5, 0xc1, 0x44, 0x26, 0x1d, 0xaa, 0x5c, 0x66, 0xc2, 0xc8, 0x54, 0x9b, 0xcc, 0xbb, 0xf0, 0x29,
	0xe9, 0x7c, 0x56, 0x13, 0xce, 0x43, 0x3e, 0x69, 0xd2, 0x7e, 0x14, 0x14, 0xd2, 0x21, 0x5d, 0xe4,
	0xb3, 0x7a, 0x14, 0x34, 0xb1, 0x7f, 0x53, 0x0e, 0x67, 0x96, 0x73, 0x9a, 0x44, 0xb4, 0x66, 0xfb,
	0xb0, 0x85, 0xb9, 0x42, 0xcb, 0x0f, 0xc3, 0xdc, 0xf2, 0x81, 0x3f, 0x12, 0x3d, 0xdd, 0x45, 0x95,
	0xa1, 0x93, 0x19, 0x3f, 0xa2, 0x64, 0xdf, 0x63, 0x1f, 0x6a, 0xe8, 0xe8, 0x4f, 0x78, 0xf2, 0xc8,
	0x9c, 0xfe, 0x85, 0xde, 0xc8, 0x55, 0x33, 0x62, 0x6f, 0xe4, 0x8a, 0x7d, 0x03, 0x5b, 0xb7, 0x98,
	0x2f, 0xea, 0x17, 0xde, 0x3f, 0x7d, 0xfe, 0x68, 0xbc, 0xad, 0xe7, 0x6b, 0x52, 0x33, 0x5f, 0x6f,
	0x9c, 0x45, 0x47, 0xdf, 0xc1, 0xce, 0x43, 0xbb, 0xfd, 0x4f, 0xe9, 0xfd, 0x76, 0xe9, 0xb8, 0xb5,
	0xfb, 0x6a, 0x9b, 0xfe, 0x04, 0xdf, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xd4, 0x82, 0x3b,
	0x2a, 0x06, 0x00, 0x00,
}
