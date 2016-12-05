package vars

import (
	"github.com/jmartin82/mmock/definition"
	"github.com/jmartin82/mmock/persist"
	"github.com/jmartin82/mmock/utils"
	"github.com/jmartin82/mmock/vars/fakedata"
)

type FillerFactory interface {
	CreateRequestFiller(req *definition.Request) Filler
	CreateFakeFiller(Fake fakedata.DataFaker) Filler
	CreateStorageFiller(Engines *persist.PersistEngineBag) Filler
	CreatePersistFiller(Engines *persist.PersistEngineBag) Filler
}

type MockFillerFactory struct{}

func (mff MockFillerFactory) CreateRequestFiller(req *definition.Request) Filler {
	return RequestVars{Request: req, RegexHelper: utils.RegexHelper{}}
}

func (mff MockFillerFactory) CreateFakeFiller(fake fakedata.DataFaker) Filler {
	return FakeVars{Fake: fake}
}

func (mff MockFillerFactory) CreateStorageFiller(engines *persist.PersistEngineBag) Filler {
	return StorageVars{Engines: engines, RegexHelper: utils.RegexHelper{}}
}

func (mff MockFillerFactory) CreatePersistFiller(engines *persist.PersistEngineBag) Filler {
	return PersistVars{Engines: engines, RegexHelper: utils.RegexHelper{}}
}