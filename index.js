require("dotenv").config()
const express = require('express')
const bodyParser = require('body-parser')
const cors = require('cors')
const morgan = require('morgan')
const cookieParser = require('cookie-parser');

const corsConfig = {
  origin: true,
  credentials: true,
};

const app = express()
app.use(morgan('combined'))
app.use(bodyParser.json())
app.use(cors(corsConfig))
app.use(cookieParser());

require('./controllers')(app)
app.listen(process.env.PORT || 3003)