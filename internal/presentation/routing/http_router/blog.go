package http_router

import (
	"github.com/gin-gonic/gin"
	"site_builder_backend/internal/adapters/http/blog_controller"
	"site_builder_backend/internal/application/use_cases/blog_use_case"
	"site_builder_backend/internal/infrastructures/impl/db/mysql/blog_repo"
	"site_builder_backend/internal/presentation/routing"
)

type BlogRouter struct {
	article  *gin.RouterGroup
	category *gin.RouterGroup
	*routing.Services
}

func NewBlogRouter(services *routing.Services) *BlogRouter {
	return &BlogRouter{
		article:  services.Router.Group("Blog"),
		category: services.Router.Group("Category"),
		Services: services,
	}
}

func (r *BlogRouter) ArticleRegister() {

	articleReadRepo := blog_repo.NewArticleReadRepository(r.PostgresClient.DB, r.Logger)
	articleWriteRepo := blog_repo.NewArticleWriteRepository(r.PostgresClient.DB, r.Logger)
	useCase := blog_use_case.NewArticleCommandUseCase(articleReadRepo, articleWriteRepo, r.Logger)
	controller := blog_controller.NewArticleController(useCase, r.Logger)

	r.article.POST("Create", controller.CreateArticle)
}

func (r *BlogRouter) CategoryRegister() {

	categoryReadRepo := blog_repo.NewCategoryReadRepository(r.PostgresClient.DB, r.Logger)
	categoryWriteRepo := blog_repo.NewCategoryWriteRepository(r.PostgresClient.DB, r.Logger)
	useCase := blog_use_case.NewCategoryUseCase(categoryReadRepo, categoryWriteRepo, r.Logger)
	controller := blog_controller.NewCategoryController(useCase, r.Logger)

	r.category.POST("Create", controller.CreateCategory)
}
