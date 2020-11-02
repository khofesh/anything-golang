package main

import "testing"

func TestDecode(t *testing.T) {
	post, err := Decode("post.json")
	if err != nil {
		t.Error(err)
	}

	if post.ID != 1 {
		t.Error("wrong id, was expecting 1 but got", post.ID)
	}

	if post.Content != "hello world!" {
		t.Error("wrong content, was expecting 'hello world!' but got", post.Content)
	}
}

func TestEncode(t *testing.T) {
	t.Skip("Skipping encoding for now")
}
