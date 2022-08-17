package message

// NewResponse creates new SIP Response
func NewResponse(request SipReq, via []SipVia, to SipTo, from SipFrom, callID SipVal, maxfor SipVal) *SipMsg {
	r := new(SipMsg)
	r.Req = request
	for i, _ := range via {
		r.Via = append(r.Via, via[i])
	}
	r.To = to
	r.From = from
	r.CallID = callID
	r.MaxFwd = maxfor
	return r
}
