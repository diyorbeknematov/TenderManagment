package casbin

import (
	"github.com/casbin/casbin/v2"
	"log/slog"
)

func CasbinEnforcer(logger *slog.Logger) (*casbin.Enforcer, error) {
	enforcer, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		logger.Error("Error creating Casbin enforcer", "error", err.Error())
		return nil, err
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		logger.Error("Error loading Casbin policy", "error", err.Error())
		return nil, err
	}

	return enforcer, nil
}
