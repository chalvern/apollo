package migrations

import (
	"github.com/chalvern/sugar"
	"github.com/jinzhu/gorm"

	"gopkg.in/gormigrate.v1"
)

var (
	// DefaultMigrator the default migrator
	// which implement the Migrator interface
	DefaultMigrator = &migrator{}

	migrations = []*gormigrate.Migration{}
)

// Migrator migrator
type Migrator interface {
	Migrate()
	MigrateTo(migrationID string)
	RollbackLast()
	RollbackTo(migrationID string)
}

type migrator struct{}

// Run migration
func (m *migrator) Migrate(db *gorm.DB) {
	gm := gormigrate.New(db, gormigrate.DefaultOptions, migrations)

	if err := gm.Migrate(); err != nil {
		sugar.Fatalf("Could not migrate: %v", err)
	}
	sugar.Info("Migration did run successfully")
}

// MigrateTo executes all migrations that did not run yet up to
// the migration that matches `migrationID`.
func (m *migrator) MigrateTo(db *gorm.DB, migrationID string) {
	gm := gormigrate.New(db, gormigrate.DefaultOptions, migrations)

	if err := gm.MigrateTo(migrationID); err != nil {
		sugar.Fatalf("Could not migrate to %s: %v", err)
	}
	sugar.Infof("Migration to %s did run successfully", migrationID)
}

// RollbackLast undo the last migration
func (m *migrator) RollbackLast(db *gorm.DB) {
	gm := gormigrate.New(db, gormigrate.DefaultOptions, migrations)
	if err := gm.RollbackLast(); err != nil {
		sugar.Fatalf("Could not rollback: %v", err)
	}
	sugar.Info("Rollback successfully")
}

// RollbackTo undoes migrations up to the given migration that matches the `migrationID`.
// Migration with the matching `migrationID` is not rolled back.
func (m *migrator) RollbackTo(db *gorm.DB, migrationID string) {
	gm := gormigrate.New(db, gormigrate.DefaultOptions, migrations)
	if err := gm.RollbackTo(migrationID); err != nil {
		sugar.Fatalf("Could not rollback to %s: %v", migrationID, err)
	}
	sugar.Infof("Rollback to %s successfully")
}
