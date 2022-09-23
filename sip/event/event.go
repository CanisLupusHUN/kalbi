package event

import (
	"fmt"

	"github.com/CanisLupusHUN/kalbi/interfaces"
	logger "github.com/CanisLupusHUN/kalbi/logger"
	"github.com/CanisLupusHUN/kalbi/sip/message"
)

// SipEvent object that gets passed to the SipListener
type SipEvent struct {
	sipmsg *message.SipMsg
	tx     interfaces.Transaction
	lp     interfaces.ListeningPoint
}

// GetSipMessage returns message that created this event
func (se *SipEvent) GetSipMessage() *message.SipMsg {
	return se.sipmsg
}

// SetSipMessage sets message that created this event
func (se *SipEvent) SetSipMessage(msg *message.SipMsg) {
	se.sipmsg = msg
	logger.Debug(fmt.Sprintln(msg.String()))
}

// GetTransaction returns transaction related to the SIP message that created this event
func (se *SipEvent) GetTransaction() interfaces.Transaction {
	return se.tx
}

// SetTransaction sets transaction related to the SIP message that created this event
func (se *SipEvent) SetTransaction(tx interfaces.Transaction) {
	se.tx = tx
}

// SetListeningPoint gives ability to set interfaces.ListeningPoint
func (se *SipEvent) SetListeningPoint(lp interfaces.ListeningPoint) {
	se.lp = lp
}

// GetListeningPoint returns interfaces.ListeningPoint
func (se *SipEvent) GetListeningPoint() interfaces.ListeningPoint {
	return se.lp
}
