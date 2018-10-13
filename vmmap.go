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

void vmmap() {
	pid_t pid;
	mach_port_t	task;
	kern_return_t err;
	mach_vm_address_t address;
	mach_vm_size_t size;
	vm_region_basic_info_data_t info;
	mach_msg_type_number_t count;
	mach_port_t object_name;
	char buffer[PATH_MAX];

	pid = getpid();
	err = task_for_pid(mach_task_self(), pid, &task);
	if (err != KERN_SUCCESS) {
		fprintf(stderr, "task_for_pid: error %d - %s\n", err, mach_error_string(err));
		return;
	}

	count = VM_REGION_BASIC_INFO_COUNT_64;
	for (address = VM_MIN_ADDRESS;; address += size) {
		err = mach_vm_region(task, &address, &size, VM_REGION_BASIC_INFO, (vm_region_info_t)&info, &count, &object_name);
		if (err != KERN_SUCCESS) {
			return;
		}

		if ((info.protection & VM_PROT_EXECUTE) == 0) {
			continue;
		}

		memset(buffer, 0, sizeof(buffer));
		proc_regionfilename(pid, address, &buffer, PATH_MAX);
		if (strlen(buffer) == 0) {
			continue;
		}

		printf("%016llx-%016llx %08x %s\n", address, address+size, 0, buffer);
	}
}
*/
import "C"
import (
	"log"
	"runtime"
)

func main() {
	if runtime.GOOS != "darwin" {
		log.Fatalf("OS %s not supported\n", runtime.GOOS)
	}
	C.vmmap()
}
