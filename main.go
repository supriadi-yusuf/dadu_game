package main

import (
	"dadu_game/game"
	"flag"
)

func main() {

	player := flag.Int("pemain", 0, "menentuka jumlah pemain")
	dadu := flag.Int("dadu", 1, "menentukan jumlah dadu")

	flag.Parse()

	daduFactory := game.CreateDaduFactory()
	playerFactory := game.CreatePlayerFactory()

	gameFactory := game.CreateGameFactory()
	mygame := gameFactory.CreateGame()
	mygame.SetDaduFactory(daduFactory)
	mygame.SetPlayerFactory(playerFactory)
	mygame.Play(*player, *dadu)
}
