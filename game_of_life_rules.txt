// -- design features
// :: 2D grid of square cells
// :: 2 states - dead or alive
// :: they interact with their surrouding neighbours
// horizontally, vertically and diagonally for every world tick
// b b b
// b a b
// b b b

// -- rules per tick
// :: any live cell with 2 or 3 live neighbours lives
// :: any dead cell with 3 live neighbours becomes a live cell
// :: all other live cells dies and dead cells stay dead
