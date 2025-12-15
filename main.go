// main.go
package main

import (
	"html/template"
	"log"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"tronroll21-dev/yudoksystem/controllers" // Assuming the controllers package is in the same directory
	"tronroll21-dev/yudoksystem/models"      // Assuming the models package is in the same directory

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func formatNumber(n int) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%d", n)
}

// Main function to set up the Gin router and start the server
func main() {

	err := godotenv.Load()
	if err != nil {
		// Log a fatal error if the .env file cannot be loaded, as the app
		// cannot run without its configuration.
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Create a new Gin router with the default middleware
	router := gin.Default()
	router.Static("/assets", "./assets")

	// 1. Define the FuncMap
	funcMap := template.FuncMap{
		"format": formatNumber, // The function from step 1
	}

	// 2. Set the FuncMap directly on the router ENGINE
	// This tells LoadHTMLGlob what functions to include.
	router.SetFuncMap(funcMap) // ✅ CORRECT METHOD

	// Load HTML templates from the 'templates' directory
	// Gin uses the standard Go template package, so we can use partials
	router.LoadHTMLGlob("templates/*")

	// Initialize the mock database in the models package
	// In a real application, you would establish a real connection here
	err = models.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Define the route for the main page
	router.GET("/", func(c *gin.Context) {
		// Serve the static index.html from the assets folder
		c.File("./assets/index.html")
	})

	// Define the route for the main page
	router.GET("/ranges", func(c *gin.Context) {
		// Serve the static index.html from the assets folder
		c.File("./assets/ranges.html")
	})

	// Use the controller handler for /menubetsuuriage
	router.GET("/api/menubetsuuriage", controllers.MenubetsuUriageHandler)

	// Define the HTMX endpoint to fetch sales data based on the date
	router.GET("/api/sales-data", controllers.SalesDataHandler)

	// Define the POST endpoint to insert or update sales data
	router.POST("/api/sales-data", controllers.PostSalesDataHandler)

	// 新しいGETリクエストハンドラ：日次報告レポートをPDFで生成
	router.GET("/api/sales-data-report", controllers.SalesDataReportHandler)

	router.GET("/api/uriage-nikkei-report", controllers.UriageNikkeiReportHandler)

	router.GET("/api/nyuuyokushasuushuukei", controllers.NyuuyokushasuuShuukeiHandler)

	router.GET("/api/jissekiyosoku", controllers.JissekiyosokuHandler)

	router.GET("/api/tantoushas", controllers.TantoushasHandler)

	// APIルートグループ
	api := router.Group("/api")
	{
		api.POST("/ranges", controllers.CreateRange)       // POST /api/ranges -> 作成
		api.GET("/ranges", controllers.GetRanges)          // GET /api/ranges -> 全て取得
		api.GET("/ranges/:id", controllers.GetRangeByID)   // GET /api/ranges/:id -> 個別取得
		api.PUT("/ranges/:id", controllers.UpdateRange)    // PUT /api/ranges/:id -> 更新
		api.DELETE("/ranges/:id", controllers.DeleteRange) // DELETE /api/ranges/:id -> 削除
	}

	log.Println("Starting server on :8080")
	router.Run(":8080")
}
