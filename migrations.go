package scapler

import (
	"log"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (s *Scapler) MigrateUp(dsn string) error {
	m, err := migrate.New("file://"+s.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Up(); err != nil {
		log.Println("Error running up migration:", err)
		return err
	}
	return nil
}

func (s *Scapler) MigrateDownAll(dsn string) error {
	m, err := migrate.New("file://"+s.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Down(); err != nil {
		log.Println("Error running down migration:", err)
		return err
	}
	return nil
}

func (s *Scapler) Steps(n int, dsn string) error {
	m, err := migrate.New("file://"+s.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Steps(n); err != nil {
		log.Println("Error running steps migration:", err)
		return err
	}
	return nil
}

func (s *Scapler) MigrateForce(dsn string) error {
	m, err := migrate.New("file://"+s.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Force(-1); err != nil {
		log.Println("Error running force migration:", err)
		return err
	}
	return nil
}
