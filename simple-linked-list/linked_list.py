#IMPORTAR NODO
from node import Node

class LinkedList:
    
    # DEFINIR EL INICIO DE LA LISTA
    def __init__(self):
        self.head = None
    
    def insert(self, value): #INSERTAR AL ULTIMO ESPACIO
        #INSERTAR NODO
        new_node = Node(value)
        #1ER PASO
        #Verificar que la lista este vacía
        if self.head == None :
            #Asignar inicio el nuevo nodo
            self.head = new_node
        else:
            #2do Paso
            #Recorrer la lista hasta encontrar el apuntodor nulo
            temp_node = self.head
            while(temp_node.next is not None):
                temp_node = temp_node.next
            #INSERCIÓN AL ÚLTIMO
            temp_node.next = new_node

    def push(self, value): #INSERTAL AL INICIO
        #INSERTAR NODO
        new_node = Node(value)
        #1ER PASO
        #Verificar que la lista este vacía
        if self.head == None :
            #Asignar inicio el nuevo nodo
            self.head = new_node
        else:
            #2do Paso
            #Nodo inicial se 'corre' al apuntador siguiente
            temp_node = self.head # Valor de la cabeza
            self.head = new_node # Asignar la nueva cabeza
            self.head.next = temp_node # Corre el nodo viejo para el nuevo

    def toString(self):       
        temp_node = self.head
        string = ""
        while(temp_node.next is not None):
            string += str(temp_node.value) + ", "
            temp_node = temp_node.next
        string += str(temp_node.value)
        print(string)
    

    def graph(self):
        nodes = ""
        conn = ""
        temp_node = self.head
        counter = 0
        while(temp_node is not None):
            nodes += 'N'+str(counter)+'[ label="'+str(temp_node.value)+'" ]\n'
            conn += 'N'+str(counter)+' -> '
            temp_node = temp_node.next
            counter += 1
        conn += 'NULL'
        print(nodes)
        print(conn)



        
