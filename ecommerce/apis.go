package golomt

import (
	"net/http"

	"github.com/techpartners-asia/golomt-ecommerce-go/utils"
)

var (
	ECommerceInvoiceCreate = utils.API{
		Url:    "/api/invoice",
		Method: http.MethodPost,
	}

	ECommerceInquiry = utils.API{
		Url:    "/api/inquiry",
		Method: http.MethodPost,
	}

	ECommercePayByToken = utils.API{
		Url:    "/api/pay",
		Method: http.MethodPost,
	}
)
