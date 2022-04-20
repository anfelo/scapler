package main

func doMigrate(arg2, arg3 string) error {
	dsn := getDSN()

	switch arg2 {
	case "up":
		err := scap.MigrateUp(dsn)
		if err != nil {
			return err
		}
	case "down":
		if arg3 == "all" {
			err := scap.MigrateDownAll(dsn)
			if err != nil {
				return err
			}
		} else {
			err := scap.Steps(-1, dsn)
			if err != nil {
				return err
			}
		}
	case "reset":
		err := scap.MigrateDownAll(dsn)
		if err != nil {
			return err
		}
		err = scap.MigrateUp(dsn)
		if err != nil {
			return err
		}
	default:
		showHelp()
	}

	return nil
}
