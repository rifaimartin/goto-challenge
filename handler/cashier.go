package handler

import (
	"net/http"
	"pos-app/auth"
	"pos-app/cashier"
	"pos-app/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type cashierHandler struct {
	cashierService cashier.Service
	authService auth.Service
}

// NewcashierHandler e
func NewcashierHandler(cashierService cashier.Service,authService auth.Service) *cashierHandler {
	return &cashierHandler{cashierService, authService}
}


func (h *cashierHandler) Login(c *gin.Context) {
	cashierId := c.Param("cashierId")
	newCashierId, _ := strconv.Atoi(cashierId);
	var input cashier.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, false, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedincashier, err := h.cashierService.Login(input, newCashierId)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse(err.Error(), http.StatusUnauthorized, false, errorMessage)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedincashier.ID)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, false, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	settoken := make(map[string]string)
	settoken["token"] = token


	response := helper.APIResponse("Success", http.StatusOK, true, settoken)

	c.JSON(http.StatusOK, response)

}

func (h *cashierHandler) GetAllCashiers(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))

	cashiers,count, err := h.cashierService.GetAllCashiers(limit,skip)

	if err != nil {
		response := helper.APIResponse("Error to get Cashier", http.StatusBadRequest, false, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Success", http.StatusOK, true, cashier.ListFormatCashier(cashiers, limit, skip, count))
	c.JSON(http.StatusOK, response)
}


//api/v1/campaigns
func (h *cashierHandler) GetcashierByID(c *gin.Context) {
	cashierId := c.Param("cashierId")
	newCashierId, _ := strconv.Atoi(cashierId);

	cashiers, err := h.cashierService.GetcashierByID(newCashierId)
	if err != nil {
		response := helper.APIResponseFail("Cashier Not Found", http.StatusNotFound, false,  nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse("Success", http.StatusOK, true, cashier.Formatcashier(cashiers) )
	c.JSON(http.StatusOK, response)
}

//api/v1/campaigns
func (h *cashierHandler) GetPassCode(c *gin.Context) {
	cashierId := c.Param("cashierId")
	newCashierId, _ := strconv.Atoi(cashierId);

	_, err := h.cashierService.GetcashierByID(newCashierId)
	if err != nil {
		response := helper.APIResponseFail("Cashier Not Found", http.StatusNotFound, false,  nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	passcode := make(map[string]int)
	passcode["passcode"] = 123456

	response := helper.APIResponse("Success", http.StatusOK, true,  passcode)
	c.JSON(http.StatusOK, response)
}



func (h *cashierHandler) Logout(c *gin.Context) {
	cashierId := c.Param("cashierId")
	newCashierId, _ := strconv.Atoi(cashierId);
	var input cashier.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, false, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	
	_, fail := h.cashierService.Logout(input, newCashierId)
	if err != fail {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse(err.Error(), http.StatusUnauthorized, false, errorMessage)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	bodyResp := make(map[string]interface{})
	bodyResp["success"] = true
	bodyResp["message"] = "success"

	c.JSON(http.StatusOK, bodyResp)

}

func (h *cashierHandler) CreateCashier(c *gin.Context) {
	var input cashier.CasherInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse(err.Error(), http.StatusUnprocessableEntity, false, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	
	newCashier, err := h.cashierService.CreateCashier(input)
	if err != nil {
		response := helper.APIResponse("Failed to create cashier", http.StatusBadRequest, false, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success", http.StatusOK, true,  cashier.FormatCreateCashier(newCashier) )
	c.JSON(http.StatusOK, response)
}

func (h *cashierHandler) UpdateCashier(c *gin.Context) {
	var input cashier.CasherInput
	cashierId := c.Param("cashierId")
	newCashierId, _ := strconv.Atoi(cashierId);

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse(err.Error(), http.StatusUnprocessableEntity, false, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	
	_, fail := h.cashierService.UpdateCashier(newCashierId, input)
	if fail != nil {
		response := helper.APIResponseFail("Cashier Not Found", http.StatusBadRequest, false, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	bodyResp := make(map[string]interface{})
	bodyResp["success"] = true
	bodyResp["message"] = "success"

	response := helper.APIResponse("Success", http.StatusOK, true,  bodyResp )
	c.JSON(http.StatusOK, response)
}

func (h *cashierHandler) DeleteCashier(c *gin.Context) {
	cashierId := c.Param("cashierId")
	newCashierId, _ := strconv.Atoi(cashierId);

	_, err := h.cashierService.DeleteCashier(newCashierId)
	if err != nil {
		response := helper.APIResponseFail("Cashier Not Found", http.StatusNotFound, false,  nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	bodyResp := make(map[string]interface{})
	bodyResp["success"] = true
	bodyResp["message"] = "success"

	c.JSON(http.StatusOK, bodyResp)
}