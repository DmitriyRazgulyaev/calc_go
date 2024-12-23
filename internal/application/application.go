package application

import (
	"encoding/json"
	"fmt"
	"github.com/DmitriyRazgulyaev/calc_go/pkg/calculation"
	"net/http"
	"os"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

type Request struct {
	Expression string `json:"expression"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
		return
	}

	result, err := calculation.Calc(request.Expression)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return

	}

	json.NewEncoder(w).Encode(map[string]string{"result": fmt.Sprintf("%f", result)})
}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate/", CalcHandler)
	return http.ListenAndServe(":"+a.config.Addr, nil)
}
