package api

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"

	"api-gateway/api/handler"
	_ "api-gateway/docs"

	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title api gat way
// @version 1.0
// @description Auth service API documentation
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func NewGin(h *handler.Handler) *gin.Engine {
	ca, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
	if err != nil {
		panic(err)
	}
	/// WEFawef
	err = ca.LoadPolicy()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	// r.Use(middleware.NewAuth(ca))

	router := r.Group("/")
	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, url))

	a := router.Group("/account")
	{
		a.POST("/create", h.CreateAccount)
		a.GET("/list", h.ListAccounts)
		a.GET("/get/:id", h.GetAccountById)
		a.PUT("/update", h.UpdateAccount)
		a.DELETE("/delete/:id", h.DeleteAccount)
	}
	b := router.Group("/budget")
	{
		b.POST("/create", h.CreateBudget)
		b.GET("/list", h.ListBudgets)
		b.GET("/get/:id", h.GetBudgetById)
		b.PUT("/update", h.UpdateBudget)
		b.DELETE("/delete/:id", h.DeleteBudget)
	}
	c := router.Group("/category")
	{
		c.POST("/create", h.CreateCategory)
		c.GET("/list", h.ListCategories)
		c.GET("/get/:id", h.GetCategoryById)
		c.PUT("/update", h.UpdateCategory)
		c.DELETE("/delete/:id", h.DeleteCategory)
	}
	g := router.Group("/goal")
	{
		g.POST("/create", h.CreateGoal)
		g.GET("/list", h.ListGoals)
		g.GET("/get/:id", h.GetGoalById)
		g.PUT("/update", h.UpdateGoal)
		g.DELETE("/delete/:id", h.DeleteGoal)
	}

	t := router.Group("/transaction")
	{
		t.POST("/create", h.CreateTransaction)
		t.GET("/list", h.GetTransactions)
		t.GET("/get/:id", h.GetTransactionById)
		t.PUT("/update", h.UpdateTransaction)
		t.DELETE("/delete/:id", h.DeleteTransaction)
	}

	n := router.Group("/notification")
	{
		n.GET("/list", h.ListNotification)
		n.GET("/get/:user_id", h.GetNotification)
		// n.DELETE("/delete/:id", h.DeleteNotification)
	}

	return r
}
