package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"

	"github.com/bygui86/go-traces/http-client/commons"
	"github.com/bygui86/go-traces/http-client/logging"
)

// TODO reduce code duplication

func (s *Server) getProducts(writer http.ResponseWriter, request *http.Request) {
	span := opentracing.StartSpan("get-products-handler")
	defer span.Finish()

	startTimer := time.Now()

	logging.Log.Info("Get products")

	endpointUrl := &url.URL{Path: rootProductsEndpoint}
	path := s.baseURL.ResolveReference(endpointUrl)
	restRequest, reqErr := http.NewRequest(http.MethodGet, path.String(), nil)
	if reqErr != nil {
		errMsg := "Get products failed: " + reqErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)

		span.SetTag("products-found", 0)
		span.SetTag("error", errMsg)
		span.LogKV("products-found", 0, "error", errMsg)
		return
	}
	restRequest.Header.Set(headerAccept, headerApplicationJson)
	restRequest.Header.Set(headerUserAgent, headerUserAgentClient)

	// Transmit the span's TraceContext as HTTP headers on our outbound request.
	traceErr := opentracing.GlobalTracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(restRequest.Header))
	if traceErr != nil {
		errMsg := "Get products failed: " + traceErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)

		span.SetTag("products-found", 0)
		span.SetTag("error", errMsg)
		span.LogKV("products-found", 0, "error", errMsg)
		return
	}

	response, respErr := s.restClient.Do(restRequest)
	if respErr != nil {
		errMsg := "Get products failed: " + respErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)

		span.SetTag("products-found", 0)
		span.SetTag("error", errMsg)
		span.LogKV("products-found", 0, "error", errMsg)
		return
	}

	// TODO do not unmarshal response
	var products []*commons.Product
	unmarshErr := json.NewDecoder(response.Body).Decode(&products)
	if unmarshErr != nil {
		errMsg := "Get products failed: " + unmarshErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)

		span.SetTag("products-found", 0)
		span.SetTag("error", errMsg)
		span.LogKV("products-found", 0, "error", errMsg)
		return
	}
	defer response.Body.Close()

	span.SetTag("products-found", len(products))
	span.LogKV("products-found", len(products))

	sendJsonResponse(writer, http.StatusOK, products)

	IncreaseRestRequests("getProducts")
	ObserveRestRequestsTime("getProducts", float64(time.Now().Sub(startTimer).Milliseconds()))
}

func (s *Server) getProduct(writer http.ResponseWriter, request *http.Request) {
	span := opentracing.StartSpan("get-product-handler")
	defer span.Finish()

	startTimer := time.Now()

	vars := mux.Vars(request)
	id, idErr := strconv.Atoi(vars["id"])
	if idErr != nil {
		errMsg := "Get product failed: Invalid product ID"
		sendErrorResponse(writer, http.StatusBadRequest, errMsg)

		span.SetTag("error", errMsg)
		span.LogKV("error", errMsg)
		return
	}
	logging.SugaredLog.Infof("Get product by ID: %d", id)
	span.SetTag("product-id", id)

	endpointUrl := &url.URL{Path: fmt.Sprintf(productsIdServerEndpoint, id)}
	path := s.baseURL.ResolveReference(endpointUrl)
	restRequest, reqErr := http.NewRequest(http.MethodGet, path.String(), nil)
	if reqErr != nil {
		errMsg := "Get product failed: " + reqErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)

		span.SetTag("product-found", false)
		span.SetTag("error", errMsg)
		span.LogKV("product-id", id, "product-found", false, "error", errMsg)
		return
	}
	restRequest.Header.Set(headerAccept, headerApplicationJson)
	restRequest.Header.Set(headerUserAgent, headerUserAgentClient)

	// Transmit the span's TraceContext as HTTP headers on our outbound request.
	traceErr := opentracing.GlobalTracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(restRequest.Header))
	if traceErr != nil {
		errMsg := "Get product failed: " + traceErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)

		span.SetTag("product-found", false)
		span.SetTag("error", errMsg)
		span.LogKV("product-id", id, "product-found", false, "error", errMsg)
		return
	}

	response, respErr := s.restClient.Do(restRequest)
	if respErr != nil {
		errMsg := "Get product failed: " + respErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)

		span.SetTag("product-found", false)
		span.SetTag("error", errMsg)
		span.LogKV("product-id", id, "product-found", false, "error", errMsg)
		return
	}

	// TODO do not unmarshal response
	var product *commons.Product
	unmarshErr := json.NewDecoder(response.Body).Decode(&product)
	if unmarshErr != nil {
		errMsg := "Get product failed: " + unmarshErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)

		span.SetTag("product-found", false)
		span.SetTag("error", errMsg)
		span.LogKV("product-id", id, "product-found", false, "error", errMsg)
		return
	}
	defer response.Body.Close()

	span.SetTag("product-found", true)
	span.LogKV("product-id", id, "product-found", true)

	sendJsonResponse(writer, http.StatusOK, product)

	IncreaseRestRequests("getProduct")
	ObserveRestRequestsTime("getProduct", float64(time.Now().Sub(startTimer).Milliseconds()))
}

func (s *Server) createProduct(writer http.ResponseWriter, request *http.Request) {
	span := opentracing.StartSpan("create-product-handler")
	defer span.Finish()

	startTimer := time.Now()

	// TODO find a way to log body
	// requestBody, reqBosyErr := ioutil.ReadAll(request.Body)
	// if reqBosyErr != nil {
	// 	errMsg := "Create product failed: invalid request payload"
	// 	sendErrorResponse(writer, http.StatusBadRequest, errMsg)
	//
	// 	span.SetTag("product-created", false)
	// 	span.SetTag("error", errMsg)
	// 	span.LogKV("product-created", false, "error", errMsg)
	// 	return
	// }
	// defer request.Body.Close()
	// logging.SugaredLog.Infof("Create product %s", string(requestBody))
	logging.SugaredLog.Infof("Create product")

	endpointUrl := &url.URL{Path: rootProductsEndpoint}
	path := s.baseURL.ResolveReference(endpointUrl)
	restRequest, reqErr := http.NewRequest(http.MethodPost, path.String(), request.Body)
	defer request.Body.Close()
	if reqErr != nil {
		errMsg := "Create product failed: " + reqErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)

		span.SetTag("product-created", false)
		span.SetTag("error", errMsg)
		span.LogKV("product-created", false, "error", errMsg)
		return
	}
	restRequest.Header.Set(headerAccept, headerApplicationJson)
	restRequest.Header.Set(headerUserAgent, headerUserAgentClient)

	// Transmit the span's TraceContext as HTTP headers on our outbound request.
	traceErr := opentracing.GlobalTracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(restRequest.Header))
	if traceErr != nil {
		errMsg := "Create product failed: " + traceErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)

		span.SetTag("product-created", false)
		span.SetTag("error", errMsg)
		span.LogKV("product-created", false, "error", errMsg)
		return
	}

	response, respErr := s.restClient.Do(restRequest)
	if respErr != nil {
		errMsg := "Create product failed: " + respErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)

		span.SetTag("product-created", false)
		span.SetTag("error", errMsg)
		span.LogKV("product-created", false, "error", errMsg)
		return
	}

	// TODO do not unmarshal response
	var product *commons.Product
	unmarshErr := json.NewDecoder(response.Body).Decode(&product)
	if unmarshErr != nil {
		errMsg := "Create product failed: " + unmarshErr.Error()
		sendErrorResponse(writer, http.StatusBadRequest, errMsg)

		span.SetTag("product-created", false)
		span.SetTag("error", errMsg)
		span.LogKV("product-created", false, "error", errMsg)
		return
	}
	defer response.Body.Close()

	span.SetTag("product", product.String())
	span.SetTag("product-created", true)
	span.LogKV("product", product.String(), "product-created", true)

	sendJsonResponse(writer, http.StatusCreated, product)

	IncreaseRestRequests("createProduct")
	ObserveRestRequestsTime("createProduct", float64(time.Now().Sub(startTimer).Milliseconds()))
}

func (s *Server) updateProduct(writer http.ResponseWriter, request *http.Request) {
	span := opentracing.StartSpan("update-product-handler")
	defer span.Finish()

	startTimer := time.Now()

	vars := mux.Vars(request)
	id, idErr := strconv.Atoi(vars["id"])
	if idErr != nil {
		errMsg := "Update product failed: invalid product ID"
		sendErrorResponse(writer, http.StatusBadRequest, errMsg)

		span.SetTag("product-updated", false)
		span.SetTag("error", errMsg)
		span.LogKV("product-updated", false, "error", errMsg)
		return
	}

	// TODO find a way to log body
	// requestBody, reqBosyErr := ioutil.ReadAll(request.Body)
	// if reqBosyErr != nil {
	// 	errMsg := "Update product failed: invalid request payload"
	// 	sendErrorResponse(writer, http.StatusBadRequest, errMsg)
	//
	// 	span.SetTag("product-updated", false)
	// 	span.SetTag("error", errMsg)
	// 	span.LogKV("product-updated", false, "error", errMsg)
	// 	return
	// }
	// logging.SugaredLog.Infof("Update product with ID %d %s", id, string(requestBody))
	logging.SugaredLog.Infof("Update product with ID %d", id)
	span.SetTag("product-id", id)

	endpointUrl := &url.URL{Path: fmt.Sprintf(productsIdServerEndpoint, id)}
	path := s.baseURL.ResolveReference(endpointUrl)
	restRequest, reqErr := http.NewRequest(http.MethodPut, path.String(), request.Body)
	defer request.Body.Close()
	if reqErr != nil {
		errMsg := "Update product failed: " + reqErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)

		span.SetTag("product-updated", false)
		span.SetTag("error", errMsg)
		span.LogKV("product-updated", false, "error", errMsg)
		return
	}
	restRequest.Header.Set(headerAccept, headerApplicationJson)
	restRequest.Header.Set(headerUserAgent, headerUserAgentClient)

	// Transmit the span's TraceContext as HTTP headers on our outbound request.
	traceErr := opentracing.GlobalTracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(restRequest.Header))
	if traceErr != nil {
		errMsg := "Update product failed: " + traceErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)

		span.SetTag("product-updated", false)
		span.SetTag("error", errMsg)
		span.LogKV("product-updated", false, "error", errMsg)
		return
	}

	response, respErr := s.restClient.Do(restRequest)
	if respErr != nil {
		errMsg := "Update product failed: " + respErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)

		span.SetTag("product-updated", false)
		span.SetTag("error", errMsg)
		span.LogKV("product-updated", false, "error", errMsg)
		return
	}

	// TODO do not unmarshal response
	var product *commons.Product
	unmarshErr := json.NewDecoder(response.Body).Decode(&product)
	if unmarshErr != nil {
		errMsg := "Update product failed: " + unmarshErr.Error()
		sendErrorResponse(writer, http.StatusBadRequest, errMsg)

		span.SetTag("product-updated", false)
		span.SetTag("error", errMsg)
		span.LogKV("product-updated", false, "error", errMsg)
		return
	}
	defer response.Body.Close()

	span.SetTag("product", product.String())
	span.SetTag("product-updated", true)
	span.LogKV("product", product.String(), "product-updated", true)

	sendJsonResponse(writer, http.StatusCreated, product)

	IncreaseRestRequests("updateProduct")
	ObserveRestRequestsTime("updateProduct", float64(time.Now().Sub(startTimer).Milliseconds()))
}

func (s *Server) deleteProduct(writer http.ResponseWriter, request *http.Request) {
	span := opentracing.StartSpan("delete-product-handler")
	defer span.Finish()

	startTimer := time.Now()

	vars := mux.Vars(request)
	id, idErr := strconv.Atoi(vars["id"])
	if idErr != nil {
		errMsg := "Delete product failed: invalid product ID"
		sendErrorResponse(writer, http.StatusBadRequest, errMsg)

		span.SetTag("product-deleted", false)
		span.SetTag("error", errMsg)
		span.LogKV("product-deleted", false, "error", errMsg)
		return
	}

	logging.SugaredLog.Infof("Delete product by ID: %d", id)
	span.SetTag("product-id", id)

	endpointUrl := &url.URL{Path: fmt.Sprintf(productsIdServerEndpoint, id)}
	path := s.baseURL.ResolveReference(endpointUrl)
	restRequest, reqErr := http.NewRequest(http.MethodDelete, path.String(), nil)
	if reqErr != nil {
		errMsg := "Delete product failed: " + reqErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)

		span.SetTag("product-deleted", false)
		span.SetTag("error", errMsg)
		span.LogKV("product-deleted", false, "error", errMsg)
		return
	}
	restRequest.Header.Set(headerAccept, headerApplicationJson)
	restRequest.Header.Set(headerUserAgent, headerUserAgentClient)

	// Transmit the span's TraceContext as HTTP headers on our outbound request.
	traceErr := opentracing.GlobalTracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(restRequest.Header))
	if traceErr != nil {
		errMsg := "Delete product failed: " + traceErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)

		span.SetTag("product-deleted", false)
		span.SetTag("error", errMsg)
		span.LogKV("product-deleted", false, "error", errMsg)
		return
	}

	_, respErr := s.restClient.Do(restRequest)
	if respErr != nil {
		errMsg := "Update product failed: " + respErr.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)

		span.SetTag("product-updated", false)
		span.SetTag("error", errMsg)
		span.LogKV("product-updated", false, "error", errMsg)
		return
	}

	span.SetTag("product-deleted", true)
	span.LogKV("product-deleted", true)

	sendJsonResponse(writer, http.StatusOK, map[string]string{"result": "success"})

	IncreaseRestRequests("deleteProduct")
	ObserveRestRequestsTime("deleteProduct", float64(time.Now().Sub(startTimer).Milliseconds()))
}
