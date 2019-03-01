package service

import (
	"context"
	"net/http"

	"github.com/graphql-go/graphql"

	"github.com/gin-gonic/gin"
)

type GraphCore struct {
	s      store
	schema *graphql.Schema
}
type queryRequest struct {
	query string `json:"query"`
}

func NewGraphCore(s store) *GraphCore {
	return &GraphCore{
		s: s,
	}
}

//start in the goroutine
func (g *GraphCore) Start() http.Handler {
	r := Resolver{s: g.s}
	root := NewRoot(&r)
	sc, err := graphql.NewSchema(graphql.SchemaConfig{Query: root.object})
	if err != nil {
		panic(err)
	}
	g.schema = &sc
	return g.graphEntry()
}
func (g *GraphCore) graphEntry() http.Handler {
	e := gin.New()
	e.Any("/graph", g.handleQuery)
	return e
}

func (g *GraphCore) handleQuery(ctx *gin.Context) {
	var qr queryRequest
	if err := ctx.BindJSON(&qr); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	r := exeuteQuery(ctx.Request.Context(), qr.query, *g.schema)
	ctx.JSON(http.StatusOK, r)
}

func exeuteQuery(ctx context.Context, query string, schema graphql.Schema) *graphql.Result {
	return graphql.Do(graphql.Params{Schema: schema, RequestString: query, Context: ctx})
}
