// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package present

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonB9c4285aDecodeGithubComOmegaatt36TinygoProjectPresent(in *jlexer.Lexer, out *JsonBME280) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "err":
			out.Err = string(in.String())
		case "temperature":
			out.Temperature = int32(in.Int32())
		case "pressure":
			out.Pressure = int32(in.Int32())
		case "humidity":
			out.Humidity = int32(in.Int32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonB9c4285aEncodeGithubComOmegaatt36TinygoProjectPresent(out *jwriter.Writer, in JsonBME280) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Err != "" {
		const prefix string = ",\"err\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Err))
	}
	{
		const prefix string = ",\"temperature\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Temperature))
	}
	{
		const prefix string = ",\"pressure\":"
		out.RawString(prefix)
		out.Int32(int32(in.Pressure))
	}
	{
		const prefix string = ",\"humidity\":"
		out.RawString(prefix)
		out.Int32(int32(in.Humidity))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v JsonBME280) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB9c4285aEncodeGithubComOmegaatt36TinygoProjectPresent(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v JsonBME280) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB9c4285aEncodeGithubComOmegaatt36TinygoProjectPresent(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *JsonBME280) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB9c4285aDecodeGithubComOmegaatt36TinygoProjectPresent(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *JsonBME280) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB9c4285aDecodeGithubComOmegaatt36TinygoProjectPresent(l, v)
}
