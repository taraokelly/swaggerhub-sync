/*
 * Swagger Petstore
 *
 * Pushing this to GitHub 3
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/Team_Org/API-EnabledSCMIntegrations/1.0.0/",
		Index,
	},

	Route{
		"AddPet",
		strings.ToUpper("Post"),
		"/Team_Org/API-EnabledSCMIntegrations/1.0.0/pet",
		AddPet,
	},

	Route{
		"DeletePet",
		strings.ToUpper("Delete"),
		"/Team_Org/API-EnabledSCMIntegrations/1.0.0/pet/{petId}",
		DeletePet,
	},

	Route{
		"FindPetsByStatus",
		strings.ToUpper("Get"),
		"/Team_Org/API-EnabledSCMIntegrations/1.0.0/pet/findByStatus",
		FindPetsByStatus,
	},

	Route{
		"FindPetsByTags",
		strings.ToUpper("Get"),
		"/Team_Org/API-EnabledSCMIntegrations/1.0.0/pet/findByTags",
		FindPetsByTags,
	},

	Route{
		"GetPetById",
		strings.ToUpper("Get"),
		"/Team_Org/API-EnabledSCMIntegrations/1.0.0/pet/{petId}",
		GetPetById,
	},

	Route{
		"UpdatePet",
		strings.ToUpper("Put"),
		"/Team_Org/API-EnabledSCMIntegrations/1.0.0/pet",
		UpdatePet,
	},

	Route{
		"UpdatePetWithForm",
		strings.ToUpper("Post"),
		"/Team_Org/API-EnabledSCMIntegrations/1.0.0/pet/{petId}",
		UpdatePetWithForm,
	},

	Route{
		"UploadFile",
		strings.ToUpper("Post"),
		"/Team_Org/API-EnabledSCMIntegrations/1.0.0/pet/{petId}/uploadImage",
		UploadFile,
	},

	Route{
		"DeleteOrder",
		strings.ToUpper("Delete"),
		"/Team_Org/API-EnabledSCMIntegrations/1.0.0/store/order/{orderId}",
		DeleteOrder,
	},

	Route{
		"GetInventory",
		strings.ToUpper("Get"),
		"/Team_Org/API-EnabledSCMIntegrations/1.0.0/store/inventory",
		GetInventory,
	},

	Route{
		"GetOrderById",
		strings.ToUpper("Get"),
		"/Team_Org/API-EnabledSCMIntegrations/1.0.0/store/order/{orderId}",
		GetOrderById,
	},

	Route{
		"PlaceOrder",
		strings.ToUpper("Post"),
		"/Team_Org/API-EnabledSCMIntegrations/1.0.0/store/order",
		PlaceOrder,
	},

	Route{
		"CreateUser",
		strings.ToUpper("Post"),
		"/Team_Org/API-EnabledSCMIntegrations/1.0.0/user",
		CreateUser,
	},

	Route{
		"CreateUsersWithArrayInput",
		strings.ToUpper("Post"),
		"/Team_Org/API-EnabledSCMIntegrations/1.0.0/user/createWithArray",
		CreateUsersWithArrayInput,
	},

	Route{
		"CreateUsersWithListInput",
		strings.ToUpper("Post"),
		"/Team_Org/API-EnabledSCMIntegrations/1.0.0/user/createWithList",
		CreateUsersWithListInput,
	},

	Route{
		"DeleteUser",
		strings.ToUpper("Delete"),
		"/Team_Org/API-EnabledSCMIntegrations/1.0.0/user/{username}",
		DeleteUser,
	},

	Route{
		"GetUserByName",
		strings.ToUpper("Get"),
		"/Team_Org/API-EnabledSCMIntegrations/1.0.0/user/{username}",
		GetUserByName,
	},

	Route{
		"LoginUser",
		strings.ToUpper("Get"),
		"/Team_Org/API-EnabledSCMIntegrations/1.0.0/user/login",
		LoginUser,
	},

	Route{
		"LogoutUser",
		strings.ToUpper("Get"),
		"/Team_Org/API-EnabledSCMIntegrations/1.0.0/user/logout",
		LogoutUser,
	},

	Route{
		"UpdateUser",
		strings.ToUpper("Put"),
		"/Team_Org/API-EnabledSCMIntegrations/1.0.0/user/{username}",
		UpdateUser,
	},
}
