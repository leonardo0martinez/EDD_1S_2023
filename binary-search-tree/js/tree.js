
// CLASE NODO
class Tnode{
    constructor(value){
        this.value = value;
        // APUNTADORES
        
        this.right = null;
        this.left = null;
    }
}


let connections = "";
let nodes = "";

class Tree{
    constructor(){
        this.root = null;
    }

    add(value){
        const node = new Tnode(value);
        // VERIFICA SI ESTA VACIA
        if(this.root == null){
            this.root = node;
        }else{
            this.#recursive(this.root, node);
        }
    }

    #recursive(curr, newNode){
        if(newNode.value < curr.value ){
            if(curr.left != null){
                this.#recursive(curr.left, newNode);
            }else{
                curr.left = newNode;
            }
        }else if(newNode.value > curr.value){
            if(curr.right != null){
                this.#recursive(curr.right, newNode);
            }else{
                curr.right = newNode;
            }
        }else{
            // NO SE HACE NADA
            console.log("El valor ya está en el árbol");
        }
    }

    treeGraph(){
        connections = "";
        nodes = "";
        this.#graphRecursive(this.root);
        console.log(nodes, connections);
    }

    #graphRecursive(current){
        if(current.left != null){
            this.#graphRecursive(current.left);
            connections += `S${current.value} -> S${current.left.value}\n`;
        }
        nodes += `S${current.value}[label=${current.value}];`
        if(current.right != null){
            this.#graphRecursive(current.right);
            connections += `S${current.value} -> S${current.right.value}\n`;
        }
    }   


    inOrder(){ // IZQUIERDA - RAÍZ - DERECHA
        connections = "";
        nodes = "";
        this.#inOrderRecursive(this.root);
        console.log(nodes, connections);
    }
    
    #inOrderRecursive(current){
        if(current.left != null){
            this.#inOrderRecursive(current.left);
            connections += ` -> `;
        }
        nodes += `S${current.value}[label=${current.value}];\n`
        connections += `S${current.value}`
        if(current.right != null){
            connections += ` -> `;
            this.#inOrderRecursive(current.right);
        }
    }

    postOrder(){ // IZQUIERDA - DERECHA - RAIZ
        connections = "";
        nodes = "";
        this.#postOrderRecursive(this.root);
        console.log(nodes, connections);
    }

    #postOrderRecursive(current){
        if(current.left != null){
            this.#postOrderRecursive(current.left);
            connections += ` -> `;
        }
        if(current.right != null){
            this.#postOrderRecursive(current.right);
            connections += ` -> `;
        }
        nodes += `S${current.value}[label=${current.value}];\n`
        connections += `S${current.value}`
    }

    preOrder(){
        connections = "";
        nodes = "";
        this.#preOrderRecursive(this.root);
        console.log(nodes, connections);
    }
    
    #preOrderRecursive(current){
        nodes += `S${current.value}[label=${current.value}];\n`
        connections += `S${current.value}`
        if(current.left != null){
            connections += ` -> `;
            this.#preOrderRecursive(current.left);
        }
        if(current.right != null){
            connections += ` -> `;
            this.#preOrderRecursive(current.right);
        }
    }
}

module.exports = Tree;