package RedisClient

import (
	"encoding/json"
	"note-gin/model"
)

func AddBook(book model.MyBook) {
	client := RedisInit()
	bookStr, _ := json.Marshal(book)
	client.HSetNX("my_book", string(book.ID), bookStr)
}

func GetAllBook() []model.MyBook {
	client := RedisInit()
	m := client.HGetAll("my_book").Val()
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

func DeleteBook(id int) {
	client := RedisInit()
	client.HDel("my_book", string(id))
}
