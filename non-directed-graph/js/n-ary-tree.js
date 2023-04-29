
// CLASE NODO 
class Tnode{
    
    constructor(folderName, weight){
        this.folderName = folderName;
        this.files = [];
        this.children = []; // TODOS LOS NODOS HIJOS
        this.id = null; // PARA GENERAR LA GRÁFICA
        this.weight = weight;
    }
}


class Tree{
    constructor(){
        // nodo raiz
        this.root = new Tnode('/', 1);
        this.root.id = 0;
        this.size = 1; // Para generar los ids
    }

    insert(folderName, fatherPath){ 
        let {node:fatherNode, weight} = this.getFolder(fatherPath);
        let newNode =  new Tnode(folderName, weight);
        // console.log(newNode)
        if(fatherNode){
            this.size += 1;
            newNode.id = this.size;
            fatherNode.children.push(newNode);
        }else{
            console.log("Ruta no existe");
        }
    }


    getFolder(path){
        let weight = 2;
        if(path == this.root.folderName){
            return {node: this.root, weight: weight};
        }else{
            let temp = this.root;
            let folders = path.split('/');
            folders = folders.filter( str => str !== '');
            let folder = null;

            while(folders.length > 0){
                let currentFolder = folders.shift()
                folder = temp.children.find(child => child.folderName == currentFolder);
                if(typeof folder == 'undefined' || folder == null){
                    return null;
                }
                temp = folder;
                weight++;
            }
            return {node: temp, weight: weight}; 
        }
    }

    graph(){
        let nodes = "";
        let connections = "";

        let node = this.root;
        let queue = [];
        queue.push(node);
        while(queue.length !== 0){
            let len = queue.length;
            for(let i = 0; i < len; i ++){
                let node = queue.shift();
                nodes += `S_${node.id}[label="${node.folderName}"];\n`;
                node.children.forEach( item => {
                    connections += `S_${node.id} -> S_${item.id} [label="${node.weight}"];\n`
                    queue.push(item);
                });
            }
        }
        // sep="+10,20"; \noverlap=scalexy;
        return  '\nlayout=neato; \nedge[dir=none];\n' + nodes +'\n'+ connections;
    }

    getHTML(path){
        let { node } = this.getFolder(path);
        let code = "";
        node.children.map(child => {
            code += ` <div class="col-2 folder" onclick="entrarCarpeta('${child.folderName}')">
                        <img src="./imgs/folder.png" width="100%"/>
                        <p class="h6 text-center">${child.folderName}</p>
                    </div>`
        })
        // console.log(node.files)
        node.files.map(file => {
            if(file.type === 'text/plain'){
                let archivo = new Blob([file.content], file.type);
                const url = URL.createObjectURL(archivo);
                code += `
                        <div class="col-2 folder">
                        <img src="./imgs/file.png" width="100%"/>
                        <p class="h6 text-center">
                            <a href="${url}" download>
                                ${file.name}
                            </a>
                        </p>
                    </div>
                `
            }else{
                code += ` <div class="col-2 folder">
                        <img src="./imgs/file.png" width="100%"/>
                        <p class="h6 text-center">
                            <a href="${file.content}" download>
                                ${file.name}
                            </a>
                        </p>
                    </div>`
            }
        })
        return code;
    }
}