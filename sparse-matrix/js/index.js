
let matrix = new SparseMatrix();


function handleSubmit(e){

    
    e.preventDefault();
    const formData = new FormData(e.target);
    const form = Object.fromEntries(formData);
    
    if(localStorage.getItem("matriz") !== null){
        let temp = JSON.retrocycle(JSON.parse(localStorage.getItem("matriz")));
        matrix.head = temp.head;    
    }

    try{
        matrix.insert(Number(form.xpos), Number(form.ypos), form.value);
        alert("Todo bien :)")
    }catch(error){
        alert("Error en la conversion")
        console.log(error);
    }

    // MANEJAR AUTOREFERENCIA DE LAS ESTRUCTURAS
    localStorage.setItem("matriz", JSON.stringify(JSON.decycle(matrix)));
    
    //localStorage.clear()
}

function showGraph(){
    let temp = JSON.retrocycle(JSON.parse(localStorage.getItem("matriz")));
    matrix.head = temp.head;
    let url = 'https://quickchart.io/graphviz?graph=';
    let body = `digraph G { ${matrix.graph()} }`
    $("#graph").attr("src", url + body);
}






