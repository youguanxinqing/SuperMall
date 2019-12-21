// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

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

func easyjson7573338fDecodeSupermallModel(in *jlexer.Lexer, out *Img) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "path":
			out.Image = string(in.String())
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
func easyjson7573338fEncodeSupermallModel(out *jwriter.Writer, in Img) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"path\":"
		out.RawString(prefix[1:])
		out.String(string(in.Image))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Img) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7573338fEncodeSupermallModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Img) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7573338fEncodeSupermallModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Img) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7573338fDecodeSupermallModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Img) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7573338fDecodeSupermallModel(l, v)
}
func easyjson7573338fDecodeSupermallModel1(in *jlexer.Lexer, out *HyperImg) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "image":
			out.Image = string(in.String())
		case "link":
			out.Link = string(in.String())
		case "title":
			out.Title = string(in.String())
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
func easyjson7573338fEncodeSupermallModel1(out *jwriter.Writer, in HyperImg) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"image\":"
		out.RawString(prefix[1:])
		out.String(string(in.Image))
	}
	{
		const prefix string = ",\"link\":"
		out.RawString(prefix)
		out.String(string(in.Link))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HyperImg) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7573338fEncodeSupermallModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HyperImg) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7573338fEncodeSupermallModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HyperImg) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7573338fDecodeSupermallModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HyperImg) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7573338fDecodeSupermallModel1(l, v)
}
