package httpApi

import (
	"github.com/phhphc/nft-marketplace-back-end/internal/controllers"
	mdw "github.com/phhphc/nft-marketplace-back-end/internal/middlewares"
)

func (s *httpServer) applyRoutes(
	controller controllers.Controller,
) {
	apiV1 := s.Echo.Group("/api/v0.1")
	nftRoute := apiV1.Group("/nft")
	nftRoute.GET("", controller.GetNFTsWithListings)
	nftRoute.PATCH("/:token/:identifier", controller.UpdateNftStatus, mdw.IsLoggedIn)

	orderRoute := apiV1.Group("/order")
	orderRoute.GET("", controller.GetOrder)
	orderRoute.POST("", controller.PostOrder, mdw.IsLoggedIn)

	collectionRoute := apiV1.Group("/collection")
	collectionRoute.POST("", controller.PostCollection, mdw.IsLoggedIn)
	collectionRoute.GET("", controller.GetCollection)
	collectionRoute.GET("/:category", controller.GetCollectionWithCategory)

	profileRoute := apiV1.Group("/profile")
	profileRoute.GET("/:address", controller.GetProfile)
	profileRoute.POST("", controller.PostProfile, mdw.IsLoggedIn)
	profileRoute.GET("/offer", controller.GetOffer)

	eventRoute := apiV1.Group("/event")
	eventRoute.GET("", controller.GetEvent)

	notificationRoute := apiV1.Group("/notification")
	notificationRoute.GET("", controller.GetNotification, mdw.IsLoggedIn)
	notificationRoute.POST("", controller.UpdateNotification, mdw.IsLoggedIn)

	authenticationRoute := apiV1.Group("/auth")
	authenticationRoute.GET("/:address/nonce", controller.GetUserNonce)
	authenticationRoute.POST("/login", controller.Login)
	authenticationRoute.POST("/test", controller.Test, mdw.IsLoggedIn)

	userRoute := apiV1.Group("/user", mdw.IsLoggedIn)
	userRoute.GET("", controller.GetUsers, mdw.Or(mdw.IsModerator, mdw.IsAdmin))
	userRoute.GET("/:address", controller.GetUser)
	userRoute.PATCH("/:address/block", controller.UpdateBlockState)
	userRoute.POST("/role", controller.CreateUserRole, mdw.IsAdmin)
	userRoute.DELETE("/role", controller.DeleteUserRole, mdw.IsAdmin)

	roleRoute := apiV1.Group("/role", mdw.IsLoggedIn, mdw.Or(mdw.IsModerator, mdw.IsAdmin))
	roleRoute.GET("", controller.GetRoles)

	settingsRoute := apiV1.Group("/settings")
	settingsRoute.GET("", controller.GetMarketplaceSettings, mdw.IsLoggedIn)
	settingsRoute.POST("", controller.CreateMarketplaceSettings, mdw.IsLoggedIn, mdw.IsAdmin)

	searchRoute := apiV1.Group("/search")
	searchRoute.GET("", controller.SearchNFTs)
}
