# Principio de segregacion de interfaz

El principio establece que muchas interfaces específicas del cliente son mejores que una interfaz de propósito general. 
No se debe obligar a los clientes a implementar una función que no necesitan.

## En la implementacion

Se tiene una interfaz llamada _*machine*_ que cuenta con diversos metodos (print, scan, etc...). A su vez esta es
implementada por una estructura llamada _*multifunctions*_ que al tener la capacidad de poder hacer muchas tareas, 
puede implementar todos los metodos que trae la interfaz. Sin embargo para el caso de las demas estructuras que
implementan a la interfaz, asi como lo es _old_fashioned_printer_ que no puede hacer todas las operaciones que tiene 
esta interfaz, por lo tanto segun el principio, es mucho mejor separar las responsabilidades de esa interfaz. Con lo que
se separaron los metodos de la interfaz original para llevarlos a dos interfaces mas acotadas (_Printer_ y _Scanner_)
las cuales como en el caso de photocopier, estan asi lo son.
