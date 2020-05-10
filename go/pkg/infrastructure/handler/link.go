package handler

import (
	"homepage/pkg/infrastructure/server/response"
	"net/http"
)

func LinkHandler(w http.ResponseWriter, r *http.Request) {
	header := &response.HeaderData{
		IsLogin:  false,
		PageType: "",
	}

	labLinks := []*link{
		&link{
			Label: "鈴木 研究室",
			URL:   "http://www.suzuki.sist.chukyo-u.ac.jp/",
		},
		&link{
			Label: "濱川 研究室",
			URL:   "http://hamakawalab.sist.chukyo-u.ac.jp/",
		},
		&link{
			Label: "MDLAB(目加田・道満 研究室)",
			URL:   "https://md.sist.chukyo-u.ac.jp/index.html",
		},
		&link{
			Label: "鬼頭 研究室",
			URL:   "http://kitolab.sist.chukyo-u.ac.jp/",
		},
		&link{
			Label: "オープンメディアラボ（宮崎・山田・中 研究室）",
			URL:   "https://www.om.sist.chukyo-u.ac.jp/",
		},
	}
	techLinks := []*link{
		&link{
			Label:       "Flutter",
			Description: "Dart言語を用いた Android／iOSアプリ開発のフレームワーク",
			URL:         "https://flutter.dev/",
		},
		&link{
			Label:       "Docker",
			Description: "軽量なコンテナ型の仮想環境を提供するオープンソースソフトウェア",
			URL:         "https://www.docker.com/",
		},
		&link{
			Label:       "Electron",
			Description: "クロスプラットホームなデスクトップアプリを開発できるフレームワーク",
			URL:         "https://www.electronjs.org/",
		},
		&link{
			Label:       "ARCore",
			Description: "Googleが提供するARフレームワーク",
			URL:         "https://developers.google.com/ar/",
		},
	}
	body := struct {
		LabLinks  []*link
		TechLinks []*link
	}{
		LabLinks:  labLinks,
		TechLinks: techLinks,
	}
	response.Success(w, "link/index.html", header, &body)
}

type link struct {
	Label       string
	Description string
	URL         string
}
