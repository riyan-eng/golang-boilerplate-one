package config

import (
	"fmt"
	"os"

	sqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/riyan-eng/golang-boilerplate-one/infrastructure"
	_ "github.com/lib/pq"
)

func NewEnforcer() *casbin.Enforcer {
	adapter, err := sqladapter.NewAdapter(infrastructure.SqlDB, "postgres", "permissions")
	if err != nil {
		fmt.Printf("casbin: failed to initialize adapter - %v \n", err)
		os.Exit(1)
	}
	enforce, err := casbin.NewEnforcer("casbin.conf", adapter)
	if err != nil {
		fmt.Printf("casbin: failed to create enforcer - %v \n", err)
		os.Exit(1)
	}
	// if hasPolicy := enforce.HasPolicy("ADMIN", "/auth/register", "(GET)|(POST)|(PUT)|(DELETE)"); !hasPolicy {
	// 	enforce.AddPolicy("ADMIN", "/auth/register", "(GET)|(POST)|(PUT)|(DELETE)")
	// }

	enforce.AddPolicy("ADMIN", "/pegawai*", "(GET)|(POST)|(PATCH)|(DELETE)")
	enforce.AddPolicy("EMPLOYEE", "/task*", "(GET)|(POST)|(PATCH)|(DELETE)")
	enforce.AddPolicy("PENGANGGARAN", "/indikator_kinerja/*", "(GET)|(POST)|(PATCH)|(PUT)|(DELETE)")
	enforce.AddPolicy("ADM_WIL", "/indikator_kinerja/*", "(GET)")
	enforce.AddPolicy("ADM_PEM", "/indikator_kinerja/*", "(GET)")
	enforce.AddPolicy("OTM_DRH", "/indikator_kinerja/*", "(GET)")
	enforce.LoadPolicy()
	return enforce
}
