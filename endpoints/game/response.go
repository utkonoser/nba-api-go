package game

import (
	"encoding/json"
	"fmt"
)

// StatsResponse represents a standard NBA Stats API response.
type StatsResponse struct {
	Resource   string      `json:"resource"`
	Parameters interface{} `json:"parameters"`
	ResultSets []ResultSet `json:"resultSets"`
}

// ResultSet represents a single result set in the response.
type ResultSet struct {
	Name    string          `json:"name"`
	Headers []string        `json:"headers"`
	RowSet  [][]interface{} `json:"rowSet"`
}

// GetDataSet returns a specific result set by name.
func (r *StatsResponse) GetDataSet(name string) (*ResultSet, error) {
	for _, rs := range r.ResultSets {
		if rs.Name == name {
			return &rs, nil
		}
	}
	return nil, fmt.Errorf("result set '%s' not found", name)
}

// ToMap converts a result set to a slice of maps.
func (rs *ResultSet) ToMap() []map[string]interface{} {
	var result []map[string]interface{}

	for _, row := range rs.RowSet {
		rowMap := make(map[string]interface{})
		for i, header := range rs.Headers {
			if i < len(row) {
				rowMap[header] = row[i]
			}
		}
		result = append(result, rowMap)
	}

	return result
}

// ToJSON converts the result set to JSON.
func (rs *ResultSet) ToJSON() (string, error) {
	data := rs.ToMap()
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal to JSON: %w", err)
	}
	return string(jsonData), nil
}

// GetRow returns a specific row as a map.
func (rs *ResultSet) GetRow(index int) (map[string]interface{}, error) {
	if index < 0 || index >= len(rs.RowSet) {
		return nil, fmt.Errorf("row index %d out of range", index)
	}

	rowMap := make(map[string]interface{})
	row := rs.RowSet[index]
	for i, header := range rs.Headers {
		if i < len(row) {
			rowMap[header] = row[i]
		}
	}

	return rowMap, nil
}

// RowCount returns the number of rows in the result set.
func (rs *ResultSet) RowCount() int {
	return len(rs.RowSet)
}

