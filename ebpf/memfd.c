#define _GNU_SOURCE
#include <stdint.h>
#include <sys/mman.h>
#include <stdlib.h>
#include <unistd.h>
#include <stdio.h>

#define errExit(msg)    do { perror(msg); exit(EXIT_FAILURE); } while (0)

int main(int argc, char *argv[])
{
	int fd;
	char *name;
	char buf[128];

	fd = memfd_create("malicious_code", MFD_ALLOW_SEALING);

	if (fd == -1)
		errExit("memfd_create");

	if (ftruncate(fd, sizeof(buf) == -1))
		errExit("truncate");

	printf("PID: %jd; fd: %d; /proc/%jd/fd/%d\n", (intmax_t) getpid(), fd, (intmax_t) getpid(), fd);
	pause();

	exit(EXIT_SUCCESS);
}

