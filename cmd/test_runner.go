package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/config"
	mylog "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/logger"
	"github.com/google/uuid"
)

func RunAPITests() error {
	mylog.Infoln("Starting API Integration Tests...")

	rest, err := config.NewRest()
	if err != nil {
		return fmt.Errorf("failed to initialize rest: %w", err)
	}
	router := rest.GetServer()

	// Helper function for making requests
	request := func(method, url string, bodyObj ...interface{}) (*httptest.ResponseRecorder, map[string]interface{}) {
		var bodyReader io.Reader
		if len(bodyObj) > 0 {
			b, _ := json.Marshal(bodyObj[0])
			bodyReader = bytes.NewBuffer(b)
		}

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(method, url, bodyReader)
		if bodyReader != nil {
			req.Header.Set("Content-Type", "application/json")
		}
		router.ServeHTTP(w, req)

		var body map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &body)

		if w.Code >= 400 && w.Code != 404 && w.Code != 401 && w.Code != 403 {
			mylog.Errorf("Request failed: %s %s -> Status %d, Body: %s", method, url, w.Code, w.Body.String())
		}

		return w, body
	}

	// Login and setup authRequest helper
	w, body := request("POST", "/api/v1/auth/login", map[string]string{"username": "admin", "password": "admin"})
	token := ""
	if w.Code == 200 {
		if data, ok := body["data"].(map[string]interface{}); ok {
			token = data["token"].(string)
		}
	}

	authRequest := func(method, url string, bodyObj ...interface{}) (*httptest.ResponseRecorder, map[string]interface{}) {
		var bodyReader io.Reader
		if len(bodyObj) > 0 {
			b, _ := json.Marshal(bodyObj[0])
			bodyReader = bytes.NewBuffer(b)
		}

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(method, url, bodyReader)
		req.Header.Set("Content-Type", "application/json")
		if token != "" {
			req.Header.Set("Authorization", "Bearer "+token)
		}
		router.ServeHTTP(w, req)

		var resBody map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &resBody)
		return w, resBody
	}

	// 1. Cabinet Info Tests
	mylog.Infoln("Testing Cabinet Info...")
	w, body = request("GET", "/api/v1/cabinet-info")
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

	// 3. Progenda Tests
	mylog.Infoln("Testing Progenda...")
	w, body = request("GET", "/api/v1/progenda")
	if w.Code != http.StatusOK {
		return fmt.Errorf("GET /progenda failed: expected 200, got %d", w.Code)
	}
	mylog.Infof("  GET /progenda: OK (found %v items)", len(body["data"].([]interface{})))

	// Test Create Progenda
	if token != "" {
		mylog.Infoln("  Testing Progenda CRUD with Auth...")
		newProgenda := map[string]interface{}{
			"name":           "Test Progenda " + uuid.New().String(),
			"goal":           "Test Goal",
			"description":    "Test Description",
			"instagram_link": "https://ig.com/test",
			"timelines": []map[string]string{
				{"event_name": "Event 1", "date": "Jan 2024"},
				{"event_name": "Event 2", "date": "Feb 2024"},
			},
		}
		w, body = authRequest("POST", "/api/v1/progenda", newProgenda)
		if w.Code != http.StatusCreated {
			return fmt.Errorf("POST /progenda failed: expected 201, got %d", w.Code)
		}
		progData := body["data"].(map[string]interface{})
		progId := progData["id"].(string)
		mylog.Infof("  POST /progenda: OK (id: %s)", progId)

		// Verify Timelines preloaded in detail
		w, body = request("GET", "/api/v1/progenda/"+progId)
		if w.Code != http.StatusOK {
			return fmt.Errorf("GET /progenda/%s failed: expected 200, got %d", progId, w.Code)
		}
		detailData := body["data"].(map[string]interface{})
		timelines := detailData["timelines"].([]interface{})
		if len(timelines) != 2 {
			return fmt.Errorf("expected 2 timelines, got %d", len(timelines))
		}
		mylog.Infoln("  GET /progenda/:id (Preload Timelines): OK")

		// Test Update
		updateReq := map[string]interface{}{
			"instagram_link": "https://ig.com/updated",
			"timelines": []map[string]string{
				{"event_name": "Revised Event", "date": "Mar 2024"},
			},
		}
		w, body = authRequest("PUT", "/api/v1/progenda/"+progId, updateReq)
		if w.Code != http.StatusOK {
			return fmt.Errorf("PUT /progenda/%s failed: expected 200, got %d", progId, w.Code)
		}
		if body["data"].(map[string]interface{})["instagram_link"] != "https://ig.com/updated" {
			return fmt.Errorf("update failed: instagram link not updated")
		}
		updatedTimelines := body["data"].(map[string]interface{})["timelines"].([]interface{})
		if len(updatedTimelines) != 1 {
			return fmt.Errorf("update failed: expected 1 timeline, got %d", len(updatedTimelines))
		}
		mylog.Infoln("  PUT /progenda/:id (Update Timelines): OK")

		// Delete
		w, _ = authRequest("DELETE", "/api/v1/progenda/"+progId)
		if w.Code != http.StatusOK {
			return fmt.Errorf("DELETE /progenda/%s failed: expected 200, got %d", progId, w.Code)
		}
		mylog.Infoln("  DELETE /progenda: OK")
	}

	// 4. News Tests with Search & Pagination
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

	// Test Autocompletion
	w, body = request("GET", "/api/v1/news/autocompletion?search=Pene")
	if w.Code != http.StatusOK {
		return fmt.Errorf("GET /news/autocompletion failed: expected 200, got %d", w.Code)
	}
	mylog.Infoln("  GET /news/autocompletion: OK")

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

	// 4. NRP Whitelist Tests
	mylog.Infoln("Testing NRP Whitelist...")
	w, body = request("GET", "/api/v1/nrp-whitelist")
	if w.Code != http.StatusOK {
		return fmt.Errorf("GET /nrp-whitelist failed: expected 200, got %d", w.Code)
	}
	mylog.Infof("  GET /nrp-whitelist: OK (found %v items)", len(body["data"].([]interface{})))

	// Test Check Whitelist (Not Found)
	w, body = request("POST", "/api/v1/nrp-whitelist", map[string]string{"nrp": "9999999999"})
	if w.Code != http.StatusForbidden {
		return fmt.Errorf("POST /nrp-whitelist (invalid) failed: expected 403, got %d", w.Code)
	}
	mylog.Infoln("  POST /nrp-whitelist (invalid): OK (403)")

	// Test Create Whitelist (Admin)
	testNrp := "5025211014"
	if token != "" {
		mylog.Infoln("  Authorized for NRP CRUD tests")

		// CLEANUP: Ensure fresh state
		authRequest("DELETE", "/api/v1/nrp-whitelist/"+testNrp)

		// Create
		w, body = authRequest("POST", "/api/v1/nrp-whitelist/add", map[string]string{"nrp": testNrp, "name": "Test User"})
		if w.Code == http.StatusOK || w.Code == http.StatusCreated {
			mylog.Infoln("  POST /nrp-whitelist/add: OK (Created)")

			// Duplicate (Expect 400)
			w, body = authRequest("POST", "/api/v1/nrp-whitelist/add", map[string]string{"nrp": testNrp, "name": "Test User"})
			if w.Code == http.StatusBadRequest {
				mylog.Infoln("  POST /nrp-whitelist/add (duplicate): OK (400)")
			} else {
				return fmt.Errorf("expected 400 for duplicate NRP, got %d", w.Code)
			}

			// Public check
			w, body = request("POST", "/api/v1/nrp-whitelist", map[string]string{"nrp": testNrp})
			if w.Code == http.StatusOK {
				mylog.Infoln("  POST /nrp-whitelist (valid check): OK")
			}

			// Final cleanup
			authRequest("DELETE", "/api/v1/nrp-whitelist/"+testNrp)
		}
	}

	// 5. Gallery Tests with Filtering
	mylog.Infoln("Testing Gallery Filtering...")
	w, body = request("GET", "/api/v1/gallery?filterby=category&filter=logo")
	if w.Code != http.StatusOK {
		return fmt.Errorf("GET /gallery?filterby=category&filter=logo failed: expected 200, got %d", w.Code)
	}
	mylog.Infoln("  GET /gallery?filterby=category&filter=logo: OK")

	mylog.Infoln("All API tests passed successfully!")
	return nil
}
