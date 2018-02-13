package docker

import (
	"fmt"
	"strings"
)

type Run struct {
	image   string
	command string
	volumes []string
	envs    []string
}

func NewRun(image string) *Run {
	return &Run{
		image: image,
	}
}

func (b *Run) String() string {
	args := make([]string, 0, 2)
	args = append(args, "docker", "run")
	for _, volume := range b.volumes {
		args = append(args, "--volume", volume)
	}
	for _, env := range b.envs {
		args = append(args, "--env", env)
	}
	args = append(args, b.image)
	args = append(args, b.command)
	return strings.Join(args, " ")
}

func (b *Run) WithVolume(volume string) *Run {
	b.volumes = append(b.volumes, volume)
	return b
}

func (b *Run) WithEnv(name, value string) *Run {
	b.envs = append(b.envs, fmt.Sprint(name, "=", value))
	return b
}

func (b *Run) WithCommand(command string, args ...string) *Run {
	b.command = fmt.Sprint(command, strings.Join(args, " "))
	return b
}
