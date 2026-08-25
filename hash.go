package main
import "golang.org/x/crypto/bcrypt"

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func register(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Could not process password", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("INSERT INTO users (username, password_hash) VALUES (?, ?)", creds.Username, string(hash))
	if err != nil {
		http.Error(w, "Username already taken", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
