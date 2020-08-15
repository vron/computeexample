package kernel

import (
	"encoding/binary"
	"math"
	"unsafe"
)

// Code generated DO NOT EDIT

type Bool struct {
	B bool
	_ [3]bool
}

var True = Bool{B: true}

var False = Bool{}

type Vec2 [2]float32

type Vec4 [4]float32

type Mat4 struct {
	Column0 Vec4
	Column1 Vec4
	Column2 Vec4
	Column3 Vec4
}

type Image2Drgba32f struct {
	Data  []float32
	_     [4]byte
	Width int32
}

type Triangle struct {
	Flag     Bool
	_        [4]byte
	Vertices [3]Vec2
}

type Example struct {
	Ts [4][2]Triangle
	M  Mat4
}

func (v Vec2) Alignof() int { return 8 }

func (v Vec2) Sizeof() int { return 8 }

func (v Vec4) Alignof() int { return 16 }

func (v Vec4) Sizeof() int { return 16 }

func (v Mat4) Alignof() int { return 16 }

func (v Mat4) Sizeof() int { return 64 }

func (v Image2Drgba32f) Alignof() int { return 8 }

func (v Image2Drgba32f) Sizeof() int { return 32 }

func (v Triangle) Alignof() int { return 8 }

func (v Triangle) Sizeof() int { return 32 }

func (v Example) Alignof() int { return 16 }

func (v Example) Sizeof() int { return 320 }

func (v *Vec2) Encode(d []byte) {
	bo.PutUint32(d[0:], math.Float32bits(v[0]))
	bo.PutUint32(d[4:], math.Float32bits(v[1]))
}

func (v *Vec4) Encode(d []byte) {
	bo.PutUint32(d[0:], math.Float32bits(v[0]))
	bo.PutUint32(d[4:], math.Float32bits(v[1]))
	bo.PutUint32(d[8:], math.Float32bits(v[2]))
	bo.PutUint32(d[12:], math.Float32bits(v[3]))
}

func (v *Mat4) Encode(d []byte) {
	(v.Column0).Encode(d[0:])
	(v.Column1).Encode(d[16:])
	(v.Column2).Encode(d[32:])
	(v.Column3).Encode(d[48:])
}

func (v *Image2Drgba32f) Encode(d []byte) {
	for i0 := 0; i0 < len(v.Data); i0++ {
		bo.PutUint32(d[0+i0*4:], math.Float32bits(v.Data[i0]))
	}
	bo.PutUint32(d[28:], uint32(v.Width))
}

func (v *Triangle) Encode(d []byte) {
	bo.PutUint32(d[0:], uint32(cBool((v.Flag).B)))
	for i0 := 0; i0 < 3; i0++ {
		(v.Vertices[i0]).Encode(d[8+i0*8:])
	}
}

func (v *Example) Encode(d []byte) {
	for i0 := 0; i0 < 4; i0++ {
		for i1 := 0; i1 < 2; i1++ {
			(v.Ts[i0][i1]).Encode(d[(0)+i0*64+i1*32:])
		}
	}
	(v.M).Encode(d[256:])
}

func encodeData(v Data) (r DataRaw) {
	r.Img = v.Img
	return
}

func (v *Vec2) Decode(d []byte) {
	v[0] = math.Float32frombits(bo.Uint32(d[0:]))
	v[1] = math.Float32frombits(bo.Uint32(d[4:]))
}

func (v *Vec4) Decode(d []byte) {
	v[0] = math.Float32frombits(bo.Uint32(d[0:]))
	v[1] = math.Float32frombits(bo.Uint32(d[4:]))
	v[2] = math.Float32frombits(bo.Uint32(d[8:]))
	v[3] = math.Float32frombits(bo.Uint32(d[12:]))
}

func (v *Mat4) Decode(d []byte) {
	(v.Column0).Decode(d[0:])
	(v.Column1).Decode(d[16:])
	(v.Column2).Decode(d[32:])
	(v.Column3).Decode(d[48:])
}

func (v *Image2Drgba32f) Decode(d []byte) {
	for i0 := 0; i0 < len(v.Data); i0++ {
		v.Data[i0] = math.Float32frombits(bo.Uint32(d[0+i0*4:]))
	}
	v.Width = int32(bo.Uint32(d[28:]))
}

func (v *Triangle) Decode(d []byte) {
	v.Flag = Bool{B: iBool(bo.Uint32(d[0:]))}
	for i0 := 0; i0 < 3; i0++ {
		(v.Vertices[i0]).Decode(d[8+i0*8:])
	}
}

func (v *Example) Decode(d []byte) {
	for i0 := 0; i0 < 4; i0++ {
		for i1 := 0; i1 < 2; i1++ {
			(v.Ts[i0][i1]).Decode(d[(0)+i0*64+i1*32:])
		}
	}
	(v.M).Decode(d[256:])
}

func decodeData(r DataRaw) (d Data) {
	d.Img = r.Img
	return
}

// ensure that the go structs mem layout match those in the C shader code
func init() {
	if unsafe.Sizeof(Vec2{}) != 8 {
		panic("sizeof(Vec2) != 8")
	}
	if unsafe.Sizeof(Vec4{}) != 16 {
		panic("sizeof(Vec4) != 16")
	}
	if unsafe.Sizeof(Mat4{}) != 64 {
		panic("sizeof(Mat4) != 64")
	}
	if unsafe.Offsetof(Mat4{}.Column0) != 0 {
		panic("offsetof(Mat4.Column0) != 0")
	}
	if unsafe.Offsetof(Mat4{}.Column1) != 16 {
		panic("offsetof(Mat4.Column1) != 16")
	}
	if unsafe.Offsetof(Mat4{}.Column2) != 32 {
		panic("offsetof(Mat4.Column2) != 32")
	}
	if unsafe.Offsetof(Mat4{}.Column3) != 48 {
		panic("offsetof(Mat4.Column3) != 48")
	}
	if unsafe.Sizeof(Image2Drgba32f{}) != 32 {
		panic("sizeof(Image2Drgba32f) != 32")
	}
	if unsafe.Offsetof(Image2Drgba32f{}.Data) != 0 {
		panic("offsetof(Image2Drgba32f.Data) != 0")
	}
	if unsafe.Offsetof(Image2Drgba32f{}.Width) != 28 {
		panic("offsetof(Image2Drgba32f.Width) != 28")
	}
	if unsafe.Sizeof(Triangle{}) != 32 {
		panic("sizeof(Triangle) != 32")
	}
	if unsafe.Offsetof(Triangle{}.Flag) != 0 {
		panic("offsetof(Triangle.Flag) != 0")
	}
	if unsafe.Offsetof(Triangle{}.Vertices) != 8 {
		panic("offsetof(Triangle.Vertices) != 8")
	}
	if unsafe.Sizeof(Example{}) != 320 {
		panic("sizeof(Example) != 320")
	}
	if unsafe.Offsetof(Example{}.Ts) != 0 {
		panic("offsetof(Example.Ts) != 0")
	}
	if unsafe.Offsetof(Example{}.M) != 256 {
		panic("offsetof(Example.M) != 256")
	}
}

var bo = binary.LittleEndian

func cBool(b bool) int32 {
	if b {
		return 1
	}
	return 0
}

func writeByte(b []byte, by byte) {
	b[0] = by
}

func readByte(b []byte) byte {
	return b[0]
}

func iBool(v uint32) bool {
	return !(v == 0)
}
