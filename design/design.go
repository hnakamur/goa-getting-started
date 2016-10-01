package design // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("cellar", func() { // API defines the microservice endpoint and
	Title("仮想ワインセラー")             // other global properties. There should be one
	Description("シンプルなgoaサービスの例") // and exactly one API definition appearing in
	Scheme("http")                // the design.
	Host("localhost:8080")
})

var _ = Resource("bottle", func() { // Resources group related API endpoints
	BasePath("/bottles")      // together. They map to REST resources for REST
	DefaultMedia(BottleMedia) // services.

	Action("show", func() { // Actions define a single API endpoint together
		Description("指定されたIDのボトルを取得する") // with its path, parameters (both path
		Routing(GET("/:bottleID"))      // parameters and querystring values) and payload
		Params(func() {                 // (shape of the request body).
			Param("bottleID", Integer, "ボトル ID")
		})
		Response(OK)       // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
	})
})

// BottleMedia defines the media type used to render bottles.
var BottleMedia = MediaType("application/vnd.goa.myexample.bottle+json", func() {
	Description("ワインボトル")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "唯一なボトルID")
		Attribute("href", String, "このボトルにリクエストを送るためのAPIのhref")
		Attribute("name", String, "ワインの名前")
		Required("id", "href", "name")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id")   // Media types may have multiple views and must
		Attribute("href") // have a "default" view.
		Attribute("name")
	})
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET") // Allow all origins to retrieve the Swagger JSON (CORS)
	})
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swaggerui/*filepath", "swaggerui/dist")
})
