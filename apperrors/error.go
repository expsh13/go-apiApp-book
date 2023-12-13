package apperrors

type MyAppError struct {
	// TODO : 独自エラーに含めるフィールドの定義
	ErrCode        // レスポンスとログに表示するエラーコード
	Message string // レスポンスに表示するエラーメッセージ
	Err     error  // エラーチェーンのための内部エラー
}

// Error メソッドの役割は「そのエラーが fmt.Print 系関数等で出力されるときにどのような文字列になるか」
func (myErr *MyAppError) Error() string {
	// myErr.Err は型としては error インターフェース
	return myErr.Err.Error()
}

// errors.Unwrap関数によって内部のエラーを取り出せるようにするには、その独自エラー構造体に Unwrap メソッドが必要
func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}
