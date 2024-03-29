package libadrsir

import (
	"errors"
	"fmt"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mock "github.com/on0z/libadrsir-go/mock"
)

func TestAdrsirGet(t *testing.T) {

}

func TestAdrsirSend(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type input struct {
		irDataStr string
	}

	type expect struct {
		err string
	}

	type mockBusTx struct {
		Err       error
		CallCount int
	}

	cases := []struct {
		name string
		input
		mockBusTx
		expect
	}{
		{
			name: "正常系",
			input: input{
				irDataStr: "00002800D00029003900160038001600120016001300160012001700120016001300160012001700380016001200170012001600130016001200170012001600130016003800160013001600380016001300160012001700120016001300160012001700120016001300160012001700120016001300160012001700120016001300160012001700120016003900160012001600390016003800160012001600390016003800160011004205",
			},
			mockBusTx: mockBusTx{
				Err:       nil,
				CallCount: 45,
			},
			expect: expect{
				err: "",
			},
		},
		{
			name: "異常系: 不適切な16進数文字列",
			input: input{
				irDataStr: "ZZ",
			},
			mockBusTx: mockBusTx{
				Err:       nil,
				CallCount: 0,
			},
			expect: expect{
				err: "strconv.ParseUint: parsing \"ZZ\": invalid syntax",
			},
		},
		{
			name: "異常系: i2c書き込み失敗",
			input: input{
				irDataStr: "00002800D00029003900160038001600120016001300160012001700120016001300160012001700380016001200170012001600130016001200170012001600130016003800160013001600380016001300160012001700120016001300160012001700120016001300160012001700120016001300160012001700120016001300160012001700120016003900160012001600390016003800160012001600390016003800160011004205",
			},
			mockBusTx: mockBusTx{
				Err:       errors.New("i2c Error"),
				CallCount: 1,
			},
			expect: expect{
				err: "i2c Error",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockBus := mock.NewMockBus(ctrl)
			mockBus.EXPECT().
				Tx(gomock.Any(), gomock.Any()).
				Times(c.mockBusTx.CallCount).
				Return(c.mockBusTx.Err)

			lib := adrsir{
				bus:          mockBus,
				waitDuration: 200 * time.Millisecond,
			}

			errCh := make(chan error)
			go func() {
				errCh <- lib.Send(c.input.irDataStr)
			}()
			err := <-errCh
			if c.expect.err != "" {
				assert.EqualError(t, err, c.expect.err)
				fmt.Println(err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
