package bamboo

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	AgentId               = `${bamboo_agentId}`
	BuildWorkingDirectory = `${bamboo_build_working_directory}`
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

func Main(s *Spec) {
	run := flag.Bool("run", false, "run tasks locally")
	flag.Parse()

	color := func(c int, v ...interface{}) string {
		args := []interface{}{fmt.Sprintf("\x1b[%dm", c)}
		args = append(args, v...)
		args = append(args, "\x1b[0m")
		return fmt.Sprint(args...)
	}
	print := func(v ...interface{}) {
		log.Print(color(32, v...))
	}
	fatal := func(v ...interface{}) {
		log.Fatal(color(31, v...))
	}

	if *run {
		subst := map[string]string{
			AgentId: "12345",
		}
		for k, v := range subst {
			print("substitute ", k, " with ", v)
		}
		for _, stage := range s.Stages {
			for _, job := range stage.Jobs {
				script := func() string {
					file, err := ioutil.TempFile("", "bamboo")
					if err != nil {
						fatal(err)
					}
					defer file.Close()
					for _, script := range job.Scripts {
						for k, v := range subst {
							script = strings.Replace(script, k, v, -1)
						}
						fmt.Fprintln(file, script)
						print(script)
					}
					return file.Name()
				}()
				print("wrote ", script)
			}
		}
		return
	}

	output, err := s.Build()
	if err != nil {
		fatal(err)
	}
	fmt.Fprint(os.Stdout, output)
}
