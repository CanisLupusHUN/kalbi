package message

// NewResponse creates new SIP Response
func NewResponse(request SipReq, via []SipVia, to SipTo, from SipFrom, callID SipVal, maxfor SipVal) *SipMsg {
	r := new(SipMsg)
	r.Req = request
	r.Via = via
	r.To = SipTo{
		UriType:  from.UriType,
		Name:     from.Name,
		User:     from.User,
		Host:     from.Host,
		Port:     from.Port,
		Tag:      from.Tag,
		UserType: from.UserType,
		Src:      from.Src,
	}
	r.From = SipFrom{
		UriType:  to.UriType,
		Name:     to.Name,
		User:     to.User,
		Host:     to.Host,
		Port:     to.Port,
		Tag:      to.Tag,
		UserType: to.UserType,
		Src:      to.Src,
	}
	r.CallID = callID
	r.MaxFwd = maxfor
	return r
}
