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

func easyjsonF3f49c6fDecodeSupermallModel(in *jlexer.Lexer, out *Goods) {
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
		case "list":
			if in.IsNull() {
				in.Skip()
				out.List = nil
			} else {
				in.Delim('[')
				if out.List == nil {
					if !in.IsDelim(']') {
						out.List = make([]Good, 0, 1)
					} else {
						out.List = []Good{}
					}
				} else {
					out.List = (out.List)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Good
					(v1).UnmarshalEasyJSON(in)
					out.List = append(out.List, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonF3f49c6fEncodeSupermallModel(out *jwriter.Writer, in Goods) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"list\":"
		out.RawString(prefix[1:])
		if in.List == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.List {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Goods) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF3f49c6fEncodeSupermallModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Goods) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF3f49c6fEncodeSupermallModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Goods) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF3f49c6fDecodeSupermallModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Goods) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF3f49c6fDecodeSupermallModel(l, v)
}
func easyjsonF3f49c6fDecodeSupermallModel1(in *jlexer.Lexer, out *Good) {
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
		case "id":
			out.ID = string(in.String())
		case "imgae":
			out.Image = string(in.String())
		case "link":
			out.Link = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "fav":
			out.Fav = uint32(in.Uint32())
		case "orgPrice":
			out.OrgPrice = uint16(in.Uint16())
		case "price":
			out.Price = uint16(in.Uint16())
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
func easyjsonF3f49c6fEncodeSupermallModel1(out *jwriter.Writer, in Good) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"imgae\":"
		out.RawString(prefix)
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
	{
		const prefix string = ",\"fav\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.Fav))
	}
	{
		const prefix string = ",\"orgPrice\":"
		out.RawString(prefix)
		out.Uint16(uint16(in.OrgPrice))
	}
	{
		const prefix string = ",\"price\":"
		out.RawString(prefix)
		out.Uint16(uint16(in.Price))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Good) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF3f49c6fEncodeSupermallModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Good) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF3f49c6fEncodeSupermallModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Good) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF3f49c6fDecodeSupermallModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Good) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF3f49c6fDecodeSupermallModel1(l, v)
}
