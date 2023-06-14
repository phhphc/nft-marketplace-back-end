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
	nftRoute.PATCH("/:token/:identifier", controller.UpdateNftStatus)

	orderRoute := apiV1.Group("/order")
	orderRoute.GET("", controller.GetOrder)
	orderRoute.POST("", controller.PostOrder)

	collectionRoute := apiV1.Group("/collection")
	collectionRoute.POST("", controller.PostCollection)
	collectionRoute.GET("", controller.GetCollection)
	collectionRoute.GET("/:category", controller.GetCollectionWithCategory)

	profileRoute := apiV1.Group("/profile")
	profileRoute.GET("/:address", controller.GetProfile)
	profileRoute.POST("", controller.PostProfile)
	profileRoute.GET("/offer", controller.GetOffer)

	eventRoute := apiV1.Group("/event")
	eventRoute.GET("", controller.GetEvent)

	notificationRoute := apiV1.Group("/notification")
	notificationRoute.GET("", controller.GetNotification)
	notificationRoute.POST("", controller.UpdateNotification)

	authenticationRoute := apiV1.Group("/auth")
	authenticationRoute.GET("/:address/nonce", controller.GetUserNonce)
	authenticationRoute.POST("/login", controller.Login)
	authenticationRoute.POST("/test", controller.Test, middlewares.IsLoggedIn)

	userRoute := apiV1.Group("/user")
	userRoute.GET("", controller.GetUsers)
	userRoute.GET("/:address", controller.GetUser)
	userRoute.PATCH("/:address", controller.UpdateBlockState)
	userRoute.POST("/role", controller.CreateUserRole)
	userRoute.DELETE("/role", controller.DeleteUserRole)

	searchRoute := apiV1.Group("/search")
	searchRoute.GET("", controller.SearchNFTs)
}
