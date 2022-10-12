package test

import (
	"context"
	"database/sql"
	"dhany007/belajar-golang-restful-api/app"
	"dhany007/belajar-golang-restful-api/controller"
	"dhany007/belajar-golang-restful-api/helper"
	"dhany007/belajar-golang-restful-api/middleware"
	"dhany007/belajar-golang-restful-api/model/domain"
	"dhany007/belajar-golang-restful-api/repository"
	"dhany007/belajar-golang-restful-api/service"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

// setup test untuk db
// memakai db_test baru
func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/belajar_golang_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

// fungsi untuk menghapus seluruh isi tabel category
func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE category")
}

func setupRouter(db *sql.DB) http.Handler {
	// dicopy yang dari main
	validate := validator.New()

	categoryRepository := repository.NewCategoryRespository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db) // menghapus isi tabel category terlebih dahulus

	requestBody := strings.NewReader(`{
		"name": "test"
	}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("api_key", "tokenrahasia")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responBody map[string]interface{}
	json.Unmarshal(body, &responBody)

	assert.Equal(t, 200, int(responBody["code"].(float64)))
	assert.Equal(t, "OK", responBody["status"])
	assert.Equal(t, "test", responBody["data"].(map[string]interface{})["name"])
}

func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db) // menghapus isi tabel category terlebih dahulus

	requestBody := strings.NewReader(`{
		"name": ""
	}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("api_key", "tokenrahasia")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responBody map[string]interface{}
	json.Unmarshal(body, &responBody)

	assert.Equal(t, 400, int(responBody["code"].(float64)))
	assert.Equal(t, "Bad Request", responBody["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db) // menghapus isi tabel category terlebih dahulus

	// add dulu data array nya

	tx, _ := db.Begin()

	categoryRepository := repository.NewCategoryRespository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	requestBody := strings.NewReader(`{
		"name": "Gadget New"
	}`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("api_key", "tokenrahasia")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responBody map[string]interface{}
	json.Unmarshal(body, &responBody)

	assert.Equal(t, 200, int(responBody["code"].(float64)))
	assert.Equal(t, "OK", responBody["status"])
	assert.Equal(t, category.Id, int(responBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Gadget New", responBody["data"].(map[string]interface{})["name"])
}

func TestUpdateCategoryFailed(t *testing.T) {

	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db) // menghapus isi tabel category terlebih dahulus

	// add dulu data array nya

	tx, _ := db.Begin()

	categoryRepository := repository.NewCategoryRespository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	requestBody := strings.NewReader(`{
		"name": ""
	}`)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("api_key", "tokenrahasia")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responBody map[string]interface{}
	json.Unmarshal(body, &responBody)

	assert.Equal(t, 400, int(responBody["code"].(float64)))
	assert.Equal(t, "Bad Request", responBody["status"])
}

func TestGetCategorySuccess(t *testing.T) {

	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db) // menghapus isi tabel category terlebih dahulus

	// add dulu data array nya

	tx, _ := db.Begin()

	categoryRepository := repository.NewCategoryRespository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("api_key", "tokenrahasia")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responBody map[string]interface{}
	json.Unmarshal(body, &responBody)

	assert.Equal(t, 200, int(responBody["code"].(float64)))
	assert.Equal(t, "OK", responBody["status"])
	assert.Equal(t, category.Id, int(responBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, category.Name, responBody["data"].(map[string]interface{})["name"])
}

func TestGetCategoryFailed(t *testing.T) {

	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db) // menghapus isi tabel category terlebih dahulus

	// data tidak usah diadd dulu, jadi otomatis not found
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/10", nil)
	request.Header.Add("api_key", "tokenrahasia")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responBody map[string]interface{}
	json.Unmarshal(body, &responBody)

	assert.Equal(t, 404, int(responBody["code"].(float64)))
	assert.Equal(t, "Not Found", responBody["status"])
}

func TestDeleteCategorySuccess(t *testing.T) {

	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db) // menghapus isi tabel category terlebih dahulus

	// add dulu data array nya

	tx, _ := db.Begin()

	categoryRepository := repository.NewCategoryRespository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("api_key", "tokenrahasia")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responBody map[string]interface{}
	json.Unmarshal(body, &responBody)

	assert.Equal(t, 200, int(responBody["code"].(float64)))
	assert.Equal(t, "OK", responBody["status"])
}

func TestDeleteCategoryFailed(t *testing.T) {

	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db) // menghapus isi tabel category terlebih dahulus

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/10", nil)
	request.Header.Add("api_key", "tokenrahasia")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responBody map[string]interface{}
	json.Unmarshal(body, &responBody)

	assert.Equal(t, 404, int(responBody["code"].(float64)))
	assert.Equal(t, "Not Found", responBody["status"])
}

func TestGetListCategoriesSuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db) // menghapus isi tabel category terlebih dahulus

	// add dulu data array nya

	tx, _ := db.Begin()

	categoryRepository := repository.NewCategoryRespository()
	category1 := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})

	category2 := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Programming",
	})
	tx.Commit()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("api_key", "tokenrahasia")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responBody map[string]interface{}
	json.Unmarshal(body, &responBody)

	assert.Equal(t, 200, int(responBody["code"].(float64)))
	assert.Equal(t, "OK", responBody["status"])

	// harus begini
	var categories = responBody["data"].([]interface{})

	categoryResponse1 := categories[0].(map[string]interface{})
	categoryResponse2 := categories[1].(map[string]interface{})

	assert.Equal(t, category1.Id, int(categoryResponse1["id"].(float64)))
	assert.Equal(t, category1.Name, categoryResponse1["name"])

	assert.Equal(t, category2.Id, int(categoryResponse2["id"].(float64)))
	assert.Equal(t, category2.Name, categoryResponse2["name"])
}

func TestUnauthorized(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db) // menghapus isi tabel category terlebih dahulus

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responBody map[string]interface{}
	json.Unmarshal(body, &responBody)

	assert.Equal(t, 401, int(responBody["code"].(float64)))
	assert.Equal(t, "Unauthorized", responBody["status"])
}
