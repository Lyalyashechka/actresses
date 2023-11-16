package context

import "context"

type Fields map[string]interface{}

func (f Fields) Extend(f2 Fields) {
	for k, v := range f2 {
		f[k] = v
	}
}

func (f Fields) Copy() Fields {
	ret := make(Fields, len(f))
	for k, v := range f {
		ret[k] = v
	}
	return ret
}

type fieldsCtxKey struct{}

func SetCtxFields(ctx context.Context, fields Fields) context.Context {
	return context.WithValue(ctx, fieldsCtxKey{}, fields)
}

func GetCtxFields(ctx context.Context) Fields {
	fields := getCtxFields(ctx)
	if fields == nil {
		return Fields{}
	}

	return fields.Copy()
}

func getCtxFields(ctx context.Context) Fields {
	fieldsVal := ctx.Value(fieldsCtxKey{})
	if fieldsVal == nil {
		return nil
	}

	return fieldsVal.(Fields)
}

func AddCtxFields(ctx context.Context, fields Fields) context.Context {
	ctxFields := getCtxFields(ctx)
	if ctxFields == nil {
		return SetCtxFields(ctx, fields)
	}

	for k, v := range fields {
		ctxFields[k] = v
	}

	return ctx
}
