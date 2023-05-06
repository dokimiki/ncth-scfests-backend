package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

type (
	SampleHandler struct {
		text string
	}
	SampleResponse struct {
		ID          string `json:"id"`
		CsvfileName string `json:"csvfilename"`
		CsvfileRow1 string `json:"csvfilerow1"`
		CsvfileRow2 string `json:"csvfilerow2"`
		Message     string `json:"message"`
		StatusCode  string `json:"returnCode"`
	}
	ErrorResponse struct {
		Message    string `json:"message"`
		StatusCode string `json:"returnCode"`
	}
)

func main() {
	http.HandleFunc("/", UsioOpen) // /が来たときに func Hello を実行する

	/* HandleFunc の第一引数のパス指定の最後には "/" を付けるべきでは? */
	/* (前回は "/end" と記述されていた
	https://github.com/dokimiki/ncth-scfests-backend/commit/8321e5efddf4b86cfcede430f056f894e8a241ac ) */
	http.HandleFunc("/webhook/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is Webhook."))
	})

	err := http.ListenAndServeTLS(":3939", "ncth-app.jp.pem", "ncth-app.jp.key", nil)
	//func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler)　error
	// errにエラーメッセージを格納する

	if err != nil { // エラーメッセージがあるとき出力
		fmt.Printf("ERROR : %s", err)
	}
}

func UsioOpen(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ファイル読み取り処理を開始します"))

	// ファイルをOpenする
	f, err := os.Open("logs/log.txt")
	// 読み取り時の例外処理
	if err != nil {
		w.Write([]byte("error\n"))
	}
	// 関数が終了した際に確実に閉じるようにする
	defer f.Close()

	// バイト型スライスの作成
	buf := make([]byte, 1024)
	for {
		// nはバイト数を示す
		n, err := f.Read(buf)
		// バイト数が0になることは、読み取り終了を示す
		if n == 0 {
			break
		}
		if err != nil {
			break
		}
		// バイト型スライスを文字列型に変換してファイルの内容を出力
		w.Write(buf[:n])
	}
	t := time.Now()
	w.Write([]byte("   now time is :" + t.Format("2006-01-02 03:04:05")))
}

// SampleHandlerの構造体にinterfaceのhttp.Handlerを設定して返す関数
// interfaceのhttp.HandlerにはServeHTTP関数が含まれており、後の処理ListenAndServe関数から呼び出される
func NewSampleHandler() http.Handler {
	return &SampleHandler{"処理は正常終了しました"}
}

// http.Handlerのinterfaceで定義されているServeHTTP関数を作成する。
// ServeHTTP関数はListenAndServe関数内で呼び出される
func (h *SampleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// http.RequestからCSVファイルを取得
	fileContent, header, err := r.FormFile("csvfile")
	if err != nil {
		ErrorResponseSend(w, "CSVAcceptError")
		return
	}
	// header.Filenameでcsvファイル名を取得できる
	csvfileName := header.Filename

	// CSVファイルの一行目を読み込む。fileContentはmultipart.File型
	reader := csv.NewReader(fileContent)
	line, err := reader.Read()
	if err != nil {
		ErrorResponseSend(w, "CSVReadError")
		return
	}

	// ステータスコード（成功）を設定
	statusCode := 200

	// httpResponseの内容を設定
	res := &SampleResponse{
		ID:          r.Form.Get("id"), // http.Requestからid情報を取得
		CsvfileName: csvfileName,
		CsvfileRow1: line[0], // csvファイルの1列名
		CsvfileRow2: line[1], // csvファイルの2列名
		Message:     h.text,
		StatusCode:  strconv.Itoa(statusCode),
	}
	// レスポンスヘッダーの設定
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// ステータスコードを設定
	w.WriteHeader(statusCode)

	// httpResponseの内容を書き込む
	buf, _ := json.MarshalIndent(res, "", "    ")
	_, _ = w.Write(buf)
}

// 処理エラーのときにResponseを返す関数
func ErrorResponseSend(w http.ResponseWriter, err string) {
	// ステータスコード（失敗）を設定
	statusCode := 500

	// httpResponseの内容を設定
	res := &ErrorResponse{
		Message:    err,
		StatusCode: strconv.Itoa(statusCode),
	}
	// レスポンスヘッダーの設定
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// ステータスコードを設定
	w.WriteHeader(statusCode)

	// httpResponseの内容を書き込む
	buf, _ := json.MarshalIndent(res, "", "    ")
	_, _ = w.Write(buf)
}
