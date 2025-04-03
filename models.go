package main

type PlayRequestBody struct {
	Name  string `json:"name"`
	Guess uint   `json:"guess"`
}

type NameRequestBody struct {
	Name string `json:"name"`
}

type ResponseBody struct {
	Response string `json:"response"`
}

type Player struct {
	Name      string
	GuessLeft uint
}

type GuessResult struct {
	Success bool
	Message string
}
