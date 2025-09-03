package migration

import (
	database "app/app/driver/core/database"
)

func JDDriverLoginIdType() {
	_, err := database.MainExec(
		// Added on 2025-09-03 11:00
		`CREATE TABLE IF NOT EXISTS jht_driver.jd_driver_login_id_type (driver_login_id_type_id CHAR(1) PRIMARY KEY); ` +
		`ALTER TABLE IF EXISTS jht_driver.jd_driver_login_id_type ADD COLUMN IF NOT EXISTS driver_login_id_type_name VARCHAR(255) NOT NULL; ` +
		`ALTER TABLE IF EXISTS jht_driver.jd_driver_login_id_type DROP CONSTRAINT IF EXISTS jd_driver_login_id_type_uq001; ` +
		`ALTER TABLE IF EXISTS jht_driver.jd_driver_login_id_type ADD CONSTRAINT jd_driver_login_id_type_uq001 UNIQUE (driver_login_id_type_name); ` +
		`INSERT INTO jht_driver.jd_driver_login_id_type VALUES ('E', 'Email') ON CONFLICT (driver_login_id_type_id) DO NOTHING; ` +
		`INSERT INTO jht_driver.jd_driver_login_id_type VALUES ('P', 'Phone') ON CONFLICT (driver_login_id_type_id) DO NOTHING; `,
	)
	if err != nil {
		panic(err)
	}
}
