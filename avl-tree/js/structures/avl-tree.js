//--------------------------------------------------------------------------
//                      CLASE NODO
//--------------------------------------------------------------------------
class AvlNode{
    constructor(item){
        this.item = item;
        this.left = null;
        this.right = null;
        this.height = 0;
    }
}

//--------------------------------------------------------------------------
//                   VARIABLES GLOBALES
//--------------------------------------------------------------------------
let = nodes = "";
let = connections = "";

//--------------------------------------------------------------------------
//                   CLASE ARBOL AVL
//--------------------------------------------------------------------------
class AvlTree{
    constructor(){
        this.root = null;
    }

    insert(item){
        this.root = this.#insertRecursive(item, this.root);
    }

    getHeight(node){
        return node === null ? -1 : node.height;
    }
    getMaxHeight(leftNode, rightNode){
        return leftNode.height > rightNode.height ? leftNode.height : rightNode.height;
    }

    //--------------------------------------------------------------------------
    //                  METODO DE INSERCIÓN
    //--------------------------------------------------------------------------
    #insertRecursive(item, node){
        if(node == null){
            node = new AvlNode(item);
        }else if(item.carnet < node.item.carnet){
            node.left = this.#insertRecursive(item, node.left);
            if(this.getHeight(node.left) - this.getHeight(node.right) == 2){
                if(item.carnet < node.left.item.carnet){
                    node = this.#rotateLeft(node);
                }else{
                    node = this.#doubleLeft(node);
                }
            }
        }else if(item.carnet > node.item.carnet){
            node.right = this.#insertRecursive(item, node.right);
            if(this.getHeight(node.right) - this.getHeight(node.left) == 2){
                if(item.carnet < node.right.item.carnet){
                    node = this.#rotateRight(node);
                }else{
                    node = this.#doubleRight(node);
                }
            }
        }else{
            alert("Elemento ya existe en el árbol");
        }
        node.height = this.getMaxHeight(this.getHeight(node.left), this.getHeight(node.right)) + 1;
        return node;
    }

    //--------------------------------------------------------------------------
    //                   ROTACIONES
    //--------------------------------------------------------------------------
    #rotateRight(node1){
        node2 = node1.right;
        node1.right = node2.left;
        node2.left = node1;
        node1.height = this.getMaxHeight(this.getHeight(node1.left), this.getHeight(node1.right)) + 1;
        node2.height = this.getMaxHeight(this.getHeight(node2.right), node1.height) + 1;
        return node2;
    }
    #rotateLeft(node2){
        node1 = node2.left;
        node2.left = node1.right;
        node1.right = node2;
        node2.height = this.getMaxHeight(this.getHeight(node2.left), this.getHeight(node2.right)) + 1;
        node1.height = this.getMaxHeight(this.getHeight(node1.left), node2.height) + 1;
        return node1;
    }
    #doubleLeft(node){
        node.left = this.#rotateRight(node.left);
        return this.#rotateLeft(node);
    }
    #doubleRight(node){
        node.right = this.#rotateLeft(node.right);
        return this.#rotateRight(node);
    }

    //--------------------------------------------------------------------------
    //                  REPORTE DEL ARBOL
    //--------------------------------------------------------------------------
    treeGraph(){       
        nodes = "";
        connections = "";
        this.#treeGraphRecursive(this.root);
        console.log(nodes,connections);
        return nodes + connections;
    }
    #treeGraphRecursive(current){
        if(current.left != null){
            this.#treeGraphRecursive(current.left);
            connections += `S_${current.item.carnet} -> S_${current.left.item.carnet};\n`;
        }
        nodes += `S_${current.item.carnet}[label="${current.item.nombre}"];`
        if(current.right != null){
            this.#treeGraphRecursive(current.right);
            connections += `S_${current.item.carnet} -> S_${current.right.item.carnet};\n`;
        }
    }
    
    //--------------------------------------------------------------------------
    //                  RECORRIDO IN ORDER
    //--------------------------------------------------------------------------
    inOrder(){
        let html = this.#inOrderRecursive(this.root);
        return html;
    }
    #inOrderRecursive(current){
        let row = "";
        if(current.left != null){
            row += this.#inOrderRecursive(current.left);
        }
        row +=`
            <tr>
                <th>${current.item.carnet}</th>
                <td>${current.item.nombre}</td>
                <td>${current.item.password}</td>
            </tr>
        `;
        if(current.right != null){
            row += this.#inOrderRecursive(current.right);
        }
        return row;
    }
    //--------------------------------------------------------------------------
    //                  RECORRIDO PRE ORDER
    //--------------------------------------------------------------------------
    preOrder(){
        let html = this.#preOrderRecursive(this.root);
        return html;
    }
    #preOrderRecursive(current){
        let row = "";
        row +=`
            <tr>
                <th>${current.item.carnet}</th>
                <td>${current.item.nombre}</td>
                <td>${current.item.password}</td>
            </tr>
        `;
        if(current.left != null){
            row += this.#inOrderRecursive(current.left);
        }
        if(current.right != null){
            row += this.#inOrderRecursive(current.right);
        }
        return row;
    }

    //--------------------------------------------------------------------------
    //                  RECORRIDO POST ORDER
    //--------------------------------------------------------------------------
    postOrder(){
        let html = this.#postOrderRecursive(this.root);
        return html;
    }
    #postOrderRecursive(current){
        let row = "";
        if(current.left != null){
            row += this.#inOrderRecursive(current.left);
        }
        if(current.right != null){
            row += this.#inOrderRecursive(current.right);
        }
        row +=`
            <tr>
                <th>${current.item.carnet}</th>
                <td>${current.item.nombre}</td>
                <td>${current.item.password}</td>
            </tr>
        `;
        return row;
    }


}