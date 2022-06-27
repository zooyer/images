package images

import (
	"image"
	"image/color"
	"math"
)

type PixHandler func(img image.Image, point image.Point, color color.Color) color.Color

// 浮点型运算灰度值（效率较慢）
//func Gray() PixelHandler {
//	return func(matrix Matrix, point image.Point, rgba color.RGBA) color.RGBA {
//		gray := uint8(float64(rgba.R)*0.299 + float64(rgba.G)*0.587 + float64(rgba.B)*0.114)
//		return color.RGBA{
//			R: gray,
//			G: gray,
//			B: gray,
//			A: rgba.A,
//		}
//	}
//}

// 整数型运算灰度值（有除法运算，效率不高）
//func Gray2() PixelHandler {
//	return func(matrix Matrix, point image.Point, rgba color.RGBA) color.RGBA {
//		gray := uint8((uint32(rgba.R)*299 + uint32(rgba.G)*587 + uint32(rgba.B)*114) / 1000)
//		return color.RGBA{
//			R: gray,
//			G: gray,
//			B: gray,
//			A: rgba.A,
//		}
//	}
//}

// 位移运算，效率高
func Gray() PixHandler {
	return func(img image.Image, point image.Point, c color.Color) color.Color {
		r, g, b, a := c.RGBA()
		gray := uint8((r*19595 + g*38469 + b*7472) >> 16)
		return color.RGBA{
			R: gray,
			G: gray,
			B: gray,
			A: uint8(a),
		}
	}
}

func GrayPS() PixHandler {
	return func(img image.Image, point image.Point, c color.Color) color.Color {
		r, g, b, a := c.RGBA()
		gray := uint8(math.Pow(math.Pow(float64(r), 2.2)*0.2973+math.Pow(float64(g), 2.2)*0.6274+math.Pow(float64(b), 2.2)*0.0753, 1/2.2))
		return color.RGBA{
			R: gray,
			G: gray,
			B: gray,
			A: uint8(a),
		}
	}
}

func GrayAvg() PixHandler {
	return func(img image.Image, point image.Point, c color.Color) color.Color {
		r, g, b, a := c.RGBA()
		avg := uint8(float64(r+g+b) / 3.0)
		return color.RGBA{
			R: avg,
			G: avg,
			B: avg,
			A: uint8(a),
		}
	}
}

func SunsetEffect(ratio float64) PixHandler {
	if ratio == 0 {
		ratio = 0.7
	}
	return func(img image.Image, point image.Point, c color.Color) color.Color {
		r, g, b, a := c.RGBA()
		return color.RGBA{
			R: uint8(float64(r) * ratio),
			G: uint8(float64(g) * ratio),
			B: uint8(b),
			A: uint8(a),
		}
	}
}

func Process(img image.Image, handler PixHandler) image.Image {
	matrix := FromImage(img)

	for i := 0; i < matrix.Height(); i++ {
		for j := 0; j < matrix.Width(); j++ {
			rgba := color.RGBA{
				R: matrix[i][j][0],
				G: matrix[i][j][1],
				B: matrix[i][j][2],
				A: matrix[i][j][3],
			}

			point := image.Point{
				X: j,
				Y: i,
			}

			r, g, b, a := handler(matrix, point, rgba).RGBA()
			matrix[i][j][0] = uint8(r)
			matrix[i][j][1] = uint8(g)
			matrix[i][j][2] = uint8(b)
			matrix[i][j][3] = uint8(a)
		}
	}

	return matrix
}

// return 0 - 1000
func similarity(p1, p2 []uint8) int {
	var diff int
	diff += abs(int(p1[0]) - int(p2[0]))
	diff += abs(int(p1[1]) - int(p2[1]))
	diff += abs(int(p1[2]) - int(p2[2]))
	diff += abs(int(p1[3]) - int(p2[3]))
	return 1000 - diff*1000/(math.MaxUint8*4)
}

// tolerance: 0 - 255
// return 0 - 1000
func Similarity(img1, img2 image.Image) int {
	m1 := FromImage(img1).Process(Gray())
	m2 := FromImage(img2).Process(Gray())
	width := min(m1.Width(), m2.Width())
	height := min(m1.Height(), m2.Height())
	var s int
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			s += similarity(m1[i][j], m2[i][j])
		}
	}
	return s / (width * height)
}

func InImage(img1, img2 image.Image) image.Point {
	var sim int
	var point image.Point

	type SubImager interface {
		SubImage(image.Rectangle) image.Image
	}

	width := img1.Bounds().Dx() - img2.Bounds().Dx()
	height := img1.Bounds().Dy() - img2.Bounds().Dy()

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			sub := img1.(SubImager).SubImage(image.Rect(j, i, img2.Bounds().Dx(), img2.Bounds().Dy()))
			if diff := Similarity(sub, img2); diff > sim {
				sim = diff
				point.X = j
				point.Y = i
			}
		}
	}

	return point
}
