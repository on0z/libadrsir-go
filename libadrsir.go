package libadrsir

import (
	"strconv"

	"github.com/pkg/errors"
)

const (
	// ADRSIRのI2Cアドレス
	ADDR = byte(0x52)

	// コマンド群
	R1_memo_no_write  = byte(0x15)
	R2_data_num_read  = byte(0x25)
	R3_data_read      = byte(0x35)
	W2_data_num_write = byte(0x29)
	W3_data_write     = byte(0x39)
	T1_trans_start    = byte(0x59)
)

type Bus interface {
	ReadReg(byte, []byte) error
	WriteReg(byte, []byte) error
}

type AdrsirAPI interface {
	Get(index int) string
	Send(irCommandStr string) error
}

type adrsir struct {
	bus Bus
}

func NewADRSIR(device Bus) AdrsirAPI {
	return &adrsir{
		bus: device,
	}
}

func (a *adrsir) Get(index int) string {
	return ""
}

func (a *adrsir) Send(irCommandStr string) error {

	// 文字列からbyte列を生成
	var irCommand []byte
	// 2文字づつ処理する
	for i := 0; i < len(irCommandStr); i += 2 {
		// 2文字を1バイトに変換する
		aByte, err := strconv.ParseUint(irCommandStr[i:i+2], 16, 8)
		if err != nil {
			return errors.WithStack(err)
		}
		// 配列に追加
		irCommand = append(irCommand, uint8(aByte))
	}

	// 送信予定のデータの長さを長さを書き込む
	irCommandLength := uint16(len(irCommand))
	sendData := []byte{byte(irCommandLength >> 8), byte(irCommandLength & 0xff)}

	err := a.bus.WriteReg(W2_data_num_write, sendData)
	if err != nil {
		return errors.WithStack(err)
	}

	// irCommandを4バイトづつ書き込む
	for i := 0; i < int(irCommandLength); i += 4 {
		a.bus.WriteReg(W3_data_write, irCommand[i:i+4])
	}

	// transferコマンドを書き込む
	err = a.bus.WriteReg(T1_trans_start, []byte{0})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
