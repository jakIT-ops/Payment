package main

//
//func main() {
//	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
//	if err != nil {
//		panic(err)
//	}
//	client := pb.NewTransactionClient(conn)
//
//	// g := gin.Default()
//	app := fiber.New()
//
//	app.Post("/account/:id/credit", func(c *fiber.Ctx) error {
//		id := c.Params("id")
//		var req pb.TransactionRequest
//		if err := c.BodyParser(&req); err != nil {
//			return c.Status(fiber.StatusBadRequest).JSON(
//				fiber.Map{
//					"success": false,
//					"error":   err.Error(),
//				},
//			)
//		}
//		req.AccountId = id
//		req.DebitAmount = 0
//		if res, err := client.CreditAmount(context.Background(), &req); err == nil {
//			return c.Status(fiber.StatusOK).JSON(fiber.Map{
//				"success": true,
//				"resp":    res,
//			})
//		}
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"success": false,
//			"error":   err.Error(),
//		})
//	})
//
//	app.Post("/account/:id/debit", func(c *fiber.Ctx) error {
//		id := c.Params("id")
//		var req pb.TransactionRequest
//		if err := c.BodyParser(&req); err != nil {
//			return c.Status(fiber.StatusBadRequest).JSON(
//				fiber.Map{
//					"success": false,
//					"error":   err.Error(),
//				},
//			)
//		}
//		req.AccountId = id
//		req.CreditAmount = 0
//		if res, err := client.DebitAmount(context.Background(), &req); err == nil {
//			return c.Status(fiber.StatusOK).JSON(fiber.Map{
//				"success": true,
//				"resp":    res,
//			})
//		}
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"success": false,
//			"error":   err.Error(),
//		})
//	})
//
//	app.Get("/account/:id/balance", func(c *fiber.Ctx) error {
//		id := c.Params("id")
//		var req pb.Account_Id
//		req.Value = id
//		// req := &pb.Account_Id{Value: id}
//		if res, err := client.GetBalance(context.Background(), &req); err == nil {
//			return c.Status(fiber.StatusOK).JSON(fiber.Map{
//				"success": true,
//				"Balance": res,
//			})
//		}
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"success": false,
//			"error":   err.Error(),
//		})
//	})
//	log.Fatal(app.Listen(":3000"))
//}
