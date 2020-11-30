package fileload

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Manager struct {
	db             *sql.DB
	jobDefinitions []JobDefinition
}

func NewManager(db *sql.DB) (m *Manager, err error) {
	m = &Manager{
		db: db,
	}

	err = m.loadJobDefinitions()

	return m, err
}

func (m *Manager) loadJobDefinitions() (err error) {
	jobDefinitionPath := filepath.Join("./fileload", "job_definitions")
	files, err := ioutil.ReadDir(jobDefinitionPath)

	if err != nil {
		return
	}
	for _, f := range files {
		ext := filepath.Ext(filepath.Join(jobDefinitionPath, f.Name()))

		if ext != ".json" {
			continue
		}

		jsonFile, err := os.Open(filepath.Join(jobDefinitionPath, f.Name()))

		if err != nil {
			return err
		}

		var jobDef parseableJobDefinition
		jsonBytes, err := ioutil.ReadAll(jsonFile)

		if err != nil {
			return err
		}

		err = json.Unmarshal(jsonBytes, &jobDef)

		if err != nil {
			return err
		}

		normalizedDef, err := jobDef.normalize()

		if err != nil {
			return err
		}

		m.jobDefinitions = append(m.jobDefinitions, normalizedDef)
	}

	return err
}
