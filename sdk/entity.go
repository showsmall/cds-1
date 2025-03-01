package sdk

import (
	"github.com/rockbears/yaml"
	"regexp"
	"time"
)

const (
	EntityTypeWorkerModelTemplate = "WorkerModelTemplate"
	EntityTypeWorkerModel         = "WorkerModel"

	EntityNamePattern = "^[a-zA-Z0-9.-_-]{1,}$"
)

type ShortEntity struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Branch string `json:"branch"`
}

type Entity struct {
	ID                  string    `json:"id" db:"id"`
	ProjectKey          string    `json:"project_key" db:"project_key"`
	ProjectRepositoryID string    `json:"project_repository_id" db:"project_repository_id"`
	Type                string    `json:"type" db:"type"`
	Name                string    `json:"name" db:"name"`
	Branch              string    `json:"branch" db:"branch"`
	Commit              string    `json:"commit" db:"commit"`
	LastUpdate          time.Time `json:"last_update" db:"last_update"`
	Data                string    `json:"data" db:"data"`
}

type Lintable interface {
	Lint() []error
	GetName() string
}

func ReadEntityFile[T Lintable](directory, fileName string, content []byte, out *[]T, t string, analysis ProjectRepositoryAnalysis) ([]Entity, MultiError) {
	namePattern, err := regexp.Compile(EntityNamePattern)
	if err != nil {
		return nil, []error{WrapError(err, "unable to compile regexp %s", namePattern)}
	}

	if err := yaml.UnmarshalMultipleDocuments(content, out); err != nil {
		return nil, []error{WrapError(err, "unable to read %s%s", directory, fileName)}
	}
	var entities []Entity
	for _, o := range *out {
		if err := o.Lint(); err != nil {
			return nil, err
		}
		entities = append(entities, Entity{
			Data:                string(content),
			Name:                o.GetName(),
			Branch:              analysis.Branch,
			Commit:              analysis.Commit,
			ProjectKey:          analysis.ProjectKey,
			ProjectRepositoryID: analysis.ProjectRepositoryID,
			Type:                t,
		})
		if !namePattern.MatchString(o.GetName()) {
			return nil, []error{WrapError(ErrInvalidData, "name %s doesn't match %s", o.GetName(), EntityNamePattern)}
		}
	}
	return entities, nil
}
