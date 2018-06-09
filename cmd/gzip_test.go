package cmd

import (
	"testing"
)

func TestGunzip(t *testing.T) {
	t.Run("gunzip "+TEST_DATE, func(t *testing.T) {
		expect := TEST_DATE
		actual := gunzipCommand("1f8b08008ed01b5b00034b492c4905007a379eaa04000000")

		Check(expect, actual, t)
	})
	t.Run("gunzip "+TEST_SEQ_AWK, func(t *testing.T) {
		expect := TEST_SEQ_AWK
		actual := gunzipCommand(`1f8b0800bdd11b5b00032b4e2d543036a8492ccf5650f70b5235b2b535a8
2e28cacc2ba9550700844045151b000000`)
		Check(expect, actual, t)
	})
}

func TestGzip(t *testing.T) {
	t.Run("gzip "+TEST_DATE, func(t *testing.T) {
		expect := "00034b492c4905007a379eaa04000000"
		actual := gzipCommand(TEST_DATE)[16:]

		Check(expect, actual, t)
	})

	t.Run("gzip "+TEST_SEQ_AWK, func(t *testing.T) {
		expect := `00032b4e2d543036a8492ccf5650f70b5235b2b535a8
2e28cacc2ba9550700844045151b000000`
		actual := gzipCommand(TEST_SEQ_AWK)[16:]

		Check(expect, actual, t)
	})
}
