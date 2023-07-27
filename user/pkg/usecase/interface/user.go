package interfaces

import (
	"context"

	"github.com/Noush-012/project-ecommerce-microservice/user-service/pkg/models"
	"github.com/Noush-012/project-ecommerce-microservice/user-service/pkg/utils/request"
	"github.com/Noush-012/project-ecommerce-microservice/user-service/pkg/utils/response"
)

type UserService interface {
	Profile(ctx context.Context, userId uint) (profile response.Profile, err error)
	Addaddress(ctx context.Context, address request.Address) error
	UpdateAddress(ctx context.Context, address request.AddressPatchReq) error
	DeleteAddress(ctx context.Context, userID, addressID uint) error
	GetAllAddress(ctx context.Context, userId uint) (address []response.Address, err error)

	SaveCartItem(ctx context.Context, addToCart request.AddToCartReq) error

	GetCartItemsbyCartId(ctx context.Context, page request.ReqPagination, userID uint) (CartItems []response.CartItemResp, err error)
	UpdateCart(ctx context.Context, cartUpadates request.UpdateCartReq) error
	RemoveCartItem(ctx context.Context, DelCartItem request.DeleteCartItemReq) error

	AddToWishlist(ctx context.Context, wishlistData request.AddToWishlist) error
	GetWishlist(ctx context.Context, userId uint) (wishlist []response.Wishlist, err error)
	DeleteFromWishlist(ctx context.Context, productId, userId uint) error

	GetWalletHistory(ctx context.Context, userId uint) (wallet []models.Wallet, err error)
}
