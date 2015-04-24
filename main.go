package main

import (
	"fmt"
	"github.com/gorilla/mux"
	// "github.com/korasdor/data"
	"github.com/korasdor/handlers"
	"net/http"
)

func GetListenAddr() string {
	return fmt.Sprintf("%s:%s", "127.0.0.1", "8000")
}

func main() {
	fmt.Println("Server run...")

	// data.RandBalls = data.CreateBalls(data.RandBalls, 80)
	// data.Coefs = data.CreateCoefs()

	// handlers.RunHub()
	// handlers.StartLototron()

	router := mux.NewRouter()
	router.HandleFunc("/game/common/balance", handlers.GetBalanceHandler)
	router.HandleFunc("/game/{game_name}/rules", handlers.GetRulesHandler)
	router.HandleFunc("/game/{game_name}/{game_id}/bets", handlers.GetBetsHandler)
	router.HandleFunc("/game/{game_name}/bets", handlers.PostBetsHandler)
	router.HandleFunc("/game/{game_name}/bets/{bet_id}/{var_type}", handlers.PostBetChangeHandler)
	router.HandleFunc("/game/{game_name}/{game_id}/win", handlers.GetWinHandler)

	http.Handle("/", router)

	// http.HandleFunc("/randstream", handlers.RandHandler)
	// http.HandleFunc("/rules/keno", handlers.RulesKenoHandler)
	// http.HandleFunc("/game/keno/bets", handlers.GameKenoBetsHandler)
	// http.HandleFunc("/game/keno/win", handlers.GameKenoWinHandler)

	err := http.ListenAndServe(GetListenAddr(), nil)
	if err != nil {
		fmt.Println(err)
	}
	// checkError(err)
}
