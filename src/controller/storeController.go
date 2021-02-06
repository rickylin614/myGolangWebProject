package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"orderbento/src/dao/storeDao"
	"orderbento/src/models"
	"orderbento/src/service/storeService"
	"orderbento/src/utils"

	"github.com/gin-gonic/gin"
)

type storeReq struct {
	Id      uint
	PhoneNo string
	Name    string
	Region  string
}

func InsertStore(ctx *gin.Context) {
	var data storeReq
	err := ctx.Bind(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp := make(gin.H)
	var store storeDao.Store
	store.Name = data.Name
	store.Phone_no, _ = strconv.Atoi(data.PhoneNo)
	store.Region = data.Region
	storeService.Insert(store)
	resp["msg"] = "新增商家成功"
	ctx.JSON(http.StatusOK, resp)
}

func UpdateStore(ctx *gin.Context) {
	var data storeReq
	err := ctx.Bind(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp := make(gin.H)
	store := storeService.QueryStoreById(data.Id)

	store.Name = data.Name
	store.Phone_no, _ = strconv.Atoi(data.PhoneNo)
	store.Region = data.Region
	storeService.Update(store)
	resp["msg"] = "修改商家成功"
	ctx.JSON(http.StatusOK, resp)
}

func QueryStoreById(ctx *gin.Context) {
	var data storeReq
	err := ctx.Bind(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	store := storeService.QueryStoreById(data.Id)
	fmt.Print(store)
	storeReq := composeOneStoreResp(store)
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "查詢商家成功",
		"data": &storeReq,
	})
}

func QueryStore(ctx *gin.Context) {
	var req map[string]interface{}
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err)
		return
	}
	stores, count := storeService.QueryStore(req)
	storeReq := composeStoreResp(stores)
	ctx.JSON(http.StatusOK, gin.H{
		"msg":       "查詢商家成功",
		"data":      &storeReq,
		"dataCount": count,
	})
}

func DeleteStore(ctx *gin.Context) {
	var data storeReq
	err := ctx.Bind(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp := make(gin.H)
	var store storeDao.Store
	store.ID = data.Id
	storeService.Delete(store)
	resp["msg"] = "修改商家成功"
	ctx.JSON(http.StatusOK, resp)
}

func composeOneStoreResp(s storeDao.Store) models.StoreResponse {

	var storesp models.StoreResponse = models.StoreResponse{
		ID:          s.ID,
		Name:        s.Name,
		PhoneNo:     strconv.Itoa(s.Phone_no),
		Region:      s.Region,
		Create_user: s.Create_user,
		CreatedAt:   utils.TimeToString(&s.CreatedAt),
	}
	return storesp
}

func composeStoreResp(s []storeDao.Store) []models.StoreResponse {
	stores := make([]models.StoreResponse, 0, len(s))
	var storesp models.StoreResponse
	for _, store := range s {
		storesp = models.StoreResponse{
			ID:          store.ID,
			Name:        store.Name,
			PhoneNo:     strconv.Itoa(store.Phone_no),
			Region:      store.Region,
			Create_user: store.Create_user,
			CreatedAt:   utils.TimeToString(&store.CreatedAt),
		}
		stores = append(stores, storesp)
	}
	return stores
}
