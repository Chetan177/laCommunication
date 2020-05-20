package restserver

type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type makeCall struct {
	To   string `json:"to"`
	From string `json:"from"`
}
