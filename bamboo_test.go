package bamboo

import (
	"testing"
)

func TestBamboo(t *testing.T) {
	NewSpec(
		NewProject("DRAGON",
			NewPlan("SLAYER", "Dragon Slayer Quest"),
		),
		NewStage(
			NewJob(
				`#!/bin/bash`,
				`echo 'Going to slay the red dragon, in bash'`,
			).WithInterpreter("shell"),
			NewJob(
				`echo 'Going to slay the red dragon, in powershell'`,
			).WithInterpreter("powershell"),
			NewJob(
				`echo 'Going to slay the red dragon, in cmd.exe'`,
			).WithInterpreter("cmd.exe"),
		),
	).Build()

	NewSpec(
		NewProject("DRAGON",
			NewPlan("SLAYER", "Dragon Slayer Quest"),
		),
		NewStage(
			NewJob(
				`echo 'Going to slay the red dragon, watch me'`,
			).WithRequirements(
				"isDragonLazy",
				"isDragonAsleep",
				"isCaveDeep",
			),
		),
	).Build()

	NewSpec(
		NewProject("DRAGON",
			NewPlan("SLAYER", "Dragon Slayer Quest"),
		),
		NewStage(
			NewJob(
				`echo 'Going to slay the red dragon, watch me'`,
			).WithArtifacts(
				NewArtifact("Blue dragon's head", "dragon/blue/head"),
				NewArtifact("Blue dragon's claw", "dragon/blue/claw"),
			),
		),
	).Build()

	NewSpec(
		NewProject("DRAGON",
			NewPlan("SLAYER", "Dragon Slayer Quest"),
		),
		NewStage(
			NewJob(
				`echo 'Going to slay the red dragon, watch me'`,
			).WithTestParsers(
				NewTestParser("mocha", "report.json"),
				NewTestParser("junit",
					`**/test-reports/*.xml`,
					`**/custom-test-reports/*.xml`),
			),
		),
	).Build()
}
