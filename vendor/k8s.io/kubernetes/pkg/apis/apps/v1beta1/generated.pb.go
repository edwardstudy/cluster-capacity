/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo.
// source: k8s.io/kubernetes/pkg/apis/apps/v1beta1/generated.proto
// DO NOT EDIT!

/*
	Package v1beta1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/pkg/apis/apps/v1beta1/generated.proto

	It has these top-level messages:
		StatefulSet
		StatefulSetList
		StatefulSetSpec
		StatefulSetStatus
*/
package v1beta1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import k8s_io_kubernetes_pkg_api_unversioned "k8s.io/kubernetes/pkg/api/unversioned"
import k8s_io_kubernetes_pkg_api_v1 "k8s.io/kubernetes/pkg/api/v1"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

func (m *StatefulSet) Reset()                    { *m = StatefulSet{} }
func (*StatefulSet) ProtoMessage()               {}
func (*StatefulSet) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *StatefulSetList) Reset()                    { *m = StatefulSetList{} }
func (*StatefulSetList) ProtoMessage()               {}
func (*StatefulSetList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *StatefulSetSpec) Reset()                    { *m = StatefulSetSpec{} }
func (*StatefulSetSpec) ProtoMessage()               {}
func (*StatefulSetSpec) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *StatefulSetStatus) Reset()                    { *m = StatefulSetStatus{} }
func (*StatefulSetStatus) ProtoMessage()               {}
func (*StatefulSetStatus) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{3} }

func init() {
	proto.RegisterType((*StatefulSet)(nil), "k8s.io.kubernetes.pkg.apis.apps.v1beta1.StatefulSet")
	proto.RegisterType((*StatefulSetList)(nil), "k8s.io.kubernetes.pkg.apis.apps.v1beta1.StatefulSetList")
	proto.RegisterType((*StatefulSetSpec)(nil), "k8s.io.kubernetes.pkg.apis.apps.v1beta1.StatefulSetSpec")
	proto.RegisterType((*StatefulSetStatus)(nil), "k8s.io.kubernetes.pkg.apis.apps.v1beta1.StatefulSetStatus")
}
func (m *StatefulSet) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StatefulSet) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Spec.Size()))
	n2, err := m.Spec.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	data[i] = 0x1a
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Status.Size()))
	n3, err := m.Status.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func (m *StatefulSetList) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StatefulSetList) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(m.ListMeta.Size()))
	n4, err := m.ListMeta.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			data[i] = 0x12
			i++
			i = encodeVarintGenerated(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *StatefulSetSpec) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StatefulSetSpec) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Replicas != nil {
		data[i] = 0x8
		i++
		i = encodeVarintGenerated(data, i, uint64(*m.Replicas))
	}
	if m.Selector != nil {
		data[i] = 0x12
		i++
		i = encodeVarintGenerated(data, i, uint64(m.Selector.Size()))
		n5, err := m.Selector.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	data[i] = 0x1a
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Template.Size()))
	n6, err := m.Template.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	if len(m.VolumeClaimTemplates) > 0 {
		for _, msg := range m.VolumeClaimTemplates {
			data[i] = 0x22
			i++
			i = encodeVarintGenerated(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	data[i] = 0x2a
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.ServiceName)))
	i += copy(data[i:], m.ServiceName)
	return i, nil
}

func (m *StatefulSetStatus) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StatefulSetStatus) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ObservedGeneration != nil {
		data[i] = 0x8
		i++
		i = encodeVarintGenerated(data, i, uint64(*m.ObservedGeneration))
	}
	data[i] = 0x10
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Replicas))
	return i, nil
}

func encodeFixed64Generated(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *StatefulSet) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *StatefulSetList) Size() (n int) {
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *StatefulSetSpec) Size() (n int) {
	var l int
	_ = l
	if m.Replicas != nil {
		n += 1 + sovGenerated(uint64(*m.Replicas))
	}
	if m.Selector != nil {
		l = m.Selector.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = m.Template.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.VolumeClaimTemplates) > 0 {
		for _, e := range m.VolumeClaimTemplates {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	l = len(m.ServiceName)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *StatefulSetStatus) Size() (n int) {
	var l int
	_ = l
	if m.ObservedGeneration != nil {
		n += 1 + sovGenerated(uint64(*m.ObservedGeneration))
	}
	n += 1 + sovGenerated(uint64(m.Replicas))
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *StatefulSet) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StatefulSet{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_kubernetes_pkg_api_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "StatefulSetSpec", "StatefulSetSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "StatefulSetStatus", "StatefulSetStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StatefulSetList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StatefulSetList{`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "k8s_io_kubernetes_pkg_api_unversioned.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Items), "StatefulSet", "StatefulSet", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StatefulSetSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StatefulSetSpec{`,
		`Replicas:` + valueToStringGenerated(this.Replicas) + `,`,
		`Selector:` + strings.Replace(fmt.Sprintf("%v", this.Selector), "LabelSelector", "k8s_io_kubernetes_pkg_api_unversioned.LabelSelector", 1) + `,`,
		`Template:` + strings.Replace(strings.Replace(this.Template.String(), "PodTemplateSpec", "k8s_io_kubernetes_pkg_api_v1.PodTemplateSpec", 1), `&`, ``, 1) + `,`,
		`VolumeClaimTemplates:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.VolumeClaimTemplates), "PersistentVolumeClaim", "k8s_io_kubernetes_pkg_api_v1.PersistentVolumeClaim", 1), `&`, ``, 1) + `,`,
		`ServiceName:` + fmt.Sprintf("%v", this.ServiceName) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StatefulSetStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StatefulSetStatus{`,
		`ObservedGeneration:` + valueToStringGenerated(this.ObservedGeneration) + `,`,
		`Replicas:` + fmt.Sprintf("%v", this.Replicas) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *StatefulSet) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatefulSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatefulSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StatefulSetList) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatefulSetList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatefulSetList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, StatefulSet{})
			if err := m.Items[len(m.Items)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StatefulSetSpec) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatefulSetSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatefulSetSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replicas", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Replicas = &v
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Selector", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Selector == nil {
				m.Selector = &k8s_io_kubernetes_pkg_api_unversioned.LabelSelector{}
			}
			if err := m.Selector.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Template", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Template.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VolumeClaimTemplates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VolumeClaimTemplates = append(m.VolumeClaimTemplates, k8s_io_kubernetes_pkg_api_v1.PersistentVolumeClaim{})
			if err := m.VolumeClaimTemplates[len(m.VolumeClaimTemplates)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceName = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StatefulSetStatus) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatefulSetStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatefulSetStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObservedGeneration", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ObservedGeneration = &v
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replicas", wireType)
			}
			m.Replicas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Replicas |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorGenerated = []byte{
	// 627 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xe3, 0xa4, 0x29, 0x61, 0x53, 0xbe, 0x96, 0x0a, 0x45, 0x11, 0x72, 0xab, 0x5c, 0x08,
	0x52, 0xbb, 0x56, 0x4a, 0x2b, 0x2a, 0x8e, 0x46, 0x02, 0x21, 0x01, 0x45, 0x0e, 0xaa, 0xa0, 0x08,
	0xa4, 0xb5, 0x33, 0x0d, 0x4b, 0x6c, 0xaf, 0xe5, 0x1d, 0xe7, 0xcc, 0x85, 0x03, 0x37, 0xde, 0x82,
	0x47, 0xe0, 0x15, 0x2a, 0x71, 0xe9, 0x91, 0x53, 0x45, 0xc3, 0x8b, 0x20, 0x6f, 0x36, 0x1f, 0x34,
	0x49, 0xa9, 0x7a, 0xf3, 0xec, 0xce, 0xff, 0x37, 0x33, 0xff, 0x1d, 0x93, 0x87, 0xbd, 0x5d, 0xc5,
	0x84, 0x74, 0x7a, 0x99, 0x0f, 0x69, 0x0c, 0x08, 0xca, 0x49, 0x7a, 0x5d, 0x87, 0x27, 0x42, 0x39,
	0x3c, 0x49, 0x94, 0xd3, 0x6f, 0xf9, 0x80, 0xbc, 0xe5, 0x74, 0x21, 0x86, 0x94, 0x23, 0x74, 0x58,
	0x92, 0x4a, 0x94, 0xf4, 0xde, 0x50, 0xc8, 0x26, 0x42, 0x96, 0xf4, 0xba, 0x2c, 0x17, 0xb2, 0x5c,
	0xc8, 0x8c, 0xb0, 0xbe, 0xd9, 0x15, 0xf8, 0x31, 0xf3, 0x59, 0x20, 0x23, 0xa7, 0x2b, 0xbb, 0xd2,
	0xd1, 0x7a, 0x3f, 0x3b, 0xd4, 0x91, 0x0e, 0xf4, 0xd7, 0x90, 0x5b, 0xdf, 0x5a, 0xd8, 0x90, 0x93,
	0x82, 0x92, 0x59, 0x1a, 0xc0, 0xd9, 0x5e, 0xea, 0x3b, 0x8b, 0x35, 0x59, 0xdc, 0x87, 0x54, 0x09,
	0x19, 0x43, 0x67, 0x46, 0xb6, 0xb1, 0x58, 0xd6, 0x9f, 0x19, 0xb8, 0xbe, 0x39, 0x3f, 0x3b, 0xcd,
	0x62, 0x14, 0xd1, 0x6c, 0x4f, 0xad, 0xf9, 0xe9, 0x19, 0x8a, 0xd0, 0x11, 0x31, 0x2a, 0x4c, 0xcf,
	0x4a, 0x1a, 0xdf, 0x8b, 0xa4, 0xda, 0x46, 0x8e, 0x70, 0x98, 0x85, 0x6d, 0x40, 0xfa, 0x86, 0x54,
	0x22, 0x40, 0xde, 0xe1, 0xc8, 0x6b, 0xd6, 0xba, 0xd5, 0xac, 0x6e, 0x35, 0xd9, 0x42, 0xd7, 0x59,
	0xbf, 0xc5, 0xf6, 0xfc, 0x4f, 0x10, 0xe0, 0x0b, 0x40, 0xee, 0xd2, 0xa3, 0x93, 0xb5, 0xc2, 0xe0,
	0x64, 0x8d, 0x4c, 0xce, 0xbc, 0x31, 0x8d, 0x1e, 0x90, 0x25, 0x95, 0x40, 0x50, 0x2b, 0x6a, 0xea,
	0x2e, 0xbb, 0xe0, 0x5b, 0xb2, 0xa9, 0xee, 0xda, 0x09, 0x04, 0xee, 0x8a, 0xa9, 0xb2, 0x94, 0x47,
	0x9e, 0x66, 0x52, 0x9f, 0x2c, 0x2b, 0xe4, 0x98, 0xa9, 0x5a, 0x49, 0xd3, 0x1f, 0x5d, 0x8a, 0xae,
	0x09, 0xee, 0x75, 0xc3, 0x5f, 0x1e, 0xc6, 0x9e, 0x21, 0x37, 0x7e, 0x5a, 0xe4, 0xc6, 0x54, 0xf6,
	0x73, 0xa1, 0x90, 0xbe, 0x9f, 0x71, 0xcb, 0x39, 0xc7, 0xad, 0xa9, 0xbd, 0x60, 0xb9, 0x5c, 0x9b,
	0x76, 0xd3, 0x94, 0xab, 0x8c, 0x4e, 0xa6, 0x2c, 0x7b, 0x4b, 0xca, 0x02, 0x21, 0x52, 0xb5, 0xe2,
	0x7a, 0xa9, 0x59, 0xdd, 0xda, 0xbe, 0xcc, 0x54, 0xee, 0x35, 0x53, 0xa0, 0xfc, 0x2c, 0x47, 0x79,
	0x43, 0x62, 0xe3, 0x47, 0xe9, 0x9f, 0x69, 0x72, 0x2f, 0x69, 0x93, 0x54, 0x52, 0x48, 0x42, 0x11,
	0x70, 0xa5, 0xa7, 0x29, 0xbb, 0x2b, 0x79, 0x63, 0x9e, 0x39, 0xf3, 0xc6, 0xb7, 0xf4, 0x03, 0xa9,
	0x28, 0x08, 0x21, 0x40, 0x99, 0x9a, 0xf7, 0xdc, 0xbe, 0xe8, 0xdc, 0xdc, 0x87, 0xb0, 0x6d, 0xb4,
	0x43, 0xfe, 0x28, 0xf2, 0xc6, 0x4c, 0xfa, 0x8e, 0x54, 0x10, 0xa2, 0x24, 0xe4, 0x08, 0xe6, 0x45,
	0x37, 0xcf, 0xdf, 0xc2, 0x57, 0xb2, 0xf3, 0xda, 0x08, 0xf4, 0x92, 0x8c, 0x5d, 0x1d, 0x9d, 0x7a,
	0x63, 0x20, 0xfd, 0x62, 0x91, 0xd5, 0xbe, 0x0c, 0xb3, 0x08, 0x1e, 0x87, 0x5c, 0x44, 0xa3, 0x0c,
	0x55, 0x5b, 0xd2, 0x2e, 0x3f, 0xf8, 0x4f, 0xa5, 0x7c, 0x14, 0x85, 0x10, 0xe3, 0xfe, 0x84, 0xe1,
	0xde, 0x35, 0xf5, 0x56, 0xf7, 0xe7, 0x80, 0xbd, 0xb9, 0xe5, 0xe8, 0x0e, 0xa9, 0x2a, 0x48, 0xfb,
	0x22, 0x80, 0x97, 0x3c, 0x82, 0x5a, 0x79, 0xdd, 0x6a, 0x5e, 0x75, 0x6f, 0x1b, 0x50, 0xb5, 0x3d,
	0xb9, 0xf2, 0xa6, 0xf3, 0x1a, 0x5f, 0x2d, 0x72, 0x6b, 0x66, 0x6b, 0xe9, 0x13, 0x42, 0xa5, 0x9f,
	0xa7, 0x41, 0xe7, 0xe9, 0xf0, 0x17, 0x17, 0x32, 0xd6, 0xaf, 0x58, 0x72, 0xef, 0x0c, 0x4e, 0xd6,
	0xe8, 0xde, 0xcc, 0xad, 0x37, 0x47, 0x41, 0x37, 0xa6, 0x76, 0xa0, 0xa8, 0x77, 0x60, 0x6c, 0xe5,
	0xec, 0x1e, 0xb8, 0xf7, 0x8f, 0x4e, 0xed, 0xc2, 0xf1, 0xa9, 0x5d, 0xf8, 0x75, 0x6a, 0x17, 0x3e,
	0x0f, 0x6c, 0xeb, 0x68, 0x60, 0x5b, 0xc7, 0x03, 0xdb, 0xfa, 0x3d, 0xb0, 0xad, 0x6f, 0x7f, 0xec,
	0xc2, 0xc1, 0x15, 0xb3, 0x92, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x28, 0x58, 0x61, 0xec, 0xf5,
	0x05, 0x00, 0x00,
}
