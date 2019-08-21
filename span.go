package main

import "github.com/tinylib/msgp/msgp"

type span struct {
	Name     string             `msg:"name"`              // operation name
	Service  string             `msg:"service"`           // service name (i.e. "grpc.server", "http.request")
	Resource string             `msg:"resource"`          // resource name (i.e. "/user?id=123", "SELECT * FROM users")
	Type     string             `msg:"type"`              // protocol associated with the span (i.e. "web", "db", "cache")
	Start    int64              `msg:"start"`             // span start time expressed in nanoseconds since epoch
	Duration int64              `msg:"duration"`          // duration of the span expressed in nanoseconds
	Meta     map[string]string  `msg:"meta,omitempty"`    // arbitrary map of metadata
	Metrics  map[string]float64 `msg:"metrics,omitempty"` // arbitrary map of numeric metrics
	SpanID   uint64             `msg:"span_id"`           // identifier of this span
	TraceID  uint64             `msg:"trace_id"`          // identifier of the root span
	ParentID uint64             `msg:"parent_id"`         // identifier of the span's direct parent
	Error    int32              `msg:"error"`             // error status of the span; 0 means no errors

	childSpans []*span `msg:"-"`
}

// DecodeMsg implements msgp.Decodable
func (z *span) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "service":
			z.Service, err = dc.ReadString()
			if err != nil {
				return
			}
		case "resource":
			z.Resource, err = dc.ReadString()
			if err != nil {
				return
			}
		case "type":
			z.Type, err = dc.ReadString()
			if err != nil {
				return
			}
		case "start":
			z.Start, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "duration":
			z.Duration, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "meta":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Meta == nil && zb0002 > 0 {
				z.Meta = make(map[string]string, zb0002)
			} else if len(z.Meta) > 0 {
				for key := range z.Meta {
					delete(z.Meta, key)
				}
			}
			for zb0002 > 0 {
				zb0002--
				var za0001 string
				var za0002 string
				za0001, err = dc.ReadString()
				if err != nil {
					return
				}
				za0002, err = dc.ReadString()
				if err != nil {
					return
				}
				z.Meta[za0001] = za0002
			}
		case "metrics":
			var zb0003 uint32
			zb0003, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Metrics == nil && zb0003 > 0 {
				z.Metrics = make(map[string]float64, zb0003)
			} else if len(z.Metrics) > 0 {
				for key := range z.Metrics {
					delete(z.Metrics, key)
				}
			}
			for zb0003 > 0 {
				zb0003--
				var za0003 string
				var za0004 float64
				za0003, err = dc.ReadString()
				if err != nil {
					return
				}
				za0004, err = dc.ReadFloat64()
				if err != nil {
					return
				}
				z.Metrics[za0003] = za0004
			}
		case "span_id":
			z.SpanID, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "trace_id":
			z.TraceID, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "parent_id":
			z.ParentID, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "error":
			z.Error, err = dc.ReadInt32()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *span) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 8 + msgp.StringPrefixSize + len(z.Service) + 9 + msgp.StringPrefixSize + len(z.Resource) + 5 + msgp.StringPrefixSize + len(z.Type) + 6 + msgp.Int64Size + 9 + msgp.Int64Size + 5 + msgp.MapHeaderSize
	if z.Meta != nil {
		for za0001, za0002 := range z.Meta {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.StringPrefixSize + len(za0002)
		}
	}
	s += 8 + msgp.MapHeaderSize
	if z.Metrics != nil {
		for za0003, za0004 := range z.Metrics {
			_ = za0004
			s += msgp.StringPrefixSize + len(za0003) + msgp.Float64Size
		}
	}
	s += 8 + msgp.Uint64Size + 9 + msgp.Uint64Size + 10 + msgp.Uint64Size + 6 + msgp.Int32Size
	return
}
