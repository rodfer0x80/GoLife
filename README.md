# gameOfLife
Golang implementation of Conway's Game of Life 

### Design features
* 2D grid of square cells
* 2 states - dead or alive
* cells interact with their surrouding neighbours
* horizontally, vertically and diagonally for every world tick 

### Rules per tick
* any live cell with 2 or 3 live neighbours lives
* any dead cell with 3 live neighbours becomes a live cell
* all other live cells dies and dead cells stay dead
