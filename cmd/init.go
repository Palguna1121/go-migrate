package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/Palguna1121/go-migrate/templates"
	"github.com/spf13/cobra"
	"golang.org/x/mod/modfile"
)

func getModuleName() (string, error) {
	goModBytes, err := ioutil.ReadFile("go.mod")
	if err != nil {
		return "", err
	}

	modName := modfile.ModulePath(goModBytes)

	return modName, nil
}

func createMigrateCMD(db string, moduleName string) error {
	if err := os.MkdirAll("cmd/migrate", os.ModePerm); err != nil {
		return fmt.Errorf("can't make directory. \n%v", err)
	}

	if err := ioutil.WriteFile("cmd/migrate/migrate.go", []byte(fmt.Sprintf(templates.MigrateTemplate, db, moduleName)), 0666); err != nil {
		return fmt.Errorf("can't create migrate file. \n%v", err)
	}

	return nil
}

func createMigrations(db string) error {
	if err := os.MkdirAll("migrations", os.ModePerm); err != nil {
		return fmt.Errorf("can't make directory. \n%v", err)
	}

	//users table
	if err := ioutil.WriteFile("migrations/2025_06_11_000000_create_users_table.go", []byte(fmt.Sprintf(templates.UserMigrationTemplate, db)), 0666); err != nil {
		return fmt.Errorf("can't create migration file. \n%v", err)
	}
	//cache table
	if err := ioutil.WriteFile("migrations/2025_06_11_000001_create_cache_table.go", []byte(fmt.Sprintf(templates.CacheMigrationTemplate, db)), 0666); err != nil {
		return fmt.Errorf("can't create migration file. \n%v", err)
	}
	//jobs table
	if err := ioutil.WriteFile("migrations/2025_06_11_000002_create_jobs_table.go", []byte(fmt.Sprintf(templates.JobsMigrationTemplate, db)), 0666); err != nil {
		return fmt.Errorf("can't create migration file. \n%v", err)
	}
	//personal access tokens table
	if err := ioutil.WriteFile("migrations/2025_06_11_000003_create_personal_access_tokens_table.go", []byte(fmt.Sprintf(templates.PersonalAccessTokensMigrationTemplate, db)), 0666); err != nil {
		return fmt.Errorf("can't create migration file. \n%v", err)
	}
	//permissions table
	if err := ioutil.WriteFile("migrations/2025_06_11_000004_create_permissions_table.go", []byte(fmt.Sprintf(templates.PermissionsMigrationTemplate, db)), 0666); err != nil {
		return fmt.Errorf("can't create migration file. \n%v", err)
	}

	return nil
}

func modTidy() error {
	if err := exec.Command("go", "mod", "tidy").Run(); err != nil {
		return err
	}

	return nil
}

func handleInit(db string) error {
	moduleName, err := getModuleName()
	if err != nil {
		return fmt.Errorf("can't get module name. \n%v", err)
	}

	if err := createMigrateCMD(db, moduleName); err != nil {
		return fmt.Errorf("can't create the migrate cmd. \n%v", err)
	}

	if err := createMigrations(db); err != nil {
		return fmt.Errorf("can't create the migration file. \n%v", err)
	}

	if err := modTidy(); err != nil {
		return fmt.Errorf("can't execute the mod tidy cmd. \n%v", err)
	}

	return nil
}

var initCmd = &cobra.Command{
	Use:   "init <db>",
	Short: "Init github.com/Palguna1121/go-migrate-cmd context in this project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db := args[0]
		fmt.Printf("Init migrate with %s.\n", db)
		if err := handleInit(db); err != nil {
			fmt.Println("ERROR: ", err)
			return
		}
		fmt.Println("Init migration ok.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
