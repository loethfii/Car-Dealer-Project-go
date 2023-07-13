package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tugas_akhir_course_net/auth"
)

//3 user biasa (auth)
//2 staff seles
//3 manager sales

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Request Access Token"})
			c.Abort()
			return
		}
		//validte token

		_, _, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "unauthorized",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func User() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Request Access Token"})
			c.Abort()
			return
		}

		_, role, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "UnAuthorized",
				"error":   "tidak dapat akses",
			})
			c.Abort()
			return
		}

		//role 3 User
		if role != 3 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Role anda tidak cocok",
				"status":  http.StatusUnauthorized,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func ManagerSales() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Request Access Token"})
			c.Abort()
			return
		}

		_, role, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "UnAuthorized",
				"error":   "tidak dapat akses",
			})
			c.Abort()
			return
		}

		//role 1 Manager Sales
		if role != 1 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Role anda tidak cocok",
				"status":  http.StatusUnauthorized,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func ManagerAndStaffSales() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Request Access Token"})
			c.Abort()
			return
		}

		_, role, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "UnAuthorized",
				"error":   "tidak dapat akses",
			})
			c.Abort()
			return
		}

		//role 1 Manager Sales
		if role != 1 && role != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Role anda tidak cocok",
				"status":  http.StatusUnauthorized,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

//func StaffSales() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		tokenString := c.GetHeader("Authorization")
//
//		if tokenString == "" {
//			c.JSON(http.StatusUnauthorized, gin.H{"error": "Request Access Token"})
//			c.Abort()
//			return
//		}
//
//		email, role, err := auth.ValidateToken(tokenString)
//		if err != nil {
//			c.JSON(http.StatusUnauthorized, gin.H{
//				"Message": "UnAuthorized",
//				"error":   "tidak dapat akses",
//			})
//			c.Abort()
//			return
//		}
//
//		fmt.Println(role)
//		fmt.Println(email)
//
//		//role 2 Staff sales
//		if role != 2 {
//			c.JSON(http.StatusUnauthorized, gin.H{
//				"message": "Role anda tidak cocok",
//				"status":  http.StatusUnauthorized,
//			})
//			c.Abort()
//			return
//		}
//		c.Next()
//	}
//}
//
