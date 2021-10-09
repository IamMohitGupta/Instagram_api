// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Id       string `json:"User_id"`
	Name     string `json:"User_name"`
	Email    string `json:"User_email"`
	Password string `json:"Password"`
}
type Post struct {
	Id               string `json:"Post_id"`
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
		User{Id: "User#1", Name: "James", Email: "mgupta060402@gmail.com", Password: "Hello"},
		User{Id: "User#2", Name: "James", Email: "mgupta060402@gmail.com", Password: "Hello"},
	}
	fmt.Print("All the users are listed")
	json.NewEncoder(w).Encode(Users)
}

type Posts []Post

func allPosts(w http.ResponseWriter, r *http.Request) {
	Posts := Posts{
		Post{Id: "Post_1", Caption: "First Post", Image_url: "http://hello.jpg", Posted_Timestamp: "12:00:00 UTC"},
	}
	fmt.Print("All the posts are listed")
	json.NewEncoder(w).Encode(Posts)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the homepageeeee!")
}
func handleRequests() {
	http.HandleFunc("/hello", homePage)
	http.HandleFunc("/users", allUsers)
	http.HandleFunc("/posts", allPosts)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
func main() {

	handleRequests()

}
