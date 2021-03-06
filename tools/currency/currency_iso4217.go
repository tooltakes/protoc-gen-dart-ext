package main

import (
	"encoding/xml"
	"flag"
	"log"
	"net/http"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/Masterminds/sprig"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/arb"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/genshared"
	"github.com/relvacode/iso8601"
	"golang.org/x/text/language"
)

const (
	emptyZero = "XXX"

	listUrl = "https://www.currency-iso.org/dam/downloads/lists/list_one.xml"
)

var (
	entity = &genshared.VersionEntity{
		LowerCamelCase: "currency",
		Version:        1,
	}
)

const protoTplStr = genshared.UnitsProtoHead + `
import "protos/l10n/l10n.proto";

// ISO 4217, publish date: {{ .Pblshd }}
enum {{ EntityV }} {
  option (l10n.enumArb).ignore = true;
	XXX = 0;
{{- range .CcyNtry }}
	{{ .Ccy }} = {{ .CcyNbr }}; // {{ .CcyNm }}
{{- end }}
}
`

const dartTplStr = genshared.DartHead + `
import './currency_format.dart';
import './units.l10n.dart';

abstract class _Valuer {
	String of(UnitsLocalization l);
}

class _XXX implements _Valuer {
  const _XXX();
  String of(UnitsLocalization l) => '';
}

{{- range .CcyNtry }}
  class _{{ .Ccy }} implements _Valuer {
    const _{{ .Ccy }}();
    String of(UnitsLocalization l) => l.{{ entity }}{{ .Ccy | powerCamel }};
  }
{{- end }}

class {{ EntityV }} {
	final String ccy;
	final int ccyNbr;
	final _Valuer _v;
	const {{ EntityV }}._(this.ccy, this.ccyNbr, this._v);
	String Function(dynamic) get format => getCurrencyFormat(ccy).format;
	String Function(dynamic) get formatSimple => getSimpleCurrencyFormat(ccy).format;
	String l10n(UnitsLocalization l10n) => l10n == null ? ccy : _v.of(l10n) ?? ccy;

	static const XXX = const {{ EntityV }}._('', 0, const _XXX());
	{{- range .CcyNtry }}
	static const {{ .Ccy }} = const {{ EntityV }}._('{{ .Ccy }}', {{ .CcyNbr }}, const _{{ .Ccy }}());
	{{- end }}
}
`

const arbTplStr = `{{ toPrettyJson .Arb }}`

var (
	entityFuncs = genshared.VersionEntityFuncs(entity)

	sharedFuncs = genshared.JoinFuncs(
		sprig.HermeticTxtFuncMap(),
		dart.Funcs,
		genshared.Funcs,
		genshared.VersionEntityFuncs(entity),
	)

	protoTpl = template.Must(template.New("proto").Funcs(entityFuncs).Parse(protoTplStr))
	dartTpl  = template.Must(template.New("dart").Funcs(sharedFuncs).Parse(dartTplStr))
	arbTpl   = template.Must(template.New("arb").Funcs(sharedFuncs).Parse(arbTplStr))
)

func main() {
	rs := genshared.NewGoTplRenderers(protoTpl, dartTpl, arbTpl)
	flag.Parse()
	rs.OpenAll()
	defer rs.Close()

	v := extractISO4217()
	filterISO4217(v)

	err := rs.Render(v)
	if err != nil {
		log.Fatalf("executing template:", err)
	}

	log.Println("Done.")
}

type CcyNtry struct {
	CtryNm     string
	CcyNm      string
	Ccy        string
	CcyNbr     int
	CcyMnrUnts string
	Unicode    []string `xml:"-"`
}

func (n *CcyNtry) DartUnicode() string {
	if len(n.Unicode) == 0 {
		return "null"
	}
	return `'\u{` + strings.Join(n.Unicode, `}\u{`) + `}'`
}

func (n *CcyNtry) ArbResource(a *arb.Arb) *arb.ArbResource {
	return arb.NewResource(a, n.Ccy, n.CcyNm, &arb.ArbAttributes{
		Type:        "text",
		Description: "ISO 4217 currency name",
	})
}

type ISO_4217 struct {
	XMLName xml.Name   `xml:"ISO_4217"`
	Pblshd  string     `xml:"Pblshd,attr"`
	CcyNtry []*CcyNtry `xml:"CcyTbl>CcyNtry"`
}

func (iso *ISO_4217) Arb() *arb.Arb {
	a := &arb.Arb{
		LastModified: iso8601.Time{time.Now()},
		Locale:       language.English,
		Entity:       entity.UpperCamel(),
		Resources:    make([]*arb.ArbResource, len(iso.CcyNtry)),
	}

	for i, n := range iso.CcyNtry {
		a.Resources[i] = n.ArbResource(a)
	}
	return a
}

func extractISO4217() *ISO_4217 {
	log.Println("Get", listUrl)
	resp, err := http.Get(listUrl)
	if err != nil {
		log.Fatalf("http: %v\n", err)
	}
	defer resp.Body.Close()

	log.Println("Parse response")
	var v ISO_4217
	err = xml.NewDecoder(resp.Body).Decode(&v)
	if err != nil {
		log.Fatalf("xml: %v\n", err)
	}
	return &v
}

func filterISO4217(v *ISO_4217) {
	total := len(v.CcyNtry)
	s := make([]*CcyNtry, 0, total)
	exist := make(map[int]bool, total)
	exist[0] = true
	for _, e := range v.CcyNtry {
		if !exist[e.CcyNbr] {
			exist[e.CcyNbr] = true
			if e.Ccy != emptyZero {
				s = append(s, e)
			}
		}
	}

	sort.Slice(s, func(i, j int) bool { return s[i].CcyNbr < s[j].CcyNbr })
	v.CcyNtry = s
}
