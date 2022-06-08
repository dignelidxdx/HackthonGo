package persistence

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/dignelidxdx/HackthonGo/wrapper/config"
	"github.com/dignelidxdx/HackthonGo/wrapper/domain"
)

func Test_employeeRepository_GetAllEmployees(t *testing.T) {
	type fields struct {
		db     *sql.DB
		config config.DBConfiguration
	}
	tests := []struct {
		name    string
		fields  fields
		want    []domain.Employee
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &employeeRepository{
				db:     tt.fields.db,
				config: tt.fields.config,
			}
			got, err := r.GetAllEmployees()
			if (err != nil) != tt.wantErr {
				t.Errorf("employeeRepository.GetAllEmployees() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("employeeRepository.GetAllEmployees() = %v, want %v", got, tt.want)
			}
		})
	}
}
