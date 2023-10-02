package company

import (
	"myapp/pkg/db"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

//#list companies
//#list company roles

type company_dao struct {
	db *gorm.DB
}

func New(db *db.Client) *company_dao {
	return &company_dao{
		db: db.DB,
	}
}

func (c company_dao) ListCompanies(r echo.Context) error {
	var companies []Company
	if err := c.db.Where("is_deleted=?", 0).Find(&companies).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch companies")
	}

	return r.JSON(http.StatusOK, companies)
}

func (c company_dao) ListCompanyRoles(r echo.Context) error {
	companyID := r.QueryParam("company_id")
	var roles []Role
	if err := c.db.Where("company_id=?", companyID).Find(&roles).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "company not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch company")
	}
	return r.JSON(http.StatusOK, roles)
}
