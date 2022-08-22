package h3revgeolocfi

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	_ "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

const dbpath = "db.db"

func loaddb() (*sql.DB, error) {
	var err error
	var fn string
	if os.Getenv("GVM_ROOT") == "" {
		fn = path.Join("serverless_function_source_code", dbpath)
	} else {
		fn = path.Join("../", dbpath)
	}
	_, err = os.Stat(fn)
	if err != nil {
		return nil, err
	}
	dbhan, err := sql.Open("sqlite3", fn)
	if err != nil {
		return nil, err
	}
	return dbhan, nil
}

func init() {
	var err error
	db, err = loaddb()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inited")
	//db.Close()
}

func query(h3id string) (string, string, error) {
	stmt, err := db.Prepare("select c,n from loc where i = ?")
	if err != nil {
		return "", "", err
	}
	defer stmt.Close()
	var c string
	var n string
	err = stmt.QueryRow(h3id).Scan(&c, &n)
	if err != nil {
		return "", "", err
	}
	return c, n, nil
}

func main() {
	ii := "890888534d3ffff"
	c, n, err := query(ii)
	if err != nil {
		panic(err)
	}
	fmt.Println(c, n)

}

type Msg struct {
	H3ID string `json:"h3id"`
}

type Resp struct {
	C string `json:"city"`
	N string `json:"neighborhood"`
}

func H3RevGeoLocFi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Connection", "keep-alive")
		w.Header().Add("Access-Control-Allow-Methods", "POST")
		w.Header().Add("Access-Control-Allow-Headers", "content-type")
		w.Header().Add("Access-Control-Max-Age", "86400")
		w.WriteHeader(http.StatusOK)
		return
	}
	d := Msg{}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c, n, err := query(d.H3ID)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(&Resp{C: c, N: n})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
