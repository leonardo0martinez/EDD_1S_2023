
const Matriz = require("./sparase-matrix");

let m = new Matriz();
m.insert(10,2,`Pos(10,2)`);
m.insert(3,1,`Pos(3,1)`);
m.insert(4,1,`Pos(3,1)`);
m.insert(3,2,`Pos(3,1)`);
m.insert(5,2,`Pos(5,2)`);
m.insert(5,6,`Pos(5,6)`);
m.insert(6,8,`Pos(6,8)`);
m.insert(8,3,`Pos(8,3)`);
m.insert(7,9,`Pos(7,9)`);
m.insert(8,4,`Pos(8,4)`);
m.insert(7,3,`Pos(7,3)`);
m.insert(3,6,`Pos(3,6)`);
m.insert(6,4,`Pos(6,4)`);
console.log("CABECERAS EN X:");
m.printX();
console.log("CABECERAS EN Y:");
m.printY();
console.log("GRAPHVIZ");
m.graph()

// console.log( Math.floor(Math.random()*(10 - 1)+1))

