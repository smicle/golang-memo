package colors

import (
	"testing"

	. "github.com/r7kamura/gospel"
)

func TestAttach(t *testing.T) {
	Describe(t, "ANSIエスケープシーケンスが正しく付与されているかのテスト", func() {
		Context("出力文字の変更テスト", func() {
			It("Boldの確認", func() {
				Expect(Attach("Bold", "test")).To(Equal, "\033[1mtest\033[0m")
			})
			It("Italicの確認", func() {
				Expect(Attach("Italic", "test")).To(Equal, "\033[3mtest\033[0m")
			})
			It("Underlineの確認", func() {
				Expect(Attach("Underline", "test")).To(Equal, "\033[4mtest\033[0m")
			})
		})

		Context("出力色の変更テスト", func() {
			It("Blackの確認", func() {
				Expect(Attach("Black", "test")).To(Equal, "\033[30mtest\033[0m")
			})
			It("Redの確認", func() {
				Expect(Attach("Red", "test")).To(Equal, "\033[31mtest\033[0m")
			})
			It("Greenの確認", func() {
				Expect(Attach("Green", "test")).To(Equal, "\033[32mtest\033[0m")
			})
			It("Yellowの確認", func() {
				Expect(Attach("Yellow", "test")).To(Equal, "\033[33mtest\033[0m")
			})
			It("Blueの確認", func() {
				Expect(Attach("Blue", "test")).To(Equal, "\033[34mtest\033[0m")
			})
			It("Magentaの確認", func() {
				Expect(Attach("Magenta", "test")).To(Equal, "\033[35mtest\033[0m")
			})
			It("Cyanの確認", func() {
				Expect(Attach("Cyan", "test")).To(Equal, "\033[36mtest\033[0m")
			})
			It("Whiteの確認", func() {
				Expect(Attach("White", "test")).To(Equal, "\033[37mtest\033[0m")
			})
		})

		Context("出力色バリエーションの変更テスト", func() {
			It("Bgの確認", func() {
				Expect(Attach("BgBlack", "test")).To(Equal, "\033[40mtest\033[0m")
			})
			It("Brightの確認", func() {
				Expect(Attach("BrightBlack", "test")).To(Equal, "\033[90mtest\033[0m")
			})
			It("BgBrightの確認", func() {
				Expect(Attach("BgBrightBlack", "test")).To(Equal, "\033[100mtest\033[0m")
			})
		})
	})
}
