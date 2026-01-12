package routes

import (
	datasource "betapa-antik-service/internal/dataSource"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, db *gorm.DB, rdb *redis.Client, cldSvc *datasource.CloudinaryService) {

}
