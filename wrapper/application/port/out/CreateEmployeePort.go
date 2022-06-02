package out

import (
	"github.com/dignelidxdx/HackthonGo/wrapper/domain"
)

type CreateEmployee interface {
	CreateEmployee(domain.Employee) (domain.Employee, error)
}
