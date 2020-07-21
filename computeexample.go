// Package computeexample is a wrapper to execute a particular GLSL compute shader
package computeexample

// #cgo darwin LDFLAGS: -L${SRCDIR} -L. build/shader.so
// #cgo linux LDFLAGS: -L${SRCDIR}/build -L. build/shader.so
// #cgo windows LDFLAGS: -L. -lshader
// #include "shared.h"
import "C"

// Code generated DO NOT EDIT

import (
	"encoding/binary"
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"unsafe"
)

type kernel struct {
	k    unsafe.Pointer
	dead bool
}

type Image2Drgba32f struct {
	Data  []byte
	Width int32
}

type Data struct {
	Img Image2Drgba32f
}

// New creates a kernel using at most numCPU+1 threads. If numCPU <= 0 the
// number of threads to use will be calculated automatically. All kernels
// must be explicitly freed using kernel.Free to avoid memory leaks.
func New(numCPU int) (k *kernel, err error) {
	k = &kernel{}
	if numCPU <= 0 {
		numCPU = runtime.NumCPU() + 2
	}
	k.k = C.cpt_new_kernel(C.int(numCPU))
	if k.k == nil {
		return nil, errors.New("failed to create kernel structure")
	}
	runtime.SetFinalizer(k, freeKernel)
	return k, nil
}

// Dispatch a kernel calculation of the specified size. The caller must ensure
// that the data provided in bind matches the kernel's assumptions and that any
// []byte field represents properly aligned data. Not data in bind must
// be accessed (read or write) until Dispatch returns.
func (k *kernel) Dispatch(bind Data, numGroupsX, numGroupsY, numGroupsZ int) error {
	if k.dead {
		panic("cannot use a kernel where Free() has been called")
	}
	cbind := C.cpt_data{
		img: C.cpt_image2Drgba32f{
			data:  unsafe.Pointer(&bind.Img.Data[0]),
			width: (C.int32_t)(bind.Img.Width),
		},
	}
	if err := ensureLength("bind.Img.Data", len(bind.Img.Data), 8, -1); err != nil {
		return err
	}

	errno := C.cpt_dispatch_kernel(k.k, cbind, C.int(numGroupsX), C.int(numGroupsY), C.int(numGroupsZ))
	if errno.code == 0 {
		return nil
	}
	errstr := C.GoString(errno.msg)
	return errors.New(strconv.Itoa(int(errno.code)) + ": " + errstr)
}

// Free dealocates any data allocated by the underlying kernel. Note that
// a kernel on which Free has been called can no longer be used.
func (k *kernel) Free() {
	freeKernel(k)
}

func freeKernel(k *kernel) {
	if k.dead {
		return
	}
	k.dead = true
	C.cpt_free_kernel(k.k)
}

func cBool(b bool) int32 {
	if b {
		return 1
	}
	return 0
}

func ensureLength(f string, l, s, arr int) error {
	if arr > 0 {
		if l != s*arr {
			return fmt.Errorf("bad data for %v, expected length %v*%v=%v but got %v", f, s, arr, s*arr, l)
		}
	}
	if arr < 0 {
		if l%s != 0 {
			return fmt.Errorf("bad data for %v, expected length to be multiple of %v but got %v", f, s, l)
		}
	}
	return nil
}

var bo = binary.LittleEndian
