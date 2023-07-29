package utils

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/Noush-012/project-ecommerce-microservice/api-gateway/pkg/utils/request"
	"github.com/gin-gonic/gin"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func StringToUint(str string) (uint, error) {
	val, err := strconv.Atoi(str)
	return uint(val), err
}

func GetUserIdFromContext(ctx *gin.Context) uint {
	userIdStr := ctx.GetString("userId")
	userIdInt, _ := strconv.Atoi(userIdStr)
	return uint(userIdInt)
}

func GenerateSKU(prod request.ProductItemReq) (string, error) {
	// var sku string

	// if prod.Brand != "" {
	// 	sku = prod.Brand + "-"
	// 	if prod.Category != "" {
	// 		sku += prod.Category + "-"
	// 		if prod.SubCategory != "" {
	// 			sku += prod.SubCategory + "-"
	// 			if prod.Name != "" {
	// 				sku += prod.Name
	// 				return sku, nil
	// 			}
	return "", nil
}

func GenerateRandomString(length int) string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rng.Intn(len(charset))]
	}
	return string(b)
}
