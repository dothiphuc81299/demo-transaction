package util

import (
	"bytes"
	"demo-transaction/config"
	"encoding/json"
	"fmt"
	"io"

	"demo-transaction/modules/database"
)

// HelperToIOReader ...
func HelperToIOReader(i interface{}) io.Reader {
	b, _ := json.Marshal(i)
	return bytes.NewReader(b)
}

// HelperConnect ...
func HelperConnect() {
	var (
		envVars = config.GetEnv()
		client  = database.GetClient()
	)

	// Set database for test ...
	db := client.Database(envVars.Database.TestName)
	fmt.Println("Database Connected to", envVars.Database.TestName)
	database.SetDB(db)
}
