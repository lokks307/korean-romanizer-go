package korom_test

import (
	"testing"

	korom "github.com/lokks307/korean-romanizer-go"
)

func TestSimple(t *testing.T) {
	korom.RomanizeDebug = true

	if korom.Romanize("안녕하세요") != "annyeonghaseyo" {
		t.Error("TestSimple failed")
	}
}

func TestSpacedText(t *testing.T) {
	korom.RomanizeDebug = true

	if korom.Romanize("아이유 방탄소년단") != "aiyu bangtansonyeondan" {
		t.Error("TestSpacedText failed")
	}
}

func TestOnset_g_d_b(t *testing.T) {
	korom.RomanizeDebug = true

	if korom.Romanize("구미") != "gumi" {
		t.Error("TestOnset_g_d_b 'g' failed")
	}
	if korom.Romanize("영동") != "yeongdong" {
		t.Error("TestOnset_g_d_b 'd' failed")
	}
	if korom.Romanize("한밭") != "hanbat" {
		t.Error("TestOnset_g_d_b 'b' failed")
	}
}

func TestCoda_g_d_b(t *testing.T) {
	korom.RomanizeDebug = true

	if korom.Romanize("밝다") != "bakda" {
		t.Error("TestCoda_g_d_b 'bakda' failed")
	}
	if korom.Romanize("바닷가") != "badatga" {
		t.Error("TestCoda_g_d_b 'badatga' failed")
	}
	if korom.Romanize("없다") != "eopda" {
		t.Error("TestCoda_g_d_b 'eopda' failed")
	}
	if korom.Romanize("앞만") != "apman" {
		t.Error("TestCoda_g_d_b 'apman' failed")
	}
	if korom.Romanize("읊다") != "eupda" {
		t.Error("TestCoda_g_d_b 'eupda' failed")
	}
}

func Test_R_L(t *testing.T) {
	korom.RomanizeDebug = true

	if out := korom.Romanize("구리"); out != "guri" {
		t.Error("Test_R_L 'guri' failed, out=" + out)
	}
	if out := korom.Romanize("설악"); out != "seorak" {
		t.Error("Test_R_L 'seorak' failed, out=" + out)
	}
	if out := korom.Romanize("울릉"); out != "ulleung" {
		t.Error("Test_R_L 'ulleung' failed, out=" + out)
	}
	if out := korom.Romanize("대관령"); out != "daegwallyeong" {
		t.Error("Test_R_L 'daegwallyeong' failed, out=", out)
	}
}

func TestNextSyllableNullInitial(t *testing.T) {
	korom.RomanizeDebug = true

	if korom.Romanize("강약") != "gangyak" {
		t.Error("TestNextSyllableNullInitial 'gangyak' failed")
	}
	if korom.Romanize("강원") != "gangwon" {
		t.Error("TestNextSyllableNullInitial 'gangwon' failed")
	}
	if out := korom.Romanize("좋아하고"); out != "joahago" {
		t.Error("TestNextSyllableNullInitial 'joahago' failed, out=", out)
	}
	if out := korom.Romanize("좋은"); out != "joeun" {
		t.Error("TestNextSyllableNullInitial 'joeun' failed, out=", out)
	}
}

func TestDoubleConsWithNextNullInitial(t *testing.T) {
	korom.RomanizeDebug = true

	if out := korom.Romanize("했었어요"); out != "haesseosseoyo" {
		t.Error("TestDoubleConsWithNextNullInitial 'haesseosseoyo' failed, out=", out)
	}
	if out := korom.Romanize("없었다"); out != "eopsseotda" {
		t.Error("TestDoubleConsWithNextNullInitial 'eopseotda' failed, out=", out)
	}
	if out := korom.Romanize("앉아봐"); out != "anjabwa" {
		t.Error("TestDoubleConsWithNextNullInitial 'anjabwa' failed, out=", out)
	}
	if korom.Romanize("닭의") != "dalgui" {
		t.Error("TestDoubleConsWithNextNullInitial 'dalgui' failed")
	}
	if korom.Romanize("밟아") != "balba" {
		t.Error("TestDoubleConsWithNextNullInitial 'balba' failed")
	}
	if korom.Romanize("닮았네") != "dalmatne" {
		t.Error("TestDoubleConsWithNextNullInitial 'dalmatne' failed")
	}
	if out := korom.Romanize("삯을"); out != "saksseul" {
		t.Error("TestDoubleConsWithNextNullInitial 'sakseul' failed, err=", out)
	}
	if out := korom.Romanize("앓았다"); out != "aratda" {
		t.Error("TestDoubleConsWithNextNullInitial 'aratda' failed, err=", out)
	}
	if korom.Romanize("읊어 보거라") != "eulpeo bogeora" {
		t.Error("TestDoubleConsWithNextNullInitial 'eulpeo bogeora' failed")
	}
	if korom.Romanize("곬이") != "golssi" {
		t.Error("TestDoubleConsWithNextNullInitial 'golssi' failed")
	}
	if korom.Romanize("훑어보다") != "hulteoboda" {
		t.Error("TestDoubleConsWithNextNullInitial 'hulteoboda' failed")
	}
}

func TestDoubleConsWithNextNotNullInitial(t *testing.T) {
	korom.RomanizeDebug = true

	if korom.Romanize("앉고싶다") != "angosipda" {
		t.Error("TestDoubleConsWithNextNotNullInitial 'angosipda' failed")
	}
	if out := korom.Romanize("뚫리다"); out != "ttullida" {
		t.Error("TestDoubleConsWithNextNotNullInitial 'ttullida' failed, out=", out)
	}
	if korom.Romanize("칡뿌리") != "chikppuri" {
		t.Error("TestDoubleConsWithNextNotNullInitial 'chikppuri' failed")
	}
}

func TestCoda_H(t *testing.T) {
	korom.RomanizeDebug = true

	if korom.Romanize("않습니다") != "ansseupnida" {
		t.Error("TestCoda_H 'ansseupnida' failed")
	}
	if korom.Romanize("앓고") != "alko" {
		t.Error("TestCoda_H 'alko' failed")
	}
}
