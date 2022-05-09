/**
 * @apiDefine ItemNotFoundError
 * @apiVersion 0.1.0
 * @apiError ItemNotFound The id of the item was not found.
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 404 Not Found
 *     {
 *       "error":  ItemNotFound"
 *     }
 */
// Item not found error

/**
 * @apiDefine InternalError
 * @apiVersion 0.1.0
 * @apiError InternalError There is an internal error.
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 505 Internal Server Error
 *     {
 *       "error":  InternalError"
 *     }
 */
// Internal Server Error

/**
 * @apiDefine admin Admin user access only
 * This optional description belong to to the group admin.
 */

/**
 * @api {get}  /items/:id Get item
 * @apiVersion 0.1.0
 * @apiName GetItem
 * @apiGroup Items
 *
 * @apiParam {Number} id Items unique ID.
 *
 * @apiSuccess {String} title  Item title.
 * @apiSuccess {Number} price  Item price.
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "id": 1,
 *       "title": "Apple",
 *       "price": "10.5"
 *     }
 *
 * @apiUse ItemNotFoundError
 * @apiUse InternalError
 */
// Get Item information

/**
 * @api {put}  /items Modify item
 * @apiVersion 0.1.0
 * @apiName PutItem
 * @apiGroup Items
 * @apiPermission admin
 *
 * @apiParam {Number} id       Items unique ID.
 * @apiParam {String} [title]  Item title.
 * @apiParam {Number} [price]  Item price.
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *
 * @apiUse InternalError
 */
// Modify Item information

/**
 * @api {delete}  /items/:id Delete item
 * @apiVersion 0.1.0
 * @apiName DeleteItem
 * @apiGroup Items
 *
 * @apiParam {Number} id       Items unique ID.
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *
 * @apiUse InternalError
 * @apiUse ItemNotFoundError
 */
// Delete Item

/**
 * @api {get} /items Get all items
 * @apiVersion 0.1.0
 * @apiName GetAllItems
 * @apiGroup Items
 *
 * @apiSuccess {Object[]} items 		Items list.
 * @apiSuccess {Number}   items.id 		Items unique ID.
 * @apiSuccess {String}   items.title 	Item title.
 * @apiSuccess {Number}   items.price 	Item price.
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     [{
 *	   	"id": 1,
 *		"title": "apple",
 *		"price": "10.5"
 *     }]
 *
 * @apiUse InternalError
 */
// Get all item information

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Item struct {
	Id    int
	Title string
	Price float32
}

type Items struct {
	Items []Item
}

func main() {
	fmt.Println("Starting server...")
	var err error

	db, err = sql.Open("postgres", "host=oleg-web.devops.rebrain.srwx.net user=api password=GQt5MTyVPuf9vsVWoWDT9YCn dbname=api sslmode=disable")
	CheckError(err)
	defer db.Close()

	http.HandleFunc("/", handler)
	http.HandleFunc("/v1/items", getItems)
	http.HandleFunc("/v1/item/", getItem)
	http.HandleFunc("/v1/item/add", addItem)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not an API command")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func getItems(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "MethodNotAllowed", http.StatusMethodNotAllowed)
	} else {
		var err error

		// insert
		// insertDynStmt := `insert into "items"("id","title", "price") values($1, $2, $3)`
		// _, err = db.Exec(insertDynStmt, 4, "tomato", 5.0)
		// CheckError(err)

		// update
		// updateStmt := `update "items" set "title"=$1, "price"=$2 where "id"=$3`
		// _, err = db.Exec(updateStmt, "bean", 3.0, 4)
		// CheckError(err)

		// Delete
		// deleteStmt := `delete from "items" where id=$1`
		// _, err = db.Exec(deleteStmt, 4)
		// CheckError(err)

		// getAll
		w_items := Items{}
		rows, err := db.Query(`SELECT "id", "title", "price" FROM "items"`)
		CheckError(err)
		defer rows.Close()

		for rows.Next() {
			w_item := Item{}
			err = rows.Scan(&w_item.Id, &w_item.Title, &w_item.Price)
			CheckError(err)
			w_items.Items = append(w_items.Items, w_item)
		}
		json.NewEncoder(w).Encode(w_items)
	}
}

func getItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var err error
		//var id string
		id := strings.TrimPrefix(r.URL.Path, "/v1/item/")
		id_int, err := strconv.Atoi(id)
		if err == nil {
			// getItem
			w_item := Item{}
			rows, err := db.Query(`SELECT "id", "title", "price" FROM "items" where "id" =` + strconv.Itoa(id_int))
			CheckError(err)
			defer rows.Close()
			for rows.Next() {
				err = rows.Scan(&w_item.Id, &w_item.Title, &w_item.Price)
				CheckError(err)
			}
			if w_item.Id != 0 {
				json.NewEncoder(w).Encode(w_item)
			} else {
				http.Error(w, "InputNotFound", http.StatusNotFound)
			}
		} else {
			//panic(err)
			http.Error(w, "InputNotAllowed", http.StatusForbidden)
		}
	} else if r.Method == "DELETE" {

		id := strings.TrimPrefix(r.URL.Path, "/v1/item/")
		id_int, err := strconv.Atoi(id)
		if err == nil {
			// getItem
			w_item := Item{}
			rows, err := db.Query(`SELECT "id", "title", "price" FROM "items" where "id" =` + strconv.Itoa(id_int))
			CheckError(err)
			defer rows.Close()
			for rows.Next() {
				err = rows.Scan(&w_item.Id, &w_item.Title, &w_item.Price)
				CheckError(err)
			}
			if w_item.Id != 0 {
				json.NewEncoder(w).Encode(w_item)
				_, err := db.Query(`DELETE FROM "items" where "id" =` + strconv.Itoa(id_int))
				CheckError(err)
				fmt.Fprintf(w, "Item DELETED!\n")
			} else {
				http.Error(w, "InputNotFound", http.StatusNotFound)
			}
		} else {
			//panic(err)
			http.Error(w, "InputNotAllowed", http.StatusForbidden)
		}
	} else {
		http.Error(w, "MethodNotAllowed", http.StatusMethodNotAllowed)
	}
}

func addItem(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "MethodNotAllowed", http.StatusMethodNotAllowed)
	} else {
		var err error
		var w_item Item
		var Id int
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&w_item)
		if err == nil {
			query := fmt.Sprintf(`insert into "items"("title", "price") values('%s', %.2f) returning id`, w_item.Title, w_item.Price)
			rows, err := db.Query(query)
			if err == nil {
				for rows.Next() {
					err = rows.Scan(&Id)
					CheckError(err)
				}
				json.NewEncoder(w).Encode(Id)
			} else {
				http.Error(w, "InputNotAllowed", http.StatusForbidden)
			}
			defer rows.Close()
		} else {
			http.Error(w, "InputNotAllowed", http.StatusForbidden)
		}

		// getAll
		w_items := Items{}
		rows, err := db.Query(`SELECT "id", "title", "price" FROM "items"`)
		CheckError(err)
		defer rows.Close()

		for rows.Next() {
			w_item := Item{}
			err = rows.Scan(&w_item.Id, &w_item.Title, &w_item.Price)
			CheckError(err)
			w_items.Items = append(w_items.Items, w_item)
		}
		json.NewEncoder(w).Encode(w_items)
	}
}
