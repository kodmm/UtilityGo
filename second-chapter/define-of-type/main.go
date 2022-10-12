package main

func main() {
	
	// 例: Stock Keeping Unit(在庫管理単位)を体系化する
	// 以下のような低レイヤーのコードを書いては行けない。桁数が異なった場合や半角英数字以外を含むようになった場合。
	//* skuCD, _ := r.URL.Query()["sku_code"]
	//* itemCD, sizeCD, colorCD := skuCD[0:5], skuCD[5:7], skuCD[7:9]

	// SKUコード
type SKUCode string

param, _ := r.URL.QUERY()["sku_code"]
skuCD := SKUCode(param)

// チェック処理
type (c SKUCode) Invalid() bool {
	// 桁数や利用可能文字のチェックを行う
}

func (c SKUCode) ItemCD() string {
	return skuCD[0:5]
}

func (c SKUCode) SizeCD() string {
	return skuCD[5:7]
}

func (c SKUCode) ColorCD() string{
	return skuCD[7:9]
}

	if skuCD.Invalid() {
		// 異常系のハンドリング
	}

	itemCD, sizeCD, colorCD, := skuCD.ItemCD(), skuCD.SizeCD(), skuCD.ColorCD()
}

type Season int

const(
	Peak Season = iota + 1 // 繁忙期
	Normal // 通常期
	Off // 閑散期
)

func (s Season) Price(price float64) bool {
	if s == Peak {
		return s + 200
	}
	return s
}

type SensorData struct {
	SensorType string
	ModelID string
	Value float32
}

// ! 悪い例
func ReadValue(r SensorData) float32 {
	if r.SensorType == "Fahrenheit" { // 華氏の場合は摂氏に変換
		return (r.Value * 9 / 5) + 32
	}
	return r.Value
}

// * 良い例
func (r SensorData) ReadValue() float32 {
	if r.SensorType == "Fahrenheit" { // 華氏の場合は摂氏に変換
		return (r.Value * 9 / 5) + 32
	}
	return r.Value
}

// * 構造体はレシーバを定義する方がテストもしやすく再利用性も良い

var i int
type ErrorCode int
var e ErrorCode

i = int(e)
e = i
