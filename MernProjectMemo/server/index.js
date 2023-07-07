import express from 'express';
import bodyParser from 'body-parser';
import mongoose from 'mongoose';
import cors from 'cors';

import postsRoutes from './routes/posts.js';

//initialise the app
const app = express();

app.use('/posts', postsRoutes);     //every route inside of postsRoutes is gonna start with posts

//general settings

app.use(bodyParser.json({ linit: "30mb", extended: true }));
app.use(bodyParser.urlencoded({ linit: "30mb", extended: true })); // setting up the bodyParser to send requests
app.use(cors());

const CONNECTION_URL = 'mongodb+srv://sumanmern:sumanmern9090@clusters.ogir9dc.mongodb.net/'
const PORT = process.env.PORT || 5000;

mongoose.connect(CONNECTION_URL, { useNewUrlParser: true, useUnifiedTopology: true }).then(() => app.listen(PORT, () => console.log(`Server is running on port: ${PORT}`))).catch((error) => console.log(error.message));             // connect with the database


// http://www.mongodb.com/cloud/atlas