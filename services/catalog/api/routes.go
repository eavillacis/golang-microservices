package api

func (a *API) initRoutes() {
	// Brands
	a.route.POST("/brands", a.CreateBrand)
}
