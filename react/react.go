package react

type reactor struct {
	cells []*cell
}

func New() *reactor {
	return &reactor{}
}

func (r *reactor) update() {
	for _, c := range r.cells {
		if c.update() {
			for _, cb := range c.callbacks {
				_cb := *cb.f
				_cb(c.Value())
			}
		}
	}
}

func (cell *cell) remove(c *canceler) {
	newCallbacks := []*canceler{}
	for _, callback := range cell.callbacks {
		if callback != c {
			newCallbacks = append(newCallbacks, callback)
		}
	}
	cell.callbacks = newCallbacks
}

func (r *reactor) CreateInput(value int) InputCell {
	c := &cell{
		reactor: r,
		value: value,
	}
	c.update = func() bool {
		state := c.shouldUpdate
		c.shouldUpdate = false
		return state
	}
	r.cells = append(r.cells, c)
	return c
}

func (r *reactor) CreateCompute1(c Cell, f func(int) int) ComputeCell {
	g := func(cells []Cell) int {
		return f(cells[0].Value())
	}

	return r.createComputeGeneral([]Cell{c}, g)
}

func (r *reactor) CreateCompute2(c1, c2 Cell, f func(int, int) int) ComputeCell {
	g := func(cells []Cell) int {
		return f(cells[0].Value(), cells[1].Value())
	}

	return r.createComputeGeneral([]Cell{c1, c2}, g)
}

func (r *reactor) createComputeGeneral(cells []Cell, f func([]Cell) int) ComputeCell {
	computeCell := r.CreateInput(0).(*cell)
	computeCell.update = func() bool {
		value := computeCell.Value()
		computeCell.value = f(cells)
		return value != computeCell.Value()
	}
	computeCell.update()
	return computeCell
}

type inputCell struct {
	value int
}

func (input *inputCell) Value() int {
	return input.value
}

func (input *inputCell) SetValue(value int) {
	input.value = value
}

type cell struct {
	reactor *reactor
	value int
	update func() bool
	shouldUpdate bool
	callbacks []*canceler
}

type callback *func(int)

func (c *cell) Value() int {
	return c.value
}

func (c *cell) SetValue(value int) {
	c.shouldUpdate = value != c.value
	c.value = value
	c.reactor.update()
}

func (c *cell) AddCallback(f func(int)) Canceler {
	_canceler := &canceler {
		cell: c,
		f: &f,
	}
	c.callbacks = append(c.callbacks, _canceler)
	return _canceler
}

type canceler struct {
	cell *cell
	f *func(int)
}

func (c *canceler) Cancel() {
	c.cell.remove(c)
}