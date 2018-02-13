//go:generate go run example.go
package main

import (
	"fmt"
	. "github.com/predmond/go-bamboo-spec"
)

func main() {
	yaml, err := NewSpec(
		NewProject("DRAGON",
			NewPlan("SLAYER", "Dragon Slayer Quest"),
		),
		NewStage(
			NewJob(
				`echo 'Going to slay the red dragon, watch me'`,
				`sleep 1`,
				`echo 'Victory!'`,
			),
		),
	).Build()
	if err != nil {
		panic(err)
	}
	fmt.Print(yaml)
}
