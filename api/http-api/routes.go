package httpApi

import (
	"github.com/phhphc/nft-marketplace-back-end/internal/controllers"
	"github.com/phhphc/nft-marketplace-back-end/internal/middlewares"
)

func (s *httpServer) applyRoutes(
	controller controllers.Controller,
) {
	apiV1 := s.Echo.Group("/api/v0.1")
	nftRoute := apiV1.Group("/nft")
	nftRoute.GET("", controller.GetNFTsWithListings)
	nftRoute.PATCH("/:token/:identifier", controller.UpdateNftStatus, middlewares.IsLoggedIn)

	orderRoute := apiV1.Group("/order")
	orderRoute.GET("", controller.GetOrder)
	orderRoute.POST("", controller.PostOrder, middlewares.IsLoggedIn)

	collectionRoute := apiV1.Group("/collection")
	collectionRoute.POST("", controller.PostCollection, middlewares.IsLoggedIn)
	collectionRoute.GET("", controller.GetCollection)
	collectionRoute.GET("/:category", controller.GetCollectionWithCategory)

	profileRoute := apiV1.Group("/profile")
	profileRoute.GET("/:address", controller.GetProfile)
	profileRoute.POST("", controller.PostProfile, middlewares.IsLoggedIn)
	profileRoute.GET("/offer", controller.GetOffer)

	eventRoute := apiV1.Group("/event")
	eventRoute.GET("", controller.GetEvent)

	notificationRoute := apiV1.Group("/notification")
	notificationRoute.GET("", controller.GetNotification, middlewares.IsLoggedIn)
	notificationRoute.POST("", controller.UpdateNotification, middlewares.IsLoggedIn)

	authenticationRoute := apiV1.Group("/auth")
	authenticationRoute.GET("/:address/nonce", controller.GetUserNonce)
	authenticationRoute.POST("/login", controller.Login)
	authenticationRoute.POST("/test", controller.Test, middlewares.IsLoggedIn)

	userRoute := apiV1.Group("/user")
	userRoute.GET("", controller.GetUsers, middlewares.IsLoggedIn, middlewares.IsModerator)
	userRoute.GET("/:address", controller.GetUser, middlewares.IsLoggedIn, middlewares.IsModerator)
	userRoute.PATCH("/:address/block", controller.UpdateBlockState, middlewares.IsLoggedIn, middlewares.IsModerator)
	userRoute.POST("/role", controller.CreateUserRole, middlewares.IsLoggedIn, middlewares.IsModerator)
	userRoute.DELETE("/role", controller.DeleteUserRole, middlewares.IsLoggedIn, middlewares.IsModerator)

	settingsRoute := apiV1.Group("/settings")
	settingsRoute.GET("", controller.GetMarketplaceSettings, middlewares.IsLoggedIn)
	settingsRoute.POST("", controller.CreateMarketplaceSettings, middlewares.IsLoggedIn, middlewares.IsAdmin)

	searchRoute := apiV1.Group("/search")
	searchRoute.GET("", controller.SearchNFTs)
}
