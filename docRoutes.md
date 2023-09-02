
# **Gianni Jiang**

API's documentation of the application


## **Routes**



### **/**

    


#### **GET**
* [_github.com/Jiang-Gianni/gianni-jiang/server.(*Server).GetIndex_](/server/index.go#L22)


<br>
<br>



### **/alpine**

    


#### **GET**
* [_github.com/Jiang-Gianni/gianni-jiang/server.(*Server).GetAlpine_](/server/alpine.go#L9)


<br>
<br>



### **/assets/***

    


#### **GET**
* [_github.com/Jiang-Gianni/gianni-jiang/server.(*Server).GetAssets_](/server/index.go#L17)


<br>
<br>



### **/gmt/result**

    


#### **GET**
* [_github.com/Jiang-Gianni/gianni-jiang/server.(*Server).GetGmtResult_](/server/index.go#L39)


<br>
<br>



### **/new/todo**

    


#### **GET**
* [_github.com/Jiang-Gianni/gianni-jiang/server.(*Server).GetNewTodo_](/server/todo.go#L80)


<br>
<br>



### **/number**

    


#### **GET**
* [_github.com/Jiang-Gianni/gianni-jiang/server.(*Server).GetNumbersApi_](/server/numbersApi.go#L17)


<br>
<br>



### **/todo/**

    


#### **GET**
* [_github.com/Jiang-Gianni/gianni-jiang/server.(*Server).GetTodo_](/server/todo.go#L21)

#### **POST**
* [_github.com/Jiang-Gianni/gianni-jiang/server.(*Server).PostTodo_](/server/todo.go#L85)


<br>
<br>



### **/todo/{todoId}/**

#### **MIDDLEWARES**

    
* [_github.com/Jiang-Gianni/gianni-jiang/server.(*Server).TodoCtx_](/server/todo.go#L25)


#### **POST**
* [_github.com/Jiang-Gianni/gianni-jiang/server.(*Server).PostTodoId_](/server/todo.go#L53)

#### **DELETE**
* [_github.com/Jiang-Gianni/gianni-jiang/server.(*Server).DeleteTodoId_](/server/todo.go#L71)

#### **GET**
* [_github.com/Jiang-Gianni/gianni-jiang/server.(*Server).GetTodoId_](/server/todo.go#L47)


<br>
<br>



### **/{project}/**

    


#### **GET**
* [_github.com/Jiang-Gianni/gianni-jiang/server.(*Server).GetProject_](/server/index.go#L26)


<br>
<br>



### **/{project}/feedback**

    


#### **POST**
* [_github.com/Jiang-Gianni/gianni-jiang/server.(*Server).PostFeedback_](/server/feedback.go#L11)


<br>
<br>



### **/{project}/{section}**

    


#### **GET**
* [_github.com/Jiang-Gianni/gianni-jiang/server.(*Server).GetProject_](/server/index.go#L26)


<br>
<br>



