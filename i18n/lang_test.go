package i18n

import "testing"

func TestI18n_Trans(t *testing.T) {
	correctTrans := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("translate test en", func(t *testing.T) {
		i := I18n{}
		i.SetLang("en")
		i.SetTransFilePath("../config/i18n")
		i.SetTransFileType("yaml")
		got := i.Trans("test.name").ToStr()
		want := "Test"
		correctTrans(t, got, want)
	})

	t.Run("translate test zh", func(t *testing.T) {
		i := I18n{}
		i.SetLang("zh")
		i.SetTransFilePath("../config/i18n")
		i.SetTransFileType("yaml")
		got := i.Trans("test.name").ToStr()
		want := "测试"
		correctTrans(t, got, want)
	})

}
