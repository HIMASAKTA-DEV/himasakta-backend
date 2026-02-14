package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/config"
	mylog "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/logger"
)

func RunAPITests() error {
	mylog.Infoln("Starting API Integration Tests...")

	rest, err := config.NewRest()
	if err != nil {
		return fmt.Errorf("failed to initialize rest: %w", err)
	}
	router := rest.GetServer()

	// Helper function for making requests
	request := func(method, url string) (*httptest.ResponseRecorder, map[string]interface{}) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(method, url, nil)
		router.ServeHTTP(w, req)

		var body map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &body)

		if w.Code >= 400 && w.Code != 404 {
			mylog.Errorf("Request failed: %s %s -> Status %d, Body: %s", method, url, w.Code, w.Body.String())
		}

		return w, body
	}

	// 1. Cabinet Info Tests
	mylog.Infoln("Testing Cabinet Info...")
	w, body := request("GET", "/api/v1/cabinet-info")
	if w.Code != 200 {
		return fmt.Errorf("GET /cabinet-info failed: expected 200, got %d", w.Code)
	}
	mylog.Infof("  GET /cabinet-info: OK (found %v items)", len(body["data"].([]interface{})))

	w, _ = request("GET", "/api/v1/current-cabinet")
	if w.Code != 200 && w.Code != 404 {
		return fmt.Errorf("GET /current-cabinet failed: expected 200 or 404, got %d", w.Code)
	}
	mylog.Infof("  GET /current-cabinet: OK (%d)", w.Code)

	// 2. Department Tests
	mylog.Infoln("Testing Department...")
	w, body = request("GET", "/api/v1/department")
	if w.Code != http.StatusOK {
		return fmt.Errorf("GET /department failed: expected 200, got %d", w.Code)
	}
	mylog.Infof("  GET /department: OK (found %v items)", len(body["data"].([]interface{})))

	// Test Get by Name
	if depts, ok := body["data"].([]interface{}); ok && len(depts) > 0 {
		firstDept := depts[0].(map[string]interface{})
		deptName := firstDept["name"].(string)
		w, _ = request("GET", "/api/v1/department/"+deptName)
		if w.Code != http.StatusOK {
			return fmt.Errorf("GET /department/%s failed: expected 200, got %d", deptName, w.Code)
		}
		mylog.Infof("  GET /department/%s: OK", deptName)
	}

	// 3. News Tests with Search & Pagination
	mylog.Infoln("Testing News (Search & Pagination)...")

	// Test Limit
	w, body = request("GET", "/api/v1/news?limit=1")
	if w.Code != http.StatusOK {
		return fmt.Errorf("GET /news?limit=1 failed: expected 200, got %d", w.Code)
	}
	data := body["data"].([]interface{})
	if len(data) > 1 {
		return fmt.Errorf("pagination failed: expected at most 1 item, got %d", len(data))
	}
	mylog.Infoln("  GET /news?limit=1: OK")

	// Test Get by Slug
	if len(data) > 0 {
		firstNews := data[0].(map[string]interface{})
		slug := firstNews["slug"].(string)
		w, _ = request("GET", "/api/v1/news/"+slug)
		if w.Code != http.StatusOK {
			return fmt.Errorf("GET /news/%s failed: expected 200, got %d", slug, w.Code)
		}
		mylog.Infof("  GET /news/%s: OK", slug)
	}

	// Test 404 Not Found
	mylog.Infoln("Testing 404 Not Found...")
	w, _ = request("GET", "/api/v1/news/this-slug-does-not-exist")
	if w.Code != http.StatusNotFound {
		return fmt.Errorf("expected 404 for missing news, got %d", w.Code)
	}
	mylog.Infoln("  GET /news/invalid: OK (404)")

	// 4. Gallery Tests with Filtering
	mylog.Infoln("Testing Gallery Filtering...")
	w, body = request("GET", "/api/v1/gallery?filterby=category&filter=logo")
	if w.Code != http.StatusOK {
		return fmt.Errorf("GET /gallery?filterby=category&filter=logo failed: expected 200, got %d", w.Code)
	}
	mylog.Infoln("  GET /gallery?filterby=category&filter=logo: OK")

	mylog.Infoln("All API tests passed successfully!")
	return nil
}
