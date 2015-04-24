package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/korasdor/data"
	"github.com/korasdor/services"
	"net/http"
)

func GetBalanceHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	response := services.GetBalanceMessage()
	fmt.Fprint(w, response)

}

func GetRulesHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(req)
	gameName := vars["game_name"]

	var response string
	if gameName == data.GameName {
		response = services.GetRulesMessage()
		fmt.Fprint(w, response)
	} else {
		response := fmt.Sprintf("%s not supproted", gameName)
		fmt.Fprint(w, response)
	}

	// response = services.GetBalanceMessage()
	// fmt.Fprint(w, response)
}

func GetBetsHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(req)
	gameName := vars["game_name"]
	gameId := vars["game_id"]

	var response string
	if gameName == data.GameName {
		response = services.GetBetsMessage(gameId)
		fmt.Fprint(w, response)
	} else {
		response := fmt.Sprintf("%s not supproted", gameName)
		fmt.Fprint(w, response)
	}
}

func PostBetsHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(req)
	gameName := vars["game_name"]

	var response string
	if gameName == data.GameName {
		if req.Method == "POST" {
			response = services.PostBetsMessage()
			fmt.Fprint(w, response)
		} else {
			response = fmt.Sprintf("%s method not supproted", req.Method)
			fmt.Fprint(w, response)
		}
	} else {
		response := fmt.Sprintf("%s not supproted", gameName)
		fmt.Fprint(w, response)
	}

}

func PostBetChangeHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(req)
	gameName := vars["game_name"]
	// gameId := vars["game_id"]
	varType := vars["var_type"]

	var response string
	if gameName == data.GameName {
		if req.Method == "POST" {

			if varType == "content" {
			} else if varType == "bet_sum" {
				response = services.GetBalanceMessage()
				fmt.Fprint(w, response)
			}

		} else {
			response = fmt.Sprintf("%s method not supproted", req.Method)
			fmt.Fprint(w, response)
		}
	} else {
		response := fmt.Sprintf("%s not supproted", gameName)
		fmt.Fprint(w, response)
	}
}

func GetWinHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(req)
	gameName := vars["game_name"]
	gameId := vars["game_id"]

	var response string
	if gameName == data.GameName {
		response = services.GetWinMessage(gameId)
		fmt.Fprint(w, response)
	} else {
		response = fmt.Sprintf("%s not supproted", gameName)
		fmt.Fprint(w, response)
	}
}
