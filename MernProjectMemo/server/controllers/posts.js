//will have all the handlers

import PostMessage from "../models/postMessage.js";

export const getPosts = async (req, res) => {
    try {
        const postMessages = await PostMessage.find();     // as finding takes time, its an asynchronous action so add await and make the funtion asynchronous by adding async in function

        res.status(200).json(postMessages);     //returns postMessages, 200 means its fine

    } catch (error) {
        res.status(404).json({ message: error.message });
    }
}

export const createPost = (req, res) => {

    const post = req.body;
    
    //create new post
    const newPost = new PostMessage(post);

    try {
        
    } catch (error) {
        
    }
}