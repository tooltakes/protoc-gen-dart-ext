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

func easyjson5011676aDecodeGithubComEmpirefoxProtocGenDartExtPkgArb(in *jlexer.Lexer, out *ArbImport) {
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
			out.Id = string(in.String())
		case "import":
			out.Ref = string(in.String())
		case "assets":
			if in.IsNull() {
				in.Skip()
				out.Assets = nil
			} else {
				in.Delim('[')
				if out.Assets == nil {
					if !in.IsDelim(']') {
						out.Assets = make([]string, 0, 4)
					} else {
						out.Assets = []string{}
					}
				} else {
					out.Assets = (out.Assets)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Assets = append(out.Assets, v1)
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
func easyjson5011676aEncodeGithubComEmpirefoxProtocGenDartExtPkgArb(out *jwriter.Writer, in ArbImport) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Id != "" {
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Id))
	}
	if in.Ref != "" {
		const prefix string = ",\"import\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Ref))
	}
	if len(in.Assets) != 0 {
		const prefix string = ",\"assets\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.Assets {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ArbImport) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5011676aEncodeGithubComEmpirefoxProtocGenDartExtPkgArb(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ArbImport) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5011676aEncodeGithubComEmpirefoxProtocGenDartExtPkgArb(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ArbImport) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5011676aDecodeGithubComEmpirefoxProtocGenDartExtPkgArb(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ArbImport) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5011676aDecodeGithubComEmpirefoxProtocGenDartExtPkgArb(l, v)
}
