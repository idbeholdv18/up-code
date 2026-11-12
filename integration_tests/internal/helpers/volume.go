package helpers

import "github.com/testcontainers/testcontainers-go"

func NewVolume(volumeName string, target testcontainers.ContainerMountTarget) (testcontainers.ContainerMount, error) {
	mnt := testcontainers.ContainerMount{
		Source: testcontainers.GenericVolumeMountSource{
			Name: volumeName,
		},
		Target: testcontainers.ContainerMountTarget(target),
	}

	return mnt, nil
}
