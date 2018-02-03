package main
import (
    "github.com/gin-gonic/gin"
    "../app"
)

func main() {
    // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //     fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })
	
	// fmt.Println("helo kitty");
    // http.ListenAndServe(":8080", nil)
    router := gin.Default()
    v1 := router.Group("/api/v1")
    {
    // v1.POST("/", createTodo)
    // v1.GET("/", fetchAllTodo)
    // v1.GET("/:id", fetchSingleTodo)
    // v1.PUT("/:id", updateTodo)
    // v1.DELETE("/:id", deleteTodo)
    v1.GET("/instructions", app.GetInstructions)
    // v1.GET("/instructions/:id", app.GetInstruction)
    // v1.POST("/instructions", app.PostInstruction)
    v1.PUT("/instructions/:id", app.UpdateInstruction)
    v1.DELETE("/instructions/:id", app.DeleteInstruction)
    }
    /* function to call for 404 no route found */
    router.NoRoute(func(c *gin.Context) {
        c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
    })
    router.Run()
}
func fetchAllTodo(c *gin.Context) {
    c.JSON(200, gin.H{"status": 200, "message": "Todo item created successfully!", "resourceId":1})
}
   