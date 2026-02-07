package cmd

import (
	"fmt"

	"os"
	"os/exec"
	"runtime"

	"github.com/HIMASAKTA-DEV/himasakta-backend/db"
	"github.com/HIMASAKTA-DEV/himasakta-backend/db/migrations"
	seeders "github.com/HIMASAKTA-DEV/himasakta-backend/db/seeder"
	mylog "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/logger"
	"gorm.io/gorm"
)

func Commands() error {
	db := db.New()
	if err := getParams(db); err != nil {
		return err
	}

	return nil
}

func getParams(db *gorm.DB) error {
	migrate := false
	seeder := false
	watch := false

	for _, arg := range os.Args[1:] {
		if arg == "--migrate" {
			migrate = true
		}
		if arg == "--seeder" {
			seeder = true
		}
		if arg == "--watch" {
			watch = true
		}
	}
	if migrate {
		if err := migrations.Migrate(db); err != nil {
			return fmt.Errorf("migration failed: %w", err)
		}
		mylog.Infof("Migration completed successfully")
	}

	if seeder {
		if err := seeders.Seeding(db); err != nil {
			return fmt.Errorf("seeding failed: %w", err)
		}
		mylog.Infof("Seeder completed successfully")
	}

	if watch {
		if err := runWatch(); err != nil {
			return fmt.Errorf("watching failed: %w", err)
		}
		mylog.Infof("Start watching program")
		os.Exit(0)
	}

	if seeder || watch || migrate {
		os.Exit(0)
	}

	return nil
}

func runWatch() error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", "air -c .air.windows.toml")
	case "linux", "darwin":
		cmd = exec.Command("bash", "-c", "air -c .air.linux.toml")
	default:
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		mylog.Errorf("Error running command: %s", err)
		return err
	}

	mylog.Infoln("Command executed successfully")
	return nil
}

