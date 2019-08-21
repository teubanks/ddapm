package main

import "github.com/tinylib/msgp/msgp"

type (
	spanList []*span

	// spanLists implements msgp.Decodable on top of a slice of spanList.
	// This type is only used in tests.
	spanLists []spanList
)

// DecodeMsg implements msgp.Decodable
func (z *spanList) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(spanList, zb0002)
	}
	for zb0001 := range *z {
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(span)
			}
			err = (*z)[zb0001].DecodeMsg(dc)
			if err != nil {
				return
			}
		}
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z spanList) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0003 := range z {
		if z[zb0003] == nil {
			s += msgp.NilSize
		} else {
			s += z[zb0003].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *spanLists) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0003 uint32
	zb0003, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0003) {
		(*z) = (*z)[:zb0003]
	} else {
		(*z) = make(spanLists, zb0003)
	}
	for zb0001 := range *z {
		var zb0004 uint32
		zb0004, err = dc.ReadArrayHeader()
		if err != nil {
			return
		}
		if cap((*z)[zb0001]) >= int(zb0004) {
			(*z)[zb0001] = ((*z)[zb0001])[:zb0004]
		} else {
			(*z)[zb0001] = make(spanList, zb0004)
		}
		for zb0002 := range (*z)[zb0001] {
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				(*z)[zb0001][zb0002] = nil
			} else {
				if (*z)[zb0001][zb0002] == nil {
					(*z)[zb0001][zb0002] = new(span)
				}
				err = (*z)[zb0001][zb0002].DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		}
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z spanLists) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0005 := range z {
		s += msgp.ArrayHeaderSize
		for zb0006 := range z[zb0005] {
			if z[zb0005][zb0006] == nil {
				s += msgp.NilSize
			} else {
				s += z[zb0005][zb0006].Msgsize()
			}
		}
	}
	return
}
