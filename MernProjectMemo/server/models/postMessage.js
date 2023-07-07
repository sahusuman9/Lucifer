import mongoose from "mongoose";

// specifies the type what the posts are gonna have.
const postSchema = mongoose.Schema({
    title: String,
    message: String,
    creater: String,
    tags: [String],
    selectedFile: String,
    likeCount: {
        type: Number,
        default: 0
    },
    createdAt: {
        type: Date,
        default: new Date()
    },
});

//turn schema into a model
const PostMessage = mongoose.model('PostMessage', postSchema);

export default PostMessage;