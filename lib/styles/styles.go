package styles

import "image/color"

func RGB(rgb uint64) color.RGBA {
	return color.RGBA{
		R: uint8((rgb & (0xFF << (8 * 2))) >> (8 * 2)),
		G: uint8((rgb & (0xFF << (8 * 1))) >> (8 * 1)),
		B: uint8((rgb & (0xFF << (8 * 0))) >> (8 * 0)),
		A: 0xFF,
	}
}

var (
	TransparentColor = color.RGBA{}
	// 主要
	PrimaryColor = RGB(0x9dd793)
	// サブ
	SecondaryColor = RGB(0x9dd793)
	// 地のテキスト
	BaseTextColor = RGB(0x000000)
	// 背景
	BackgroundColor = RGB(0x000000)

	// アイドル
	ButtonIdleColor = RGB(0xdff4ff)
)
