package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
)

type Mocks struct {
	MDConverter      *MockMDConverter
	DB               *MockDB
	Storage          *MockStorage
	HTMLEditor       *MockHTMLEditor
	HTMLMediaFetcher *MockHTMLMediaFetcher
	HTMLTOCExtractor *MockHTMLTOCExtractor
}

func NewMockDepends(t *testing.T) (*Mocks, *Impl, func()) {
	finishFuncs := []func(){}
	mocks := Mocks{}

	ctrlMDConverter := gomock.NewController(t)
	finishFuncs = append(finishFuncs, ctrlMDConverter.Finish)
	mocks.MDConverter = NewMockMDConverter(ctrlMDConverter)

	ctrlDB := gomock.NewController(t)
	finishFuncs = append(finishFuncs, ctrlDB.Finish)
	mocks.DB = NewMockDB(ctrlDB)

	ctrlStorage := gomock.NewController(t)
	finishFuncs = append(finishFuncs, ctrlStorage.Finish)
	mocks.Storage = NewMockStorage(ctrlStorage)

	ctrlHTMLEditor := gomock.NewController(t)
	finishFuncs = append(finishFuncs, ctrlHTMLEditor.Finish)
	mocks.HTMLEditor = NewMockHTMLEditor(ctrlHTMLEditor)

	ctrlHTMLMediaFetcher := gomock.NewController(t)
	finishFuncs = append(finishFuncs, ctrlHTMLMediaFetcher.Finish)
	mocks.HTMLMediaFetcher = NewMockHTMLMediaFetcher(ctrlHTMLMediaFetcher)

	ctrlHTMLTOCExtractor := gomock.NewController(t)
	finishFuncs = append(finishFuncs, ctrlHTMLTOCExtractor.Finish)
	mocks.HTMLTOCExtractor = NewMockHTMLTOCExtractor(ctrlHTMLTOCExtractor)

	closeFunc := func() {
		for _, fc := range finishFuncs {
			fc()
		}
	}
	impl := Impl{
		MDConverter:      mocks.MDConverter,
		DB:               mocks.DB,
		Storage:          mocks.Storage,
		HTMLEditor:       mocks.HTMLEditor,
		HTMLMediaFetcher: mocks.HTMLMediaFetcher,
		HTMLTOCExtractor: mocks.HTMLTOCExtractor,
	}
	return &mocks, &impl, closeFunc
}
