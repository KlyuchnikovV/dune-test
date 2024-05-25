package main

import (
	"bytes"
	_ "embed"
	"io"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/render"
	tpls "github.com/go-echarts/go-echarts/v2/templates"
)

var (
	//go:embed templates/header.tpl
	HeaderTpl string
	//go:embed templates/base.tpl
	BaseTpl string
)

type Renderer struct {
	render.BaseRender
	*charts.Bar
	Table [][]string
}

func NewRenderer(bar *charts.Bar, table ...[]string) render.Renderer {
	return &Renderer{
		Bar:   bar,
		Table: table,
	}
}

func (r *Renderer) Render(w io.Writer) error {
	r.Bar.Validate()

	var (
		contents = []string{HeaderTpl, BaseTpl, tpls.ChartTpl}
		tpl      = render.MustTemplate("chart", contents)
	)

	var buf bytes.Buffer
	if err := tpl.ExecuteTemplate(&buf, "chart", r); err != nil {
		return err
	}

	_, err := w.Write(buf.Bytes())
	return err
}
