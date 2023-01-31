package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// The handleProducts function is a HTTP handler function that routes requests to the appropriate function based on the request method.
// It starts by checking the request method, it could be "POST", "PUT", "DELETE" and "GET",
// if the request method is "POST" it will call the createProduct(w, r) function,
// if the request method is "PUT" it will call the handleUpdateProduct(w, r) function,
// if the request method is "DELETE" it will call the handleDeleteProduct(w, r) function,
// and if the request method is "GET" it will call the handleGetProducts(w, r) function.
func handleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		createProduct(w, r)
	case "PUT":
		handleUpdateProduct(w, r)
	case "DELETE":
		handleDeleteProduct(w, r)
	case "GET":
		handleGetProducts(w, r)
	}
}

// This function is a handler for HTTP requests to create a new product.
// It creates a new product by decoding a JSON object from the request body, inserting it into a PostgreSQL database,
// and returning the product's ID and other details in the response body.
// The function starts by checking if the request method is "POST", and if not,
// it returns an error with a status code of "405 Method Not Allowed".
// Then, it creates a new Product struct variable and decodes the request body into that
// struct using json.NewDecoder. If there is an error with decoding, it returns an error with a status code of "400 Bad Request".
// It then constructs a connection string to a PostgreSQL database using the "host", "port", "user", "password", and "dbname" variables, and
// opens a connection to the database using sql.Open. If there is an error connecting to the database, it returns an error with a status code of "500 Internal Server Error".
// It then perform a "ping" to the database to check the connection,
// if it fails returns an error with a status code of "500 Internal Server Error"
// Then it creates an SQL query to insert the product details into the "products" table and retrieve the newly-assigned ID.
//  It uses the product's name, description, and price in the query.
// Finally, it sets the "Content-Type" header of the response to "application/json" and writes the product's
// ID and other details to the response body using json.NewEncoder

// It also sets the status code of the response to "201 Created" if there is no error.
func createProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db = GetDb()

	query := `
		INSERT INTO products (name, description, price)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	var id int
	if err := db.QueryRow(query, product.Name, product.Description, product.Price).Scan(&id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	product.ID = id

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// This  function handleUpdateProduct which is a HTTP handler function that
// updates an existing product in a PostgreSQL database.
// It starts by decoding a JSON object from the request body, this JSON object should contain the updated product details.
// If there is an error with decoding, it returns an error with a status code of "400 Bad Request".
// It then constructs a connection string to a PostgreSQL database using the "host", "port", "user", "password", and "dbname" variables, and opens a connection
// to the database using sql.Open. If there is an error connecting to the database, it returns an error with a status code of "500 Internal Server Error".
// It then perform a "ping" to the database to check the connection,
// if it fails returns an error with a status code of "500 Internal Server Error"
// Then it creates an SQL query to update the product details in the "products" table using the ID, Name
// , Description, and Price fields of the product struct in the request body. It uses the db.Exec to execute the query
// If there is an error executing the query, it returns an error with a
// status code of "500 Internal Server Error"
// If everything goes well it sets the
//
//	status code of the response to "200 OK"
//
// It is worth noting that, the ID field must be present in the request body and must match
// an existing product ID in the database in order to update the product correctly.
func handleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db = GetDb()
	fmt.Println("DB CONNECTED")
	query := `
		UPDATE products
		SET name = $2, description = $3, price = $4
		WHERE id = $1
	`
	if _, err := db.Exec(query, product.ID, product.Name, product.Description, product.Price); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// This function handleDeleteProduct which is a HTTP handler function that delete an existing
//  product from a PostgreSQL database.

// It starts by decoding a JSON object from the request body, this JSON object should contain
//  the ID of the product that needs to be deleted. If there is an error with decoding, it returns an error with a status code of "400 Bad Request".

// It then constructs a connection string to a PostgreSQL database using the "host", "port", "user", "password", and "dbname" variables,
//  and opens a connection to the database using sql.Open. If there is an error connecting to the database,
//  it returns an error with a status code of "500 Internal Server Error".

// It then perform a "ping" to the database to check the connection,
// if it fails returns an error with a status code of "500 Internal Server Error"

// Then it creates an SQL query to delete the product from the "products" table using the ID of the product struct in the request body.
// It uses the db.Exec to execute the query
// If there is an error executing the query, it returns an error with a status code of "500 Internal Server Error"
// If everything goes well it sets the status code of the response to "200 OK"

// It is worth noting that, the ID field must be present in the request body and must match an existing product ID in the database in order to delete the product correctly.

func handleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db = GetDb()
	query := "DELETE FROM products WHERE id = $1"
	if _, err := db.Exec(query, product.ID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// handleGetProducts is a HTTP handler function that retrieves a list of products from a PostgreSQL database,
// and returns the product details in the response body as a JSON object. It starts by constructing a
// connection string to the PostgreSQL database using the "host", "port", "user", "password", and "dbname" variables,
// and opens a connection to the database using sql.Open. It then checks the connection by pinging the database,
// if it fails returns an error with a status code of "500 Internal Server Error"

// It then creates an SQL query to select all products from the "products" table,
//  and executes that query using the db.Query method. If there is an error executing the query,
// it returns an error with a status code of "500 Internal Server Error"

// It then scans the query results row by row, and creates a new Product struct variable for each row, with ID, Name, Description,
// and Price fields populated from the corresponding columns in the current row.

// It then set the "Content-Type" header of the response to "application/json" and write the array
//  of product struct in the response body using json.NewEncoder

func handleGetProducts(w http.ResponseWriter, r *http.Request) {
	db = GetDb()
	query := "SELECT id, name, description, price FROM products"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// This function handleSignUp which is a HTTP handler function that allows new users to sign up for an account.
// It starts by decoding a JSON object from the request body, this JSON object should contain the new user details
// such as the username, email, and password. If there is an error with decoding, it returns an error with a status code of "400 Bad Request".

// It then creates a salted and hashed version of the user's password using the bcrypt library,
// which is a widely-used library for password hashing. This is done to protect against password
// cracking attacks by making it computationally infeasible to recover the original plaintext password from
// the hashed version stored in the database.

// Then it creates a connection string to a PostgreSQL database using the "host", "port", "user", "password", and "dbname" variables,
// and opens a connection to the database using sql.Open. If there is an error connecting to the database,
//  it returns an error with a status code of "500 Internal Server Error".

// It then Inserts the new user into the database with an INSERT INTO statement, using the db.Exec method.
// It performs a check for the presence of email in the database and returns an error with a status code of "400 Bad Request

// It then generates a JSON web token(JWT) for the authenticated user with a signing method of HS256 and
// claims such as the "sub"(subject) and "exp"(expiration) set to the user's ID and 24 hours respectively.

// It signs the JWT with a secret key and sets the JWT as an HTTP cookie, with the token
// string as the value, it expires after 24 hours and path is set to '/'

func handleSignUp(w http.ResponseWriter, r *http.Request) {
	var newuser User
	if err := json.NewDecoder(r.Body).Decode(&newuser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("IN SIGN UP")
	// Validate the user data and hash the password
	// You can use the bcrypt library to hash the password before storing it in the database.
	// https://godoc.org/golang.org/x/crypto/bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newuser.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newuser.Password = string(hashedPassword)

	db = GetDb()

	// Insert the new user into the database
	_, err = db.Exec("INSERT INTO users (username, email, password) SELECT $1, $2, $3 WHERE NOT EXISTS (SELECT 1 FROM users WHERE email=$2)", newuser.Username, newuser.Email, newuser.Password)
	if err, ok := err.(*pq.Error); ok && err.Code == "23505" {
		http.Error(w, "Email already exists", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "Email already exists", http.StatusBadRequest)
		return
	}

	// Generate a JWT for the new user
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = newuser.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign the JWT with a secret key
	key := []byte("my-secret-key")
	tokenString, err := token.SignedString(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})

}

// This function handleLogin which is a HTTP handler function that allows existing users to log in to their account.

// It starts by decoding a JSON object from the request body, this JSON object should contain
// the login details such as the username and password. If there is an error with decoding,
//  it returns an error with a status code of "400 Bad Request".

// Then it creates a connection string to a PostgreSQL database using the "host", "port", "user", "password", and "dbname" variables
// , and opens a connection to the database using sql.Open. If there is an error connecting to the database, it returns an error
// with a status code of "500 Internal Server Error".

// It queries the user's data from the database using the db.QueryRow method, with the username and password as
//  the query parameters. It scans the row returned by the query for the password, which is then stored in the storedPassword variable.

// The function then compares the hashed password stored in the database with the plain text password provided
// by the user using the bcrypt library's CompareHashAndPassword method. If the hashed password doesn't match the provided password,
// then it returns an error with a status code of "401 Unauthorized".

// It then generates a JSON web token(JWT) for the authenticated user with a signing method of HS256 and
// claims such as the "sub"(subject) and "exp"(expiration) set to the user's ID and 24 hours respectively.

// It signs the JWT with a secret key and sets the JWT as an HTTP cookie, with the token string as the value,
//  it expires after 24 hours and path is set to '/'
// It then returns the token in a JSON object with a header of "201 Created"

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var newuser User
	if err := json.NewDecoder(r.Body).Decode(&newuser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db = GetDb()
	var storedPassword string
	err = db.QueryRow("SELECT password FROM users WHERE username=$1", newuser.Username).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Incorrect username or password", http.StatusUnauthorized)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Compare the hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(newuser.Password)); err != nil {
		http.Error(w, "Incorrect username or password", http.StatusUnauthorized)
		return
	}

	// Generate a JWT for the authenticated user
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = newuser.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign the JWT with a secret key
	key := []byte("my-secret-key")
	tokenString, err := token.SignedString(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the JWT as an HTTP cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Value:   tokenString,
		Expires: time.Now().Add(time.Hour * 24),
		Path:    "/",
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}

// This function handleAddCreditCard is handling the adding of a credit card for an authenticated user. It does this by:

// Extracting the JWT token from the request's Authorization header and validating it using the jwt library.
// Extracting the user ID from the claims of the parsed JWT and storing it in a variable.
// Decoding the JSON body of the request into a CreditCard struct
// Setting the user_id of the creditCard struct to the authenticated user's ID
// Connecting to the PostgreSQL database using the sql package and executing an insert statement
// to insert the credit card to the credit_cards table of the database same as previously
// If there are no errors, it sends the added card in JSON format as the response.
// If any error occurs, it sends an error response with appropriate HTTP status code.
func handleAddCreditCard(w http.ResponseWriter, r *http.Request) {
	// Get the authenticated user ID from the request context
	tokenString := r.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// verify the signing algorithm and secret key
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("my-secret-key"), nil
	})
	if err != nil {
		http.Error(w, "Invalid JWT", http.StatusUnauthorized)
		return
	}

	// extract the user ID from the JWT
	userID, ok := token.Claims.(jwt.MapClaims)["user_id"]
	if !ok {
		http.Error(w, "Invalid JWT", http.StatusUnauthorized)
		return
	}

	var creditCard CreditCard
	if err := json.NewDecoder(r.Body).Decode(&creditCard); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Print(userID)
	// Set the user ID of the credit card to the authenticated user's ID
	creditCard.UserID = 1

	// Connect to the database and insert the credit card
	db = GetDb()
	_, err = db.Exec("INSERT INTO credit_cards (user_id, card_number, expiry_month, expiry_year, cvv, name_on_card) VALUES ($1, $2, $3, $4, $5, $6)", creditCard.UserID, creditCard.CardNumber, creditCard.ExpiryMonth, creditCard.ExpiryYear, creditCard.CVV, creditCard.NameOnCard)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(creditCard)
}

// func createCharge(productID, token string) error {
// 	// Set your secret key: remember to change this to your live secret key in production
// 	// See your keys here: https://dashboard.stripe.com/account/apikeys
// 	stripe.Key = "sk_test_4eC39HqLyjWDarjtT1zdp7dc"

// 	// Get the product details
// 	product, err := getProduct(productID)
// 	if err != nil {
// 		return err
// 	}

// 	// Convert the price to cents
// 	priceInCents := product.Price * 100

// 	// Create the charge
// 	chargeParams := &stripe.ChargeParams{
// 		Amount:   stripe.Float64(priceInCents),
// 		Currency: stripe.String("usd"),
// 		Source:   &stripe.SourceParams{Token: token},
// 		Metadata: map[string]string{"product_id": productID},
// 	}
// 	chargeParams.SetDescription(product.Name)
// 	_, err = charge.New(chargeParams)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// The getProduct Method is used by handle Purchase
// it queries the database and get all products based on product id
func getProduct(productID string) (*Product, error) {
	// Connect to the database
	db = GetDb()
	// Execute the SELECT statement
	row := db.QueryRow("SELECT * FROM products WHERE product_id = $1", productID)

	// Scan the result into a product struct
	var product Product
	err = row.Scan(&product.ID, &product.Name, &product.Quantity, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Product not found")
		}
		return nil, err
	}

	return &product, nil
}

// The updateProduct Method is used by handle Purchase
// // it queries the database and update the products
func updateProduct(product *Product) error {
	// Connect to the database
	db = GetDb()
	// Execute the UPDATE statement
	_, err = db.Exec("UPDATE products SET name = $1, quantity = $2, price = $3 WHERE product_id = $4", product.Name, product.Quantity, product.Price, product.ID)
	if err != nil {
		return err
	}

	return nil
}

// This handlePurchase function appears to handle a request to purchase a product. It does this by:

// Getting the product ID, payment token and amount from the request body.
// It creates a charge by calling createCharge function which uses Stripe Library
//

// It get the product from database by calling getProduct function

// If the product is available, it deducts the quantity from the inventory
// and updates the product in the database by calling updateProduct function.
// At the end it returns a success response to the user.
func handlePurchase(w http.ResponseWriter, r *http.Request) {
	// Get the product ID and the payment details from the request body
	productID := r.FormValue("productID")
	if productID == "" {
		http.Error(w, "Missing productID parameter", http.StatusBadRequest)
		return
	}
	token := r.FormValue("token")
	if token == "" {
		http.Error(w, "Missing token parameter", http.StatusBadRequest)
		return
	}
	amount := r.FormValue("amount")
	if amount == "" {
		http.Error(w, "Missing amount parameter", http.StatusBadRequest)
		return
	}

	// Create a charge using the Stripe API
	// _, err := createCharge(productID, token)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// Get the product from the database
	product, err := getProduct(productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check the availability of the product
	if product.Quantity <= 0 {
		http.Error(w, "Product out of stock", http.StatusBadRequest)
		return
	}

	// Deduct the quantity from the inventory
	product.Quantity--

	// Update the product in the database
	err = updateProduct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	fmt.Fprintf(w, "Product purchased successfully")
}
