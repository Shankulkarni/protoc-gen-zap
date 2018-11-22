package main

import (
	"github.com/lyft/protoc-gen-star"
	"github.com/lyft/protoc-gen-star/lang/go"
	"github.com/shankulkarni/protoc-gen-zap/zap"
)

type zapGen struct {
	pgs.ModuleBase
	pgsgo.Context
}

func (*zapGen) Name() string {
	return "zap"
}

func (m *zapGen) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.Context = pgsgo.InitContext(c.Parameters())
}

func (m *zapGen) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {

	for _, f := range targets {

		name := m.Context.OutputPath(f).SetExt(".zap.go").String()
		fm := fileModel{PackageName: m.Context.PackageName(f).String()}

		for _, msg := range f.AllMessages() {

			fields := msg.Fields()
			mp := messageModel{}
			mp.Name = msg.Name().UpperCamelCase().String()

			list := make([]zapField, len(f.AllMessages()))

			for _, v := range fields {

				redact := false
				_, err := v.Extension(zappb.E_Redact, &redact)
				if err != nil {
					m.Log(err)
				}

				r := zapField{
					Redact:   redact,
					Name:     v.Name().UpperCamelCase().String(),
					Type:     v.Descriptor().Type.String(),
					Label:    v.Descriptor().GetLabel().String(),
					TypeName: v.Descriptor().GetTypeName(),
				}
				list = append(list, r)

			}
			mp.Fields = list
			fm.Messages = append(fm.Messages, mp)
		}

		m.OverwriteGeneratorTemplateFile(
			name,
			T.Lookup("File"),
			&fm,
		)
	}

	return m.Artifacts()

}

type zapField struct {
	Name     string
	Type     string
	Redact   bool
	Label    string
	TypeName string
}

type messageModel struct {
	Name   string
	Fields []zapField
}

type fileModel struct {
	PackageName string
	Messages    []messageModel
}
