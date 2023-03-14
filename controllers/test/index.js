const { Router } = require("express");
const router = Router();

router.get("/", (req, res, next) => {
  console.log("test");
  res.status(200).send({ message: "Hello World!" });
});

module.exports = router;
