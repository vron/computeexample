// Package kernel is a wrapper to execute a particular GLSL compute shader
package kernel

/*
#cgo darwin LDFLAGS: -L${SRCDIR} -L. build/shader.so
#cgo linux LDFLAGS: -L${SRCDIR}/build -L. build/shader.so
#cgo windows LDFLAGS: -L. -lshader

#include "shared.h"
struct cpt_error_t wrap_dispatch(void *k,
                                 cpt_image2Drgba32f	 img,
                                 int32_t x, int32_t y, int32_t z) {
	cpt_data d;
	d.img = img;
	return cpt_dispatch_kernel(k, d, x, y, z);
}
*/
import "C"
import (
	"errors"
	"runtime"
	"strconv"
	"unsafe"
)

// Code generated DO NOT EDIT

// AlignedSlice returns a byte slice where the first element has a minimum
// alignment of align and a length if size.
func AlignedSlice(size, align int) (b []byte) {
	if align < 1 {
		panic("align must be > 0")
	}
	b = make([]byte, size+align-1)
	adr := uintptr(unsafe.Pointer(&b[0]))
	diff := 0
	if int(adr)%align != 0 {
		diff = align - int(adr)%align
	}
	return b[diff : diff+size]
}

type Kernel struct {
	k    unsafe.Pointer
	dead bool
}

// New creates a Kernel using at most numCPU+1 threads. If numCPU <= 0 the
// number of threads to use will be calculated automatically. All kernels
// must be explicitly freed using Kernel.Free to avoid memory leaks.
func New(numCPU int, stackSize int) (k *Kernel, err error) {
	k = &Kernel{}
	if numCPU <= 0 {
		numCPU = runtime.NumCPU() + 2
	}
	k.k = C.cpt_new_kernel(C.int32_t(numCPU), C.int32_t(stackSize))
	if k.k == nil {
		return nil, errors.New("failed to create kernel structure")
	}
	runtime.SetFinalizer(k, freeKernel)
	return k, nil
}

// Free dealocates any data allocated by the underlying Kernel. Note that
// a kernel on which Free has been called can no longer be used.
func (k *Kernel) Free() {
	freeKernel(k)
}

func freeKernel(k *Kernel) {
	if k.dead {
		return
	}
	k.dead = true
	C.cpt_free_kernel(k.k)
}
func (v Image2Drgba32f) toC() C.cpt_image2Drgba32f {
	return C.cpt_image2Drgba32f{
		data:  (*C.float)(&v.Data[0]),
		width: (C.int32_t)(v.Width),
	}
}

type Data struct {
	Img Image2Drgba32f
}

type DataRaw struct {
	Img Image2Drgba32f
}

// Dispatch a Kernel calculation of the specified size. The caller must ensure
// that the data provided in bind matches the kernel's assumptions and that any
// []byte field represents properly aligned data. Not data in bind must
// be accessed (read or write) until Dispatch returns.
func (k *Kernel) Dispatch(bind Data, numGroupsX, numGroupsY, numGroupsZ int) error {
	if k.dead {
		panic("cannot use a Kernel where Free() has been called")
	}
	errno := C.wrap_dispatch(k.k,
		bind.Img.toC(), C.int(numGroupsX), C.int(numGroupsY), C.int(numGroupsZ))
	if errno.code == 0 {
		return nil
	}
	errstr := C.GoString(errno.msg)
	return errors.New(strconv.Itoa(int(errno.code)) + ": " + errstr)
}

func (k *Kernel) DispatchRaw(bind DataRaw, numGroupsX, numGroupsY, numGroupsZ int) error {
	if k.dead {
		panic("cannot use a Kernel where Free() has been called")
	}
	errno := C.wrap_dispatch(k.k,
		bind.Img.toC(), C.int(numGroupsX), C.int(numGroupsY), C.int(numGroupsZ))
	if errno.code == 0 {
		return nil
	}
	errstr := C.GoString(errno.msg)
	return errors.New(strconv.Itoa(int(errno.code)) + ": " + errstr)
}
