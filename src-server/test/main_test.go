package test

import (
	_ "go.beyondstorage.io/services/fs/v4"
	"go.beyondstorage.io/v5/services"
	"os"
	"testing"
)

func TestStorage(t *testing.T) {
	Storager, err := services.NewStoragerFromString("fs://C:/Projects/go/SparkWayCore/home")
	if err != nil {
		t.Log(err)
	} else {
		t.Log(Storager.String())
		file, e := os.Open("C:\\Projects\\go\\SparkWayCore\\home\\upload\\images\\cd\\cd42fe1e51e3511767f4926d3ea8c9b0.jpg")
		fileinfo, e := file.Stat()
		if e != nil {
			t.Log(e)
		} else {
			n, e2 := Storager.Write("dd/a/b/c.jpg", file, fileinfo.Size())
			if e2 != nil {
				t.Log(e2)
			} else {
				t.Log("n = ", n)
			}
		}

	}
}
