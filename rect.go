package images

import "image"

func RectMove(rect image.Rectangle, point image.Point) image.Rectangle {
	rect.Min = rect.Min.Add(point)
	rect.Max = rect.Max.Add(point)
	return rect
}

func RectResize(rect image.Rectangle, width, height int) image.Rectangle {
	rect.Max.X = rect.Min.X + width
	rect.Max.Y = rect.Min.Y + height
	return rect
}
