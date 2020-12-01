package game

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	baris = "=============================================================================="
)

// IGame is interface describe what this game does
type IGame interface {
	SetDaduFactory(factory IDaduFactory)
	SetPlayerFactory(factory IPlayerFactory)
	Play(totalPlayer int, totalDadu int)
}

// IGameFactory is interface to create game
type IGameFactory interface {
	CreateGame() (mygame IGame)
}

type myGame struct {
	Players       []IPlayer
	ActivePlayers []IPlayer
	daduFactory   IDaduFactory
	playerFactory IPlayerFactory
}

func (mygame *myGame) SetDaduFactory(factory IDaduFactory) {

	mygame.daduFactory = factory
}

func (mygame *myGame) SetPlayerFactory(factory IPlayerFactory) {

	mygame.playerFactory = factory
}

func (mygame *myGame) lemparDaduPlayers() {
	wg := new(sync.WaitGroup)
	wg.Add(len(mygame.ActivePlayers))

	for _, player := range mygame.ActivePlayers {

		go func(myPlayer IPlayer) {
			myPlayer.LemparDadu()
			wg.Done()
		}(player)
	}

	wg.Wait()
}

func (mygame *myGame) evaluasiDaduPlayers() {
	wg := new(sync.WaitGroup)
	wg.Add(len(mygame.ActivePlayers))

	jobDoneMap := make(map[int]chan int)
	dadu1Map := make(map[int]DaduChanel)
	for _, player := range mygame.ActivePlayers {
		jobDoneMap[player.GetPlayerID()] = make(chan int)
		dadu1Map[player.GetPlayerID()] = make(DaduChanel)
	}

	for _, player := range mygame.ActivePlayers {

		go func(myPlayer IPlayer) {
			dadu1, _ := myPlayer.EvaluasiDadu()

			dadu1Map[myPlayer.GetPlayerID()] <- dadu1
			close(dadu1Map[myPlayer.GetPlayerID()])

			jobDoneMap[myPlayer.GetPlayerID()] <- 1
			close(jobDoneMap[myPlayer.GetPlayerID()])
			//wg.Done()
		}(player)

		go func(myPlayer IPlayer) {
			dadu1 := <-dadu1Map[myPlayer.GetPlayerID()]
			//fmt.Println(<-jobDoneMap[myPlayer.GetNextPlayer().GetPlayerID()])
			<-jobDoneMap[myPlayer.GetNextPlayer().GetPlayerID()] //wait until job done
			myPlayer.GetNextPlayer().TambahDaduBaru(dadu1...)

			wg.Done()
		}(player)
	}

	wg.Wait()
}

func (mygame *myGame) createDadu(totalDadu int) (slcDadu []IDadu) {
	arrDadu := make([]IDadu, 0)

	for i := 0; i < totalDadu; i++ {
		arrDadu = append(arrDadu, mygame.daduFactory.CreateDadu())
	}

	return arrDadu
}

func (mygame *myGame) displayGame() {

	for _, player := range mygame.Players {

		arrStrDadu := make([]string, 0)
		for _, dadu := range player.LihatDadu() {
			arrStrDadu = append(arrStrDadu, fmt.Sprintf("%d", dadu.Lihat()))
		}

		daduListStr := strings.Join(arrStrDadu, ",")
		if daduListStr == "" {
			daduListStr = "_(Berhenti bermain karena tak memiliki dadu)"
		}

		fmt.Printf("\tPemain #%d (%d): %s\n", player.GetPlayerID(), player.LihatJumlahPoin(),
			daduListStr)
	}

}

func (mygame *myGame) Play(totalPlayer int, totalDadu int) {

	fmt.Printf("Pemain = %d, Dadu = %d\n", totalPlayer, totalDadu)
	fmt.Println(baris)

	//create players
	mygame.Players = make([]IPlayer, 0)
	mygame.ActivePlayers = make([]IPlayer, 0)

	for i := 1; i <= totalPlayer; i++ {

		arrDadu := mygame.createDadu(totalDadu)
		player := mygame.playerFactory.CreatePlayer(i)
		player.SetDadu(arrDadu...)

		mygame.Players = append(mygame.Players, player)
		mygame.ActivePlayers = append(mygame.ActivePlayers, player)
		//fmt.Println(player.GetPlayerID())
	}

	giliran := 1
	for len(mygame.ActivePlayers) > 1 {

		// set next player for every player
		i := 1
		for _, player := range mygame.ActivePlayers {
			if i < len(mygame.ActivePlayers) {
				player.SetNextPlayer(mygame.ActivePlayers[i])
			} else {
				player.SetNextPlayer(mygame.ActivePlayers[0])
			}

			//fmt.Println(player.GetPlayerID(), " >> ", player.GetNextPlayer().GetPlayerID())
			i++
		}

		mygame.lemparDaduPlayers()

		// display dadu and point for every player
		fmt.Printf("Giliran %d lempar dadu\n", giliran)
		mygame.displayGame()

		mygame.evaluasiDaduPlayers()

		fmt.Println("Setelah evaluasi :")

		// display dadu and point for every player
		mygame.displayGame()
		fmt.Println(baris)

		// reset active player
		activePlayers := make([]IPlayer, 0)
		for _, player := range mygame.ActivePlayers {
			if player.IsActive() {
				activePlayers = append(activePlayers, player)
			}
		}

		mygame.ActivePlayers = activePlayers

		giliran++
	}

	// search for winner
	if len(mygame.Players) < 1 {
		return
	}

	winner := mygame.Players[0]
	winners := []IPlayer{winner}
	for _, player := range mygame.Players[1:] {
		if winner.LihatJumlahPoin() == player.LihatJumlahPoin() {
			winners = append(winners, player)
		} else if winner.LihatJumlahPoin() < player.LihatJumlahPoin() {
			winner = player
			winners = []IPlayer{winner}
		}
	}

	if len(winners) > 1 {

		fmt.Print("winners : ")
		for i, player := range winners {

			if i > 0 {
				fmt.Print(", ")
			}

			fmt.Printf("pemain #%d", player.GetPlayerID())
		}

	} else {
		fmt.Printf("winner adalah pemain #%d", winner.GetPlayerID())
	}

	fmt.Println()
}

func createGame() (mygame IGame) {
	return new(myGame)
}

type gameFactory struct{}

func (factory *gameFactory) CreateGame() (mygame IGame) {
	return createGame()
}

// CreateGameFactory is function to create game factory
func CreateGameFactory() (factory IGameFactory) {
	return new(gameFactory)
}
