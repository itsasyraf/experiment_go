package migration

import (
	database "app/app/driver/core/database"
)

func JDDriver() {
	_, err := database.MainExec(
		// Added on 2025-09-03 11:00
		`CREATE TABLE IF NOT EXISTS jht_driver.jd_driver (driver_id uuid PRIMARY KEY); ` +
		`ALTER TABLE IF EXISTS jht_driver.jd_driver ADD COLUMN IF NOT EXISTS driver_login_id VARCHAR(255) NULL;` +
		`ALTER TABLE IF EXISTS jht_driver.jd_driver DROP CONSTRAINT IF EXISTS jd_driver_uq001;` +
		`ALTER TABLE IF EXISTS jht_driver.jd_driver ADD CONSTRAINT jd_driver_uq001 UNIQUE (driver_login_id);` +
		`ALTER TABLE IF EXISTS jht_driver.jd_driver ADD COLUMN IF NOT EXISTS driver_login_id_type_id CHAR(1) NOT NULL;` +
		`ALTER TABLE IF EXISTS jht_driver.jd_driver DROP CONSTRAINT IF EXISTS jd_driver_fk001;` +
		`ALTER TABLE IF EXISTS jht_driver.jd_driver ADD CONSTRAINT jd_driver_fk001 FOREIGN KEY (driver_login_id_type_id) REFERENCES jht_driver.jd_driver_login_id_type(driver_login_id_type_id);` +
		`ALTER TABLE IF EXISTS jht_driver.jd_driver ADD COLUMN IF NOT EXISTS driver_login_pin VARCHAR(255) NULL;` +
		`ALTER TABLE IF EXISTS jht_driver.jd_driver ADD COLUMN IF NOT EXISTS driver_name VARCHAR(255) NULL;` +
		`ALTER TABLE IF EXISTS jht_driver.jd_driver ADD COLUMN IF NOT EXISTS driver_phone VARCHAR(255) NULL;` +
		`ALTER TABLE IF EXISTS jht_driver.jd_driver ADD COLUMN IF NOT EXISTS driver_email VARCHAR(255) NULL;`,
	)
	if err != nil {
		panic(err)
	}
}
