package cryptoconditions

type Exception struct {
	Msg string
}
func (e *Exception) Error() string{
	return e.Msg
}

type ParsingError Exception

type PrefixError Exception

type UnsupportedError Exception

type ValidationError Exception

type UnknownEncodingError Exception

type MissingDataError Exception

type ASN1EncodeError Exception

type ASN1DecodeError Exception
