package HomeAssist

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type HealthCheckResult struct {
	Status              string  `json:"status"`
	Checks              []Check `json:"checks"`
	TotalResponseTimeMS float32 `json:"total_response_time_ms,omitempty"`
	StartTime           string  `json:"start_time,omitempty"`
}

type Check struct {
	Key            string  `json:"key"`
	Status         string  `json:"status"`
	Value          string  `json:"value,omitempty"`
	ResponseTimeMS float32 `json:"response_time_ms,omitempty"`
}

func GetHC(w http.ResponseWriter, db *sql.DB) {
	startTime := time.Now()
	health := "healthy"
	checks := make([]Check, 0)

	start := time.Now()
	err := checkDatabase(db)
	elapsed := time.Since(start).Seconds() * 1000

	if err != nil {
		checks = append(checks, Check{
			Key:            "failed_db_check",
			Status:         "unhealthy",
			Value:          err.Error(),
			ResponseTimeMS: float32(elapsed),
		})
		health = "unhealthy"
	} else {
		checks = append(checks, Check{
			Key:            "successful_db_check",
			Status:         "healthy",
			Value:          "No Errors :)",
			ResponseTimeMS: float32(elapsed),
		})
	}

	totalResponseTime := float32(time.Since(startTime).Seconds() * 1000)
	healthCheckResult := HealthCheckResult{
		Status:              health,
		Checks:              checks,
		TotalResponseTimeMS: totalResponseTime,
		StartTime:           startTime.Format(time.RFC3339Nano),
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(healthCheckResult)
	if err != nil {
		http.Error(w, "Error encoding JSON response: "+err.Error(), http.StatusInternalServerError)
	}
}

func checkDatabase(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		log.Printf("Error connecting to the database: %v", err)
		return err
	}
	return nil
}
