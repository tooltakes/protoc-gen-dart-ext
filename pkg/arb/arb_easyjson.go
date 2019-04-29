// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package arb

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

func easyjsonAf4d4e0fDecodeGithubComEmpirefoxProtocGenDartExtPkgArb(in *jlexer.Lexer, out *Arb) {
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
		case "@@last_modified":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.LastModified).UnmarshalJSON(data))
			}
		case "@@context":
			out.Context = string(in.String())
		case "@@locale":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.Locale).UnmarshalText(data))
			}
		case "@@author":
			out.Author = string(in.String())
		case "@@x-entity":
			out.Entity = string(in.String())
		case "@@x-imports":
			if in.IsNull() {
				in.Skip()
				out.Imports = nil
			} else {
				in.Delim('[')
				if out.Imports == nil {
					if !in.IsDelim(']') {
						out.Imports = make([]*ArbImport, 0, 8)
					} else {
						out.Imports = []*ArbImport{}
					}
				} else {
					out.Imports = (out.Imports)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *ArbImport
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(ArbImport)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Imports = append(out.Imports, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "@@x-nils":
			if in.IsNull() {
				in.Skip()
				out.Nils = nil
			} else {
				in.Delim('[')
				if out.Nils == nil {
					if !in.IsDelim(']') {
						out.Nils = make([]string, 0, 4)
					} else {
						out.Nils = []string{}
					}
				} else {
					out.Nils = (out.Nils)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.Nils = append(out.Nils, v2)
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
func easyjsonAf4d4e0fEncodeGithubComEmpirefoxProtocGenDartExtPkgArb(out *jwriter.Writer, in Arb) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		const prefix string = ",\"@@last_modified\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.LastModified).MarshalJSON())
	}
	if in.Context != "" {
		const prefix string = ",\"@@context\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Context))
	}
	if true {
		const prefix string = ",\"@@locale\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.RawText((in.Locale).MarshalText())
	}
	if in.Author != "" {
		const prefix string = ",\"@@author\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Author))
	}
	if in.Entity != "" {
		const prefix string = ",\"@@x-entity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Entity))
	}
	if len(in.Imports) != 0 {
		const prefix string = ",\"@@x-imports\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v3, v4 := range in.Imports {
				if v3 > 0 {
					out.RawByte(',')
				}
				if v4 == nil {
					out.RawString("null")
				} else {
					(*v4).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.Nils) != 0 {
		const prefix string = ",\"@@x-nils\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.Nils {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Arb) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAf4d4e0fEncodeGithubComEmpirefoxProtocGenDartExtPkgArb(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Arb) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAf4d4e0fDecodeGithubComEmpirefoxProtocGenDartExtPkgArb(l, v)
}