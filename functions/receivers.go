package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Name    string `json:"name"`
	Surname string `json:""`
	Likes   []Like
}

type Like struct {
	ID   string
	Date time.Time
}

func (u *User) Like() {
	if u.Likes == nil {
		u.Likes = make([]Like, 0)
	}

	l := Like{
		ID:   "id",
		Date: time.Time{},
	}

	u.Likes = append(u.Likes, l)
}

func (u User) LikeCount() int {
	return len(u.Likes)
}

func timer(c <-chan time.Time) {
	for {
		select {
		case <-c:
			return
		default:
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	timer(time.After(2 * time.Second))
	u1 := User{
		Name:    "go",
		Surname: "turkiye",
	}

	u1.Like()
	u1.Like()
	u1.Like()
	u1.Like()

	arr, _ := json.Marshal(u1)
	fmt.Println(string(arr))

	fmt.Println(u1.LikeCount())
}
