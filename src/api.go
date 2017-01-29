package main

import (
  "net/http"

  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)
type User struct{
 Firstname string `json:"firstname,omitempty"`
 Lastname string `json:"lastname,omitempty"`
 Username string `json:"username,omitempty"`
 Password string `json:"password,omitempty"`
}
func getUserByID(c echo.Context) error {
  id := c.Param("id")
  return c.JSON(http.StatusOK, id)
}

func index(c echo.Context) error {
  return c.JSON(http.StatusOK, "Hello World")
}
func getUsers(c echo.Context) error {
 mint := User{
  Firstname: "Pornrika",
  Lastname:"Panyayuang",
  Username:"merxer",

 }
  return c.JSON(http.StatusOK, mint)
}
func saveUser(c echo.Context) error {
  user := new(User)
  err := c.Bind(user)
  if err != nil {
    return c.JSON(http.StatusBadRequest, nil)
  }
  return c.JSON(http.StatusOK, user)
}
func main() {
  e := echo.New()
  e.Use(middleware.Logger())
  e.GET("/users/:id", getUserByID)
  e.GET("/", index)
  e.GET("/users/get", getUsers)
  e.POST("/users", saveUser)
  e.Logger.Fatal(e.Start(":1323"))
}
