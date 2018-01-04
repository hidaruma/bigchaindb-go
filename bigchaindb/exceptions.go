package bigchaindb

type ValueError struct {
	Msg string
}
func (ve *ValueError) Error() string {
	return ve.Msg
}
type BigchainDBError struct {
	Msg string
}

func (b *BigchainDBError) Error() string {
	return b.Msg
}
type CriticalDoubleSpend struct {
	Msg string
}
func (cds *CriticalDoubleSpend) Error() string {
	return cds.Msg
}
type CriticalDoubleInclusion struct {
	Msg string
}
func (cdi *CriticalDoubleInclusion) Error() string {
	return cdi.Msg
}
type CriticalDuplicateVote struct {
	Msg string
}

func (cdv *CriticalDuplicateVote) Error() string {
	return cdv.Msg
}