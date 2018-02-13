package bamboo

import (
	"gopkg.in/yaml.v2"
)

type Plan struct {
	Key  string
	Name string
}

func NewPlan(key, name string) *Plan {
	return &Plan{
		Key:  key,
		Name: name,
	}
}

type Project struct {
	Key  string
	Plan *Plan
}

func NewProject(key string, plan *Plan) *Project {
	return &Project{
		Key:  key,
		Plan: plan,
	}
}

type Job struct {
	Scripts      []string
	Interpreter  string        `yaml:",omitempty"`
	Requirements []string      `yaml:",omitempty"`
	Artifacts    []*Artifact   `yaml:",omitempty"`
	TestParsers  []*TestParser `yaml:"testParsers,omitempty"`
}

func NewJob(scripts ...string) *Job {
	return &Job{
		Scripts: scripts,
	}
}

func (j *Job) WithInterpreter(s string) *Job {
	j.Interpreter = s
	return j
}

func (j *Job) WithRequirements(v ...string) *Job {
	j.Requirements = append(j.Requirements, v...)
	return j
}

type Artifact struct {
	Name string
	Path string
}

func NewArtifact(name, path string) *Artifact {
	return &Artifact{
		Name: name,
		Path: path,
	}
}

func (j *Job) WithArtifacts(v ...*Artifact) *Job {
	j.Artifacts = append(j.Artifacts, v...)
	return j
}

type TestParser struct {
	Type        string
	TestResults []string `yaml:"testResults"`
}

func NewTestParser(typeName string, testResults ...string) *TestParser {
	return &TestParser{
		Type:        typeName,
		TestResults: testResults,
	}
}

func (j *Job) WithTestParsers(v ...*TestParser) *Job {
	j.TestParsers = append(j.TestParsers, v...)
	return j
}

type Stage struct {
	Jobs []*Job
}

func NewStage(jobs ...*Job) *Stage {
	return &Stage{
		Jobs: jobs,
	}
}

type Spec struct {
	Project *Project
	Stages  []*Stage
}

func NewSpec(project *Project, stages ...*Stage) *Spec {
	return &Spec{
		Project: project,
		Stages:  stages,
	}
}

func (s *Spec) Build() (string, error) {
	data, err := yaml.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
