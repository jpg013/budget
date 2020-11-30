package ingest

import "time"

type JobType string

const (
	CSVImportJob JobType = "csv_import"
)

type Job struct {
	Name        string
	Type        JobType
	CreatedAt   *time.Time
	StartedAt   *time.Time
	CompletedAt *time.Time
	// Status      JobStatus
}

// job_name: Discover All Activity
// file_patten: Discover-AllAvailable-[0-9]+.csv
// job_type: csv
// file_columns:
//   - name: Posted Date
//     key: posted_date
//     ordinal: 2
//     type: timestamp
//     args:
//       timestamp_format: 01/02/2006
//   - name: Description
//     key: description
//     ordinal: 3
//     type: string
//   - name: Amount
//     key: amount
//     ordinal: 4
//     type: float64
//   - name: Category
//     key: category
//     ordinal: 5
//     type: string
