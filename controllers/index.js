const path = require("path")
const fs = require('fs')

const routes = {
    "/register": require("./register"),
    "/test": require("./test")
}

module.exports = (app) => {
    for (const [ routeName, route ] of Object.entries(routes)) {
        app.use(routeName, route)
    }
}