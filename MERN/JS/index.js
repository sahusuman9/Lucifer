//objects
var marks = {
    ravi: 34,
    subham: 78,
    harry: 99.977
}

console.log(marks);

//arrays

var arr = [1,2,3,4,5]
console.log(arr);

// Logical operators
console.log(true && true); // true
console.log(true && false);// false
console.log(false && true);// false
console.log(false && false);// false


console.log(true || true); // true
console.log(true || false);// true
console.log(false || true);// true
console.log(false || false);// false


//functions

function avg(a, b) {
    return (a + b)/2;
}

c1 = avg(4, 6)
c2 = avg(4, 6)

// function invoke

function avg(a, b) {
    c = (a + b)/2;
    return c; //function invoke
}
//function call
c3 = avg(4, 6)
c4 = avg(4, 6)


//conditions in JS
var arr1 = [1,2,3,4,5,6,7];
console.log(arr1);

//1st method
for(var i = 0; i < arr1.length; i++) {
    console.log(arr1[i]);
}

//2wnd method
arr1.forEach(function(element) {
    console.log(element);
})