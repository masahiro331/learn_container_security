docker run \
	--name tracee --rm -it \
	--pid=host --cgroupns=host --privileged \
	-v /etc/os-release:/etc/os-release-host:ro \
	-e LIBBPFGO_OSRELEASE_FILE=/etc/os-release-host \
	aquasec/tracee:latest
