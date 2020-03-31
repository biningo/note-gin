package RedisClient

import (
	"encoding/json"
	"note-gin/model"
)

func AddBook(book model.MyBook) {
	client := RedisInit()
	defer client.Close()


	bookStr, _ := json.Marshal(book)
	client.HSet("my_book", string(book.ID), bookStr)
}

func GetAllBook() []model.MyBook {
	client := RedisInit()
	defer client.Close()
	m := client.HGetAll("my_book").Val()

	//
	IsExist := client.Exists("my_book").Val()
	//
	if IsExist == 0 {
		books := model.MyBook{}.GetAll()
		for _, v := range books {
			bStr, _ := json.Marshal(v)
			client.HSetNX("my_book", string(v.ID), bStr)
		}
		return books
	} else {
		books := make([]model.MyBook, len(m))
		book := model.MyBook{}

		c := 0
		for _, v := range m {
			json.Unmarshal([]byte(v), &book)
			books[c] = book
			c++
		}
		return books
	}

}

func DeleteBook(id int) {
	client := RedisInit()
	defer client.Close()
	client.HDel("my_book", string(id))
}
