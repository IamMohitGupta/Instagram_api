// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Id       int    `json:"User_id"`
	Name     string `json:"User_name"`
	Email    string `json:"User_email"`
	Password string `json:"Password"`
}
type Post struct {
	Id               int    `json:"Post_id"`
	Caption          string `json:"Post_caption"`
	Image_url        string `json:"Post_url"`
	Posted_Timestamp string `json:"Post_time"`
}

// func main() {
// 	// fmt.Println("Hello")
// 	resp, err := http.Get("http://webcode.me")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(string(body))
// 	user := User{Id: "User#1", Name: "James", Email: "mgupta060402@gmail.com", Password: "Hello"}
// 	post := Posts{Id: "Post_1", Caption: "First Post", Image_url: "http://heelo.jpg", Posted_Timestamp: "12:00:00 UTC"}
// 	users, err := json.MarshalIndent(user, "", "")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(users))
// 	posts, err := json.MarshalIndent(post, "", "")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(posts))

// }

type Users []User

func allUsers(w http.ResponseWriter, r *http.Request) {
	Users := Users{
		User{Id: 1, Name: "James", Email: "mgupta060402@gmail.com", Password: "Hello"},
		User{Id: 2, Name: "James", Email: "mgupta060402@gmail.com", Password: "Hello"},
	}
	fmt.Print("All the users are listed")
	json.NewEncoder(w).Encode(Users)
}

type Posts []Post

func allPosts(w http.ResponseWriter, r *http.Request) {
	Posts := Posts{
		Post{Id: 1, Caption: "First Post", Image_url: "http://hello.jpg", Posted_Timestamp: "12:00:00 UTC"},
	}
	fmt.Print("All the posts are listed")
	json.NewEncoder(w).Encode(Posts)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the homepageeeee!")
}

// func handleRequests() {
// 	http.HandleFunc("/hello", homePage)
// 	http.HandleFunc("/users", allUsers)
// 	http.HandleFunc("/posts", allPosts)
// 	log.Fatal(http.ListenAndServe(":8000", nil))
// }
// func createEvent(w http.ResponseWriter, r *http.Request) {
// 	var newEvent event
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
// 	}

// 	json.Unmarshal(reqBody, &newEvent)
// 	events = append(events, newEvent)
// 	w.WriteHeader(http.StatusCreated)

// 	json.NewEncoder(w).Encode(newEvent)
// }
// func getOneEvent(w http.ResponseWriter, r *http.Request) {
// 	eventID := mux.Vars(r)["id"]

// 	for _, singleEvent := range events {
// 		if singleEvent.ID == eventID {
// 			json.NewEncoder(w).Encode(singleEvent)
// 		}
// 	}
// }

// func main() {

// 	handleRequests()

// }
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/users", allUsers)
	router.HandleFunc("/posts", allPosts)
	log.Fatal(http.ListenAndServe(":8000", router))
}

// http.Get("/foo", homeLink)

// http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
// })

// log.Fatal(http.ListenAndServe(":8080", nil))

// resp, err := http.Get("http://example.com/")
// if err != nil {
// 	// handle error
// }
// defer resp.Body.Close()
// body, err := io.ReadAll(resp.Body)
