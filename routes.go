// routes.go
package main

import (
	"tronroll21-dev/yudoksystem/controllers"
	"tronroll21-dev/yudoksystem/middlewares"

	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine) {
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")

	// Auth
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)
	router.GET("/login", func(c *gin.Context) {
		c.File("./assets/login.html")
	})

	// Protected routes
	authorized := router.Group("/")
	authorized.Use(middlewares.RequireAuth)
	{
		authorized.GET("/", func(c *gin.Context) { c.File("./assets/index.html") })
		authorized.GET("/api/user/me", controllers.GetCurrentUser)
		authorized.GET("/power-readings", func(c *gin.Context) { c.File("./assets/power-readings/index.html") })
		authorized.GET("/api/power-readings", controllers.PowerReadingsHandler)
		authorized.POST("/api/power-readings", controllers.SavePowerReadingHandler)
		authorized.GET("/gas-readings", func(c *gin.Context) { c.File("./assets/gas-readings/index.html") })
		authorized.GET("/api/gas-readings", controllers.GasReadingsHandler)
		authorized.POST("/api/gas-readings", controllers.SaveGasReadingHandler)
	}

	// Shiire
	shiire := router.Group("/api/shiire")
	{
		shiire.GET("", controllers.ShiiremeisaiHandler)
		shiire.POST("", controllers.SaveShiiremeisaiHandler)
		shiire.DELETE("", controllers.DeleteShiiremeisaiHandler)
		shiire.POST("/shouhizei", controllers.SaveShiireshouhizeiHandler)
		shiire.POST("/init", controllers.InitializeShiiremeisaiHandler)
		shiire.GET("/available-contractors", controllers.GetAvailableContractorsHandler)
		shiire.GET("/fixed_contractors", controllers.GetFixedContractorsHandler)
		shiire.POST("/add-contractor", controllers.AddContractorHandler)
	}
	router.GET("/shiire", func(c *gin.Context) { c.File("./assets/shiire/index.html") })

	// Sales / reports
	salesAPI := router.Group("/api")
	{
		salesAPI.GET("/sales-data", controllers.SalesDataHandler)
		salesAPI.POST("/sales-data", controllers.PostSalesDataHandler)
		salesAPI.GET("/sales-data-report", controllers.SalesDataReportHandler)
		salesAPI.GET("/sales-data-report-pdf", controllers.SalesDataReportHandlerPDF)
		salesAPI.GET("/sales-data-report-svg", controllers.SalesDataReportSVGHandler)
		salesAPI.GET("/uriage-nikkei-report", controllers.UriageNikkeiReportHandler)
		salesAPI.GET("/menubetsuuriage", controllers.MenubetsuUriageHandler)
		salesAPI.POST("/menubetsu-uriage", controllers.SaveMenubetsuUriage)
		salesAPI.GET("/nyuuyokushasuushuukei", controllers.NyuuyokushasuuShuukeiHandler)
		salesAPI.GET("/jissekiyosoku", controllers.JissekiyosokuHandler)
		salesAPI.GET("/analysis-data", controllers.AnalysisDataHandler)
		salesAPI.GET("/tantoushas", controllers.TantoushasHandler)
		salesAPI.GET("/rearegi-details", controllers.GetRearegiDetailHandler)
		salesAPI.GET("/rearegi-report", controllers.GetRearegiReportHandler)
		salesAPI.POST("/rearegi-details", controllers.SaveRearegiDetailHandler)
	}

	// Ranges
	router.GET("/ranges", func(c *gin.Context) { c.File("./assets/ranges.html") })
	ranges := router.Group("/api/ranges")
	{
		ranges.POST("", controllers.CreateRange)
		ranges.GET("", controllers.GetRanges)
		ranges.GET("/:id", controllers.GetRangeByID)
		ranges.PUT("/:id", controllers.UpdateRange)
		ranges.DELETE("/:id", controllers.DeleteRange)
	}
}
