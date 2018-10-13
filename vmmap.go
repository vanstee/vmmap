package main

/*
#include <errno.h>
#include <libproc.h>
#include <mach/mach_error.h>
#include <mach/mach_init.h>
#include <mach/mach_vm.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/syslimits.h>
#include <sys/types.h>
#include <unistd.h>

char buffer[PATH_MAX];

kern_return_t VmmapRegion(mach_vm_address_t *address, mach_vm_size_t *size, char *buffer) {
	pid_t pid;
	mach_port_t	task;
	kern_return_t err;
	vm_region_basic_info_data_t info;
	mach_msg_type_number_t count;
	mach_port_t object_name;

	memset(buffer, 0, PATH_MAX);

	pid = getpid();
	err = task_for_pid(mach_task_self(), pid, &task);
	if (err != KERN_SUCCESS) {
		return err;
	}

	count = VM_REGION_BASIC_INFO_COUNT_64;
	err = mach_vm_region(task, address, size, VM_REGION_BASIC_INFO, (vm_region_info_t)&info, &count, &object_name);
	if (err != KERN_SUCCESS) {
		return err;
	}

	if ((info.protection & VM_PROT_EXECUTE) == 0) {
		return KERN_SUCCESS;
	}

	proc_regionfilename(pid, *address, buffer, PATH_MAX);
	return KERN_SUCCESS;
}
*/
import "C"
import (
	"fmt"
	"log"
	"runtime"
)

var (
	address C.mach_vm_address_t
	size    C.mach_vm_size_t
	ret     C.kern_return_t
)

func main() {
	if runtime.GOOS != "darwin" {
		log.Fatalf("OS %s not supported\n", runtime.GOOS)
	}

	for address = C.VM_MIN_ADDRESS; ret != C.KERN_INVALID_ADDRESS; address = address + size {
		ret = C.VmmapRegion(&address, &size, &C.buffer[0])
		file := C.GoString(&C.buffer[0])
		if file != "" {
			fmt.Printf("%016x-%016x %08x %s\n", address, address+size, 0, file)
		}
	}
}
