package game

import (
	"sync"
)

// DaduChanel is data type for chanel containing dadu collection
type DaduChanel chan []IDadu

// IPlayer is interface describing what a player can do in the game
type IPlayer interface {
	LemparDadu()
	LihatDadu() (dadu []IDadu)
	EvaluasiDadu() (dadu1 []IDadu, dadu6 []IDadu)
	LihatJumlahPoin() (jmlPoint int)
	SetDadu(slcDadu ...IDadu)
	TambahDaduBaru(daduBaru ...IDadu)
	IsActive() (status bool)

	GetPlayerID() (id int)
	SetNextPlayer(nextPlayer IPlayer)
	GetNextPlayer() (nextPlayer IPlayer)
}

// IPlayerFactory is interface to create player
type IPlayerFactory interface {
	CreatePlayer(id int) (player IPlayer)
}

type cPlayer struct {
	ID         int
	NextPlayer IPlayer
	SlcDadu    []IDadu
	Point      int
	//ChanStoreDadu DaduChanel
}

func (player *cPlayer) LemparDadu() {
	wg := new(sync.WaitGroup)
	wg.Add(len(player.SlcDadu))

	for _, dadu := range player.SlcDadu {

		go func(prcDadu IDadu) {
			prcDadu.Acak()
			wg.Done()
		}(dadu)

	}

	wg.Wait()
}

func (player *cPlayer) LihatDadu() (dadu []IDadu) {
	return player.SlcDadu
}

func (player *cPlayer) EvaluasiDadu() (dadu1 []IDadu, dadu6 []IDadu) {

	newSlcDadu := make([]IDadu, 0)
	for _, dadu := range player.SlcDadu {

		if dadu.Lihat() == 6 {
			player.Point++ //tambah poin
			dadu6 = append(dadu6, dadu)
			continue
		}

		if dadu.Lihat() == 1 {
			dadu1 = append(dadu1, dadu)
			continue
		}

		newSlcDadu = append(newSlcDadu, dadu)
	}

	player.SetDadu(newSlcDadu...) //reset dadu

	//player.NextPlayer.storeDaduBaru(dadu1...) //store dadu1 to next player

	//daduBaru := <-player.ChanStoreDadu // get new dadu (dadu1) from this channel

	//player.SlcDadu = append(player.SlcDadu, daduBaru...) // update dadu
	return
}

func (player *cPlayer) LihatJumlahPoin() (jmlPoint int) {

	return player.Point
}

func (player *cPlayer) SetDadu(slcDadu ...IDadu) {
	player.SlcDadu = slcDadu
}

func (player *cPlayer) TambahDaduBaru(daduBaru ...IDadu) {

	player.SlcDadu = append(player.SlcDadu, daduBaru...)
}

func (player *cPlayer) IsActive() (status bool) {

	return len(player.SlcDadu) > 0
}

func (player *cPlayer) GetPlayerID() (id int) {
	return player.ID
}

func (player *cPlayer) SetNextPlayer(nextPlayer IPlayer) {
	player.NextPlayer = nextPlayer
}

func (player *cPlayer) GetNextPlayer() (nextPlayer IPlayer) {
	return player.NextPlayer
}

func createPlayer(id int) (player IPlayer) {
	newPlayer := new(cPlayer)
	newPlayer.ID = id
	//newPlayer.ChanStoreDadu = make(DaduChanel)

	return newPlayer
}

type cPlayerFactory struct{}

func (factory *cPlayerFactory) CreatePlayer(id int) (player IPlayer) {
	return createPlayer(id)
}

// CreatePlayerFactory is function to create player factory
func CreatePlayerFactory() (factory IPlayerFactory) {
	return new(cPlayerFactory)
}
