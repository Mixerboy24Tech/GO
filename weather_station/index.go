package main
 
import (
    "fmt"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)
 
func main() {
    db, err := sql.Open("mysql", "username:password@tcp(host.tdl)/weather")
    defer db.Close()
 
    if err != nil {
        log.Fatal(err.Error())
    }
 
    err = db.Ping()
    if err != nil {
        log.Fatal(err.Error())
    }
 
 
    rows, err := db.Query("SELECT AMBIENT_TEMPERATURE FROM WEATHER_MEASUREMENT ORDER BY ID DESC LIMIT 1")
    defer rows.Close()
 
    if err != nil {
        log.Fatal(err.Error())
    }
 
    for rows.Next() {
        var AMBIENT_TEMPERATURE  string
        if err := rows.Scan(&AMBIENT_TEMPERATURE); err != nil {
                log.Fatal(err.Error())
        }
        fmt.Println("Lampotila: ", AMBIENT_TEMPERATURE," astetta")
    }
    if err := rows.Err(); err != nil {
        log.Fatal(err.Error())
    }
 
 
}
