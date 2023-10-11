package parse

type ParseContext struct {
	ParentContext *ParseContext
	//指针类型
	Instance any
}

func ToParseContext(instance any, parentContext *ParseContext) *ParseContext {
	return &ParseContext{
		ParentContext: parentContext,
		Instance:      instance,
	}
}
