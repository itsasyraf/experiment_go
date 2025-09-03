package apidoc

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3gen"
	"github.com/gofiber/fiber/v2"
)

type Route struct {
	Method, Path, Summary string
	ReqModel              interface{}
	ResModel              interface{}
	ResExample            map[string]interface{}
	Tags                  []string
}

type ResponseDef struct {
    Model   interface{}
    Example map[string]interface{}
}

var routes []Route

// Add simple route
func Add(app *fiber.App, method, path, summary string, handler fiber.Handler, tags ...string) {
	routes = append(routes, Route{Method: method, Path: path, Summary: summary, Tags: tags})
	app.Add(method, path, handler)
}
// Register a route with request/response models + example
func AddWithModels(app *fiber.App, method, path, summary string,
	handler fiber.Handler,
	reqModel interface{}, resModel interface{}, resExample map[string]interface{}, tags ...string,
) {
	routes = append(routes, Route{method, path, summary, reqModel, resModel, resExample, tags})
	app.Add(method, path, handler)
}

func BuildSwagger() *openapi3.T {
    doc := &openapi3.T{
        OpenAPI: "3.0.0",
        Info: &openapi3.Info{
            Title:   "Dynamic Fiber API",
            Version: "1.0.0",
        },
        Paths:      openapi3.NewPaths(), // ✅ correct
        Components: &openapi3.Components{Schemas: make(map[string]*openapi3.SchemaRef)},
    }

    for _, r := range routes {
        op := &openapi3.Operation{
            Summary:   r.Summary,
            Responses: openapi3.NewResponses(), // ✅ correct
            Tags:      r.Tags,
        }

        // Request body
        if r.ReqModel != nil {
            schemaRef, _ := openapi3gen.NewSchemaRefForValue(r.ReqModel, nil)
            op.RequestBody = &openapi3.RequestBodyRef{
                Value: &openapi3.RequestBody{
                    Required: true,
                    Content: openapi3.NewContentWithJSONSchemaRef(schemaRef),
                },
            }
        }

        // Response
        for code, res := range r.ResExample {
            schemaRef, _ := openapi3gen.NewSchemaRefForValue(r.ResModel, nil)
            if res  != nil {
                schemaRef.Value.Example = res
            }

            desc := "Response " + code
            op.Responses.Set(code, &openapi3.ResponseRef{
                Value: &openapi3.Response{
                    Description: &desc,
                    Content: openapi3.NewContentWithJSONSchemaRef(schemaRef),
                },
            })
        }

        item := &openapi3.PathItem{}
        switch r.Method {
        case fiber.MethodGet:
            item.Get = op
        case fiber.MethodPost:
            item.Post = op
        case fiber.MethodPut:
            item.Put = op
        case fiber.MethodDelete:
            item.Delete = op
        }

        doc.Paths.Set(r.Path, item) // ✅ correct
    }

    return doc
}
