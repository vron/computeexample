package computeexample

// #cgo LDFLAGS: -L${SRCDIR} kernel.a
// #include "shared.h"
import "C"

// Code generated DO NOT EDIT

import (
	"errors"
	"runtime"
	"strconv"
	"unsafe"
)

type kernel struct {
	k unsafe.Pointer
	dead bool
}

type Data struct {
	ImgData *float32
	ImgWidth uint
}

// New creates a new kernel instance that may retain memory created
// using malloc. In order to ensure this memory is deallocated please
// ensure to call k.Free(). If numCPU <= 0 the number of threads to use
// will be calculated automatically.
func New(numCPU int) (k *kernel, err error) {
	k = &kernel{}
	if numCPU <= 0 {
		numCPU = runtime.NumCPU()+2
	}
	k.k = C.cpt_new_kernel(C.int(numCPU));
	if k.k == nil {
		return nil, errors.New("failed to create kernel structure")
	}
	runtime.SetFinalizer(k, freeKernel)
	return k, nil
}

// Dispatch a kernel calculation, with the given global work group sizes
// in x, y, and z direction respectively. The data proviced in bind is bound
// to the kernel during this call. It is the callers responsibility that the
// data provided in bind matches the kernel's assumptions given the work
// group size.
func (k *kernel) Dispatch(bind Data, numx, numy, numz int) error {
	if k.dead {
		panic("cannot use a kernel where Free() has been called")
	}
	cbind := C.kernel_data{
		imgData: (*C.float)(bind.ImgData),
		imgWidth: (C.uint)(bind.ImgWidth),
	}
	errno := C.cpt_dispatch_kernel(k.k, cbind, C.int(numx), C.int(numy), C.int(numz))
	return mapErrno(int(errno))
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
	C.cpt_free_kernel(k.k);
}

var dispatchErrors = map[int]string{
	1: "unspecified error",
}

func mapErrno(errno int) error {
	if errno == 0 {
		return nil
	}
	v, ok :=dispatchErrors[errno]
	if !ok {
		v = "unknown error code"
	}
	return errors.New(v + ": " + strconv.Itoa(errno))
}
