package images

import (
	"github.com/Comdex/imgo"
	"testing"
)

func TestSimilarity(t *testing.T) {
	img, err := imgo.Read("C:\\Users\\Administrator\\Desktop\\邮件箱.png")
	if err != nil {
		t.Fatal(err)
	}
	img2, err := imgo.Read("C:\\Users\\Administrator\\Desktop\\邮件箱3.png")
	if err != nil {
		t.Fatal(err)
	}
	diff := Similarity(img, img2) / 10
	t.Log(diff)
}
