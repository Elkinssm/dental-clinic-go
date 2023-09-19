package middlewares

import "github.com/gin-gonic/gin"

type AuthMiddleware struct {
	publicKey, privateKey string
}

func NewAuthMiddleware(publicKey, privateKey string) *AuthMiddleware {
	return &AuthMiddleware{
		publicKey:  publicKey,
		privateKey: privateKey,
	}
}

func (a *AuthMiddleware) AuthHeader(ctx *gin.Context) {
	if a.publicKey != ctx.GetHeader("PUBLIC-KEY") || a.privateKey != ctx.GetHeader("PRIVATE-KEY") {
		ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
		return
	}
}