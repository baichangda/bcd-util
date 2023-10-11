package parse

type Writeable interface {
	Write(_byteBuf *ByteBuf, _parentParseContext *ParseContext)
}
