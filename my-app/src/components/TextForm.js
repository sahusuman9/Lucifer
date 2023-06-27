import React, {useState} from 'react'


export default function TextForm(props) {

const handleUpClick = ()=> {
    //console.log("UpperCase was clicked" + text);
    let newText = text.toUpperCase();
    setText(newText)
}

const handleLowClick = ()=> {
    //console.log("UpperCase was clicked" + text);
    let newText = text.toLowerCase();
    setText(newText)
}

const handleOnChange = (event)=> {
    //console.log("OnChange");
    setText(event.target.value);
}


    //hooks and state, to change the state use setText in this case.
    //setText("Enter enother text");
const [text, setText] = useState('');
  return (
    <>
    <div className="container">
        <h1>{props.heading} </h1>
        <div className="mb-3">
        <textarea className="form-control" value={text} onChange={handleOnChange} id="mybox" rows="8"></textarea>
        </div>
        <button className="btn btn-primary mx-2" onClick={handleUpClick}>Convert To Uppercase</button>
        <button className="btn btn-primary" onClick={handleLowClick}>Convert To LowerCase</button>
    </div>
    <div className="container">
        <h2>Your text Summary</h2>
        <p>{text.split(" ").length} words and {text.length} characters</p>
        <p>{0.008 * text.split(" ").length} Minutes Reading</p>
        <h2>Preview</h2>
        <p>{text}</p>
    </div>
    </>
  )
}
