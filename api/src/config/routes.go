package routes

import (
    "api/src/controller/books"
    "api/src/controller/categories"
)

func Route(url string) string {
	routesDict := getRouteDict()
    callback, isExist := routesDict[url]
    if !isExist {
        return "error"
    }
    return callback()
}

func getRouteDict() (map[string]func() string) {
    return map[string]func() string {
        "/": func() string {
            return "a"
        },
        "/books": func() string {
            return booksController.GetIndex()
        },
        "/categories": func() string {
        	return categoriesController.GetIndex()
        },
    }
}
