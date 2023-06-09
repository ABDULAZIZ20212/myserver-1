package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Product struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Calories    int     `json:"calories"`
}

func getHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	return hostname
}

func health(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("x-COMING_FROM_HOST", getHostName())
	resp := make(map[string]string)
	resp["status"] = "Healthy-abdulaziz-change"
	resp["host"] = getHostName()
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	fmt.Fprintf(w, string(jsonResp))
}
func about(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("x-COMING_FROM_HOST", getHostName())
	resp := make(map[string]string)
	resp["name"] = "abdulaziz"
	resp["course"] = "CPIT-632"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	fmt.Fprintf(w, string(jsonResp))
}

func fatal(w http.ResponseWriter, req *http.Request) {
// 	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("x-COMING_FROM_HOST", getHostName())
	resp := make(map[string]string)
	resp["status"] = "aborted"
	resp["host"] = getHostName()
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	fmt.Fprintf(w, string(jsonResp))
	log.Fatal("Internal server error")
}

func menu(w http.ResponseWriter, req *http.Request) {
	var products []Product = []Product{
		Product{Name: "Baklava", Description: "Traditional sweet layered pastry dessert filled with chopped nuts, and sweetened with our signature syrup.", Price: 2.50, Calories: 560},
		Product{Name: "Sourdough Baguette", Description: "Crisp and light, with a crackly brown crust, baguettes ",
			Price: 2.99, Calories: 300},
		Product{Name: "Cheese Roll", Description: "Super soft and fluffy grilled cheese roll.",
			Price: 3.39, Calories: 440},
		Product{Name: "Three Cheese Artisan Bread", Description: "Speciality artisan bread combines the rich flavor of three classic Italian cheeses: Asiago, Parmesan and Romano.",
			Price: 5.50, Calories: 620},
		Product{Name: "Cheese Croissant", Description: "Classic cheese croissant.",
			Price: 2.50, Calories: 480},
		Product{Name: "Cranberry Bagel", Description: "Special Cranberry Bagels made with real cranberries for a delicious taste you can't resist.", Price: 1.99, Calories: 380},
	}
	if req.URL.Path == "/us/menu" {
		products = append(products, Product{Name: "Broccoli Cheddar Soup", Description: "Broccoli Cheddar Soup served in a sourdough bread bowl.",
			Price: 8.50, Calories: 820})
	}
	productsJson, _ := json.Marshal(&products)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("x-COMING_FROM_HOST", getHostName())
	fmt.Fprintf(w, string(productsJson))
}

func main() {
	host := "0.0.0.0"
	port := "80"
	if os.Getenv("HOST") != "" {
		host = os.Getenv("HOST")
	}
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	http.HandleFunc("/health", health)
	http.HandleFunc("/about", about)
	http.HandleFunc("/sa/menu", menu)
	http.HandleFunc("/us/menu", menu)
	http.HandleFunc("/fatal", fatal)
	fmt.Println("Listening on http://" + host + ":" + port)
	err := http.ListenAndServe(host+":"+port, nil)
	if err != nil {
		log.Fatalf("Error starting the web server. Err: %s", err)
	}
}
