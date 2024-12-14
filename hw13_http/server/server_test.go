package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandlerGet(t *testing.T) {
	testCases := []struct {
		desc      string
		r         *http.Request
		resStatus int
		resHeader string // "Content-Type"
		resBody   User
	}{
		{
			desc: "Positive",
			r: &http.Request{
				Method: http.MethodGet,
			},
			resStatus: http.StatusOK,
			resHeader: "application/json",
			resBody: User{
				Name:    "Iva",
				Surname: "Inov",
				Age:     30,
			},
		},
		{
			desc: "Negative_WrongMethod",
			r: &http.Request{
				Method: http.MethodPost,
			},
			resStatus: http.StatusMethodNotAllowed,
			resHeader: "",
			resBody:   User{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			handleGet(w, tC.r)

			gotBody := User{}
			json.Unmarshal(w.Body.Bytes(), &gotBody)
			gotHeader := w.Header().Get("Content-Type")

			require.Equal(t, tC.resStatus, w.Code)
			require.Equal(t, tC.resHeader, gotHeader)
			require.Equal(t, tC.resBody, gotBody)
		})
	}
}

var BodyToSend = User{
	Name:    "Lenn",
	Surname: "Lashev",
	Age:     54,
}

func (b *User) CreateReaderCloser() io.ReadCloser {
	BodyToSendJSON, _ := json.Marshal(b)
	BodyToSendReader := strings.NewReader(string(BodyToSendJSON))
	BodyToSendReadCloser := io.NopCloser(BodyToSendReader)
	return BodyToSendReadCloser
}

func TestHandlerPost(t *testing.T) {
	testCases := []struct {
		desc      string
		r         *http.Request
		resStatus int
		resHeader string // "Content-Type"
		resBody   User
	}{
		{
			desc: "Positive",
			r: &http.Request{
				Method: http.MethodPost,
				Body:   BodyToSend.CreateReaderCloser(),
			},
			resStatus: http.StatusCreated,
			resHeader: "application/json",
			resBody: User{
				Name:    "Lenn",
				Surname: "Lashev",
				Age:     54,
			},
		},
		{
			desc: "Negative_WrongMethod",
			r: &http.Request{
				Method: http.MethodGet,
			},
			resStatus: http.StatusMethodNotAllowed,
			resHeader: "",
			resBody:   User{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			handlePost(w, tC.r)

			gotBody := User{}
			json.Unmarshal(w.Body.Bytes(), &gotBody)
			gotHeader := w.Header().Get("Content-Type")

			require.Equal(t, tC.resStatus, w.Code)
			require.Equal(t, tC.resHeader, gotHeader)
			require.Equal(t, tC.resBody, gotBody)
		})
	}
}
