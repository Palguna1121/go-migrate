package templates

var JobsMigrationTemplate = `package migrations

import (
	"github.com/Palguna1121/go-migrate/config"
	"github.com/Palguna1121/go-migrate/pkg/interfaces"
	"github.com/Palguna1121/go-migrate/pkg/lib/%[1]s"
)

func init() {
	config.Migrations = append(config.Migrations, CreateJobsTable())
}

type JobsTable struct{}

func CreateJobsTable() interfaces.Migration {
	return &JobsTable{}
}

func (t *JobsTable) Up() error {
	if err := mysql.Schema.Create("jobs", func(table interfaces.Blueprint) {
		table.Id("id", 11)
		table.String("queue", 255).Index()
		table.Text("payload")
		table.Integer("attempts", 11).Default(0)
		table.DateTime("reserved_at").Nullable()
		table.DateTime("available_at")
		table.DateTime("created_at")
		table.Timestamps()
	}); err != nil {
		return err
	}

	if err := mysql.Schema.Create("job_batches", func(table interfaces.Blueprint) {
		table.Id("id", 11)
		table.String("name", 255)
		table.Integer("total_jobs", 11)
		table.Integer("pending_jobs", 11)
		table.Integer("failed_jobs", 11)
		table.Text("failed_job_ids")
		table.Text("options").Nullable()
		table.DateTime("cancelled_at").Nullable()
		table.DateTime("created_at")
		table.DateTime("finished_at").Nullable()
		table.Timestamps()
	}); err != nil {
		return err
	}

	if err := mysql.Schema.Create("failed_jobs", func(table interfaces.Blueprint) {
		table.Id("id", 11)
		table.String("uuid", 36).Unique()
		table.String("connection", 255)
		table.String("queue", 255)
		table.Text("payload")
		table.Text("exception")
		table.DateTime("failed_at")
		table.Timestamps()
	}); err != nil {
		return err
	}

	return nil
}

func (t *JobsTable) Down() error {
	if err := mysql.Schema.DropIfExists("jobs"); err != nil {
		return err
	}
	if err := mysql.Schema.DropIfExists("job_batches"); err != nil {
		return err
	}
	if err := mysql.Schema.DropIfExists("failed_jobs"); err != nil {
		return err
	}
	return nil
}

`
