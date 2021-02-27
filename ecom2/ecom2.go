// ecom2
package main

//testing
//https://www.golangprograms.com/example-of-golang-crud-using-mysql-from-scratch.html

//////////
import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

//$keyword1 = $_GET['keyword'];
//$titleOfSelectedDropDown = $_GET['val1'];

var titleOfSelectedDropDown = "computers"

//$gKeyword1 = "";
//$gKeyword2 = "";
//$gKeyword3 = "";

//$key1ID = "";
//$key2ID = "";
//$key3ID = "";

var string1 = ""

type employee struct {
	gKeyword1           string
	gKeyword2           string
	gKeyword3           string
	ProductName         string
	ProductID           int
	ProductdDescription string
	ProductCost         int
	ProductQuantity     int
	ProductCatTitle     string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "ecommerce"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func index(w http.ResponseWriter, r *http.Request) {

	db := dbConn()

	//var q = "ELECT products.ProductKeyword1, products.Product"
	var q = "SELECT products.ProductKeyword1, products.ProductKeyword2, products.ProductKeyword3, products.ProductName, products.ProductID, " +
		"products.ProductDescription, products.ProductCost, products.ProductQuantity, products.ProductCatTitle " +
		"FROM products WHERE " +
		"((products.ProductKeyWord1 = \"apple1\") OR " +
		"(products.ProductKeyWord2 = \"apple1\") or (products.ProductKeyWord3 = \"apple1\" ))"

		//and products.ProductCatTitle = \"$titleOfSelectedDropDown\"  "

	selDB, err := db.Query(q)
	if err != nil {
		panic(err.Error())
	}
	//emp := employee{}
	//res := []employee{}

	counter := 0

	for selDB.Next() {

		var ProductID, ProductCost, ProductQuantity int
		var gKeyword1, gKeyword2, gKeyword3, ProductName, ProductDescription, ProductCatTitle string

		err = selDB.Scan(&gKeyword1, &gKeyword2, &gKeyword3, &ProductName, &ProductID, &ProductDescription, &ProductCost, &ProductQuantity, &ProductCatTitle)

		if err != nil {
			panic(err.Error())
		}

		//var title1 = "testinga this"

		counter = counter + 1
		str := strconv.Itoa(counter)
		//deleteFlag := 1
		var titleID = "titleID" + str
		var descID = "descID" + str
		var costID = "costID" + str
		var quantityID = "quantityID" + str
		var key1ID = "key1ID" + str
		var key2ID = "key2ID" + str
		var key3ID = "key3ID" + str
		//var var100
		//var var100 = "true"
		//var var99 = 1
		//var mainDiv = var1 + str

		//var titleID = 0

		/*
			///////////////
			var testvar = "isitseen"
			string1 = "<div class=\"container\">" +

				"<div class=\"row\" >" +
				"<div class=\"col\">" +
				"<p id =\"\"  >Title</p>" +
				"</div>" +
				"<div class=\"col\">" +
				"<p id =\"\"  >Title</p>" +
				"</div> " +
				"<div class=\"col\">" +
				"<p> <input id = value =\"\" placeholder=" + testvar + " ></p>" +
				"<p id = \"link1\">  product id   : " + strconv.Itoa(ProductID) + " </p>" +
				"</div></div></div> "

			//////////
		*/

		string1 = "<p id = \"link1\">product id   : " + strconv.Itoa(ProductID) + " </p>" +
			"<p>category id  : " + ProductCatTitle + "</p>" +

			"<div class=\"container\">" +
			"<div class=\"row\" >" +

			"<div class=\"col\">" +
			"<h4><center><p id = \"\">Image</p></center></h4>" +
			"<center><p id = \"\"> image </p></center>" +
			"</div>" +

			"<div class=\"col\">" +
			"<h4><center><p id =\"\"  >Title</p></center></h4>" +

			"<center>      <p  >      <input id = " + titleID + " value = " + ProductName + " type=\"text\" name=\"title\" placeholder=\"\"></p></center>" +
			"</div>" +

			"<div class=\"col\">" +

			"<h4><center><p id = \"\">Desc</p></center></h4>" +

			"<center><textarea wrap id = " + descID + "   value = " + ProductDescription + "  type=\"text\" rows=\"5\" cols=\"34\"></textarea></center>" +
			"</div>" +

			"<div class=\"col\">" +
			"<h4><center><p id = \"\" >Cost</p></center></h4>" +
			"<center><p>	<input id = " + costID + " value =  " + strconv.Itoa(ProductCost) + "   type=\"number\" name=\"title\" placeholder=\"\">		</p></center>" +

			"</div>" +

			"</div>" +
			"</div>" +

			"<div class=\"container\">" +
			"<div class=\"row\" >" +

			"<div class=\"col\">" +
			"<h4><center><p id = \"\" >Quantity</p></center></h4>" +
			"<center><p> <input id = " + quantityID + " value = " + strconv.Itoa(ProductQuantity) + "  type=\"number\" name=\"title\" placeholder=\"\">	</p></center>" +
			"</div>" +

			"<div class=\"col\">" +
			"<h4><center><p id = \"\" >Keyword 1</p></center></h4>" +
			"<center><p>	<input id = " + key1ID + " value = " + gKeyword1 + " type=\"text\" name=\"title\" placeholder=\"\">		</p></center>" +

			"</div>" +

			"<div class=\"col\">" +
			"<h4><center><p id = \"\" >Keyword 2</p></center></h4>" +
			"<center><p>	<input id = " + key2ID + " value = " + gKeyword2 + " type=\"text\" name=\"title\" placeholder=\"\">		</p></center>" +

			"</div>" +

			"<div class=\"col\">" +
			"<h4><center><p id = \"\" >Keyword 3</p></center></h4>" +
			"<center><p>	<input id =   " + key3ID + "  value = " + gKeyword3 + " type=\"text\" name=\"title\" placeholder=\"\">		</p></center>" +

			"</div>" +

			"<br><br>" +
			"</div>" +
			"</div>" +

			"<div class=\"container\">" +
			"<div class=\"row\" >" +

			"<div class=\"col\">" +

			//<!--
			//deleteFlag + "); , mainDiv , strconv.Itoa(ProductID) , titleID , descID,costID,quantityID ,key1ID,key2ID,key3ID
			"<center><button id = \"\" onclick = \"SaveProductItems(" + strconv.Itoa(ProductID) + ", " + titleID + ", " + descID + " )\">Submit</button></center>" +

			//<!--flag for determining if record delete will effect" sessioncount, 0 is no.-->
			//"<center><button id = \"\" onclick = \"deleteRecord( 1,  mainDiv, strconv.Itoa(ProductID) )\">Delete</button></center>" +
			//"<p><a href=\"#add\">To Add</a></p>" +
			//-->
			"</div>" +
			"<br><br>" +

			"</div>" +
			"</div>" +

			" <br><br><br><br>"

	} //for selDB.Next()

	receiveAjax(w, r)
}

func receiveAjax(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		data := r.FormValue("post_data")
		fmt.Println("Receive ajax post data string ", data)
		w.Header().Add("Content-Type", "application/html")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Write([]byte(string1))

	}
}

func main() {
	mux := http.NewServeMux()
	//mux.HandleFunc("/receive", receiveAjax)

	mux.HandleFunc("/receive", index)

	http.ListenAndServe(":8080", mux)
}

//	"<center><button id = \"\" onclick = \"SaveProductItems( " + deleteFlag + "); , mainDiv , strconv.Itoa(ProductID) , titleID , descID,costID,quantityID ,key1ID,key2ID,key3ID
