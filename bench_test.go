package computeexample

import (
	"runtime"
	"testing"
)

func createData(nop int) (Data, []float32) {
	img := make([]float32, nop*nop*8*8*4)
	d := Data{
		ImgData:  &(img[0]),
		ImgWidth: uint(nop * 8),
	}
	return d, img
}

func TestSame(t *testing.T) {
	// Test to ensure that we get the same data from both the
	// cpu implementation and when calling the core shader
	nop := 8
	d1, i1 := createData(nop)
	_, i2 := createData(nop)

	// run it over the shader
	s, e := New(1)
	if e != nil {
		t.Error(e)
		t.FailNow()
	}
	e = s.Dispatch(d1, nop, nop, 1)
	if e != nil {
		t.Error(e)
		t.FailNow()
	}
	s.Free()

	// run the cpu implementation
	cpuImplementation(nop, i2)

	// compare to ensure that the results are identical
	for i := range i1 {
		if i1[i] != i2[i] {
			t.Error("results not identical at", i, i1[i], i2[i])
			t.FailNow()
		}
	}
}

func BenchmarkGo128x128(b *testing.B) {
	nop := 128 / 8
	_, i1 := createData(nop)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpuImplementation(nop, i1)
	}
	b.StopTimer()

	runtime.KeepAlive(i1)
}

func BenchmarkShader128x128(b *testing.B) {
	nop := 128 / 8
	d1, i1 := createData(nop)

	s, _ := New(1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Dispatch(d1, nop, nop, 1)
	}
	b.StopTimer()

	s.Free()
	runtime.KeepAlive(i1)
}

func BenchmarkGo2048x2048(b *testing.B) {
	nop := 2048 / 8
	_, i1 := createData(nop)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpuImplementation(nop, i1)
	}
	b.StopTimer()

	runtime.KeepAlive(i1)
}

func BenchmarkShader2048x2048(b *testing.B) {
	nop := 2048 / 8
	d1, i1 := createData(nop)

	s, _ := New(1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Dispatch(d1, nop, nop, 1)
	}
	b.StopTimer()

	s.Free()
	runtime.KeepAlive(i1)
}

func cpuImplementation(nop int, arr []float32) {
	for gy := 0; gy < nop; gy++ {
		for gx := 0; gx < nop; gx++ {
			for ly := 0; ly < 8; ly++ {
				for lx := 0; lx < 8; lx++ {
					px := gx*8 + lx
					py := gy*8 + ly
					x := int32(px)
					y := int32(py)
					z := int32(px) - 5
					w := int32(py) - 5
					for i := 0; i < 100; i++ {
						x = x * (x + 1)
						y = y * (y + 1)
						z = z * (z + 1)
						w = w * (w + 1)
					}
					index := py*4*(nop*8) + px*4
					if px%2+py%2 == 0 {
						arr[index] = float32(x)
						arr[index+1] = float32(y)
						arr[index+2] = float32(z)
						arr[index+3] = float32(w)
					} else {
						arr[index] = 0
						arr[index+1] = 0
						arr[index+2] = 0
						arr[index+3] = 1
					}
				}
			}
		}
	}
}
