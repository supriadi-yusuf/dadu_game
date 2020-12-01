package game

import "math/rand"

// IDadu is interface describing what dadu can do
type IDadu interface {
	Acak()             //acak dadu
	Lihat() (sisi int) //lihat permukaan dadu
}

// IDaduFactory is interface to create dadu
type IDaduFactory interface {
	CreateDadu() (dadu IDadu)
}

type cDadu struct {
	sisi int
}

func (dadu *cDadu) Acak() {

	dadu.sisi = rand.Intn(6) + 1
}

func (dadu *cDadu) Lihat() (sisi int) {
	return dadu.sisi
}

func createDadu() (dadu IDadu) {
	newDadu := new(cDadu)
	newDadu.sisi = 1
	return newDadu
}

type cDaduFactory struct{}

func (factory *cDaduFactory) CreateDadu() (dadu IDadu) {
	return createDadu()
}

// CreateDaduFactory is function to create dadu factory
func CreateDaduFactory() (factory IDaduFactory) {
	return new(cDaduFactory)
}
