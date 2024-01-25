package datacenter

// import (
// 	//"crypto/sha256"

// 	_ "github.com/mattn/go-sqlite3"
// )

// var createTableQuery = `
// CREATE TABLE IF NOT EXISTS triggers (
// Location TEXT,
// Cookies TEXT,
// Referrer TEXT,
// UserAgent TEXT,
// BrowserTime TEXT,
// Origin TEXT,
// DOM TEXT,
// IP TEXT,
// Method TEXT,
// Screenshot TEXT,
// HASH TEXT,
// Collection TEXT,
// Timestamp INTEGER,
// Visible INTEGER
// );
// `

// var insertQuery = `
//   INSERT INTO triggers (
//     Location,
//     Cookies,
//     Referrer,
//     UserAgent,
//     BrowserTime,
//     Origin,
//     DOM,
//     IP,
//     Method,
//     Screenshot,
//     HASH,
//     Collection,
//     Timestamp,
//     Visible
//   ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
// `

// var removeQuery = `
//   		DELETE FROM triggers
//       WHERE HASH = ?;
//   	`

// var readQuery = `
//   		SELECT * FROM triggers
//   		WHERE Visible = 1 AND Timestamp > ?;
//   	`

// var dataPath string = misc.DataDir + "/" + "data"

// func CreateDatabase() {
// 	db := opendb()
// 	defer db.Close()

// 	_, err := db.Exec(createTableQuery)
// 	if err != nil {
// 		misc.ErrorLog.Printf("Error creating table in database: %s", err)
// 	}
// }

// func saveNewData(data *Data) {
// 	db := opendb()
// 	defer db.Close()

// 	_, err := db.Exec(insertQuery, data.Location, data.Cookies, data.Referrer, data.UserAgent, data.BrowserTime, data.Origin, data.DOM, data.IP, data.Method, data.Screenshot, data.HASH, data.Collection, data.Timestamp, 1)
// 	if err != nil {
// 		misc.ErrorLog.Printf("Error inserting data into database: %s", err)
// 	}
// }

// func readData(timestamp string) []*Data {
// 	db := opendb()
// 	defer db.Close()

// 	// Query the data from the table
// 	rows, err := db.Query(readQuery, timestamp)
// 	if err != nil {
// 		misc.ErrorLog.Printf("Error querying database: %s", err)
// 	}
// 	defer rows.Close()

// 	var collection []*Data

// 	for rows.Next() {
// 		data := &Data{}
// 		err := rows.Scan(
// 			&data.Location,
// 			&data.Cookies,
// 			&data.Referrer,
// 			&data.UserAgent,
// 			&data.BrowserTime,
// 			&data.Origin,
// 			&data.DOM,
// 			&data.IP,
// 			&data.Method,
// 			&data.Screenshot,
// 			&data.HASH,
// 			&data.Collection,
// 			&data.Timestamp,
// 			&data.Visible,
// 		)
// 		if err != nil {
// 			misc.ErrorLog.Printf("Error scanning database row: %s", err)
// 		}
// 		collection = append([]*Data{data}, collection...)
// 	}

// 	if err := rows.Err(); err != nil {
// 		misc.ErrorLog.Printf("Error retrieving database row: %s", err)
// 	}
// 	return collection
// }

// func removeData(hash string) {
// 	db := opendb()
// 	defer db.Close()

// 	_, err := db.Exec(removeQuery, hash)
// 	if err != nil {
// 		misc.ErrorLog.Printf("Error deleting data: %s", err)
// 	}
// }

// func opendb() *sql.DB {
// 	db, err := sql.Open("sqlite3", misc.DataDir+"/"+"wwdatabase.db")
// 	if err != nil {
// 		misc.ErrorLog.Printf("Error opening database: %s", err)
// 	}
// 	return db
// }
