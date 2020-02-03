package nsenter

/*
#define _GNU_SOURCE
#include <unistd.h>
#include <errno.h>
#include <sched.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <fcntl.h>

// 一旦这个包被引用，则这个函数就会被自动执行
__attribute__((constructor)) void enter_namespace(void) {
	char *docker_pid;
	docker_pid = getenv("docker_pid");
	if (docker_pid) {
		//fprintf(stdout, "got docker_pid=%s\n", docker_pid);
	} else {
		//fprintf(stdout, "missing docker_pid env skip nsenter");
		return;
	}
	char *docker_cmd;
	docker_cmd = getenv("docker_cmd");
	if (docker_cmd) {
		//fprintf(stdout, "got docker_cmd=%s\n", docker_cmd);
	} else {
		//fprintf(stdout, "missing docker_cmd env skip nsenter");
		return;
	}
	int i;
	char nspath[1024];
	char *namespaces[] = { "ipc", "uts", "net", "pid", "mnt" };
	for (i=0; i<5; i++) {
		sprintf(nspath, "/proc/%s/ns/%s", docker_pid, namespaces[i]);
		int fd = open(nspath, O_RDONLY);
		close(fd);
	}
	int res = system(docker_cmd);
	exit(0);
	return;
}
*/
import "C"
