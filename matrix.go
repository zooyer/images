package images

import (
	"image"
	"image/color"
)

type Matrix [][][]uint8

func (m Matrix) Width() int {
	return len(m[0])
}

func (m Matrix) Height() int {
	return len(m)
}

func (m Matrix) Pixel() int {
	return m.Width() * m.Height()
}

func (m Matrix) Clone() Matrix {
	var matrix = NewMatrix(m.Width(), m.Height())
	for i := 0; i < matrix.Height(); i++ {
		for j := 0; j < matrix.Width(); j++ {
			copy(matrix[i][j], m[i][j])
		}
	}
	return matrix
}

func (m Matrix) Process(handler PixHandler) Matrix {
	copy(m, Process(m, handler).(Matrix))
	return m
}

func (m Matrix) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Matrix) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.Width(), m.Height())
}

func (m Matrix) At(x, y int) color.Color {
	pix := m[y][x]
	return color.RGBA{
		R: pix[0],
		G: pix[1],
		B: pix[2],
		A: pix[3],
	}
}

func (m Matrix) SubImage(rect image.Rectangle) image.Image {
	return m
}

func NewMatrix(width, height int) Matrix {
	var matrix = make(Matrix, height)
	for i := range matrix {
		matrix[i] = make([][]uint8, width)
		for j := range matrix[i] {
			matrix[i][j] = make([]uint8, 4)
		}
	}
	return matrix
}
