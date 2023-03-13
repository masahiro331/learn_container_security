#include <fcntl.h>
#include <stdio.h>
#include <unistd.h>
#include <string.h>
#include <linux/seccomp.h>
#include <sys/prctl.h>


int main(int argc, char **argv)
{
	printf("output.txt をopen");
	int output = open("output.txt", O_WRONLY);
	const char *val = "test";

	printf("seccompをstrictモードで設定\n");
	printf("strictモードは exit, read, write, sigreturn のみ許可\n");
	prctl(PR_SET_SECCOMP, SECCOMP_MODE_STRICT);

	printf("output.txtにtestと書き込む(write)\n");
	write(output, val, strlen(val)+1);

	printf("output.txtを再度開く(open)\n");
	open("output.txt", O_RDONLY);

	printf("プログラムを終了\n");
}

