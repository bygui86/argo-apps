package main_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	
	"github.com/bygui86/go-postgres/database"
	"github.com/bygui86/go-postgres/logging"
	"github.com/bygui86/go-postgres/rest"
)

var server *rest.Server

// *** TESTS ***

func TestMain(m *testing.M) {
	// setup
	envVarsErr := setEnvVars()
	if envVarsErr != nil {
		logging.SugaredLog.Errorf("Setup environment variables failed: %s", envVarsErr.Error())
		os.Exit(501)
	}
	
	var err error
	server, err = rest.NewServer()
	if err != nil {
		logging.SugaredLog.Errorf("REST server creation failed: %s", err.Error())
		os.Exit(501)
	}
	ensureTableExists()
	
	// execute
	code := m.Run()
	
	// teardown
	clearTable()
	os.Exit(code) // does not respect defer statements
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	request, reqErr := http.NewRequest("GET", "/products", nil)
	if reqErr != nil {
		t.Errorf("Create request failed: %s", reqErr.Error())
	}
	response := executeRequest(request)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetNonExistentProduct(t *testing.T) {
	clearTable()

	request, reqErr := http.NewRequest("GET", "/products/11", nil)
	if reqErr != nil {
		t.Errorf("Create request failed: %s", reqErr.Error())
	}
	response := executeRequest(request)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Product not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Product not found'. Got '%s'", m["error"])
	}
}

func TestGetProduct(t *testing.T) {
	clearTable()
	addProducts(1)

	request, reqErr := http.NewRequest("GET", "/products/1", nil)
	if reqErr != nil {
		t.Errorf("Create request failed: %s", reqErr.Error())
	}
	response := executeRequest(request)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestCreateProduct(t *testing.T) {
	clearTable()

	productName := "test product"
	productPrice := 11.22
	product := &database.Product{
		Name:  productName,
		Price: productPrice,
	}
	payload, marshErr := json.Marshal(product)
	if marshErr != nil {
		t.Errorf("Product marshal error: %s", marshErr.Error())
	}
	
	request, reqErr := http.NewRequest("POST", "/products", bytes.NewBuffer(payload))
	if reqErr != nil {
		t.Errorf("Create request failed: %s", reqErr.Error())
	}
	response := executeRequest(request)

	checkResponseCode(t, http.StatusCreated, response.Code)
	
	var responseProduct database.Product
	unmarshErr := json.Unmarshal(response.Body.Bytes(), &responseProduct)
	if unmarshErr != nil {
		t.Errorf("Product unmarshal error: %s", unmarshErr.Error())
	}
	
	expectedProductId := 1
	if responseProduct.ID != expectedProductId {
		t.Errorf("Expected product ID to be '%d'. Got '%d'", expectedProductId, responseProduct.ID)
	}
	if responseProduct.Name != productName {
		t.Errorf("Expected product name to be '%s'. Got '%s'", productName, responseProduct.Name)
	}
	if responseProduct.Price != productPrice {
		t.Errorf("Expected product price to be '%f'. Got '%f'", productPrice, responseProduct.Price)
	}
}

func TestUpdateProduct(t *testing.T) {
	clearTable()
	addProducts(1)

	request, reqErr := http.NewRequest("GET", "/products/1", nil)
	if reqErr != nil {
		t.Errorf("Create request failed: %s", reqErr.Error())
	}
	response := executeRequest(request)
	var originalProduct database.Product
	unmarshErr := json.Unmarshal(response.Body.Bytes(), &originalProduct)
	if unmarshErr != nil {
		t.Errorf("Product unmarshal error: %s", unmarshErr.Error())
	}
	
	updatedProductName := "updated product"
	updatedProductPrice := 9.42
	updatedProduct := &database.Product{
		ID:    originalProduct.ID,
		Name:  updatedProductName,
		Price: updatedProductPrice,
	}
	payload, marshErr := json.Marshal(updatedProduct)
	if marshErr != nil {
		t.Errorf("Product marshal error: %s", marshErr.Error())
	}
	
	request, _ = http.NewRequest("PUT", "/products/1", bytes.NewBuffer(payload))
	response = executeRequest(request)

	checkResponseCode(t, http.StatusOK, response.Code)
	
	var responseProduct database.Product
	unmarshErr = json.Unmarshal(response.Body.Bytes(), &responseProduct)
	if unmarshErr != nil {
		t.Errorf("Product unmarshal error: %s", unmarshErr.Error())
	}
	
	if responseProduct.ID != originalProduct.ID {
		t.Errorf("Expected product ID to remain the same '%d'. Got '%d'", originalProduct.ID, responseProduct.ID)
	}
	if responseProduct.Name != updatedProductName {
		t.Errorf("Expected product name to change from '%s' to '%s'. Got '%s'", originalProduct.Name, updatedProductName, responseProduct.Name)
	}
	if responseProduct.Price != updatedProductPrice {
		t.Errorf("Expected product price to change from '%f' to '%f'. Got '%f'", originalProduct.Price, updatedProductPrice, responseProduct.Price)
	}
}

func TestDeleteProduct(t *testing.T) {
	clearTable()
	addProducts(1)

	request, reqErr := http.NewRequest("GET", "/products/1", nil)
	if reqErr != nil {
		t.Errorf("Create request failed: %s", reqErr.Error())
	}
	response := executeRequest(request)
	checkResponseCode(t, http.StatusOK, response.Code)

	request, reqErr = http.NewRequest("DELETE", "/products/1", nil)
	if reqErr != nil {
		t.Errorf("Create request failed: %s", reqErr.Error())
	}
	response = executeRequest(request)
	checkResponseCode(t, http.StatusOK, response.Code)

	request, reqErr = http.NewRequest("GET", "/products/1", nil)
	if reqErr != nil {
		t.Errorf("Create request failed: %s", reqErr.Error())
	}
	response = executeRequest(request)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

// *** UTILS ***

func setEnvVars() error {
	userErr := os.Setenv("DB_USERNAME", "postgres")
	if userErr != nil {
		return userErr
	}
	pwErr:=os.Setenv("DB_PASSWORD", "supersecret")
	if pwErr != nil {
		return pwErr
	}
	nameErr:=os.Setenv("DB_NAME", "postgres")
	if nameErr != nil {
		return nameErr
	}
	return nil
}

func ensureTableExists() {
	_, err := server.DbConnection.Exec(database.CreateTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	server.DbConnection.Exec("DELETE FROM products")
	server.DbConnection.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
	server.Router.ServeHTTP(responseRecorder, req)
	return responseRecorder
}

func addProducts(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		server.DbConnection.Exec("INSERT INTO products(name, price) VALUES($1, $2)", "Product "+strconv.Itoa(i), (i+1.0)*10)
	}
}
