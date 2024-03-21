package handler

// func signUpHandler(c *fiber.Ctx) error {
// 	user := User{}
// 	err := c.BodyParser(&user)
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).SendString("Invalid data")
// 	}

// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MaxCost)
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).SendString("Error hashing password")
// 	}

// 	user.Password = string(hashedPassword)

// 	stmt, err := db.Prepare("INSERT INTO users (username, password, email) VALUES (?, ?, ?)")
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).SendString("Error preparing statement")
// 	}

// 	_, err = stmt.Exec(user.Username, user.Password, user.Email)
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).SendString("Error saving user")
// 	}

// 	return c.SendString("Success")
// }
