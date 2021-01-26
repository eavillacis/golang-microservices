package cmd

import (
	"github.com/eavillacis/velociraptor/services/catalog/conf"	
	migrator "github.com/golang-migrate/migrate"
	log "github.com/sirupsen/logrus"

	"github.com/golang-migrate/migrate/database/postgres"

	// golang-migrate file dependency
	_ "github.com/golang-migrate/migrate/source/file"
	
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd, dropCmd, resetCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "db:migrate",
	Short: "Migrate catalog database structure.",
	Long:  "Migrate catalog database structure.",
	Run: func(cmd *cobra.Command, args []string) {
		migrate()
	},
}

var dropCmd = &cobra.Command{
	Use:   "db:drop",
	Short: "Warning: Drop database.",
	Long:  "Warning: Drop database.",
	Run: func(cmd *cobra.Command, args []string) {
		drop()
	},
}

var resetCmd = &cobra.Command{
	Use:   "db:reset",
	Short: "Reset database.",
	Long:  "Reset database.",
	Run: func(cmd *cobra.Command, args []string) {
		drop()
		migrate()
	},
}

// NewMigrator create a new instance of database with our custom config
func NewMigrator(db *gorm.DB) (*migrator.Migrate, error) {
	driver, err := postgres.WithInstance(db.DB(), &postgres.Config{})
	if err != nil {
		return nil, err
	}

	return migrator.NewWithDatabaseInstance("file://services/catalog/migrations", "postgres", driver)
}

func migrate() {
	config, err := conf.LoadConfig(configFile)
	if err != nil {
		log.WithError(err).Error("Failed to load configuration")
		return
	}

	conn, err := gorm.Open("postgres", config.DB.URL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %+v", err)
	}
	defer conn.Close()

	migrator, err := NewMigrator(conn)
	if err != nil {
		log.WithError(err).Error("Error opening database")
		return
	}

	log.Info("Migrating catalog...")

	err = migrator.Up()
	if err != nil {
		log.WithError(err).Error("Error migrating database")
		return
	}

	log.Info("Migrate finished.")
}

func drop() {
	config, err := conf.LoadConfig(configFile)
	if err != nil {
		log.WithError(err).Error("Failed to load configuration")
		return
	}

	conn, err := gorm.Open("postgres", config.DB.URL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %+v", err)
	}
	defer conn.Close()

	migrator, err := NewMigrator(conn)
	if err != nil {
		log.WithError(err).Error("Error opening database")
		return
	}

	log.Info("Dropping catalog...")

	err = migrator.Down()
	if err != nil {
		log.WithError(err).Error("Error dropping database")
		return
	}

	log.Info("Drop finished.")
}
