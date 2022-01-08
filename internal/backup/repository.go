package internal

import (
	"errors"
	"log"

	"github.com/dignelidxdx/HackthonGo/internal/models"
	"github.com/dignelidxdx/HackthonGo/pkg/db"
)

type ConditionEnum int

const (
	Locked   ConditionEnum = 0
	Inactive               = 1
	Active                 = 2
)

type BackUpRepository interface {
	isLocked(name string) (bool, error)
	SaveToLock(name string, id int) (bool, error)
}

type backUpRepository struct {
}

func NewBackUpRepository() BackUpRepository {
	return &backUpRepository{}
}

func (r *backUpRepository) isLocked(name string) (bool, error) {
	db := db.StorageDB

	backup := models.BackUp{}
	rows, err := db.Query("SELECT id, is_update_data, field FROM backups WHERE field = ?", name)

	if err != nil {
		log.Fatal(err)
		return true, err
	}

	for rows.Next() {

		err = rows.Scan(&backup.ID, &backup.IsUpdatedData, &backup.Field)
		if err != nil {
			log.Fatal(err)
			return true, err
		}
	}

	if backup.IsUpdatedData == "1" {
		return true, nil
	} else {
		return false, nil
	}
}

func (r *backUpRepository) SaveToLock(name string, id int) (bool, error) {
	db := db.StorageDB

	stmt, err := db.Prepare("UPDATE backups SET is_update_data = 1 WHERE field = ? AND id = ?")
	if err != nil {
		log.Fatal("err", err)
	}

	defer stmt.Close()
	result, err := stmt.Exec(name, id)
	if err != nil {
		return false, err
	}
	filasActualizadas, _ := result.RowsAffected()

	if filasActualizadas == 0 {
		return false, errors.New("No se encontro el codigo")
	}
	return true, nil
}
