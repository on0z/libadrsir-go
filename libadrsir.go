package libadrsir

const (
	// ADRSIRのI2Cアドレス
	ADDR = 0x52

	// コマンド群
	R1_memo_no_write  = 0x15
	R2_data_num_read  = 0x25
	R3_data_read      = 0x35
	W2_data_num_write = 0x29
	W3_data_write     = 0x39
	T1_trans_start    = 0x59
)

type bus interface {
	ReadReg(byte, []byte) error
	WriteReg(byte, []byte) error
}

type AdrsirAPI interface {
	Get(index int) string
	Send(irData string) error
}

type adrsir struct {
	bus bus
}

func NewADRSIR(device bus) AdrsirAPI {
	return &adrsir{
		bus: device,
	}
}

func (a *adrsir) Get(index int) string {
	return ""
}

func (a *adrsir) Send(irData string) error {
	return nil
}
